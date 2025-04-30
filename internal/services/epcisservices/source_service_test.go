package epcisservices

import (
	"context"
	"reflect"
	"testing"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
)

func TestEpcisService_CreateSource(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:pgln:4012345.00225"
	source.EventId = uint32(1)
	source.TypeOfEvent = "ObjectEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateSourceRequest
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
				in:  &source,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sourceResp, err := tt.es.CreateSource(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateSource() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, sourceResp)
		sourceResult := sourceResp.Source
		assert.Equal(t, sourceResult.SourceType, "location", "they should be equal")
		assert.Equal(t, sourceResult.Source, "urn:epc:id:pgln:4012345.00225", "they should be equal")
	}
}

func TestEpcisService_GetSources(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	source3, err := GetSource("location", "urn:epc:id:sgln:4012345.00225.0", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	source2, err := GetSource("possessing_party", "urn:epc:id:pgln:4012345.00225", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	source1, err := GetSource("owning_party", "urn:epc:id:pgln:4012345.00225", uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	sources := []*epcisproto.Source{}
	sources = append(sources, source3, source2, source1)

	sourcesResponse := epcisproto.GetSourcesResponse{}
	sourcesResponse.Sources = sources

	form := epcisproto.GetSourcesRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetSourcesRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetSourcesResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &sourcesResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sourceResponse, err := tt.es.GetSources(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetSources() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(sourceResponse, tt.want) {
			t.Errorf("EpcisService.GetSources() = %v, want %v", sourceResponse, tt.want)
		}
		sourceResult := sourceResponse.Sources[1]
		assert.NotNil(t, sourceResult)
		assert.Equal(t, sourceResult.SourceType, "possessing_party", "they should be equal")
		assert.Equal(t, sourceResult.Source, "urn:epc:id:pgln:4012345.00225", "they should be equal")
	}
}

func GetSource(sourceType string, src string, eventId uint32, typeOfEvent string) (*epcisproto.Source, error) {
	source := epcisproto.Source{}
	source.SourceType = sourceType
	source.Source = src
	source.EventId = eventId
	source.TypeOfEvent = typeOfEvent

	return &source, nil
}
