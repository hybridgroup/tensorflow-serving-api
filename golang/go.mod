module github.com/hybridgroup/tensorflow-serving-api-go

go 1.15

replace (
	github.com/tensorflow/tensorflow => ./github.com/tensorflow/tensorflow
	tensorflow_serving => ./tensorflow_serving
)

require (
	github.com/northvolt/firefly-code-scan v0.0.0-20201105154142-8c47ae58e862
	github.com/tensorflow/tensorflow v0.0.0-00010101000000-000000000000
	gocv.io/x/gocv v0.25.0
	google.golang.org/grpc v1.33.1
	tensorflow_serving v0.0.0-00010101000000-000000000000
)
