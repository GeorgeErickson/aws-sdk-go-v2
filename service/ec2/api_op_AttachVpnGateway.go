// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for AttachVpnGateway.
type AttachVpnGatewayInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the VPC.
	//
	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	// The ID of the virtual private gateway.
	//
	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachVpnGatewayInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachVpnGatewayInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachVpnGatewayInput"}

	if s.VpcId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpcId"))
	}

	if s.VpnGatewayId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of AttachVpnGateway.
type AttachVpnGatewayOutput struct {
	_ struct{} `type:"structure"`

	// Information about the attachment.
	VpcAttachment *VpcAttachment `locationName:"attachment" type:"structure"`
}

// String returns the string representation
func (s AttachVpnGatewayOutput) String() string {
	return awsutil.Prettify(s)
}

const opAttachVpnGateway = "AttachVpnGateway"

// AttachVpnGatewayRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Attaches a virtual private gateway to a VPC. You can attach one virtual private
// gateway to one VPC at a time.
//
// For more information, see AWS Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html)
// in the AWS Site-to-Site VPN User Guide.
//
//    // Example sending a request using AttachVpnGatewayRequest.
//    req := client.AttachVpnGatewayRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AttachVpnGateway
func (c *Client) AttachVpnGatewayRequest(input *AttachVpnGatewayInput) AttachVpnGatewayRequest {
	op := &aws.Operation{
		Name:       opAttachVpnGateway,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachVpnGatewayInput{}
	}

	req := c.newRequest(op, input, &AttachVpnGatewayOutput{})

	return AttachVpnGatewayRequest{Request: req, Input: input, Copy: c.AttachVpnGatewayRequest}
}

// AttachVpnGatewayRequest is the request type for the
// AttachVpnGateway API operation.
type AttachVpnGatewayRequest struct {
	*aws.Request
	Input *AttachVpnGatewayInput
	Copy  func(*AttachVpnGatewayInput) AttachVpnGatewayRequest
}

// Send marshals and sends the AttachVpnGateway API request.
func (r AttachVpnGatewayRequest) Send(ctx context.Context) (*AttachVpnGatewayResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AttachVpnGatewayResponse{
		AttachVpnGatewayOutput: r.Request.Data.(*AttachVpnGatewayOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AttachVpnGatewayResponse is the response type for the
// AttachVpnGateway API operation.
type AttachVpnGatewayResponse struct {
	*AttachVpnGatewayOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AttachVpnGateway request.
func (r *AttachVpnGatewayResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
