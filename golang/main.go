package main

import (
	"context"
	"errors"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	apis "tensorflow_serving/apis"
	"time"

	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	"github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
	"google.golang.org/grpc"

	"github.com/northvolt/firefly-code-scan/lib/sections"
	"gocv.io/x/gocv"
)

var (
	// Tensorflow serving grpc address.
	address = "127.0.0.1:8500"
)

const (
	defaultStride    = 32
	defaultPatchSize = 64
	defaultBatchSize = 256
)

func main() {
	// load image
	rawdata, err := ioutil.ReadFile("/Users/deadprogram/src/github.com/northvolt/tfgotest/2020-07-24T13-15-40.466535072Z.jpg") //openTestDataFile("~/src/github.com/northvolt/tfgotest/2020-07-24T13-15-40.466535072Z.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	// prepare data
	data, err := prepareImage(rawdata)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a grpc connection.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		log.Fatalf("couldn't connect: %s", err.Error())
	}
	defer conn.Close()

	// Wrap the grpc uri with client.
	client := apis.NewPredictionServiceClient(conn)

	for _, batch := range data {
		// Create a grpc request.
		request := &apis.PredictRequest{
			ModelSpec: &apis.ModelSpec{},
			Inputs:    make(map[string]*tensor_go_proto.TensorProto),
		}
		request.ModelSpec.Name = "anomaly-detector"
		request.ModelSpec.SignatureName = "serving_default"
		// request.ModelSpec.VersionChoice = &apis.ModelSpec_VersionLabel{VersionLabel: "stable"}
		// request.ModelSpec.VersionChoice = &apis.ModelSpec_Version{Version: &wrappers.Int64Value{Value: 0}}
		request.Inputs["inputs"] = &tensor_go_proto.TensorProto{
			Dtype: types_go_proto.DataType_DT_FLOAT,
			TensorShape: &tensor_shape_go_proto.TensorShapeProto{
				Dim: []*tensor_shape_go_proto.TensorShapeProto_Dim{
					&tensor_shape_go_proto.TensorShapeProto_Dim{
						Size: int64(256),
					},
					&tensor_shape_go_proto.TensorShapeProto_Dim{
						Size: int64(64),
					},
					&tensor_shape_go_proto.TensorShapeProto_Dim{
						Size: int64(64),
					},
					&tensor_shape_go_proto.TensorShapeProto_Dim{
						Size: int64(1),
					},
				},
			},
			StringVal: [][]byte{batch},
		}

		// Send the grpc request.
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := client.Predict(ctx, request)
		if err != nil {
			log.Fatalf("couldn't get response: %v", err)
		}
		log.Printf("%+v", response)
	}
}

func findSections(img gocv.Mat) (sections.Sections, error) {
	return sections.FindInMat(img, sections.NewSectionParams())
}

func prepareImage(data []byte) ([][]byte, error) {
	img, err := gocv.IMDecode(data, gocv.IMReadGrayScale)
	if err != nil {
		return nil, err
	}

	// left_foil, right_foil = self.parse_sections(d['sections'])
	sections, err := findSections(img)
	if err != nil {
		return nil, err
	}

	// img = self.parse_buffer(img_data)
	coated := img.Region(sections.Coated)
	coatedFloat := gocv.NewMat()
	defer coatedFloat.Close()
	coated.ConvertTo(&coatedFloat, gocv.MatTypeCV32F)

	// img = self.preprocess_img(img, left_foil, right_foil)
	//img = img.astype(np.float32)
	//img /= 255.0
	coatedFloat.DivideFloat(255.0)

	//img = cv.resize(img, img_size, interpolation=cv.INTER_AREA)
	resized := gocv.NewMat()
	defer resized.Close()

	gocv.Resize(coatedFloat, &resized, image.Pt(4096, 256), 0.0, 0.0, gocv.InterpolationArea)

	// img = np.expand_dims(img, axis=[0, -1])

	// patches = self.extract_patches(img)
	patches, err := extractPatches(resized, defaultPatchSize, defaultStride)
	if err != nil {
		return nil, err
	}
	fmt.Println("patches", len(patches))

	// _, n_rows, n_cols, _ = patches.shape
	// patches = tf.reshape(patches, [-1, PATCH_SIZE, PATCH_SIZE, 1])
	// patches_ds = tf.data.Dataset.from_tensor_slices(patches).batch(FLAGS.batch_size)

	batches, err := batchPatches(patches, defaultBatchSize)
	if err != nil {
		return nil, err
	}

	for _, batch := range batches {
		fmt.Println(len(batch))
	}

	// turn all batches into []byte
	results := make([][]byte, 0)
	for _, batch := range batches {
		result := make([]byte, 0)
		for _, mat := range batch {
			result = append(result, mat.ToBytes()...)
		}
		results = append(results, result)
	}

	return results, nil
}

func extractPatches(img gocv.Mat, sz, stride int) ([]gocv.Mat, error) {
	results := make([]gocv.Mat, 0)
	sizes := img.Size()

	for y := 0; y < sizes[0]-sz; y += stride {
		for x := 0; x < sizes[1]-sz; x += stride {
			// do something with image.Region()
			rect := image.Rect(x, y, x+sz, y+sz)
			results = append(results, img.Region(rect))
		}
	}

	return results, nil
}

func batchPatches(imgs []gocv.Mat, batchSize int) ([][]gocv.Mat, error) {
	results := make([][]gocv.Mat, 0)

	fmt.Println("len(imgs)/batchSize", len(imgs)/batchSize)

	for i := 0; i < len(imgs); i += batchSize {
		sz := i + batchSize
		if i+batchSize > len(imgs) {
			sz = len(imgs)
		}
		results = append(results, imgs[i:sz])
	}

	return results, nil
}

func openTestDataFile(filename string) ([]byte, error) {
	// check the ENV var that points to nv-vision-testdata
	if os.Getenv("NV_TEST_DATA") == "" {
		return nil, errors.New("must set NV_TEST_DATA env var to run tests")
	}
	log.Print("Opening ", filename)
	return ioutil.ReadFile(os.Getenv("NV_TEST_DATA") + filename)
}

// stride = ..
// kernel_size = .. # dimension of image patch
// n_x = .. # number of image patches in x direction
// n_y = .. # number of image patches in y direction
// patches = list()
// for j in range(n_y):
//   start_y = j * stride
//   end_y = start_y + kernel_size
//   for i in range(n_x):
//     start_x = i * stride
//     end_x = start_x + kernel_size
//     image_patch = image[start_y:end_y, start_x:end_x, :]
//     patches.append(image_patch)
