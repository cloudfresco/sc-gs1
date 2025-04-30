package orderworkflows

import (
	"context"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type OrderActivities struct {
	OrderServiceClient orderproto.OrderServiceClient
}

// CreateOrderActivity - Create Order activity
func (oa *OrderActivities) CreateOrderActivity(ctx context.Context, form *orderproto.CreateOrderRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*orderproto.CreateOrderResponse, error) {
	orderServiceClient := oa.OrderServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	order, err := orderServiceClient.CreateOrder(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return order, nil
}

// UpdateOrderActivity - update Order activity
func (oa *OrderActivities) UpdateOrderActivity(ctx context.Context, form *orderproto.UpdateOrderRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	orderServiceClient := oa.OrderServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := orderServiceClient.UpdateOrder(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
