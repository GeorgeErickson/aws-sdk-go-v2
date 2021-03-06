// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemaker

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeUserProfileInput struct {
	_ struct{} `type:"structure"`

	// The domain ID.
	//
	// DomainId is a required field
	DomainId *string `type:"string" required:"true"`

	// The user profile name.
	//
	// UserProfileName is a required field
	UserProfileName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeUserProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeUserProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeUserProfileInput"}

	if s.DomainId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainId"))
	}

	if s.UserProfileName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserProfileName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeUserProfileOutput struct {
	_ struct{} `type:"structure"`

	// The creation time.
	CreationTime *time.Time `type:"timestamp"`

	// The domain ID.
	DomainId *string `type:"string"`

	// The failure reason.
	FailureReason *string `type:"string"`

	// The home Amazon Elastic File System (EFS) Uid.
	HomeEfsFileSystemUid *string `type:"string"`

	// The last modified time.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The SSO user identifier.
	SingleSignOnUserIdentifier *string `type:"string"`

	// The SSO user value.
	SingleSignOnUserValue *string `type:"string"`

	// The status.
	Status UserProfileStatus `type:"string" enum:"true"`

	// The user profile Amazon Resource Name (ARN).
	UserProfileArn *string `type:"string"`

	// The user profile name.
	UserProfileName *string `type:"string"`

	// A collection of settings.
	UserSettings *UserSettings `type:"structure"`
}

// String returns the string representation
func (s DescribeUserProfileOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeUserProfile = "DescribeUserProfile"

// DescribeUserProfileRequest returns a request value for making API operation for
// Amazon SageMaker Service.
//
// Describes the user profile.
//
//    // Example sending a request using DescribeUserProfileRequest.
//    req := client.DescribeUserProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sagemaker-2017-07-24/DescribeUserProfile
func (c *Client) DescribeUserProfileRequest(input *DescribeUserProfileInput) DescribeUserProfileRequest {
	op := &aws.Operation{
		Name:       opDescribeUserProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeUserProfileInput{}
	}

	req := c.newRequest(op, input, &DescribeUserProfileOutput{})

	return DescribeUserProfileRequest{Request: req, Input: input, Copy: c.DescribeUserProfileRequest}
}

// DescribeUserProfileRequest is the request type for the
// DescribeUserProfile API operation.
type DescribeUserProfileRequest struct {
	*aws.Request
	Input *DescribeUserProfileInput
	Copy  func(*DescribeUserProfileInput) DescribeUserProfileRequest
}

// Send marshals and sends the DescribeUserProfile API request.
func (r DescribeUserProfileRequest) Send(ctx context.Context) (*DescribeUserProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeUserProfileResponse{
		DescribeUserProfileOutput: r.Request.Data.(*DescribeUserProfileOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeUserProfileResponse is the response type for the
// DescribeUserProfile API operation.
type DescribeUserProfileResponse struct {
	*DescribeUserProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeUserProfile request.
func (r *DescribeUserProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
