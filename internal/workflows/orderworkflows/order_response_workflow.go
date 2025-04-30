package orderworkflows

import (
	"time"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

// CreateOrderResponseWorkflow - Create OrderResponse workflow
func CreateOrderResponseWorkflow(ctx workflow.Context, form *orderproto.CreateOrderResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*orderproto.CreateOrderResponseResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *OrderResponseActivities
	var orderResponse orderproto.CreateOrderResponseResponse
	err := workflow.ExecuteActivity(ctx, da.CreateOrderResponseActivity, form, tokenString, user, log).Get(ctx, &orderResponse)
	if err != nil {
		logger.Error("Failed to CreateOrderResponseWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &orderResponse, nil
}

// UpdateOrderResponseWorkflow - update OrderResponse workflow
func UpdateOrderResponseWorkflow(ctx workflow.Context, form *orderproto.UpdateOrderResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var da *OrderResponseActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, da.UpdateOrderResponseActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateOrderResponseWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
