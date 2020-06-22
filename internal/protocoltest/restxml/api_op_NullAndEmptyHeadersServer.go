// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Null and empty headers are not sent over the wire.
func (c *Client) NullAndEmptyHeadersServer(ctx context.Context, params *NullAndEmptyHeadersServerInput, optFns ...func(*Options)) (*NullAndEmptyHeadersServerOutput, error) {
	stack := middleware.NewStack("NullAndEmptyHeadersServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "NullAndEmptyHeadersServer",
			Err:           err,
		}
	}
	out := result.(*NullAndEmptyHeadersServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NullAndEmptyHeadersServerInput struct {
	A *string
	B *string
	C []*string
}

type NullAndEmptyHeadersServerOutput struct {
	A *string
	B *string
	C []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}