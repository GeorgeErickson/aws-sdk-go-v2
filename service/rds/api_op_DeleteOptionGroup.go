// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
)

type DeleteOptionGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the option group to be deleted.
	//
	// You can't delete default option groups.
	//
	// OptionGroupName is a required field
	OptionGroupName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteOptionGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteOptionGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteOptionGroupInput"}

	if s.OptionGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("OptionGroupName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteOptionGroupOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteOptionGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeleteOptionGroup = "DeleteOptionGroup"

// DeleteOptionGroupRequest returns a request value for making API operation for
// Amazon Relational Database Service.
//
// Deletes an existing option group.
//
//    // Example sending a request using DeleteOptionGroupRequest.
//    req := client.DeleteOptionGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rds-2014-10-31/DeleteOptionGroup
func (c *Client) DeleteOptionGroupRequest(input *DeleteOptionGroupInput) DeleteOptionGroupRequest {
	op := &aws.Operation{
		Name:       opDeleteOptionGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteOptionGroupInput{}
	}

	req := c.newRequest(op, input, &DeleteOptionGroupOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)

	return DeleteOptionGroupRequest{Request: req, Input: input, Copy: c.DeleteOptionGroupRequest}
}

// DeleteOptionGroupRequest is the request type for the
// DeleteOptionGroup API operation.
type DeleteOptionGroupRequest struct {
	*aws.Request
	Input *DeleteOptionGroupInput
	Copy  func(*DeleteOptionGroupInput) DeleteOptionGroupRequest
}

// Send marshals and sends the DeleteOptionGroup API request.
func (r DeleteOptionGroupRequest) Send(ctx context.Context) (*DeleteOptionGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteOptionGroupResponse{
		DeleteOptionGroupOutput: r.Request.Data.(*DeleteOptionGroupOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteOptionGroupResponse is the response type for the
// DeleteOptionGroup API operation.
type DeleteOptionGroupResponse struct {
	*DeleteOptionGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteOptionGroup request.
func (r *DeleteOptionGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
