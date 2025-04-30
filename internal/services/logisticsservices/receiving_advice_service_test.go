package logisticsservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestReceivingAdviceService_CreateReceivingAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)
	receivingAdvice := logisticsproto.CreateReceivingAdviceRequest{}
	receivingAdvice.ReportingCode = "FULL_DETAILS"
	receivingAdvice.TotalAcceptedAmount = float64(0)
	receivingAdvice.TaaCodeListVersion = ""
	receivingAdvice.TaaCurrencyCode = ""
	receivingAdvice.TotalDepositAmount = float64(0)
	receivingAdvice.TdaCodeListVersion = ""
	receivingAdvice.TdaCurrencyCode = ""
	receivingAdvice.TotalNumberOfLines = uint32(0)
	receivingAdvice.TotalOnHoldAmount = float64(0)
	receivingAdvice.TohaCodeListVersion = ""
	receivingAdvice.TohaCurrencyCode = ""
	receivingAdvice.TotalRejectedAmount = float64(0)
	receivingAdvice.TraCodeListVersion = ""
	receivingAdvice.TraCurrencyCode = ""
	receivingAdvice.ReceivingAdviceTransportInformation = uint32(0)
	receivingAdvice.BillOfLadingNumber = uint32(0)
	receivingAdvice.Buyer = uint32(0)
	receivingAdvice.Carrier = uint32(0)
	receivingAdvice.ConsignmentIdentification = uint32(0)
	receivingAdvice.DeliveryNote = uint32(0)
	receivingAdvice.DespatchAdvice = uint32(0)
	receivingAdvice.InventoryLocation = uint32(0)
	receivingAdvice.PurchaseOrder = uint32(0)
	receivingAdvice.Receiver = uint32(2)
	receivingAdvice.ReceivingAdviceIdentification = uint32(0)
	receivingAdvice.Seller = uint32(0)
	receivingAdvice.ShipFrom = uint32(0)
	receivingAdvice.ShipmentIdentification = uint32(0)
	receivingAdvice.Shipper = uint32(1)
	receivingAdvice.ShipTo = uint32(3)
	receivingAdvice.DespatchAdviceDeliveryDateTimeBegin = "04/11/2011"
	receivingAdvice.DespatchAdviceDeliveryDateTimeEnd = "04/11/2011"
	receivingAdvice.PaymentDateTimeBegin = "04/11/2011"
	receivingAdvice.PaymentDateTimeEnd = "04/11/2011"
	receivingAdvice.ReceivingDateTimeBegin = "04/11/2011"
	receivingAdvice.ReceivingDateTimeEnd = "04/11/2011"
	receivingAdvice.UserId = "auth0|673ee1a719dd4000cd5a3832"
	receivingAdvice.UserEmail = "sprov300@gmail.com"
	receivingAdvice.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateReceivingAdviceRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.CreateReceivingAdviceResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx: ctx,
				in:  &receivingAdvice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceResponse, err := tt.ras.CreateReceivingAdvice(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.CreateReceivingAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		receivingAdviceResult := receivingAdviceResponse.ReceivingAdvice
		assert.NotNil(t, receivingAdviceResult)
		assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestReceivingAdviceService_GetReceivingAdvices(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)
	receivingAdvices := []*logisticsproto.ReceivingAdvice{}

	receivingAdvice, err := GetReceivingAdvice(uint32(1), []byte{162, 169, 105, 42, 107, 244, 65, 193, 160, 176, 226, 118, 145, 179, 101, 201}, "a2a9692a-6bf4-41c1-a0b0-e27691b365c9", "FULL_DETAILS", float64(0), "", "", float64(0), "", "", uint32(0), float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-13T11:45:00Z", "2011-04-13T11:45:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdvices = append(receivingAdvices, receivingAdvice)

	form := logisticsproto.GetReceivingAdvicesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	receivingAdvicesResponse := logisticsproto.GetReceivingAdvicesResponse{ReceivingAdvices: receivingAdvices, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *logisticsproto.GetReceivingAdvicesRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.GetReceivingAdvicesResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &receivingAdvicesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceResponse, err := tt.ras.GetReceivingAdvices(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.GetReceivingAdvices() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(receivingAdviceResponse, tt.want) {
			t.Errorf("ReceivingAdviceService.GetReceivingAdvices() = %v, want %v", receivingAdviceResponse, tt.want)
		}
		receivingAdviceResult := receivingAdviceResponse.ReceivingAdvices[0]
		assert.NotNil(t, receivingAdviceResult)
		assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestReceivingAdviceService_GetReceivingAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)
	receivingAdvice, err := GetReceivingAdvice(uint32(1), []byte{162, 169, 105, 42, 107, 244, 65, 193, 160, 176, 226, 118, 145, 179, 101, 201}, "a2a9692a-6bf4-41c1-a0b0-e27691b365c9", "FULL_DETAILS", float64(0), "", "", float64(0), "", "", uint32(0), float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-13T11:45:00Z", "2011-04-13T11:45:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdviceResponse := logisticsproto.GetReceivingAdviceResponse{}
	receivingAdviceResponse.ReceivingAdvice = receivingAdvice

	form := logisticsproto.GetReceivingAdviceRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "a2a9692a-6bf4-41c1-a0b0-e27691b365c9"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetReceivingAdviceRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.GetReceivingAdviceResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &receivingAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceResponse, err := tt.ras.GetReceivingAdvice(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.GetReceivingAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(receivingAdviceResponse, tt.want) {
			t.Errorf("ReceivingAdviceService.GetReceivingAdvice() = %v, want %v", receivingAdviceResponse, tt.want)
		}
		receivingAdviceResult := receivingAdviceResponse.ReceivingAdvice
		assert.NotNil(t, receivingAdviceResult)
		assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestReceivingAdviceService_GetReceivingAdviceByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	receivingAdviceService := NewReceivingAdviceService(log, dbService, redisService, userServiceClient)
	receivingAdvice, err := GetReceivingAdvice(uint32(1), []byte{162, 169, 105, 42, 107, 244, 65, 193, 160, 176, 226, 118, 145, 179, 101, 201}, "a2a9692a-6bf4-41c1-a0b0-e27691b365c9", "FULL_DETAILS", float64(0), "", "", float64(0), "", "", uint32(0), float64(0), "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-11T23:00:00Z", "2011-04-13T11:45:00Z", "2011-04-13T11:45:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	receivingAdviceResponse := logisticsproto.GetReceivingAdviceByPkResponse{}
	receivingAdviceResponse.ReceivingAdvice = receivingAdvice

	form := logisticsproto.GetReceivingAdviceByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetReceivingAdviceByPkRequest
	}
	tests := []struct {
		ras     *ReceivingAdviceService
		args    args
		want    *logisticsproto.GetReceivingAdviceByPkResponse
		wantErr bool
	}{
		{
			ras: receivingAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &receivingAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		receivingAdviceResponse, err := tt.ras.GetReceivingAdviceByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReceivingAdviceService.GetReceivingAdviceByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(receivingAdviceResponse, tt.want) {
			t.Errorf("ReceivingAdviceService.GetReceivingAdviceByPk() = %v, want %v", receivingAdviceResponse, tt.want)
		}
		receivingAdviceResult := receivingAdviceResponse.ReceivingAdvice
		assert.NotNil(t, receivingAdviceResult)
		assert.Equal(t, receivingAdviceResult.ReceivingAdviceD.ReportingCode, "FULL_DETAILS", "they should be equal")
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, receivingAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func GetReceivingAdvice(id uint32, uuid4 []byte, idS string, reportingCode string, totalAcceptedAmount float64, taaCodeListVersion string, taaCurrencyCode string, totalDepositAmount float64, tdaCodeListVersion string, tdaCurrencyCode string, totalNumberOfLines uint32, totalOnHoldAmount float64, tohaCodeListVersion string, tohaCurrencyCode string, totalRejectedAmount float64, traCodeListVersion string, traCurrencyCode string, receivingAdviceTransportInformation uint32, billOfLadingNumber uint32, buyer uint32, carrier uint32, consignmentIdentification uint32, deliveryNote uint32, despatchAdvice uint32, inventoryLocation uint32, purchaseOrder uint32, receiver uint32, receivingAdviceIdentification uint32, seller uint32, shipFrom uint32, shipmentIdentification uint32, shipper uint32, shipTo uint32, despatchAdviceDeliveryDateTimeBegin string, despatchAdviceDeliveryDateTimeEnd string, paymentDateTimeBegin string, paymentDateTimeEnd string, receivingDateTimeBegin string, receivingDateTimeEnd string, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*logisticsproto.ReceivingAdvice, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	despatchAdviceDeliveryDateTimeBegin1, err := common.ConvertTimeToTimestamp(Layout, despatchAdviceDeliveryDateTimeBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	despatchAdviceDeliveryDateTimeEnd1, err := common.ConvertTimeToTimestamp(Layout, despatchAdviceDeliveryDateTimeEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	paymentDateTimeBegin1, err := common.ConvertTimeToTimestamp(Layout, paymentDateTimeBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	paymentDateTimeEnd1, err := common.ConvertTimeToTimestamp(Layout, paymentDateTimeEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	receivingDateTimeBegin1, err := common.ConvertTimeToTimestamp(Layout, receivingDateTimeBegin)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	receivingDateTimeEnd1, err := common.ConvertTimeToTimestamp(Layout, receivingDateTimeEnd)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	receivingAdviceD := logisticsproto.ReceivingAdviceD{}
	receivingAdviceD.Id = id
	receivingAdviceD.Uuid4 = uuid4
	receivingAdviceD.IdS = idS
	receivingAdviceD.ReportingCode = reportingCode
	receivingAdviceD.TotalAcceptedAmount = totalAcceptedAmount
	receivingAdviceD.TaaCodeListVersion = taaCodeListVersion
	receivingAdviceD.TaaCurrencyCode = taaCurrencyCode
	receivingAdviceD.TotalDepositAmount = totalDepositAmount
	receivingAdviceD.TdaCodeListVersion = tdaCodeListVersion
	receivingAdviceD.TdaCurrencyCode = tdaCurrencyCode
	receivingAdviceD.TotalNumberOfLines = totalNumberOfLines
	receivingAdviceD.TotalOnHoldAmount = totalOnHoldAmount
	receivingAdviceD.TohaCodeListVersion = tohaCodeListVersion
	receivingAdviceD.TohaCurrencyCode = tohaCurrencyCode
	receivingAdviceD.TotalRejectedAmount = totalRejectedAmount
	receivingAdviceD.TraCodeListVersion = traCodeListVersion
	receivingAdviceD.TraCurrencyCode = traCurrencyCode
	receivingAdviceD.ReceivingAdviceTransportInformation = receivingAdviceTransportInformation
	receivingAdviceD.BillOfLadingNumber = billOfLadingNumber
	receivingAdviceD.Buyer = buyer
	receivingAdviceD.Carrier = carrier
	receivingAdviceD.ConsignmentIdentification = consignmentIdentification
	receivingAdviceD.DeliveryNote = deliveryNote
	receivingAdviceD.DespatchAdvice = despatchAdvice
	receivingAdviceD.InventoryLocation = inventoryLocation
	receivingAdviceD.PurchaseOrder = purchaseOrder
	receivingAdviceD.Receiver = receiver
	receivingAdviceD.ReceivingAdviceIdentification = receivingAdviceIdentification
	receivingAdviceD.Seller = seller
	receivingAdviceD.ShipFrom = shipFrom
	receivingAdviceD.ShipmentIdentification = shipmentIdentification
	receivingAdviceD.Shipper = shipper
	receivingAdviceD.ShipTo = shipTo

	receivingAdviceT := logisticsproto.ReceivingAdviceT{}
	receivingAdviceT.DespatchAdviceDeliveryDateTimeBegin = despatchAdviceDeliveryDateTimeBegin1
	receivingAdviceT.DespatchAdviceDeliveryDateTimeEnd = despatchAdviceDeliveryDateTimeEnd1
	receivingAdviceT.PaymentDateTimeBegin = paymentDateTimeBegin1
	receivingAdviceT.PaymentDateTimeEnd = paymentDateTimeEnd1
	receivingAdviceT.ReceivingDateTimeBegin = receivingDateTimeBegin1
	receivingAdviceT.ReceivingDateTimeEnd = receivingDateTimeEnd1

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	receivingAdviceResult := logisticsproto.ReceivingAdvice{ReceivingAdviceD: &receivingAdviceD, ReceivingAdviceT: &receivingAdviceT, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &receivingAdviceResult, nil
}
