package orderservices

import (
	"context"
	"testing"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestOrderResponseService_CreateOrderResponseLineItemDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)

	orderResponseLineItemDetail := orderproto.CreateOrderResponseLineItemDetailRequest{}
	orderResponseLineItemDetail.ConfirmedQuantity = float64(20)
	orderResponseLineItemDetail.CqMeasurementUnitCode = "EA"
	orderResponseLineItemDetail.CqCodeListVersion = ""
	orderResponseLineItemDetail.ReturnReasonCode = "Product Out Of Stock"
	orderResponseLineItemDetail.OrderResponseLineItemId = uint32(3)

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderResponseLineItemDetailRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.CreateOrderResponseLineItemDetailResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx: ctx,
				in:  &orderResponseLineItemDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		orderResponseLineItemDetailResp, err := tt.o.CreateOrderResponseLineItemDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.CreateOrderResponseLineItemDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, orderResponseLineItemDetailResp)
		orderResponseLineItemDetailResult := orderResponseLineItemDetailResp.OrderResponseLineItemDetail
		assert.Equal(t, orderResponseLineItemDetailResult.ConfirmedQuantity, float64(20), "they should be equal")
		assert.Equal(t, orderResponseLineItemDetailResult.CqMeasurementUnitCode, "EA", "they should be equal")
	}
}
