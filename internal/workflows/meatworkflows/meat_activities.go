package meatworkflows

import (
	"context"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/metadata"
)

type MeatActivities struct {
	MeatServiceClient meatproto.MeatServiceClient
}

// CreateMeatActivityHistoryActivity - Create MeatActivityHistory activity
func (ma *MeatActivities) CreateMeatActivityHistoryActivity(ctx context.Context, form *meatproto.CreateMeatActivityHistoryRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatActivityHistoryResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatActivityHistory, err := meatServiceClient.CreateMeatActivityHistory(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatActivityHistory, nil
}

// CreateMeatAcidityActivity - Create MeatAcidity activity
func (ma *MeatActivities) CreateMeatAcidityActivity(ctx context.Context, form *meatproto.CreateMeatAcidityRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatAcidityResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatAcidity, err := meatServiceClient.CreateMeatAcidity(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatAcidity, nil
}

// CreateMeatTestActivity - Create MeatTest activity
func (ma *MeatActivities) CreateMeatTestActivity(ctx context.Context, form *meatproto.CreateMeatTestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatTestResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatTest, err := meatServiceClient.CreateMeatTest(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatTest, nil
}

// CreateMeatBreedingDetailActivity - Create MeatBreedingDetail activity
func (ma *MeatActivities) CreateMeatBreedingDetailActivity(ctx context.Context, form *meatproto.CreateMeatBreedingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatBreedingDetailResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatBreedingDetail, err := meatServiceClient.CreateMeatBreedingDetail(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatBreedingDetail, nil
}

// CreateMeatCuttingDetailActivity - Create MeatCuttingDetail activity
func (ma *MeatActivities) CreateMeatCuttingDetailActivity(ctx context.Context, form *meatproto.CreateMeatCuttingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatCuttingDetailResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatCuttingDetail, err := meatServiceClient.CreateMeatCuttingDetail(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatCuttingDetail, nil
}

// CreateMeatFatteningDetailActivity - Create MeatFatteningDetail activity
func (ma *MeatActivities) CreateMeatFatteningDetailActivity(ctx context.Context, form *meatproto.CreateMeatFatteningDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatFatteningDetailResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatFatteningDetail, err := meatServiceClient.CreateMeatFatteningDetail(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatFatteningDetail, nil
}

// CreateMeatMincingDetailActivity - Create MeatMincingDetail activity
func (ma *MeatActivities) CreateMeatMincingDetailActivity(ctx context.Context, form *meatproto.CreateMeatMincingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatMincingDetailResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatMincingDetail, err := meatServiceClient.CreateMeatMincingDetail(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatMincingDetail, nil
}

// CreateMeatProcessingPartyActivity - Create MeatProcessingParty activity
func (ma *MeatActivities) CreateMeatProcessingPartyActivity(ctx context.Context, form *meatproto.CreateMeatProcessingPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatProcessingPartyResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatProcessingParty, err := meatServiceClient.CreateMeatProcessingParty(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatProcessingParty, nil
}

// CreateMeatSlaughteringDetailActivity - Create MeatSlaughteringDetail activity
func (ma *MeatActivities) CreateMeatSlaughteringDetailActivity(ctx context.Context, form *meatproto.CreateMeatSlaughteringDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatSlaughteringDetailResponse, error) {
	meatServiceClient := ma.MeatServiceClient
	md := metadata.Pairs("authorization", "Bearer "+tokenString)
	ctxNew := metadata.NewOutgoingContext(ctx, md)
	meatSlaughteringDetail, err := meatServiceClient.CreateMeatSlaughteringDetail(ctxNew, form)
	if err != nil {
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return meatSlaughteringDetail, nil
}
