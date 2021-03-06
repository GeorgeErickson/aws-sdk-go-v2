// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutExternalModelInput struct {
	_ struct{} `type:"structure"`

	// The model endpoint input configuration.
	//
	// InputConfiguration is a required field
	InputConfiguration *ModelInputConfiguration `locationName:"inputConfiguration" type:"structure" required:"true"`

	// The model endpoints name.
	//
	// ModelEndpoint is a required field
	ModelEndpoint *string `locationName:"modelEndpoint" type:"string" required:"true"`

	// The model endpoint’s status in Amazon Fraud Detector.
	//
	// ModelEndpointStatus is a required field
	ModelEndpointStatus ModelEndpointStatus `locationName:"modelEndpointStatus" type:"string" required:"true" enum:"true"`

	// The source of the model.
	//
	// ModelSource is a required field
	ModelSource ModelSource `locationName:"modelSource" type:"string" required:"true" enum:"true"`

	// The model endpoint output configuration.
	//
	// OutputConfiguration is a required field
	OutputConfiguration *ModelOutputConfiguration `locationName:"outputConfiguration" type:"structure" required:"true"`

	// The IAM role used to invoke the model endpoint.
	//
	// Role is a required field
	Role *Role `locationName:"role" type:"structure" required:"true"`
}

// String returns the string representation
func (s PutExternalModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutExternalModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutExternalModelInput"}

	if s.InputConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("InputConfiguration"))
	}

	if s.ModelEndpoint == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelEndpoint"))
	}
	if len(s.ModelEndpointStatus) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ModelEndpointStatus"))
	}
	if len(s.ModelSource) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ModelSource"))
	}

	if s.OutputConfiguration == nil {
		invalidParams.Add(aws.NewErrParamRequired("OutputConfiguration"))
	}

	if s.Role == nil {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}
	if s.InputConfiguration != nil {
		if err := s.InputConfiguration.Validate(); err != nil {
			invalidParams.AddNested("InputConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.OutputConfiguration != nil {
		if err := s.OutputConfiguration.Validate(); err != nil {
			invalidParams.AddNested("OutputConfiguration", err.(aws.ErrInvalidParams))
		}
	}
	if s.Role != nil {
		if err := s.Role.Validate(); err != nil {
			invalidParams.AddNested("Role", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutExternalModelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutExternalModelOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutExternalModel = "PutExternalModel"

// PutExternalModelRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates or updates an Amazon SageMaker model endpoint. You can also use this
// action to update the configuration of the model endpoint, including the IAM
// role and/or the mapped variables.
//
//    // Example sending a request using PutExternalModelRequest.
//    req := client.PutExternalModelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/PutExternalModel
func (c *Client) PutExternalModelRequest(input *PutExternalModelInput) PutExternalModelRequest {
	op := &aws.Operation{
		Name:       opPutExternalModel,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutExternalModelInput{}
	}

	req := c.newRequest(op, input, &PutExternalModelOutput{})

	return PutExternalModelRequest{Request: req, Input: input, Copy: c.PutExternalModelRequest}
}

// PutExternalModelRequest is the request type for the
// PutExternalModel API operation.
type PutExternalModelRequest struct {
	*aws.Request
	Input *PutExternalModelInput
	Copy  func(*PutExternalModelInput) PutExternalModelRequest
}

// Send marshals and sends the PutExternalModel API request.
func (r PutExternalModelRequest) Send(ctx context.Context) (*PutExternalModelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutExternalModelResponse{
		PutExternalModelOutput: r.Request.Data.(*PutExternalModelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutExternalModelResponse is the response type for the
// PutExternalModel API operation.
type PutExternalModelResponse struct {
	*PutExternalModelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutExternalModel request.
func (r *PutExternalModelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
