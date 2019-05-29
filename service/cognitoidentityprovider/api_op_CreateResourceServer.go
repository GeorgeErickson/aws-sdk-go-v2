// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateResourceServerRequest
type CreateResourceServerInput struct {
	_ struct{} `type:"structure"`

	// A unique resource server identifier for the resource server. This could be
	// an HTTPS endpoint where the resource server is located. For example, https://my-weather-api.example.com.
	//
	// Identifier is a required field
	Identifier *string `min:"1" type:"string" required:"true"`

	// A friendly name for the resource server.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// A list of scopes. Each scope is map, where the keys are name and description.
	Scopes []ResourceServerScopeType `type:"list"`

	// The user pool ID for the user pool.
	//
	// UserPoolId is a required field
	UserPoolId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateResourceServerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateResourceServerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateResourceServerInput"}

	if s.Identifier == nil {
		invalidParams.Add(aws.NewErrParamRequired("Identifier"))
	}
	if s.Identifier != nil && len(*s.Identifier) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Identifier", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if s.UserPoolId == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserPoolId"))
	}
	if s.UserPoolId != nil && len(*s.UserPoolId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserPoolId", 1))
	}
	if s.Scopes != nil {
		for i, v := range s.Scopes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Scopes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateResourceServerResponse
type CreateResourceServerOutput struct {
	_ struct{} `type:"structure"`

	// The newly created resource server.
	//
	// ResourceServer is a required field
	ResourceServer *ResourceServerType `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateResourceServerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateResourceServer = "CreateResourceServer"

// CreateResourceServerRequest returns a request value for making API operation for
// Amazon Cognito Identity Provider.
//
// Creates a new OAuth2.0 resource server and defines custom scopes in it.
//
//    // Example sending a request using CreateResourceServerRequest.
//    req := client.CreateResourceServerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cognito-idp-2016-04-18/CreateResourceServer
func (c *Client) CreateResourceServerRequest(input *CreateResourceServerInput) CreateResourceServerRequest {
	op := &aws.Operation{
		Name:       opCreateResourceServer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateResourceServerInput{}
	}

	req := c.newRequest(op, input, &CreateResourceServerOutput{})
	return CreateResourceServerRequest{Request: req, Input: input, Copy: c.CreateResourceServerRequest}
}

// CreateResourceServerRequest is the request type for the
// CreateResourceServer API operation.
type CreateResourceServerRequest struct {
	*aws.Request
	Input *CreateResourceServerInput
	Copy  func(*CreateResourceServerInput) CreateResourceServerRequest
}

// Send marshals and sends the CreateResourceServer API request.
func (r CreateResourceServerRequest) Send(ctx context.Context) (*CreateResourceServerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateResourceServerResponse{
		CreateResourceServerOutput: r.Request.Data.(*CreateResourceServerOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateResourceServerResponse is the response type for the
// CreateResourceServer API operation.
type CreateResourceServerResponse struct {
	*CreateResourceServerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateResourceServer request.
func (r *CreateResourceServerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}