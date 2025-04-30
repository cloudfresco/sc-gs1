package epcisservices

import (
	"context"
	"reflect"
	"testing"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/cloudfresco/sc-gs1/test"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestEpcisService_CreateAggregationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	aggregationEvent := epcisproto.CreateAggregationEventRequest{}
	aggregationEvent.EventId = "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0"
	aggregationEvent.EventTimeZoneOffset = "+02:00"
	aggregationEvent.Certification = "https://accreditation-council.example.org/certificate/ABC12345"
	aggregationEvent.EventTime = "04/05/2005"
	aggregationEvent.Reason = "incorrect_data"
	aggregationEvent.DeclarationTime = "07/11/2005"
	aggregationEvent.ParentId = "urn:epc:id:sscc:0614141.1234567890"
	aggregationEvent.Action = "OBSERVE"
	aggregationEvent.BizStep = "receiving"
	aggregationEvent.Disposition = "in_progress"
	aggregationEvent.ReadPoint = "urn:epc:id:sgln:0037000.00729.8202"
	aggregationEvent.BizLocation = "urn:epc:id:sgln:0037000.00729.8202"
	aggregationEvent.UserId = "auth0|673ee1a719dd4000cd5a3832"
	aggregationEvent.UserEmail = "sprov300@gmail.com"
	aggregationEvent.RequestId = "bks1m1g91jau4nkks2f0"

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "AggregationEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"
	persistentDispositions := []*epcisproto.CreatePersistentDispositionRequest{}
	persistentDispositions = append(persistentDispositions, &persistentDisposition)
	aggregationEvent.PersistentDispositions = persistentDispositions

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:sgtin:0614141.107346.2017"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "AggregationEvent"
	epc.TypeOfEpc = "output"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	epcs := []*epcisproto.CreateEpcRequest{}
	epcs = append(epcs, &epc)
	aggregationEvent.ChildEpcs = epcs

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "po"
	bizTransaction.BizTransaction = "urn:epc:id:gdti:0614141.00001.1618034"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "AggregationEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	bizTransactions := []*epcisproto.CreateBizTransactionRequest{}
	bizTransactions = append(bizTransactions, &bizTransaction)
	aggregationEvent.BizTransactionList = bizTransactions

	quantityElement := epcisproto.CreateQuantityElementRequest{}
	quantityElement.EpcClass = "urn:epc:idpat:sgtin:4012345.098765"
	quantityElement.Quantity = float64(200)
	quantityElement.Uom = "KGM"
	quantityElement.EventId = uint32(1)
	quantityElement.TypeOfEvent = "AggregationEvent"
	quantityElement.TypeOfQuantity = "output"
	quantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	quantityElement.UserEmail = "sprov300@gmail.com"
	quantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	quantityElements := []*epcisproto.CreateQuantityElementRequest{}
	quantityElements = append(quantityElements, &quantityElement)
	aggregationEvent.ChildQuantityList = quantityElements

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:sgln:4012345.00225.0"
	source.EventId = uint32(1)
	source.TypeOfEvent = "AggregationEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"
	sources := []*epcisproto.CreateSourceRequest{}
	sources = append(sources, &source)
	aggregationEvent.SourceList = sources

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777.0"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "AggregationEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	destinations := []*epcisproto.CreateDestinationRequest{}
	destinations = append(destinations, &destination)
	aggregationEvent.DestinationList = destinations

	sensorElement := epcisproto.CreateSensorElementRequest{}
	sensorElement.DeviceId = "urn:epc:id:giai:4000001.111"
	sensorElement.DeviceMetadata = "https://id.gs1.org/8004/4000001111"
	sensorElement.RawData = "https://example.org/8004/401234599999"
	sensorElement.DataProcessingMethod = "https://example.com/253/4012345000054987"
	sensorElement.BizRules = "https://example.com/253/4012345000054987"
	sensorElement.SensorTime = "04/11/2023"
	sensorElement.StartTime = "04/11/2023"
	sensorElement.EndTime = "04/11/2023"
	sensorElement.EventId = uint32(1)
	sensorElement.TypeOfEvent = "AggregationEvent"
	sensorElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	sensorElement.UserEmail = "sprov300@gmail.com"
	sensorElement.RequestId = "bks1m1g91jau4nkks2f0"

	sensorReport := epcisproto.CreateSensorReportRequest{}
	sensorReport.SensorReportType = "Temperature"
	sensorReport.DeviceId = "urn:epc:id:giai:4000001.111"
	sensorReport.RawData = "https://example.org/8004/401234599999"
	sensorReport.DataProcessingMethod = "https://example.com/253/4012345000054987"
	sensorReport.Microorganism = "https://www.ncbi.nlm.nih.gov/taxonomy/1126011"
	sensorReport.ChemicalSubstance = "https://identifiers.org/inchikey:CZMRCDWAGMRECN-UGDNZRGBSA-N"
	sensorReport.SensorValue = float64(26)
	sensorReport.Component = "example:x"
	sensorReport.StringValue = "SomeString"
	sensorReport.BooleanValue = true
	sensorReport.HexBinaryValue = "f0f0f0"
	sensorReport.UriValue = "https://id.gs1.org/8004/4000001111"
	sensorReport.MinValue = float64(26)
	sensorReport.MaxValue = float64(26.2)
	sensorReport.MeanValue = float64(13.2)
	sensorReport.PercRank = float64(50)
	sensorReport.PercValue = float64(12.7)
	sensorReport.Uom = "CEL"
	sensorReport.SDev = float64(0.1)
	sensorReport.DeviceMetadata = "https://id.gs1.org/8004/4000001111"
	sensorReport.SensorReportTime = "04/11/2023"
	sensorReport.UserId = "auth0|673ee1a719dd4000cd5a3832"
	sensorReport.UserEmail = "sprov300@gmail.com"
	sensorReport.RequestId = "bks1m1g91jau4nkks2f0"

	sensorReports := []*epcisproto.CreateSensorReportRequest{}
	sensorReports = append(sensorReports, &sensorReport)
	sensorElement.SensorReports = sensorReports
	sensorElements := []*epcisproto.CreateSensorElementRequest{}
	sensorElements = append(sensorElements, &sensorElement)
	aggregationEvent.SensorElementList = sensorElements

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateAggregationEventRequest
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
				in:  &aggregationEvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		aggregationEventResp, err := tt.es.CreateAggregationEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateAggregationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, aggregationEventResp)
		aggregationEventResult := aggregationEventResp.AggregationEvent
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventId, "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Action, "OBSERVE", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.ReadPoint, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizLocation, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
	}
}

