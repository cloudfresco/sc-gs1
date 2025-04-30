package logisticsworkflows

import (
	"context"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type DespatchAdviceActivities struct {
	DespatchAdviceServiceClient logisticsproto.DespatchAdviceServiceClient
}

// CreateDespatchAdviceActivity - Create DespatchAdvice activity
func (da *DespatchAdviceActivities) CreateDespatchAdviceActivity(ctx context.Context, form *logisticsproto.CreateDespatchAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*logisticsproto.CreateDespatchAdviceResponse, error) {
	despatchAdviceServiceClient := da.DespatchAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	despatchAdvice, err := despatchAdviceServiceClient.CreateDespatchAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return despatchAdvice, nil
}

// UpdateDespatchAdviceActivity - update DespatchAdvice activity
func (da *DespatchAdviceActivities) UpdateDespatchAdviceActivity(ctx context.Context, form *logisticsproto.UpdateDespatchAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	despatchAdviceServiceClient := da.DespatchAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := despatchAdviceServiceClient.UpdateDespatchAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
