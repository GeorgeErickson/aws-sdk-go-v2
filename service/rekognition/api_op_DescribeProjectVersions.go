// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeProjectVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per paginated call. The largest value
	// you can specify is 100. If you specify a value greater than 100, a ValidationException
	// error occurs. The default value is 100.
	MaxResults *int64 `min:"1" type:"integer"`

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in
	// the response. You can use this pagination token to retrieve the next set
	// of results.
	NextToken *string `type:"string"`

	// The Amazon Resource Name (ARN) of the project that contains the models you
	// want to describe.
	//
	// ProjectArn is a required field
	ProjectArn *string `min:"20" type:"string" required:"true"`

	// A list of model version names that you want to describe. You can add up to
	// 10 model version names to the list. If you don't specify a value, all model
	// descriptions are returned.
	VersionNames []string `min:"1" type:"list"`
}

// String returns the string representation
func (s DescribeProjectVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeProjectVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeProjectVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if s.ProjectArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProjectArn"))
	}
	if s.ProjectArn != nil && len(*s.ProjectArn) < 20 {
		invalidParams.Add(aws.NewErrParamMinLen("ProjectArn", 20))
	}
	if s.VersionNames != nil && len(s.VersionNames) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionNames", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeProjectVersionsOutput struct {
	_ struct{} `type:"structure"`

	// If the previous response was incomplete (because there is more results to
	// retrieve), Amazon Rekognition Custom Labels returns a pagination token in
	// the response. You can use this pagination token to retrieve the next set
	// of results.
	NextToken *string `type:"string"`

	// A list of model descriptions. The list is sorted by the creation date and
	// time of the model versions, latest to earliest.
	ProjectVersionDescriptions []ProjectVersionDescription `type:"list"`
}

// String returns the string representation
func (s DescribeProjectVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeProjectVersions = "DescribeProjectVersions"

// DescribeProjectVersionsRequest returns a request value for making API operation for
// Amazon Rekognition.
//
// Lists and describes the models in an Amazon Rekognition Custom Labels project.
// You can specify up to 10 model versions in ProjectVersionArns. If you don't
// specify a value, descriptions for all models are returned.
//
// This operation requires permissions to perform the rekognition:DescribeProjectVersions
// action.
//
//    // Example sending a request using DescribeProjectVersionsRequest.
//    req := client.DescribeProjectVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) DescribeProjectVersionsRequest(input *DescribeProjectVersionsInput) DescribeProjectVersionsRequest {
	op := &aws.Operation{
		Name:       opDescribeProjectVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeProjectVersionsInput{}
	}

	req := c.newRequest(op, input, &DescribeProjectVersionsOutput{})

	return DescribeProjectVersionsRequest{Request: req, Input: input, Copy: c.DescribeProjectVersionsRequest}
}

// DescribeProjectVersionsRequest is the request type for the
// DescribeProjectVersions API operation.
type DescribeProjectVersionsRequest struct {
	*aws.Request
	Input *DescribeProjectVersionsInput
	Copy  func(*DescribeProjectVersionsInput) DescribeProjectVersionsRequest
}

// Send marshals and sends the DescribeProjectVersions API request.
func (r DescribeProjectVersionsRequest) Send(ctx context.Context) (*DescribeProjectVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeProjectVersionsResponse{
		DescribeProjectVersionsOutput: r.Request.Data.(*DescribeProjectVersionsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeProjectVersionsRequestPaginator returns a paginator for DescribeProjectVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeProjectVersionsRequest(input)
//   p := rekognition.NewDescribeProjectVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeProjectVersionsPaginator(req DescribeProjectVersionsRequest) DescribeProjectVersionsPaginator {
	return DescribeProjectVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeProjectVersionsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeProjectVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeProjectVersionsPaginator struct {
	aws.Pager
}

func (p *DescribeProjectVersionsPaginator) CurrentPage() *DescribeProjectVersionsOutput {
	return p.Pager.CurrentPage().(*DescribeProjectVersionsOutput)
}

// DescribeProjectVersionsResponse is the response type for the
// DescribeProjectVersions API operation.
type DescribeProjectVersionsResponse struct {
	*DescribeProjectVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeProjectVersions request.
func (r *DescribeProjectVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
