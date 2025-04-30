package orderservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestOrderResponseService_CreateOrderResponse(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)

	orderResponse := orderproto.CreateOrderResponseRequest{}
	orderResponse.OrderResponseReasonCode = "DISCONTINUED_LINE"
	orderResponse.ResponseStatusCode = "REJECTED"
	orderResponse.TotalMonetaryAmountExcludingTaxes = float64(1257)
	orderResponse.TmaetCodeListVersion = ""
	orderResponse.TmaetCurrencyCode = ""
	orderResponse.TotalMonetaryAmountIncludingTaxes = float64(0)
	orderResponse.TmaitCodeListVersion = ""
	orderResponse.TmaitCurrencyCode = ""
	orderResponse.TotalTaxAmount = float64(450)
	orderResponse.TtaCodeListVersion = ""
	orderResponse.TtaCurrencyCode = ""
	orderResponse.AmendedDateTimeValue = uint32(0)
	orderResponse.BillTo = uint32(0)
	orderResponse.Buyer = uint32(2)
	orderResponse.OrderResponseIdentification = uint32(4)
	orderResponse.OriginalOrder = uint32(0)
	orderResponse.SalesOrder = uint32(0)
	orderResponse.Seller = uint32(1)
	orderResponse.ShipTo = uint32(0)
	orderResponse.UserId = "auth0|673ee1a719dd4000cd5a3832"
	orderResponse.UserEmail = "sprov300@gmail.com"
	orderResponse.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *orderproto.CreateOrderResponseRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.CreateOrderResponseResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx: ctx,
				in:  &orderResponse,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {

		orderRes, err := tt.o.CreateOrderResponse(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.CreateOrderResponse() error = %v, wantErr %v", err, tt.wantErr)
			return
		}

		assert.NotNil(t, orderRes)
		orderRespResult := orderRes.OrderResponse
		assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.TotalMonetaryAmountExcludingTaxes, float64(1257), "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.TotalTaxAmount, float64(450), "they should be equal")
	}
}

func TestOrderResponseService_GetOrderResponses(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)
	orderResponses := []*orderproto.OrderResponse{}
	orderResponse1, err := GetOrderResponse(uint32(2), []byte{5, 66, 5, 107, 127, 180, 74, 121, 158, 14, 230, 91, 219, 7, 162, 162}, "0542056b-7fb4-4a79-9e0e-e65bdb07a2a2", "DISCONTINUED_LINE", "REJECTED", float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(2), uint32(4), uint32(0), uint32(0), uint32(1), uint32(0), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponse2, err := GetOrderResponse(uint32(3), []byte{169, 207, 75, 196, 191, 75, 64, 19, 149, 211, 211, 187, 76, 122, 249, 82}, "a9cf4bc4-bf4b-4013-95d3-d3bb4c7af952", "", "MODIFIED", float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(2), uint32(5), uint32(0), uint32(0), uint32(1), uint32(0), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResponses = append(orderResponses, orderResponse2, orderResponse1)

	form := orderproto.GetOrderResponsesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MQ=="
	ordersResps := orderproto.GetOrderResponsesResponse{OrderResponses: orderResponses, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *orderproto.GetOrderResponsesRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.GetOrderResponsesResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &ordersResps,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		ordersResp, err := tt.o.GetOrderResponses(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.GetOrderResponses() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(ordersResp, tt.want) {
			t.Errorf("OrderResponseService.GetOrderResponses() = %v, want %v", ordersResp, tt.want)
		}
		assert.NotNil(t, ordersResp)
		orderRespResult := ordersResp.OrderResponses[1]
		assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
	}
}

func TestOrderResponseService_GetOrderResponse(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)
	orderResponse, err := GetOrderResponse(uint32(2), []byte{5, 66, 5, 107, 127, 180, 74, 121, 158, 14, 230, 91, 219, 7, 162, 162}, "0542056b-7fb4-4a79-9e0e-e65bdb07a2a2", "DISCONTINUED_LINE", "REJECTED", float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(2), uint32(4), uint32(0), uint32(0), uint32(1), uint32(0), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	form := orderproto.GetOrderResponseRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "0542056b-7fb4-4a79-9e0e-e65bdb07a2a2"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	orderResp := orderproto.GetOrderResponseResponse{}
	orderResp.OrderResponse = orderResponse

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderResponseRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.GetOrderResponseResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderRes, err := tt.o.GetOrderResponse(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.GetOrderResponse() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderRes, tt.want) {
			t.Errorf("OrderResponseService.GetOrderResponse() = %v, want %v", orderRes, tt.want)
		}
		assert.NotNil(t, orderRes)
		orderRespResult := orderRes.OrderResponse
		assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")

	}
}

func TestOrderResponseService_GetOrderResponseByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	orderResponseService := NewOrderResponseService(log, dbService, redisService, userServiceClient)

	orderResponse, err := GetOrderResponse(uint32(2), []byte{5, 66, 5, 107, 127, 180, 74, 121, 158, 14, 230, 91, 219, 7, 162, 162}, "0542056b-7fb4-4a79-9e0e-e65bdb07a2a2", "DISCONTINUED_LINE", "REJECTED", float64(0), "", "", float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(2), uint32(4), uint32(0), uint32(0), uint32(1), uint32(0), "2019-03-11T10:04:26Z", "2019-03-11T10:04:26Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	orderResp := orderproto.GetOrderResponseByPkResponse{}
	orderResp.OrderResponse = orderResponse

	form := orderproto.GetOrderResponseByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(2)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *orderproto.GetOrderResponseByPkRequest
	}
	tests := []struct {
		o       *OrderResponseService
		args    args
		want    *orderproto.GetOrderResponseByPkResponse
		wantErr bool
	}{
		{
			o: orderResponseService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &orderResp,
			wantErr: false,
		},
	}
	for _, tt := range tests {

		orderRes, err := tt.o.GetOrderResponseByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("OrderResponseService.GetOrderResponseByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(orderRes, tt.want) {
			t.Errorf("OrderResponseService.GetOrderResponseByPk() = %v, want %v", orderRes, tt.want)
		}
		assert.NotNil(t, orderRes)
		orderRespResult := orderRes.OrderResponse
		assert.Equal(t, orderRespResult.OrderResponseD.OrderResponseReasonCode, "DISCONTINUED_LINE", "they should be equal")
		assert.Equal(t, orderRespResult.OrderResponseD.ResponseStatusCode, "REJECTED", "they should be equal")
	}
}

