package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatCuttingDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatCuttingDetail := meatproto.CreateMeatCuttingDetailRequest{}
	meatCuttingDetail.MeatProfileCode = "1164353010400017999"

	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatCuttingDetailRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatCuttingDetailResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatCuttingDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		meatCuttingDetailResp, err := tt.ms.CreateMeatCuttingDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatCuttingDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatCuttingDetailResult := meatCuttingDetailResp.MeatCuttingDetail
		assert.NotNil(t, meatCuttingDetailResult)
		assert.Equal(t, meatCuttingDetailResult.MeatProfileCode, "1164353010400017999", "they should be equal")

	}
}
