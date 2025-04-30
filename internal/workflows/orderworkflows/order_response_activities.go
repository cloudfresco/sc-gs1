package orderworkflows

import (
	"context"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type OrderResponseActivities struct {
	OrderResponseServiceClient orderproto.OrderResponseServiceClient
}

// CreateOrderResponseActivity - Create OrderResponse activity
func (ora *OrderResponseActivities) CreateOrderResponseActivity(ctx context.Context, form *orderproto.CreateOrderResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*orderproto.CreateOrderResponseResponse, error) {
	orderResponseServiceClient := ora.OrderResponseServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	orderResponse, err := orderResponseServiceClient.CreateOrderResponse(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return orderResponse, nil
}

// UpdateOrderResponseActivity - update OrderResponse activity
func (ora *OrderResponseActivities) UpdateOrderResponseActivity(ctx context.Context, form *orderproto.UpdateOrderResponseRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	orderResponseServiceClient := ora.OrderResponseServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := orderResponseServiceClient.UpdateOrderResponse(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
