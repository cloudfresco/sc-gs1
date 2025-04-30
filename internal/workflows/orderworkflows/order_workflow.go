package orderworkflows

import (
	"time"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "gs1"
)

// CreateOrderWorkflow - Create Order workflow
func CreateOrderWorkflow(ctx workflow.Context, form *orderproto.CreateOrderRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*orderproto.CreateOrderResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var oa *OrderActivities
	var order orderproto.CreateOrderResponse
	err := workflow.ExecuteActivity(ctx, oa.CreateOrderActivity, form, tokenString, user, log).Get(ctx, &order)
	if err != nil {
		logger.Error("Failed to CreateOrderWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &order, nil
}

// UpdateOrderWorkflow - update Order workflow
func UpdateOrderWorkflow(ctx workflow.Context, form *orderproto.UpdateOrderRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var oa *OrderActivities
	var resp string
	err := workflow.ExecuteActivity(ctx, oa.UpdateOrderActivity, form, tokenString, user, log).Get(ctx, &resp)
	if err != nil {
		logger.Error("Failed to UpdateOrderWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return resp, nil
}
