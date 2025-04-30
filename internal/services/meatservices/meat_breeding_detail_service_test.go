package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatBreedingDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatBreedingDetail := meatproto.CreateMeatBreedingDetailRequest{}
	meatBreedingDetail.AnimalTypeCode = ""
	meatBreedingDetail.BreedCode = "Jersey"
	meatBreedingDetail.BreedOfFatherCode = "Jersey"
	meatBreedingDetail.BreedOfMotherCode = "Jersey"
	meatBreedingDetail.CrossBreedIndicator = false
	meatBreedingDetail.FeedingSystemCode = "organic"
	meatBreedingDetail.GenderCode = "male"
	meatBreedingDetail.HousingSystemCode = "predominantly barn"
	meatBreedingDetail.BreedingDNATest = uint32(0)
	meatBreedingDetail.DateOfBirth = "11/13/2023"
	meatBreedingDetail.UserId = "auth0|673ee1a719dd4000cd5a3832"
	meatBreedingDetail.UserEmail = "sprov300@gmail.com"
	meatBreedingDetail.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatBreedingDetailRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatBreedingDetailResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatBreedingDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		meatBreedingDetailResp, err := tt.ms.CreateMeatBreedingDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatBreedingDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatBreedingDetailResult := meatBreedingDetailResp.MeatBreedingDetail
		assert.NotNil(t, meatBreedingDetailResult)
		assert.Equal(t, meatBreedingDetailResult.MeatBreedingDetailD.BreedCode, "Jersey", "they should be equal")
		assert.Equal(t, meatBreedingDetailResult.MeatBreedingDetailD.FeedingSystemCode, "organic", "they should be equal")
		assert.Equal(t, meatBreedingDetailResult.MeatBreedingDetailD.GenderCode, "male", "they should be equal")
		assert.Equal(t, meatBreedingDetailResult.MeatBreedingDetailD.HousingSystemCode, "predominantly barn", "they should be equal")
	}
}
