package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestFishService_CreateFishDespatchAdviceLineItemExtension(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	fishService := NewFishService(log, dbService, redisService, userServiceClient)
	fishDespatchAdviceLineItemExtension := meatproto.CreateFishDespatchAdviceLineItemExtensionRequest{}
	fishDespatchAdviceLineItemExtension.AquaticSpeciesCode = "COD"
	fishDespatchAdviceLineItemExtension.AquaticSpeciesName = "Gadus morhua"
	fishDespatchAdviceLineItemExtension.FishPresentationCode = "FIL"
	fishDespatchAdviceLineItemExtension.FpCodeListAgencyName = "European Communities"
	fishDespatchAdviceLineItemExtension.FpCodeListVersion = "1.2"
	fishDespatchAdviceLineItemExtension.FpCodeListName = "EuropeanFishPresentationCode"
	fishDespatchAdviceLineItemExtension.FishSizeCode = "3"
	fishDespatchAdviceLineItemExtension.FsCodeListAgencyName = "European Communities"
	fishDespatchAdviceLineItemExtension.FsCodeListVersion = ""
	fishDespatchAdviceLineItemExtension.FsCodeListName = "EuropeanFishSizeCode"
	fishDespatchAdviceLineItemExtension.QualityGradeCode = "A"
	fishDespatchAdviceLineItemExtension.QgCodeListAgencyName = "European Communities"
	fishDespatchAdviceLineItemExtension.QgCodeListVersion = ""
	fishDespatchAdviceLineItemExtension.QgCodeListName = "EuropeanQualityGradeCode"
	fishDespatchAdviceLineItemExtension.StorageStateCode = "PREVIOUSLY_FROZEN"
	fishDespatchAdviceLineItemExtension.AquaCultureProductionUnit = uint32(0)
	fishDespatchAdviceLineItemExtension.FishingVessel = uint32(0)
	fishDespatchAdviceLineItemExtension.PlaceOfSlaughter = uint32(0)
	fishDespatchAdviceLineItemExtension.PortOfLanding = uint32(0)
	fishDespatchAdviceLineItemExtension.FishCatchOrProductionDateId = uint32(1)
	fishDespatchAdviceLineItemExtension.FishCatchOrProductionId = uint32(1)
	fishDespatchAdviceLineItemExtension.DateOfLanding = "11/13/2023"
	fishDespatchAdviceLineItemExtension.DateOfSlaughter = "11/14/2023"
	fishDespatchAdviceLineItemExtension.UserId = "auth0|673ee1a719dd4000cd5a3832"
	fishDespatchAdviceLineItemExtension.UserEmail = "sprov300@gmail.com"
	fishDespatchAdviceLineItemExtension.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *meatproto.CreateFishDespatchAdviceLineItemExtensionRequest
	}
	tests := []struct {
		fis     *FishService
		args    args
		want    *meatproto.CreateFishDespatchAdviceLineItemExtensionResponse
		wantErr bool
	}{
		{
			fis: fishService,
			args: args{
				ctx: ctx,
				in:  &fishDespatchAdviceLineItemExtension,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		fishDespatchAdviceLineItemExtensionResp, err := tt.fis.CreateFishDespatchAdviceLineItemExtension(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("FishService.CreateFishDespatchAdviceLineItemExtension() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		fishDespatchAdviceLineItemExtensionResult := fishDespatchAdviceLineItemExtensionResp.FishDespatchAdviceLineItemExtension
		assert.NotNil(t, fishDespatchAdviceLineItemExtensionResult)
		assert.Equal(t, fishDespatchAdviceLineItemExtensionResult.FishDespatchAdviceLineItemExtensionD.AquaticSpeciesCode, "COD", "they should be equal")
	}
}
