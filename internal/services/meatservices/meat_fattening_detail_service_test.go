package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatFatteningDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatFatteningDetail := meatproto.CreateMeatFatteningDetailRequest{}
	meatFatteningDetail.FeedingSystemCode = "Intensive"
	meatFatteningDetail.HousingSystemCode = "predominantly pasture"
	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatFatteningDetailRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatFatteningDetailResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatFatteningDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		meatFatteningDetailResp, err := tt.ms.CreateMeatFatteningDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatFatteningDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatFatteningDetailResult := meatFatteningDetailResp.MeatFatteningDetail
		assert.NotNil(t, meatFatteningDetailResult)
		assert.Equal(t, meatFatteningDetailResult.FeedingSystemCode, "Intensive", "they should be equal")
		assert.Equal(t, meatFatteningDetailResult.HousingSystemCode, "predominantly pasture", "they should be equal")
	}
}
