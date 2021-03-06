// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input for ListStreams.
type ListStreamsInput struct {
	_ struct{} `type:"structure"`

	// The name of the stream to start the list with.
	ExclusiveStartStreamName *string `min:"1" type:"string"`

	// The maximum number of streams to list.
	Limit *int64 `min:"1" type:"integer"`
}

// String returns the string representation
func (s ListStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListStreamsInput"}
	if s.ExclusiveStartStreamName != nil && len(*s.ExclusiveStartStreamName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ExclusiveStartStreamName", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output for ListStreams.
type ListStreamsOutput struct {
	_ struct{} `type:"structure"`

	// If set to true, there are more streams available to list.
	//
	// HasMoreStreams is a required field
	HasMoreStreams *bool `type:"boolean" required:"true"`

	// The names of the streams that are associated with the AWS account making
	// the ListStreams request.
	//
	// StreamNames is a required field
	StreamNames []string `type:"list" required:"true"`
}

// String returns the string representation
func (s ListStreamsOutput) String() string {
	return awsutil.Prettify(s)
}

const opListStreams = "ListStreams"

// ListStreamsRequest returns a request value for making API operation for
// Amazon Kinesis.
//
// Lists your Kinesis data streams.
//
// The number of streams may be too large to return from a single call to ListStreams.
// You can limit the number of returned streams using the Limit parameter. If
// you do not specify a value for the Limit parameter, Kinesis Data Streams
// uses the default limit, which is currently 10.
//
// You can detect if there are more streams available to list by using the HasMoreStreams
// flag from the returned output. If there are more streams available, you can
// request more streams by using the name of the last stream returned by the
// ListStreams request in the ExclusiveStartStreamName parameter in a subsequent
// request to ListStreams. The group of stream names returned by the subsequent
// request is then added to the list. You can continue this process until all
// the stream names have been collected in the list.
//
// ListStreams has a limit of five transactions per second per account.
//
//    // Example sending a request using ListStreamsRequest.
//    req := client.ListStreamsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/kinesis-2013-12-02/ListStreams
func (c *Client) ListStreamsRequest(input *ListStreamsInput) ListStreamsRequest {
	op := &aws.Operation{
		Name:       opListStreams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"ExclusiveStartStreamName"},
			OutputTokens:    []string{"StreamNames[-1]"},
			LimitToken:      "Limit",
			TruncationToken: "HasMoreStreams",
		},
	}

	if input == nil {
		input = &ListStreamsInput{}
	}

	req := c.newRequest(op, input, &ListStreamsOutput{})

	return ListStreamsRequest{Request: req, Input: input, Copy: c.ListStreamsRequest}
}

// ListStreamsRequest is the request type for the
// ListStreams API operation.
type ListStreamsRequest struct {
	*aws.Request
	Input *ListStreamsInput
	Copy  func(*ListStreamsInput) ListStreamsRequest
}

// Send marshals and sends the ListStreams API request.
func (r ListStreamsRequest) Send(ctx context.Context) (*ListStreamsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListStreamsResponse{
		ListStreamsOutput: r.Request.Data.(*ListStreamsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListStreamsRequestPaginator returns a paginator for ListStreams.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListStreamsRequest(input)
//   p := kinesis.NewListStreamsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListStreamsPaginator(req ListStreamsRequest) ListStreamsPaginator {
	return ListStreamsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListStreamsInput
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

// ListStreamsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListStreamsPaginator struct {
	aws.Pager
}

func (p *ListStreamsPaginator) CurrentPage() *ListStreamsOutput {
	return p.Pager.CurrentPage().(*ListStreamsOutput)
}

// ListStreamsResponse is the response type for the
// ListStreams API operation.
type ListStreamsResponse struct {
	*ListStreamsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListStreams request.
func (r *ListStreamsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
