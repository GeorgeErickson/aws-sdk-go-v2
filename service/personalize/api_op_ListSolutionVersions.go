// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package personalize

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListSolutionVersionsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of solution versions to return.
	MaxResults *int64 `locationName:"maxResults" min:"1" type:"integer"`

	// A token returned from the previous call to ListSolutionVersions for getting
	// the next set of solution versions (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`

	// The Amazon Resource Name (ARN) of the solution.
	SolutionArn *string `locationName:"solutionArn" type:"string"`
}

// String returns the string representation
func (s ListSolutionVersionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListSolutionVersionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListSolutionVersionsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListSolutionVersionsOutput struct {
	_ struct{} `type:"structure"`

	// A token for getting the next set of solution versions (if they exist).
	NextToken *string `locationName:"nextToken" type:"string"`

	// A list of solution versions describing the version properties.
	SolutionVersions []SolutionVersionSummary `locationName:"solutionVersions" type:"list"`
}

// String returns the string representation
func (s ListSolutionVersionsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListSolutionVersions = "ListSolutionVersions"

// ListSolutionVersionsRequest returns a request value for making API operation for
// Amazon Personalize.
//
// Returns a list of solution versions for the given solution. When a solution
// is not specified, all the solution versions associated with the account are
// listed. The response provides the properties for each solution version, including
// the Amazon Resource Name (ARN). For more information on solutions, see CreateSolution.
//
//    // Example sending a request using ListSolutionVersionsRequest.
//    req := client.ListSolutionVersionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/personalize-2018-05-22/ListSolutionVersions
func (c *Client) ListSolutionVersionsRequest(input *ListSolutionVersionsInput) ListSolutionVersionsRequest {
	op := &aws.Operation{
		Name:       opListSolutionVersions,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"nextToken"},
			OutputTokens:    []string{"nextToken"},
			LimitToken:      "maxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListSolutionVersionsInput{}
	}

	req := c.newRequest(op, input, &ListSolutionVersionsOutput{})

	return ListSolutionVersionsRequest{Request: req, Input: input, Copy: c.ListSolutionVersionsRequest}
}

// ListSolutionVersionsRequest is the request type for the
// ListSolutionVersions API operation.
type ListSolutionVersionsRequest struct {
	*aws.Request
	Input *ListSolutionVersionsInput
	Copy  func(*ListSolutionVersionsInput) ListSolutionVersionsRequest
}

// Send marshals and sends the ListSolutionVersions API request.
func (r ListSolutionVersionsRequest) Send(ctx context.Context) (*ListSolutionVersionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListSolutionVersionsResponse{
		ListSolutionVersionsOutput: r.Request.Data.(*ListSolutionVersionsOutput),
		response:                   &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListSolutionVersionsRequestPaginator returns a paginator for ListSolutionVersions.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListSolutionVersionsRequest(input)
//   p := personalize.NewListSolutionVersionsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListSolutionVersionsPaginator(req ListSolutionVersionsRequest) ListSolutionVersionsPaginator {
	return ListSolutionVersionsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListSolutionVersionsInput
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

// ListSolutionVersionsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListSolutionVersionsPaginator struct {
	aws.Pager
}

func (p *ListSolutionVersionsPaginator) CurrentPage() *ListSolutionVersionsOutput {
	return p.Pager.CurrentPage().(*ListSolutionVersionsOutput)
}

// ListSolutionVersionsResponse is the response type for the
// ListSolutionVersions API operation.
type ListSolutionVersionsResponse struct {
	*ListSolutionVersionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListSolutionVersions request.
func (r *ListSolutionVersionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
