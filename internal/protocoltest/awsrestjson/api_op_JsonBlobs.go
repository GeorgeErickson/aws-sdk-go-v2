// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Blobs are base64 encoded
func (c *Client) JsonBlobs(ctx context.Context, params *JsonBlobsInput, optFns ...func(*Options)) (*JsonBlobsOutput, error) {
	stack := middleware.NewStack("JsonBlobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	stack.Serialize.Add(&awsRestjson1_serializeOpJsonBlobs{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpJsonBlobs{}, middleware.After)
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
			OperationName: "JsonBlobs",
			Err:           err,
		}
	}
	out := result.(*JsonBlobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type JsonBlobsInput struct {
	Data []byte
}

type JsonBlobsOutput struct {
	Data []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}