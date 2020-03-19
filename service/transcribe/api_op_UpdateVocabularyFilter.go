// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package transcribe

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateVocabularyFilterInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 location of a text file used as input to create the vocabulary
	// filter. Only use characters from the character set defined for custom vocabularies.
	// For a list of character sets, see Character Sets for Custom Vocabularies
	// (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	//
	// The specified file must be less than 50 KB of UTF-8 characters.
	//
	// If you provide the location of a list of words in the VocabularyFilterFileUri
	// parameter, you can't use the Words parameter.
	VocabularyFilterFileUri *string `min:"1" type:"string"`

	// The name of the vocabulary filter to update.
	//
	// VocabularyFilterName is a required field
	VocabularyFilterName *string `min:"1" type:"string" required:"true"`

	// The words to use in the vocabulary filter. Only use characters from the character
	// set defined for custom vocabularies. For a list of character sets, see Character
	// Sets for Custom Vocabularies (https://docs.aws.amazon.com/transcribe/latest/dg/how-vocabulary.html#charsets).
	//
	// If you provide a list of words in the Words parameter, you can't use the
	// VocabularyFilterFileUri parameter.
	Words []string `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateVocabularyFilterInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateVocabularyFilterInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateVocabularyFilterInput"}
	if s.VocabularyFilterFileUri != nil && len(*s.VocabularyFilterFileUri) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFilterFileUri", 1))
	}

	if s.VocabularyFilterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VocabularyFilterName"))
	}
	if s.VocabularyFilterName != nil && len(*s.VocabularyFilterName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VocabularyFilterName", 1))
	}
	if s.Words != nil && len(s.Words) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Words", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateVocabularyFilterOutput struct {
	_ struct{} `type:"structure"`

	// The language code of the words in the vocabulary filter.
	LanguageCode LanguageCode `type:"string" enum:"true"`

	// The date and time that the vocabulary filter was updated.
	LastModifiedTime *time.Time `type:"timestamp"`

	// The name of the updated vocabulary filter.
	VocabularyFilterName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateVocabularyFilterOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateVocabularyFilter = "UpdateVocabularyFilter"

// UpdateVocabularyFilterRequest returns a request value for making API operation for
// Amazon Transcribe Service.
//
// Updates a vocabulary filter with a new list of filtered words.
//
//    // Example sending a request using UpdateVocabularyFilterRequest.
//    req := client.UpdateVocabularyFilterRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/transcribe-2017-10-26/UpdateVocabularyFilter
func (c *Client) UpdateVocabularyFilterRequest(input *UpdateVocabularyFilterInput) UpdateVocabularyFilterRequest {
	op := &aws.Operation{
		Name:       opUpdateVocabularyFilter,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateVocabularyFilterInput{}
	}

	req := c.newRequest(op, input, &UpdateVocabularyFilterOutput{})
	return UpdateVocabularyFilterRequest{Request: req, Input: input, Copy: c.UpdateVocabularyFilterRequest}
}

// UpdateVocabularyFilterRequest is the request type for the
// UpdateVocabularyFilter API operation.
type UpdateVocabularyFilterRequest struct {
	*aws.Request
	Input *UpdateVocabularyFilterInput
	Copy  func(*UpdateVocabularyFilterInput) UpdateVocabularyFilterRequest
}

// Send marshals and sends the UpdateVocabularyFilter API request.
func (r UpdateVocabularyFilterRequest) Send(ctx context.Context) (*UpdateVocabularyFilterResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateVocabularyFilterResponse{
		UpdateVocabularyFilterOutput: r.Request.Data.(*UpdateVocabularyFilterOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateVocabularyFilterResponse is the response type for the
// UpdateVocabularyFilter API operation.
type UpdateVocabularyFilterResponse struct {
	*UpdateVocabularyFilterOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateVocabularyFilter request.
func (r *UpdateVocabularyFilterResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}