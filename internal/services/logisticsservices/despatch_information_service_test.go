package logisticsservices

import (
	"context"
	"testing"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestDespatchAdviceService_CreateDespatchInformation(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchInformation := logisticsproto.CreateDespatchInformationRequest{}
	despatchInformation.DespatchAdviceId = uint32(1)
	despatchInformation.ActualShipDateTime = "04/11/2011"
	despatchInformation.DespatchDateTime = "04/11/2011"
	despatchInformation.EstimatedDeliveryDateTime = "04/11/2011"
	despatchInformation.EstimatedDeliveryDateTimeAtUltimateConsignee = "04/11/2011"
	despatchInformation.LoadingDateTime = "04/11/2011"
	despatchInformation.PickUpDateTime = "04/11/2011"
	despatchInformation.ReleaseDateTimeOfSupplier = "04/11/2011"
	despatchInformation.EstimatedDeliveryPeriodBegin = "04/11/2011"
	despatchInformation.EstimatedDeliveryPeriodEnd = "04/11/2011"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateDespatchInformationRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.CreateDespatchInformationResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx: ctx,
				in:  &despatchInformation,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchInformationResponse, err := tt.das.CreateDespatchInformation(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.CreateDespatchInformation() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		despatchInformationResult := despatchInformationResponse.DespatchInformation
		assert.NotNil(t, despatchInformationResult)
	}
}
