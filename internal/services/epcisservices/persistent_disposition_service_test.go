package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreatePersistentDisposition(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "ObjectEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreatePersistentDispositionRequest
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
				in:  &persistentDisposition,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		persistentDispositionResp, err := tt.es.CreatePersistentDisposition(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreatePersistentDisposition() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, persistentDispositionResp)
		persistentDispositionResult := persistentDispositionResp.PersistentDisposition
		assert.Equal(t, persistentDispositionResult.SetDisp, "completeness_verified", "they should be equal")
		assert.Equal(t, persistentDispositionResult.UnsetDisp, "completeness_inferred", "they should be equal")
	}
}

func TestEpcisService_GetPersistentDispositions(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	persistentDisposition, err := GetPersistentDisposition("completeness_verified", "completeness_inferred", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	persistentDispositions := []*epcisproto.PersistentDisposition{}
	persistentDispositions = append(persistentDispositions, persistentDisposition)

	persistentDispositionsResponse := epcisproto.GetPersistentDispositionsResponse{}
	persistentDispositionsResponse.PersistentDispositions = persistentDispositions

	form := epcisproto.GetPersistentDispositionsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetPersistentDispositionsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetPersistentDispositionsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &persistentDispositionsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		persistentDispositionResponse, err := tt.es.GetPersistentDispositions(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetPersistentDispositions() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(persistentDispositionResponse, tt.want) {
			t.Errorf("EpcisService.GetPersistentDispositions() = %v, want %v", persistentDispositionResponse, tt.want)
		}
		persistentDispositionResult := persistentDispositionResponse.PersistentDispositions[0]
		assert.NotNil(t, persistentDispositionResult)
		assert.Equal(t, persistentDispositionResult.SetDisp, "completeness_verified", "they should be equal")
		assert.Equal(t, persistentDispositionResult.UnsetDisp, "completeness_inferred", "they should be equal")
	}
}

func GetPersistentDisposition(setDisp string, unsetDisp string, eventId uint32, typeOfEvent string) (*epcisproto.PersistentDisposition, error) {
	persistentDisposition := epcisproto.PersistentDisposition{}
	persistentDisposition.SetDisp = setDisp
	persistentDisposition.UnsetDisp = unsetDisp
	persistentDisposition.EventId = eventId
	persistentDisposition.TypeOfEvent = typeOfEvent

	return &persistentDisposition, nil
}
