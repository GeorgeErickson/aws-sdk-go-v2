// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iot

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeProvisioningTemplateVersionInput struct {
	_ struct{} `type:"structure"`

	// The template name.
	//
	// TemplateName is a required field
	TemplateName *string `location:"uri" locationName:"templateName" min:"1" type:"string" required:"true"`

	// The fleet provisioning template version ID.
	//
	// VersionId is a required field
	VersionId *int64 `location:"uri" locationName:"versionId" type:"integer" required:"true"`
}

// String returns the string representation
func (s DescribeProvisioningTemplateVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProvisioningTemplateVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProvisioningTemplateVersionInput"}

	if s.TemplateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateName"))
	}
	if s.TemplateName != nil && len(*s.TemplateName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateName", 1))
	}

	if s.VersionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VersionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProvisioningTemplateVersionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.TemplateName != nil {
		v := *s.TemplateName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "templateName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "versionId", protocol.Int64Value(v), metadata)
	}
	return nil
}

type DescribeProvisioningTemplateVersionOutput struct {
	_ struct{} `type:"structure"`

	// The date when the fleet provisioning template version was created.
	CreationDate *time.Time `locationName:"creationDate" type:"timestamp"`

	// True if the fleet provisioning template version is the default version.
	IsDefaultVersion *bool `locationName:"isDefaultVersion" type:"boolean"`

	// The JSON formatted contents of the fleet provisioning template version.
	TemplateBody *string `locationName:"templateBody" type:"string"`

	// The fleet provisioning template version ID.
	VersionId *int64 `locationName:"versionId" type:"integer"`
}

// String returns the string representation
func (s DescribeProvisioningTemplateVersionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeProvisioningTemplateVersionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.CreationDate != nil {
		v := *s.CreationDate

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "creationDate",
			protocol.TimeValue{V: v, Format: protocol.UnixTimeFormatName, QuotedFormatTime: true}, metadata)
	}
	if s.IsDefaultVersion != nil {
		v := *s.IsDefaultVersion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "isDefaultVersion", protocol.BoolValue(v), metadata)
	}
	if s.TemplateBody != nil {
		v := *s.TemplateBody

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "templateBody", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VersionId != nil {
		v := *s.VersionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "versionId", protocol.Int64Value(v), metadata)
	}
	return nil
}

const opDescribeProvisioningTemplateVersion = "DescribeProvisioningTemplateVersion"

// DescribeProvisioningTemplateVersionRequest returns a request value for making API operation for
// AWS IoT.
//
// Returns information about a fleet provisioning template version.
//
//    // Example sending a request using DescribeProvisioningTemplateVersionRequest.
//    req := client.DescribeProvisioningTemplateVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeProvisioningTemplateVersionRequest(input *DescribeProvisioningTemplateVersionInput) DescribeProvisioningTemplateVersionRequest {
	op := &aws.Operation{
		Name:       opDescribeProvisioningTemplateVersion,
		HTTPMethod: "GET",
		HTTPPath:   "/provisioning-templates/{templateName}/versions/{versionId}",
	}

	if input == nil {
		input = &DescribeProvisioningTemplateVersionInput{}
	}

	req := c.newRequest(op, input, &DescribeProvisioningTemplateVersionOutput{})

	return DescribeProvisioningTemplateVersionRequest{Request: req, Input: input, Copy: c.DescribeProvisioningTemplateVersionRequest}
}

// DescribeProvisioningTemplateVersionRequest is the request type for the
// DescribeProvisioningTemplateVersion API operation.
type DescribeProvisioningTemplateVersionRequest struct {
	*aws.Request
	Input *DescribeProvisioningTemplateVersionInput
	Copy  func(*DescribeProvisioningTemplateVersionInput) DescribeProvisioningTemplateVersionRequest
}

// Send marshals and sends the DescribeProvisioningTemplateVersion API request.
func (r DescribeProvisioningTemplateVersionRequest) Send(ctx context.Context) (*DescribeProvisioningTemplateVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProvisioningTemplateVersionResponse{
		DescribeProvisioningTemplateVersionOutput: r.Request.Data.(*DescribeProvisioningTemplateVersionOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeProvisioningTemplateVersionResponse is the response type for the
// DescribeProvisioningTemplateVersion API operation.
type DescribeProvisioningTemplateVersionResponse struct {
	*DescribeProvisioningTemplateVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProvisioningTemplateVersion request.
func (r *DescribeProvisioningTemplateVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
