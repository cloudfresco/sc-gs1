package orderservices

import (
	"context"
	"testing"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestOrderService_CreateOrderLogisticalInformation(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderService := NewOrderService(log, dbService, redisService, userServiceClient)
	orderLogisticalInformation := orderproto.CreateOrderLogisticalInformationRequest{}
	orderLogisticalInformation.CommodityTypeCode = ""
	orderLogisticalInformation.ShipmentSplitMethodCode = ""
	orderLogisticalInformation.IntermediateDeliveryParty = uint32(0)
	orderLogisticalInformation.InventoryLocation = uint32(1)
	orderLogisticalInformation.ShipFrom = uint32(1)
	orderLogisticalInformation.ShipTo = uint32(3)
	orderLogisticalInformation.UltimateConsignee = uint32(0)
	orderLogisticalInformation.OrderId = uint32(1)

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderLogisticalInformationRequest
	}
	tests := []struct {
		o       *OrderService
		args    args
		want    *orderproto.CreateOrderLogisticalInformationResponse
		wantErr bool
	}{
		{
			o: orderService,
			args: args{
				ctx: ctx,
				in:  &orderLogisticalInformation,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderLogisticalInformationResp, err := tt.o.CreateOrderLogisticalInformation(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderService.CreateOrderLogisticalInformation() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, orderLogisticalInformationResp)

	}
}
