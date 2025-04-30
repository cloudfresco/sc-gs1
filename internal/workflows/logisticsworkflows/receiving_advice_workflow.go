package logisticsworkflows

import (
	"time"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "gs1"
)

// CreateReceivingAdviceWorkflow - Create ReceivingAdvice workflow
func CreateReceivingAdviceWorkflow(ctx workflow.Context, form *logisticsproto.CreateReceivingAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*logisticsproto.CreateReceivingAdviceResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ra *ReceivingAdviceActivities
	var receivingAdvice logisticsproto.CreateReceivingAdviceResponse
	err := workflow.ExecuteActivity(ctx, ra.CreateReceivingAdviceActivity, form, tokenString, user, log).Get(ctx, &receivingAdvice)
	if err != nil {
		logger.Error("Failed to CreateReceivingAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &receivingAdvice, nil
}

// UpdateReceivingAdviceWorkflow - update ReceivingAdvice workflow
func UpdateReceivingAdviceWorkflow(ctx workflow.Context, form *logisticsproto.UpdateReceivingAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ra *ReceivingAdviceActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, ra.UpdateReceivingAdviceActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateReceivingAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
