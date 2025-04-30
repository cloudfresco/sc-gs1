package invoiceservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestDebitCreditAdviceService_CreateDebitCreditAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdvice := invoiceproto.CreateDebitCreditAdviceRequest{}
	debitCreditAdvice.DebitCreditIndicatorCode = "CREDIT"
	debitCreditAdvice.TotalAmount = float64(100)
	debitCreditAdvice.TaCodeListVersion = ""
	debitCreditAdvice.TaCurrencyCode = "EUR"
	debitCreditAdvice.BillTo = uint32(0)
	debitCreditAdvice.Buyer = uint32(2)
	debitCreditAdvice.Carrier = uint32(0)
	debitCreditAdvice.DebitCreditAdviceIdentification = uint32(0)
	debitCreditAdvice.Seller = uint32(1)
	debitCreditAdvice.ShipFrom = uint32(0)
	debitCreditAdvice.ShipTo = uint32(0)
	debitCreditAdvice.UltimateConsignee = uint32(0)
	debitCreditAdvice.UserId = "auth0|673ee1a719dd4000cd5a3832"
	debitCreditAdvice.UserEmail = "sprov300@gmail.com"
	debitCreditAdvice.RequestId = "bks1m1g91jau4nkks2f0"

	debitCreditAdviceLineItem := invoiceproto.CreateDebitCreditAdviceLineItemRequest{}
	debitCreditAdviceLineItem.AdjustmentAmount = float64(100)
	debitCreditAdviceLineItem.AaCodeListVersion = ""
	debitCreditAdviceLineItem.AaCurrencyCode = "EUR"
	debitCreditAdviceLineItem.DebitCreditIndicatorCode = "CREDIT"
	debitCreditAdviceLineItem.FinancialAdjustmentReasonCode = "17"
	debitCreditAdviceLineItem.LineItemNumber = uint32(1)
	debitCreditAdviceLineItem.ParentLineItemNumber = uint32(0)
	debitCreditAdviceLineItem.DebitCreditAdviceId = uint32(1)
	debitCreditAdviceLineItem.UserId = "auth0|673ee1a719dd4000cd5a3832"
	debitCreditAdviceLineItem.UserEmail = "sprov300@gmail.com"
	debitCreditAdviceLineItem.RequestId = "bks1m1g91jau4nkks2f0"

	debitCreditAdviceLineItems := []*invoiceproto.CreateDebitCreditAdviceLineItemRequest{}
	debitCreditAdviceLineItems = append(debitCreditAdviceLineItems, &debitCreditAdviceLineItem)
	debitCreditAdvice.DebitCreditAdviceLineItems = debitCreditAdviceLineItems

	type args struct {
		ctx context.Context
		in  *invoiceproto.CreateDebitCreditAdviceRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx: ctx,
				in:  &debitCreditAdvice,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceResponse, err := tt.ds.CreateDebitCreditAdvice(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.CreateDebitCreditAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		debitCreditAdviceResult := debitCreditAdviceResponse.DebitCreditAdvice
		assert.NotNil(t, debitCreditAdviceResult)
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestDebitCreditAdviceService_GetDebitCreditAdvices(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdvices := []*invoiceproto.DebitCreditAdvice{}
	debitCreditAdvice1, err := GetDebitCreditAdvice(uint32(1), []byte{247, 169, 242, 99, 161, 225, 67, 7, 137, 144, 111, 126, 77, 148, 223, 8}, "f7a9f263-a1e1-4307-8990-6f7e4d94df08", "CREDIT", float64(100), "", "EUR", uint32(0), uint32(2), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), "2005-04-13T11:00:00Z", "2005-04-13T11:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	debitCreditAdvices = append(debitCreditAdvices, debitCreditAdvice1)

	form := invoiceproto.GetDebitCreditAdvicesRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	debitCreditAdvicesResponse := invoiceproto.GetDebitCreditAdvicesResponse{DebitCreditAdvices: debitCreditAdvices, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *invoiceproto.GetDebitCreditAdvicesRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		want    *invoiceproto.GetDebitCreditAdvicesResponse
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &debitCreditAdvicesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceResponse, err := tt.ds.GetDebitCreditAdvices(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdvices() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(debitCreditAdviceResponse, tt.want) {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdvices() = %v, want %v", debitCreditAdviceResponse, tt.want)
		}
		debitCreditAdviceResult := debitCreditAdviceResponse.DebitCreditAdvices[0]
		assert.NotNil(t, debitCreditAdviceResult)
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestDebitCreditAdviceService_GetDebitCreditAdvice(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdvice, err := GetDebitCreditAdvice(uint32(1), []byte{247, 169, 242, 99, 161, 225, 67, 7, 137, 144, 111, 126, 77, 148, 223, 8}, "f7a9f263-a1e1-4307-8990-6f7e4d94df08", "CREDIT", float64(100), "", "EUR", uint32(0), uint32(2), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), "2005-04-13T11:00:00Z", "2005-04-13T11:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	debitCreditAdviceResponse := invoiceproto.GetDebitCreditAdviceResponse{}
	debitCreditAdviceResponse.DebitCreditAdvice = debitCreditAdvice

	form := invoiceproto.GetDebitCreditAdviceRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "f7a9f263-a1e1-4307-8990-6f7e4d94df08"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetDebitCreditAdviceRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		want    *invoiceproto.GetDebitCreditAdviceResponse
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &debitCreditAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceResponse, err := tt.ds.GetDebitCreditAdvice(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdvice() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(debitCreditAdviceResponse, tt.want) {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdvice() = %v, want %v", debitCreditAdviceResponse, tt.want)
		}
		debitCreditAdviceResult := debitCreditAdviceResponse.DebitCreditAdvice
		assert.NotNil(t, debitCreditAdviceResult)
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func TestDebitCreditAdviceService_GetDebitCreditAdviceByPk(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	debitCreditAdviceService := NewDebitCreditAdviceService(log, dbService, redisService, userServiceClient)
	debitCreditAdvice, err := GetDebitCreditAdvice(uint32(1), []byte{247, 169, 242, 99, 161, 225, 67, 7, 137, 144, 111, 126, 77, 148, 223, 8}, "f7a9f263-a1e1-4307-8990-6f7e4d94df08", "CREDIT", float64(100), "", "EUR", uint32(0), uint32(2), uint32(0), uint32(0), uint32(1), uint32(0), uint32(0), uint32(0), "2005-04-13T11:00:00Z", "2005-04-13T11:00:00Z", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	debitCreditAdviceResponse := invoiceproto.GetDebitCreditAdviceByPkResponse{}
	debitCreditAdviceResponse.DebitCreditAdvice = debitCreditAdvice

	form := invoiceproto.GetDebitCreditAdviceByPkRequest{}
	gform := commonproto.GetByIdRequest{}
	gform.Id = uint32(1)
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetByIdRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *invoiceproto.GetDebitCreditAdviceByPkRequest
	}
	tests := []struct {
		ds      *DebitCreditAdviceService
		args    args
		want    *invoiceproto.GetDebitCreditAdviceByPkResponse
		wantErr bool
	}{
		{
			ds: debitCreditAdviceService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &debitCreditAdviceResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		debitCreditAdviceResponse, err := tt.ds.GetDebitCreditAdviceByPk(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdviceByPk() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(debitCreditAdviceResponse, tt.want) {
			t.Errorf("DebitCreditAdviceService.GetDebitCreditAdviceByPk() = %v, want %v", debitCreditAdviceResponse, tt.want)
		}
		debitCreditAdviceResult := debitCreditAdviceResponse.DebitCreditAdvice
		assert.NotNil(t, debitCreditAdviceResult)
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.DebitCreditIndicatorCode, "CREDIT", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TaCurrencyCode, "EUR", "they should be equal")
		assert.Equal(t, debitCreditAdviceResult.DebitCreditAdviceD.TotalAmount, float64(100), "they should be equal")
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.CreatedAt)
		assert.NotNil(t, debitCreditAdviceResult.CrUpdTime.UpdatedAt)
	}
}

func GetDebitCreditAdvice(id uint32, uuid4 []byte, idS string, debitCreditIndicatorCode string, totalAmount float64, taCodeListVersion string, taCurrencyCode string, billTo uint32, buyer uint32, carrier uint32, debitCreditAdviceIdentification uint32, seller uint32, shipFrom uint32, shipTo uint32, ultimateConsignee uint32, createdAt string, updatedAt string, createdByUserId string, updatedByUserId string) (*invoiceproto.DebitCreditAdvice, error) {
	createdAt1, err := common.ConvertTimeToTimestamp(Layout, createdAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	updatedAt1, err := common.ConvertTimeToTimestamp(Layout, updatedAt)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	debitCreditAdviceD := invoiceproto.DebitCreditAdviceD{}
	debitCreditAdviceD.Id = id
	debitCreditAdviceD.Uuid4 = uuid4
	debitCreditAdviceD.IdS = idS
	debitCreditAdviceD.DebitCreditIndicatorCode = debitCreditIndicatorCode
	debitCreditAdviceD.TotalAmount = totalAmount
	debitCreditAdviceD.TaCodeListVersion = taCodeListVersion
	debitCreditAdviceD.TaCurrencyCode = taCurrencyCode
	debitCreditAdviceD.BillTo = billTo
	debitCreditAdviceD.Buyer = buyer
	debitCreditAdviceD.Carrier = carrier
	debitCreditAdviceD.DebitCreditAdviceIdentification = debitCreditAdviceIdentification
	debitCreditAdviceD.Seller = seller
	debitCreditAdviceD.ShipFrom = shipFrom
	debitCreditAdviceD.ShipTo = shipTo
	debitCreditAdviceD.UltimateConsignee = ultimateConsignee

	crUpdTime := new(commonproto.CrUpdTime)
	crUpdTime.CreatedAt = createdAt1
	crUpdTime.UpdatedAt = updatedAt1

	crUpdUser := new(commonproto.CrUpdUser)
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	debitCreditAdviceResult := invoiceproto.DebitCreditAdvice{DebitCreditAdviceD: &debitCreditAdviceD, CrUpdUser: crUpdUser, CrUpdTime: crUpdTime}
	return &debitCreditAdviceResult, nil
}
