// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package health

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeHealthServiceStatusForOrganizationInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeHealthServiceStatusForOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeHealthServiceStatusForOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// Information about the status of enabling or disabling AWS Health Organizational
	// View in your organization.
	//
	// Valid values are ENABLED | DISABLED | PENDING.
	HealthServiceAccessStatusForOrganization *string `locationName:"healthServiceAccessStatusForOrganization" type:"string"`
}

// String returns the string representation
func (s DescribeHealthServiceStatusForOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeHealthServiceStatusForOrganization = "DescribeHealthServiceStatusForOrganization"

// DescribeHealthServiceStatusForOrganizationRequest returns a request value for making API operation for
// AWS Health APIs and Notifications.
//
// This operation provides status information on enabling or disabling AWS Health
// to work with your organization. To call this operation, you must sign in
// as an IAM user, assume an IAM role, or sign in as the root user (not recommended)
// in the organization's master account.
//
//    // Example sending a request using DescribeHealthServiceStatusForOrganizationRequest.
//    req := client.DescribeHealthServiceStatusForOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/health-2016-08-04/DescribeHealthServiceStatusForOrganization
func (c *Client) DescribeHealthServiceStatusForOrganizationRequest(input *DescribeHealthServiceStatusForOrganizationInput) DescribeHealthServiceStatusForOrganizationRequest {
	op := &aws.Operation{
		Name:       opDescribeHealthServiceStatusForOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHealthServiceStatusForOrganizationInput{}
	}

	req := c.newRequest(op, input, &DescribeHealthServiceStatusForOrganizationOutput{})

	return DescribeHealthServiceStatusForOrganizationRequest{Request: req, Input: input, Copy: c.DescribeHealthServiceStatusForOrganizationRequest}
}

// DescribeHealthServiceStatusForOrganizationRequest is the request type for the
// DescribeHealthServiceStatusForOrganization API operation.
type DescribeHealthServiceStatusForOrganizationRequest struct {
	*aws.Request
	Input *DescribeHealthServiceStatusForOrganizationInput
	Copy  func(*DescribeHealthServiceStatusForOrganizationInput) DescribeHealthServiceStatusForOrganizationRequest
}

// Send marshals and sends the DescribeHealthServiceStatusForOrganization API request.
func (r DescribeHealthServiceStatusForOrganizationRequest) Send(ctx context.Context) (*DescribeHealthServiceStatusForOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeHealthServiceStatusForOrganizationResponse{
		DescribeHealthServiceStatusForOrganizationOutput: r.Request.Data.(*DescribeHealthServiceStatusForOrganizationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeHealthServiceStatusForOrganizationResponse is the response type for the
// DescribeHealthServiceStatusForOrganization API operation.
type DescribeHealthServiceStatusForOrganizationResponse struct {
	*DescribeHealthServiceStatusForOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeHealthServiceStatusForOrganization request.
func (r *DescribeHealthServiceStatusForOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
