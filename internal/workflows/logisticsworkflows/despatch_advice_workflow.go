package logisticsworkflows

import (
	"time"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateDespatchAdviceWorkflow - Create DespatchAdvice workflow
func CreateDespatchAdviceWorkflow(ctx workflow.Context, form *logisticsproto.CreateDespatchAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*logisticsproto.CreateDespatchAdviceResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *DespatchAdviceActivities
	var despatchAdvice logisticsproto.CreateDespatchAdviceResponse
	err := workflow.ExecuteActivity(ctx, da.CreateDespatchAdviceActivity, form, tokenString, user, log).Get(ctx, &despatchAdvice)
	if err != nil {
		logger.Error("Failed to CreateDespatchAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &despatchAdvice, nil
}

// UpdateDespatchAdviceWorkflow - update DespatchAdvice workflow
func UpdateDespatchAdviceWorkflow(ctx workflow.Context, form *logisticsproto.UpdateDespatchAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *DespatchAdviceActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, da.UpdateDespatchAdviceActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateDespatchAdviceWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
