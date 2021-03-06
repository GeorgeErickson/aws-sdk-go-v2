// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateTagOptionInput struct {
	_ struct{} `type:"structure"`

	// The updated active state.
	Active *bool `type:"boolean"`

	// The TagOption identifier.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// The updated value.
	Value *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateTagOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTagOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTagOptionInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}
	if s.Value != nil && len(*s.Value) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Value", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTagOptionOutput struct {
	_ struct{} `type:"structure"`

	// Information about the TagOption.
	TagOptionDetail *TagOptionDetail `type:"structure"`
}

// String returns the string representation
func (s UpdateTagOptionOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateTagOption = "UpdateTagOption"

// UpdateTagOptionRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Updates the specified TagOption.
//
//    // Example sending a request using UpdateTagOptionRequest.
//    req := client.UpdateTagOptionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/UpdateTagOption
func (c *Client) UpdateTagOptionRequest(input *UpdateTagOptionInput) UpdateTagOptionRequest {
	op := &aws.Operation{
		Name:       opUpdateTagOption,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateTagOptionInput{}
	}

	req := c.newRequest(op, input, &UpdateTagOptionOutput{})

	return UpdateTagOptionRequest{Request: req, Input: input, Copy: c.UpdateTagOptionRequest}
}

// UpdateTagOptionRequest is the request type for the
// UpdateTagOption API operation.
type UpdateTagOptionRequest struct {
	*aws.Request
	Input *UpdateTagOptionInput
	Copy  func(*UpdateTagOptionInput) UpdateTagOptionRequest
}

// Send marshals and sends the UpdateTagOption API request.
func (r UpdateTagOptionRequest) Send(ctx context.Context) (*UpdateTagOptionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTagOptionResponse{
		UpdateTagOptionOutput: r.Request.Data.(*UpdateTagOptionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTagOptionResponse is the response type for the
// UpdateTagOption API operation.
type UpdateTagOptionResponse struct {
	*UpdateTagOptionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTagOption request.
func (r *UpdateTagOptionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
