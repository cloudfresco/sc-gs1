package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatSlaughteringDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatSlaughteringDetail := meatproto.CreateMeatSlaughteringDetailRequest{}
	meatSlaughteringDetail.AgeOfAnimal = uint32(2)
	meatSlaughteringDetail.FatContentPercent = float64(40)
	meatSlaughteringDetail.FatCoverCode = ""
	meatSlaughteringDetail.MeatCategoryCode = "young intact male"
	meatSlaughteringDetail.MeatColourCode = ""
	meatSlaughteringDetail.MeatConformationCode = "E"
	meatSlaughteringDetail.MeatProfileCode = ""
	meatSlaughteringDetail.SlaughteringSystemCode = "halal"
	meatSlaughteringDetail.SlaughteringWeight = float64(280.3)
	meatSlaughteringDetail.SWCodeListVersion = ""
	meatSlaughteringDetail.SWMeasurementUnitCode = ""
	meatSlaughteringDetail.BseTestId = uint32(1)
	meatSlaughteringDetail.SlaughteringDNATestId = uint32(1)
	meatSlaughteringDetail.DateOfSlaughtering = "11/12/2023"
	meatSlaughteringDetail.OptimumMaturationDate = "12/15/2023"
	meatSlaughteringDetail.UserId = "auth0|673ee1a719dd4000cd5a3832"
	meatSlaughteringDetail.UserEmail = "sprov300@gmail.com"
	meatSlaughteringDetail.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatSlaughteringDetailRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatSlaughteringDetailResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatSlaughteringDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		meatSlaughteringDetailResp, err := tt.ms.CreateMeatSlaughteringDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatSlaughteringDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatSlaughteringDetailResult := meatSlaughteringDetailResp.MeatSlaughteringDetail
		assert.NotNil(t, meatSlaughteringDetailResult)
		assert.Equal(t, meatSlaughteringDetailResult.MeatSlaughteringDetailD.FatContentPercent, float64(40), "they should be equal")
		assert.Equal(t, meatSlaughteringDetailResult.MeatSlaughteringDetailD.MeatCategoryCode, "young intact male", "they should be equal")
		assert.Equal(t, meatSlaughteringDetailResult.MeatSlaughteringDetailD.MeatConformationCode, "E", "they should be equal")
		assert.Equal(t, meatSlaughteringDetailResult.MeatSlaughteringDetailD.SlaughteringSystemCode, "halal", "they should be equal")
		assert.Equal(t, meatSlaughteringDetailResult.MeatSlaughteringDetailD.SlaughteringWeight, float64(280.3), "they should be equal")

	}
}
