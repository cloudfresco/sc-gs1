package invoiceworkflows

import (
	"time"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateDebitCreditAdviceWorkflow - Create DebitCreditAdvice workflow
func CreateDebitCreditAdviceWorkflow(ctx workflow.Context, form *invoiceproto.CreateDebitCreditAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*invoiceproto.CreateDebitCreditAdviceResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *DebitCreditAdviceActivities
	var debitCreditAdvice invoiceproto.CreateDebitCreditAdviceResponse
	err := workflow.ExecuteActivity(ctx, da.CreateDebitCreditAdviceActivity, form, tokenString, user, log).Get(ctx, &debitCreditAdvice)
	if err != nil {
		logger.Error("Failed to CreateDebitCreditAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &debitCreditAdvice, nil
}

// UpdateDebitCreditAdviceWorkflow - update DebitCreditAdvice workflow
func UpdateDebitCreditAdviceWorkflow(ctx workflow.Context, form *invoiceproto.UpdateDebitCreditAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *DebitCreditAdviceActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, da.UpdateDebitCreditAdviceActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateDebitCreditAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
