package logisticsworkflows

import (
	"context"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type ReceivingAdviceActivities struct {
	ReceivingAdviceServiceClient logisticsproto.ReceivingAdviceServiceClient
}

// CreateReceivingAdviceActivity - Create Receipt Advice Header activity
func (ra *ReceivingAdviceActivities) CreateReceivingAdviceActivity(ctx context.Context, form *logisticsproto.CreateReceivingAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*logisticsproto.CreateReceivingAdviceResponse, error) {
	receivingAdviceServiceClient := ra.ReceivingAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	receivingAdvice, err := receivingAdviceServiceClient.CreateReceivingAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return receivingAdvice, nil
}

// UpdateReceivingAdviceActivity - update ReceivingAdvice activity
func (ra *ReceivingAdviceActivities) UpdateReceivingAdviceActivity(ctx context.Context, form *logisticsproto.UpdateReceivingAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	receivingAdviceServiceClient := ra.ReceivingAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := receivingAdviceServiceClient.UpdateReceivingAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
