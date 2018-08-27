package mura

import (
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/medtune/capsul/pkg/request"
)

var (
	// Model name
	Model = "mura_inception_resnet_v2"
	// Signature name
	Signature = "predict_images"
	// Version int
	Version = 1
)

// SetSpecs .
func SetSpecs(model, signature string, version int) {
	Model = model
	Signature = signature
	Version = version
}

// New return inference proto message as struct
func New(model, signature string, version int, imageBytes []byte) *pb.PredictRequest {
	return &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          model,
			SignatureName: signature,
			Version: &google_protobuf.Int64Value{
				Value: int64(1),
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype:     tf_core_framework.DataType_DT_STRING,
				StringVal: [][]byte{imageBytes},
			},
		},
	}
}

// Default return default inference proto message as struct
func Default(imageBytes []byte) *pb.PredictRequest {
	return &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          Model,
			SignatureName: Signature,
			Version: &google_protobuf.Int64Value{
				Value: int64(Version),
			},
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"images": &tf_core_framework.TensorProto{
				Dtype:     tf_core_framework.DataType_DT_STRING,
				StringVal: [][]byte{imageBytes},
			},
		},
	}
}

// Status request
func Status() *pb.GetModelStatusRequest {
	return request.Status(Model, Version)
}
