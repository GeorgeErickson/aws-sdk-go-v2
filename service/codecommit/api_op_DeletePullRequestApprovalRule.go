// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codecommit

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePullRequestApprovalRuleInput struct {
	_ struct{} `type:"structure"`

	// The name of the approval rule you want to delete.
	//
	// ApprovalRuleName is a required field
	ApprovalRuleName *string `locationName:"approvalRuleName" min:"1" type:"string" required:"true"`

	// The system-generated ID of the pull request that contains the approval rule
	// you want to delete.
	//
	// PullRequestId is a required field
	PullRequestId *string `locationName:"pullRequestId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePullRequestApprovalRuleInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePullRequestApprovalRuleInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePullRequestApprovalRuleInput"}

	if s.ApprovalRuleName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApprovalRuleName"))
	}
	if s.ApprovalRuleName != nil && len(*s.ApprovalRuleName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApprovalRuleName", 1))
	}

	if s.PullRequestId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PullRequestId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeletePullRequestApprovalRuleOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the deleted approval rule.
	//
	// If the approval rule was deleted in an earlier API call, the response is
	// 200 OK without content.
	//
	// ApprovalRuleId is a required field
	ApprovalRuleId *string `locationName:"approvalRuleId" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePullRequestApprovalRuleOutput) String() string {
	return awsutil.Prettify(s)
}

const opDeletePullRequestApprovalRule = "DeletePullRequestApprovalRule"

// DeletePullRequestApprovalRuleRequest returns a request value for making API operation for
// AWS CodeCommit.
//
// Deletes an approval rule from a specified pull request. Approval rules can
// be deleted from a pull request only if the pull request is open, and if the
// approval rule was created specifically for a pull request and not generated
// from an approval rule template associated with the repository where the pull
// request was created. You cannot delete an approval rule from a merged or
// closed pull request.
//
//    // Example sending a request using DeletePullRequestApprovalRuleRequest.
//    req := client.DeletePullRequestApprovalRuleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codecommit-2015-04-13/DeletePullRequestApprovalRule
func (c *Client) DeletePullRequestApprovalRuleRequest(input *DeletePullRequestApprovalRuleInput) DeletePullRequestApprovalRuleRequest {
	op := &aws.Operation{
		Name:       opDeletePullRequestApprovalRule,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeletePullRequestApprovalRuleInput{}
	}

	req := c.newRequest(op, input, &DeletePullRequestApprovalRuleOutput{})

	return DeletePullRequestApprovalRuleRequest{Request: req, Input: input, Copy: c.DeletePullRequestApprovalRuleRequest}
}

// DeletePullRequestApprovalRuleRequest is the request type for the
// DeletePullRequestApprovalRule API operation.
type DeletePullRequestApprovalRuleRequest struct {
	*aws.Request
	Input *DeletePullRequestApprovalRuleInput
	Copy  func(*DeletePullRequestApprovalRuleInput) DeletePullRequestApprovalRuleRequest
}

// Send marshals and sends the DeletePullRequestApprovalRule API request.
func (r DeletePullRequestApprovalRuleRequest) Send(ctx context.Context) (*DeletePullRequestApprovalRuleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeletePullRequestApprovalRuleResponse{
		DeletePullRequestApprovalRuleOutput: r.Request.Data.(*DeletePullRequestApprovalRuleOutput),
		response:                            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeletePullRequestApprovalRuleResponse is the response type for the
// DeletePullRequestApprovalRule API operation.
type DeletePullRequestApprovalRuleResponse struct {
	*DeletePullRequestApprovalRuleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeletePullRequestApprovalRule request.
func (r *DeletePullRequestApprovalRuleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
