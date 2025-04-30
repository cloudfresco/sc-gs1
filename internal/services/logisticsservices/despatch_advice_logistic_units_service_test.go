package logisticsservices

import (
	"context"
	"testing"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestDespatchAdviceService_CreateDespatchAdviceLogisticUnit(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)

	despatchAdviceLogisticUnit := logisticsproto.CreateDespatchAdviceLogisticUnitRequest{}
	despatchAdviceLogisticUnit.AdditionalLogisiticUnitIdentification = "PE (pallet, modular)"
	despatchAdviceLogisticUnit.AdditionalLogisticUnitIdentificationTypeCode = ""
	despatchAdviceLogisticUnit.CodeListVersion = ""
	despatchAdviceLogisticUnit.Sscc = "409876506700001010"
	despatchAdviceLogisticUnit.UltimateConsignee = uint32(0)
	despatchAdviceLogisticUnit.DespatchAdviceId = uint32(1)
	despatchAdviceLogisticUnit.EstimatedDeliveryDateTimeAtUltimateConsignee = "04/11/2011"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateDespatchAdviceLogisticUnitRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.CreateDespatchAdviceLogisticUnitResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx: ctx,
				in:  &despatchAdviceLogisticUnit,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceLogisticUnitResponse, err := tt.das.CreateDespatchAdviceLogisticUnit(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.CreateDespatchAdviceLogisticUnit() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		despatchAdviceLogisticUnitResult := despatchAdviceLogisticUnitResponse.DespatchAdviceLogisticUnit
		assert.NotNil(t, despatchAdviceLogisticUnitResult)
	}
}
