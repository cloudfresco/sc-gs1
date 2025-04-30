package invoiceworkflows

import (
	"context"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type DebitCreditAdviceActivities struct {
	DebitCreditAdviceServiceClient invoiceproto.DebitCreditAdviceServiceClient
}

// CreateDebitCreditAdviceActivity - Create DebitCreditAdvice activity
func (da *DebitCreditAdviceActivities) CreateDebitCreditAdviceActivity(ctx context.Context, form *invoiceproto.CreateDebitCreditAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*invoiceproto.CreateDebitCreditAdviceResponse, error) {
	debitCreditAdviceServiceClient := da.DebitCreditAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	debitCreditAdvice, err := debitCreditAdviceServiceClient.CreateDebitCreditAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return debitCreditAdvice, nil
}

// UpdateDebitCreditAdviceActivity - update DebitCreditAdvice activity
func (da *DebitCreditAdviceActivities) UpdateDebitCreditAdviceActivity(ctx context.Context, form *invoiceproto.UpdateDebitCreditAdviceRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (string, error) {
	debitCreditAdviceServiceClient := da.DebitCreditAdviceServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	_, err := debitCreditAdviceServiceClient.UpdateDebitCreditAdvice(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return "", err
	}
	return "Updated Successfully", nil
}
