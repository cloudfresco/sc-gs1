package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatMincingDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatMincingDetail := meatproto.CreateMeatMincingDetailRequest{}
	meatMincingDetail.FatContentPercent = float64(40)
	meatMincingDetail.MincingTypeCode = "MincingTypeCode"

	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatMincingDetailRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatMincingDetailResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatMincingDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		meatMincingDetailResp, err := tt.ms.CreateMeatMincingDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatMincingDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatMincingDetailResult := meatMincingDetailResp.MeatMincingDetail
		assert.NotNil(t, meatMincingDetailResult)
		assert.Equal(t, meatMincingDetailResult.FatContentPercent, float64(40), "they should be equal")
		assert.Equal(t, meatMincingDetailResult.MincingTypeCode, "MincingTypeCode", "they should be equal")

	}
}