func TestEpcisService_GetAggregationEvents(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	aggregationEvents := []*epcisproto.AggregationEvent{}

	aggregationEvent, err := GetAggregationEvent(uint32(1), []byte{233, 210, 57, 150, 113, 115, 73, 53, 157, 104, 230, 4, 187, 159, 48, 158}, "e9d23996-7173-4935-9d68-e604bb9f309e", "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0", "+02:00", "https://accreditation-council.example.org/certificate/ABC12345", "2013-06-08T14:58:56Z", "incorrect_data", "2013-11-07T14:00:00Z", "urn:epc:id:sscc:0614141.1234567890", "OBSERVE", "receiving", "in_progress", "urn:epc:id:sgln:0037000.00729.8202", "urn:epc:id:sgln:0037000.00729.8202", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	aggregationEvents = append(aggregationEvents, aggregationEvent)

	form := epcisproto.GetAggregationEventsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	aggregationEventsResponse := epcisproto.GetAggregationEventsResponse{AggregationEvents: aggregationEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *epcisproto.GetAggregationEventsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetAggregationEventsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &aggregationEventsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		aggregationEventResp, err := tt.es.GetAggregationEvents(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetAggregationEvents() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(aggregationEventResp, tt.want) {
			t.Errorf("EpcisService.GetAggregationEvents() = %v, want %v", aggregationEventResp, tt.want)
		}
		assert.NotNil(t, aggregationEventResp)
		aggregationEventResult := aggregationEventResp.AggregationEvents[0]
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventId, "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Action, "OBSERVE", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.ReadPoint, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizLocation, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
	}
}

func TestEpcisService_GetAggregationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	aggregationEvent, err := GetAggregationEvent(uint32(1), []byte{233, 210, 57, 150, 113, 115, 73, 53, 157, 104, 230, 4, 187, 159, 48, 158}, "e9d23996-7173-4935-9d68-e604bb9f309e", "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0", "+02:00", "https://accreditation-council.example.org/certificate/ABC12345", "2013-06-08T14:58:56Z", "incorrect_data", "2013-11-07T14:00:00Z", "urn:epc:id:sscc:0614141.1234567890", "OBSERVE", "receiving", "in_progress", "urn:epc:id:sgln:0037000.00729.8202", "urn:epc:id:sgln:0037000.00729.8202", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	aggregationEventResponse := epcisproto.GetAggregationEventResponse{}
	aggregationEventResponse.AggregationEvent = aggregationEvent

	form := epcisproto.GetAggregationEventRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "e9d23996-7173-4935-9d68-e604bb9f309e"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *epcisproto.GetAggregationEventRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetAggregationEventResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &aggregationEventResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		aggregationEventResp, err := tt.es.GetAggregationEvent(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetAggregationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(aggregationEventResp, tt.want) {
			t.Errorf("EpcisService.GetAggregationEvent() = %v, want %v", aggregationEventResp, tt.want)
		}
		assert.NotNil(t, aggregationEventResp)
		aggregationEventResult := aggregationEventResp.AggregationEvent
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventId, "ni:///sha-256;cd834b5a08e76778617369c29c9ecc1007508a0ae5dcf063e48b6bf05eb10097?ver=CBV2.0", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.EventTimeZoneOffset, "+02:00", "they should be equal")
		assert.Equal(t, aggregationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Action, "OBSERVE", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizStep, "receiving", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.ReadPoint, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
		assert.Equal(t, aggregationEventResult.AggregationEventD.BizLocation, "urn:epc:id:sgln:0037000.00729.8202", "they should be equal")
	}
}

func GetAggregationEvent(id uint32, uuid4 []byte, idS string, eventId string, eventTimeZoneOffset string, certification string, eventTime string, reason string, declarationTime string, parentId string, action string, bizStep string, disposition string, readPoint string, bizLocation string, statusCode string, createdByUserId string, updatedByUserId string) (*epcisproto.AggregationEvent, error) {
	eventTime1, err := common.ConvertTimeToTimestamp(Layout, eventTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	declarationTime1, err := common.ConvertTimeToTimestamp(Layout, declarationTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	epcisEventD := commonproto.EpcisEventD{}
	epcisEventD.EventId = eventId
	epcisEventD.EventTimeZoneOffset = eventTimeZoneOffset
	epcisEventD.Certification = certification

	epcisEventT := commonproto.EpcisEventT{}
	epcisEventT.EventTime = eventTime1

	errorDeclarationD := commonproto.ErrorDeclarationD{}
	errorDeclarationD.Reason = reason

	errorDeclarationT := commonproto.ErrorDeclarationT{}
	errorDeclarationT.DeclarationTime = declarationTime1

	aggregationEventD := epcisproto.AggregationEventD{}
	aggregationEventD.Id = id
	aggregationEventD.Uuid4 = uuid4
	aggregationEventD.IdS = idS
	aggregationEventD.ParentId = parentId
	aggregationEventD.Action = action
	aggregationEventD.BizStep = bizStep
	aggregationEventD.Disposition = disposition
	aggregationEventD.ReadPoint = readPoint
	aggregationEventD.BizLocation = bizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = statusCode
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	aggregationEvent := epcisproto.AggregationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, AggregationEventD: &aggregationEventD, CrUpdUser: &crUpdUser}

	return &aggregationEvent, nil
}
