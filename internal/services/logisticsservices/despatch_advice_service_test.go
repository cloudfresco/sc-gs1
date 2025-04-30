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

func TestDespatchAdviceService_CreateDespatchAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdvice := logisticsproto.CreateDespatchAdviceRequest{}
	despatchAdvice.DeliveryTypeCode = ""
	despatchAdvice.RackIdAtPickUpLocation = ""
	despatchAdvice.TotalDepositAmount = float64(0)
	despatchAdvice.TdaCodeListVersion = ""
	despatchAdvice.TdaCurrencyCode = ""
	despatchAdvice.TotalNumberOfLines = uint32(0)
	despatchAdvice.BlanketOrder = uint32(0)
	despatchAdvice.Buyer = uint32(0)
	despatchAdvice.Carrier = uint32(0)
	despatchAdvice.Contract = uint32(0)
	despatchAdvice.CustomerDocumentReference = uint32(0)
	despatchAdvice.DeclarantsCustomsIdentity = uint32(0)
	despatchAdvice.DeliveryNote = uint32(0)
	despatchAdvice.DeliverySchedule = uint32(0)
	despatchAdvice.DespatchAdviceIdentification = uint32(0)
	despatchAdvice.FreightForwarder = uint32(0)
	despatchAdvice.InventoryLocation = uint32(0)
	despatchAdvice.Invoice = uint32(0)
	despatchAdvice.Invoicee = uint32(0)
	despatchAdvice.LogisticServiceProvider = uint32(0)
	despatchAdvice.OrderResponse = uint32(0)
	despatchAdvice.PickUpLocation = uint32(0)
	despatchAdvice.ProductCertification = uint32(0)
	despatchAdvice.PromotionalDeal = uint32(0)
	despatchAdvice.PurchaseConditions = uint32(0)
	despatchAdvice.PurchaseOrder = uint32(0)
	despatchAdvice.Receiver = uint32(2)
	despatchAdvice.ReturnsInstruction = uint32(0)
	despatchAdvice.Seller = uint32(0)
	despatchAdvice.ShipFrom = uint32(0)
	despatchAdvice.Shipper = uint32(1)
	despatchAdvice.ShipTo = uint32(3)
	despatchAdvice.Specification = uint32(0)
	despatchAdvice.TransportInstruction = uint32(0)
	despatchAdvice.UltimateConsignee = uint32(0)
	despatchAdvice.UserId = "auth0|673ee1a719dd4000cd5a3832"
	despatchAdvice.UserEmail = "sprov300@gmail.com"
	despatchAdvice.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *logisticsproto.CreateDespatchAdviceRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.CreateDespatchAdviceResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx: ctx,
				in:  &despatchAdvice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceResponse, err := tt.das.CreateDespatchAdvice(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.CreateDespatchAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		despatchAdviceResult := despatchAdviceResponse.DespatchAdvice
		assert.NotNil(t, despatchAdviceResult)
		assert.NotNil(t, despatchAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, despatchAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestDespatchAdviceService_GetDespatchAdvices(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdvice, err := GetDespatchAdvice(uint32(1), []byte{67, 169, 56, 210, 69, 80, 77, 193, 128, 216, 0, 245, 253, 19, 76, 202}, "43a938d2-4550-4dc1-80d8-00f5fd134cca", "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), uint32(0), uint32(0), uint32(0), "2011-04-11T09:00:00Z", "2011-04-11T09:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}
	despatchAdvices := []*logisticsproto.DespatchAdvice{}
	despatchAdvices = append(despatchAdvices, despatchAdvice)

	form := logisticsproto.GetDespatchAdvicesRequest{}
	form.Limit = "8"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	despatchAdvicesResponse := logisticsproto.GetDespatchAdvicesResponse{DespatchAdvices: despatchAdvices, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *logisticsproto.GetDespatchAdvicesRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.GetDespatchAdvicesResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &despatchAdvicesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdvicesResp, err := tt.das.GetDespatchAdvices(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.GetDespatchAdvices() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(despatchAdvicesResp, tt.want) {
			t.Errorf("DespatchAdviceService.GetDespatchAdvices() = %v, want %v", despatchAdvicesResp, tt.want)
		}
		assert.NotNil(t, despatchAdvicesResp)
	}
}

func TestDespatchAdviceService_GetDespatchAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdvice, err := GetDespatchAdvice(uint32(1), []byte{67, 169, 56, 210, 69, 80, 77, 193, 128, 216, 0, 245, 253, 19, 76, 202}, "43a938d2-4550-4dc1-80d8-00f5fd134cca", "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), uint32(0), uint32(0), uint32(0), "2011-04-11T09:00:00Z", "2011-04-11T09:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}
	despatchAdviceResponse := logisticsproto.GetDespatchAdviceResponse{}
	despatchAdviceResponse.DespatchAdvice = despatchAdvice
	gform := commonproto.GetRequest{}
	gform.Id = "43a938d2-4550-4dc1-80d8-00f5fd134cca"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := logisticsproto.GetDespatchAdviceRequest{}
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetDespatchAdviceRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.GetDespatchAdviceResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &despatchAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceResp, err := tt.das.GetDespatchAdvice(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.GetDespatchAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(despatchAdviceResp, tt.want) {
			t.Errorf("DespatchAdviceService.GetDespatchAdvice() = %v, want %v", despatchAdviceResp, tt.want)
		}
		assert.NotNil(t, despatchAdviceResp)
	}
}

func TestDespatchAdviceService_GetDespatchAdviceByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	despatchAdviceService := NewDespatchAdviceService(log, dbService, redisService, userServiceClient)
	despatchAdvice, err := GetDespatchAdvice(uint32(1), []byte{67, 169, 56, 210, 69, 80, 77, 193, 128, 216, 0, 245, 253, 19, 76, 202}, "43a938d2-4550-4dc1-80d8-00f5fd134cca", "", "", float64(0), "", "", uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(0), uint32(2), uint32(0), uint32(0), uint32(0), uint32(1), uint32(3), uint32(0), uint32(0), uint32(0), "2011-04-11T09:00:00Z", "2011-04-11T09:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	despatchAdviceResponse := logisticsproto.GetDespatchAdviceByPkResponse{}
	despatchAdviceResponse.DespatchAdvice = despatchAdvice

	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"

	form := logisticsproto.GetDespatchAdviceByPkRequest{}
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *logisticsproto.GetDespatchAdviceByPkRequest
	}
	tests := []struct {
		das     *DespatchAdviceService
		args    args
		want    *logisticsproto.GetDespatchAdviceByPkResponse
		wantErr bool
	}{
		{
			das: despatchAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &despatchAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		despatchAdviceResp, err := tt.das.GetDespatchAdviceByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DespatchAdviceService.GetDespatchAdviceByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(despatchAdviceResp, tt.want) {
			t.Errorf("DespatchAdviceService.GetDespatchAdviceByPk() = %v, want %v", despatchAdviceResp, tt.want)
		}
		assert.NotNil(t, despatchAdviceResp)
	}
}

func GetDespatchAdvice(id uint32, uuid4 []byte, idS string, deliveryTypeCode string, rackIdAtPickUpLocation string, totalDepositAmount float64, tdaCodeListVersion string, tdaCurrencyCode string, totalNumberOfLines uint32, blanketOrder uint32, buyer uint32, carrier uint32, contract uint32, customerDocumentReference uint32, declarantsCustomsIdentity uint32, deliveryNote uint32, deliverySchedule uint32, despatchAdviceIdentification uint32, freightForwarder uint32, inventoryLocation uint32, invoice uint32, invoicee uint32, logisticServiceProvider uint32, orderResponse uint32, pickUpLocation uint32, productCertification uint32, promotionalDeal uint32, purchaseConditions uint32, purchaseOrder uint32, receiver uint32, returnsInstruction uint32, seller uint32, shipFrom uint32, shipper uint32, shipTo uint32, specification uint32, transportInstruction uint32, ultimateConsignee uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*logisticsproto.DespatchAdvice, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	despatchAdviceD := logisticsproto.DespatchAdviceD{}
	despatchAdviceD.Id = id
	despatchAdviceD.Uuid4 = uuid4
	despatchAdviceD.IdS = idS
	despatchAdviceD.DeliveryTypeCode = deliveryTypeCode
	despatchAdviceD.RackIdAtPickUpLocation = rackIdAtPickUpLocation
	despatchAdviceD.TotalDepositAmount = totalDepositAmount
	despatchAdviceD.TdaCodeListVersion = tdaCodeListVersion
	despatchAdviceD.TdaCurrencyCode = tdaCurrencyCode
	despatchAdviceD.TotalNumberOfLines = totalNumberOfLines
	despatchAdviceD.BlanketOrder = blanketOrder
	despatchAdviceD.Buyer = buyer
	despatchAdviceD.Carrier = carrier
	despatchAdviceD.Contract = contract
	despatchAdviceD.CustomerDocumentReference = customerDocumentReference
	despatchAdviceD.DeclarantsCustomsIdentity = declarantsCustomsIdentity
	despatchAdviceD.DeliveryNote = deliveryNote
	despatchAdviceD.DeliverySchedule = deliverySchedule
	despatchAdviceD.DespatchAdviceIdentification = despatchAdviceIdentification
	despatchAdviceD.FreightForwarder = freightForwarder
	despatchAdviceD.InventoryLocation = inventoryLocation
	despatchAdviceD.Invoice = invoice
	despatchAdviceD.Invoicee = invoicee
	despatchAdviceD.LogisticServiceProvider = logisticServiceProvider
	despatchAdviceD.OrderResponse = orderResponse
	despatchAdviceD.PickUpLocation = pickUpLocation
	despatchAdviceD.ProductCertification = productCertification
	despatchAdviceD.PromotionalDeal = promotionalDeal
	despatchAdviceD.PurchaseConditions = purchaseConditions
	despatchAdviceD.PurchaseOrder = purchaseOrder
	despatchAdviceD.Receiver = receiver
	despatchAdviceD.ReturnsInstruction = returnsInstruction
	despatchAdviceD.Seller = seller
	despatchAdviceD.ShipFrom = shipFrom
	despatchAdviceD.Shipper = shipper
	despatchAdviceD.ShipTo = shipTo
	despatchAdviceD.Specification = specification
	despatchAdviceD.TransportInstruction = transportInstruction
	despatchAdviceD.UltimateConsignee = ultimateConsignee

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	despatchAdviceResult := logisticsproto.DespatchAdvice{DespatchAdviceD: &despatchAdviceD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &despatchAdviceResult, nil
}
