package orderservices

import (
	"context"
	"testing"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestOrderService_CreateOrderLineItemDetail(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
	orderLineItemDetail := orderproto.CreateOrderLineItemDetailRequest{}
	orderLineItemDetail.RequestedQuantity = float64(20)
	orderLineItemDetail.RqMeasurementUnitCode = "EA"
	orderLineItemDetail.RqCodeListVersion = ""
	orderLineItemDetail.OrderLineItemId = uint32(4)

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderLineItemDetailRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx: ctx,
				in:  &orderLineItemDetail,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderLineItemDetailResp, err := tt.o.CreateOrderLineItemDetail(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.CreateOrderLineItemDetail() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, orderLineItemDetailResp)
		orderLineItemDetailResult := orderLineItemDetailResp.OrderLineItemDetail
		assert.Equal(t, orderLineItemDetailResult.RequestedQuantity, float64(20), "they should be equal")
		assert.Equal(t, orderLineItemDetailResult.RqMeasurementUnitCode, "EA", "they should be equal")
	}
}