func GetOrderResponse(id uint32, uuid4 []byte, idS string, orderResponseReasonCode string, responseStatusCode string, totalMonetaryAmountExcludingTaxes float64, tmaetCodeListVersion string, tmaetCurrencyCode string, totalMonetaryAmountIncludingTaxes float64, tmaitCodeListVersion string, tmaitCurrencyCode string, totalTaxAmount float64, ttaCodeListVersion string, ttaCurrencyCode string, amendedDateTimeValue uint32, billTo uint32, buyer uint32, orderResponseIdentification uint32, originalOrder uint32, salesOrder uint32, seller uint32, shipTo uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*orderproto.OrderResponse, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	orderResponseD := orderproto.OrderResponseD{}
	orderResponseD.Id = id
	orderResponseD.Uuid4 = uuid4
	orderResponseD.IdS = idS
	orderResponseD.OrderResponseReasonCode = orderResponseReasonCode
	orderResponseD.ResponseStatusCode = responseStatusCode
	orderResponseD.TotalMonetaryAmountExcludingTaxes = totalMonetaryAmountExcludingTaxes
	orderResponseD.TmaetCodeListVersion = tmaetCodeListVersion
	orderResponseD.TmaetCurrencyCode = tmaetCurrencyCode
	orderResponseD.TotalMonetaryAmountIncludingTaxes = totalMonetaryAmountIncludingTaxes
	orderResponseD.TmaitCodeListVersion = tmaitCodeListVersion
	orderResponseD.TmaitCurrencyCode = tmaitCurrencyCode
	orderResponseD.TotalTaxAmount = totalTaxAmount
	orderResponseD.TtaCodeListVersion = ttaCodeListVersion
	orderResponseD.TtaCurrencyCode = ttaCurrencyCode
	orderResponseD.AmendedDateTimeValue = amendedDateTimeValue
	orderResponseD.BillTo = billTo
	orderResponseD.Buyer = buyer
	orderResponseD.OrderResponseIdentification = orderResponseIdentification
	orderResponseD.OriginalOrder = originalOrder
	orderResponseD.SalesOrder = salesOrder
	orderResponseD.Seller = seller
	orderResponseD.ShipTo = shipTo

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	orderResponse := orderproto.OrderResponse{OrderResponseD: &orderResponseD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &orderResponse, nil
}
