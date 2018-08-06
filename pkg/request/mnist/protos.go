package mnist

import (
	tf_core_framework "tensorflow/core/framework"
	pb "tensorflow_serving/apis"

	google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
)

var (
	// Model name
	Model = "mnist"
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
func New(model, signature string, version int, image []float32) *pb.PredictRequest {
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
				Dtype: tf_core_framework.DataType_DT_FLOAT,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(784),
						},
					},
				},
				FloatVal: image,
			},
		},
	}
}

// Default return default inference proto message as struct
func Default(imageF32 []float32) *pb.PredictRequest {
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
				Dtype: tf_core_framework.DataType_DT_FLOAT,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(784),
						},
					},
				},
				FloatVal: imageF32,
			},
		},
	}
}