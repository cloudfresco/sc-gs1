package meatservices

import (
	"context"
	"testing"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestMeatService_CreateMeatActivityHistory(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	meatService := NewMeatService(log, dbService, redisService, userServiceClient)
	meatActivityHistory := meatproto.CreateMeatActivityHistoryRequest{}
	meatActivityHistory.ActivitySubStepIdentification = uint32(0)
	meatActivityHistory.CountryOfActivityCode = "Denmark"
	meatActivityHistory.CurrentStepIdentification = uint32(6)
	meatActivityHistory.MeatProcessingActivityTypeCode = "Cutting"
	meatActivityHistory.MovementReasonCode = "Cutter"
	meatActivityHistory.NextStepIdentification = uint32(0)
	meatActivityHistory.MeatMincingDetailId = uint32(0)
	meatActivityHistory.MeatFatteningDetailId = uint32(2)
	meatActivityHistory.MeatCuttingDetailId = uint32(0)
	meatActivityHistory.MeatBreedingDetailId = uint32(0)
	meatActivityHistory.MeatProcessingPartyId = uint32(1)
	meatActivityHistory.MeatWorkItemIdentificationId = uint32(0)
	meatActivityHistory.MeatSlaughteringDetailId = uint32(0)
	meatActivityHistory.MeatDespatchAdviceLineItemExtensionId = uint32(0)
	meatActivityHistory.DateOfArrival = "11/13/2023"
	meatActivityHistory.DateOfDeparture = "11/09/2023"
	meatActivityHistory.UserId = "auth0|673ee1a719dd4000cd5a3832"
	meatActivityHistory.UserEmail = "sprov300@gmail.com"
	meatActivityHistory.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *meatproto.CreateMeatActivityHistoryRequest
	}
	tests := []struct {
		ms      *MeatService
		args    args
		want    *meatproto.CreateMeatActivityHistoryResponse
		wantErr bool
	}{
		{
			ms: meatService,
			args: args{
				ctx: ctx,
				in:  &meatActivityHistory,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		meatActivityHistoryResp, err := tt.ms.CreateMeatActivityHistory(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("MeatService.CreateMeatActivityHistory() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		meatActivityHistoryResult := meatActivityHistoryResp.MeatActivityHistory
		assert.NotNil(t, meatActivityHistoryResult)
		assert.Equal(t, meatActivityHistoryResult.MeatActivityHistoryD.CountryOfActivityCode, "Denmark", "they should be equal")
		assert.Equal(t, meatActivityHistoryResult.MeatActivityHistoryD.CurrentStepIdentification, uint32(6), "they should be equal")
		assert.Equal(t, meatActivityHistoryResult.MeatActivityHistoryD.MeatProcessingActivityTypeCode, "Cutting", "they should be equal")
		assert.Equal(t, meatActivityHistoryResult.MeatActivityHistoryD.MovementReasonCode, "Cutter", "they should be equal")
	}
}
