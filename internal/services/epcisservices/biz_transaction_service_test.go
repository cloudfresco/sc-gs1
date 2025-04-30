package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreateBizTransaction(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "pedigree"
	bizTransaction.BizTransaction = "urn:epc:id:gsrn:0614141.0000010254"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "ObjectEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateBizTransactionRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &bizTransaction,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		bizTransactionResp, err := tt.es.CreateBizTransaction(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateBizTransaction() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, bizTransactionResp)
		bizTransactionResult := bizTransactionResp.BizTransaction
		assert.Equal(t, bizTransactionResult.BizTransactionType, "pedigree", "they should be equal")
		assert.Equal(t, bizTransactionResult.BizTransaction, "urn:epc:id:gsrn:0614141.0000010254", "they should be equal")
	}
}

func TestEpcisService_GetBizTransactions(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	bizTransaction1, err := GetBizTransaction("pedigree", "urn:epc:id:gsrn:0614141.0000010253", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	bizTransaction2, err := GetBizTransaction("po", "urn:epc:id:gdti:0614141.00001.1618034", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	bizTransactions := []*epcisproto.BizTransaction{}
	bizTransactions = append(bizTransactions, bizTransaction2, bizTransaction1)

	bizTransactionsResponse := epcisproto.GetBizTransactionsResponse{}
	bizTransactionsResponse.BizTransactions = bizTransactions

	form := epcisproto.GetBizTransactionsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetBizTransactionsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetBizTransactionsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &bizTransactionsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		bizTransactionResponse, err := tt.es.GetBizTransactions(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetBizTransactions() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(bizTransactionResponse, tt.want) {
			t.Errorf("EpcisService.GetBizTransactions() = %v, want %v", bizTransactionResponse, tt.want)
		}
		bizTransactionResult := bizTransactionResponse.BizTransactions[1]
		assert.NotNil(t, bizTransactionResult)
		assert.Equal(t, bizTransactionResult.BizTransactionType, "pedigree", "they should be equal")
		assert.Equal(t, bizTransactionResult.BizTransaction, "urn:epc:id:gsrn:0614141.0000010253", "they should be equal")
	}
}

func GetBizTransaction(bizTransactionType string, bTransaction string, eventId uint32, typeOfEvent string) (*epcisproto.BizTransaction, error) {
	bizTransaction := epcisproto.BizTransaction{}
	bizTransaction.BizTransactionType = bizTransactionType
	bizTransaction.BizTransaction = bTransaction
	bizTransaction.EventId = eventId
	bizTransaction.TypeOfEvent = typeOfEvent

	return &bizTransaction, nil
}
