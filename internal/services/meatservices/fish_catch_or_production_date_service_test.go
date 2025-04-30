package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestFishService_CreateFishCatchOrProductionDate(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	fishService := NewFishService(log, dbService, redisService, userServiceClient)
	fishCatchOrProductionDate := meatproto.CreateFishCatchOrProductionDateRequest{}
	fishCatchOrProductionDate.CatchEndDate = "11/13/2023"
	fishCatchOrProductionDate.CatchStartDate = "11/11/2023"
	fishCatchOrProductionDate.FirstFreezeDate = "11/12/2023"
	fishCatchOrProductionDate.CatchDateTime = "11/12/2023"

	type args struct {
		ctx context.Context
		in  *meatproto.CreateFishCatchOrProductionDateRequest
	}
	tests := []struct {
		fis     *FishService
		args    args
		want    *meatproto.CreateFishCatchOrProductionDateResponse
		wantErr bool
	}{
		{
			fis: fishService,
			args: args{
				ctx: ctx,
				in:  &fishCatchOrProductionDate,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		fishCatchOrProductionDateResp, err := tt.fis.CreateFishCatchOrProductionDate(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("FishService.CreateFishCatchOrProductionDate() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		fishCatchOrProductionDateResult := fishCatchOrProductionDateResp.FishCatchOrProductionDate
		assert.NotNil(t, fishCatchOrProductionDateResult)
	}
}
