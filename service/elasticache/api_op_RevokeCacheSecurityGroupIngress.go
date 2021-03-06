// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a RevokeCacheSecurityGroupIngress operation.
type RevokeCacheSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	// The name of the cache security group to revoke ingress from.
	//
	// CacheSecurityGroupName is a required field
	CacheSecurityGroupName *string `type:"string" required:"true"`

	// The name of the Amazon EC2 security group to revoke access from.
	//
	// EC2SecurityGroupName is a required field
	EC2SecurityGroupName *string `type:"string" required:"true"`

	// The AWS account number of the Amazon EC2 security group owner. Note that
	// this is not the same thing as an AWS access key ID - you must provide a valid
	// AWS account number for this parameter.
	//
	// EC2SecurityGroupOwnerId is a required field
	EC2SecurityGroupOwnerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RevokeCacheSecurityGroupIngressInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeCacheSecurityGroupIngressInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RevokeCacheSecurityGroupIngressInput"}

	if s.CacheSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSecurityGroupName"))
	}

	if s.EC2SecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("EC2SecurityGroupName"))
	}

	if s.EC2SecurityGroupOwnerId == nil {
		invalidParams.Add(aws.NewErrParamRequired("EC2SecurityGroupOwnerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RevokeCacheSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of one of the following operations:
	//
	//    * AuthorizeCacheSecurityGroupIngress
	//
	//    * CreateCacheSecurityGroup
	//
	//    * RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *CacheSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s RevokeCacheSecurityGroupIngressOutput) String() string {
	return awsutil.Prettify(s)
}

const opRevokeCacheSecurityGroupIngress = "RevokeCacheSecurityGroupIngress"

// RevokeCacheSecurityGroupIngressRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Revokes ingress from a cache security group. Use this operation to disallow
// access from an Amazon EC2 security group that had been previously authorized.
//
//    // Example sending a request using RevokeCacheSecurityGroupIngressRequest.
//    req := client.RevokeCacheSecurityGroupIngressRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/RevokeCacheSecurityGroupIngress
func (c *Client) RevokeCacheSecurityGroupIngressRequest(input *RevokeCacheSecurityGroupIngressInput) RevokeCacheSecurityGroupIngressRequest {
	op := &aws.Operation{
		Name:       opRevokeCacheSecurityGroupIngress,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeCacheSecurityGroupIngressInput{}
	}

	req := c.newRequest(op, input, &RevokeCacheSecurityGroupIngressOutput{})

	return RevokeCacheSecurityGroupIngressRequest{Request: req, Input: input, Copy: c.RevokeCacheSecurityGroupIngressRequest}
}

// RevokeCacheSecurityGroupIngressRequest is the request type for the
// RevokeCacheSecurityGroupIngress API operation.
type RevokeCacheSecurityGroupIngressRequest struct {
	*aws.Request
	Input *RevokeCacheSecurityGroupIngressInput
	Copy  func(*RevokeCacheSecurityGroupIngressInput) RevokeCacheSecurityGroupIngressRequest
}

// Send marshals and sends the RevokeCacheSecurityGroupIngress API request.
func (r RevokeCacheSecurityGroupIngressRequest) Send(ctx context.Context) (*RevokeCacheSecurityGroupIngressResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RevokeCacheSecurityGroupIngressResponse{
		RevokeCacheSecurityGroupIngressOutput: r.Request.Data.(*RevokeCacheSecurityGroupIngressOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RevokeCacheSecurityGroupIngressResponse is the response type for the
// RevokeCacheSecurityGroupIngress API operation.
type RevokeCacheSecurityGroupIngressResponse struct {
	*RevokeCacheSecurityGroupIngressOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RevokeCacheSecurityGroupIngress request.
func (r *RevokeCacheSecurityGroupIngressResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
