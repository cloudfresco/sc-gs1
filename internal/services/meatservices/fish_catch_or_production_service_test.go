package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestFishService_CreateFishCatchOrProduction(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	fishService := NewFishService(log, dbService, redisService, userServiceClient)
	fishCatchOrProduction := meatproto.CreateFishCatchOrProductionRequest{}
	fishCatchOrProduction.CatchArea = "21.6"
	fishCatchOrProduction.FishingGearTypeCode = "SV"
	fishCatchOrProduction.ProductionMethodForFishAndSeafoodCode = "MARINE_FISHERY"

	type args struct {
		ctx context.Context
		in  *meatproto.CreateFishCatchOrProductionRequest
	}
	tests := []struct {
		fis     *FishService
		args    args
		want    *meatproto.CreateFishCatchOrProductionResponse
		wantErr bool
	}{
		{
			fis: fishService,
			args: args{
				ctx: ctx,
				in:  &fishCatchOrProduction,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		fishCatchOrProductionResp, err := tt.fis.CreateFishCatchOrProduction(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("FishService.CreateFishCatchOrProduction() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		fishCatchOrProductionResult := fishCatchOrProductionResp.FishCatchOrProduction
		assert.NotNil(t, fishCatchOrProductionResult)
		assert.Equal(t, fishCatchOrProductionResult.CatchArea, "21.6", "they should be equal")
		assert.Equal(t, fishCatchOrProductionResult.FishingGearTypeCode, "SV", "they should be equal")
		assert.Equal(t, fishCatchOrProductionResult.ProductionMethodForFishAndSeafoodCode, "MARINE_FISHERY", "they should be equal")
	}
}
