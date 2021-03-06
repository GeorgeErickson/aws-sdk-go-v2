// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpointemail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain
// that the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the campaign. Amazon Pinpoint automatically generates
	// and assigns this identifier to a campaign. This value is not the same as
	// the campaign identifier that Amazon Pinpoint assigns to campaigns that you
	// create and manage by using the Amazon Pinpoint API or the Amazon Pinpoint
	// console.
	//
	// CampaignId is a required field
	CampaignId *string `location:"uri" locationName:"CampaignId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetDomainDeliverabilityCampaignInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetDomainDeliverabilityCampaignInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetDomainDeliverabilityCampaignInput"}

	if s.CampaignId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CampaignId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainDeliverabilityCampaignInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.CampaignId != nil {
		v := *s.CampaignId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "CampaignId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// An object that contains all the deliverability data for a specific campaign.
// This data is available for a campaign only if the campaign sent email by
// using a domain that the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
type GetDomainDeliverabilityCampaignOutput struct {
	_ struct{} `type:"structure"`

	// An object that contains the deliverability data for the campaign.
	//
	// DomainDeliverabilityCampaign is a required field
	DomainDeliverabilityCampaign *DomainDeliverabilityCampaign `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetDomainDeliverabilityCampaignOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetDomainDeliverabilityCampaignOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.DomainDeliverabilityCampaign != nil {
		v := s.DomainDeliverabilityCampaign

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DomainDeliverabilityCampaign", v, metadata)
	}
	return nil
}

const opGetDomainDeliverabilityCampaign = "GetDomainDeliverabilityCampaign"

// GetDomainDeliverabilityCampaignRequest returns a request value for making API operation for
// Amazon Pinpoint Email Service.
//
// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain
// that the Deliverability dashboard is enabled for (PutDeliverabilityDashboardOption
// operation).
//
//    // Example sending a request using GetDomainDeliverabilityCampaignRequest.
//    req := client.GetDomainDeliverabilityCampaignRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-email-2018-07-26/GetDomainDeliverabilityCampaign
func (c *Client) GetDomainDeliverabilityCampaignRequest(input *GetDomainDeliverabilityCampaignInput) GetDomainDeliverabilityCampaignRequest {
	op := &aws.Operation{
		Name:       opGetDomainDeliverabilityCampaign,
		HTTPMethod: "GET",
		HTTPPath:   "/v1/email/deliverability-dashboard/campaigns/{CampaignId}",
	}

	if input == nil {
		input = &GetDomainDeliverabilityCampaignInput{}
	}

	req := c.newRequest(op, input, &GetDomainDeliverabilityCampaignOutput{})

	return GetDomainDeliverabilityCampaignRequest{Request: req, Input: input, Copy: c.GetDomainDeliverabilityCampaignRequest}
}

// GetDomainDeliverabilityCampaignRequest is the request type for the
// GetDomainDeliverabilityCampaign API operation.
type GetDomainDeliverabilityCampaignRequest struct {
	*aws.Request
	Input *GetDomainDeliverabilityCampaignInput
	Copy  func(*GetDomainDeliverabilityCampaignInput) GetDomainDeliverabilityCampaignRequest
}

// Send marshals and sends the GetDomainDeliverabilityCampaign API request.
func (r GetDomainDeliverabilityCampaignRequest) Send(ctx context.Context) (*GetDomainDeliverabilityCampaignResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDomainDeliverabilityCampaignResponse{
		GetDomainDeliverabilityCampaignOutput: r.Request.Data.(*GetDomainDeliverabilityCampaignOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDomainDeliverabilityCampaignResponse is the response type for the
// GetDomainDeliverabilityCampaign API operation.
type GetDomainDeliverabilityCampaignResponse struct {
	*GetDomainDeliverabilityCampaignOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDomainDeliverabilityCampaign request.
func (r *GetDomainDeliverabilityCampaignResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
