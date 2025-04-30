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

func TestEpcisService_CreateAssociationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	associationEvent := epcisproto.CreateAssociationEventRequest{}
	associationEvent.EventId = "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0"
	associationEvent.EventTimeZoneOffset = "+01:00"
	associationEvent.Certification = "https://accreditation-council.example.org/certificate/ABC12345"
	associationEvent.EventTime = "04/05/2005"
	associationEvent.Reason = "incorrect_data"
	associationEvent.DeclarationTime = "07/11/2005"
	associationEvent.ParentId = "urn:epc:id:grai:4012345.55555.987"
	associationEvent.Action = "ADD"
	associationEvent.BizStep = "assembling"
	associationEvent.Disposition = "in_progress"
	associationEvent.ReadPoint = "urn:epc:id:sgln:4012345.00001.0"
	associationEvent.BizLocation = "urn:epc:id:sgln:0614141.00888.0"
	associationEvent.UserId = "auth0|673ee1a719dd4000cd5a3832"
	associationEvent.UserEmail = "sprov300@gmail.com"
	associationEvent.RequestId = "bks1m1g91jau4nkks2f0"

	persistentDisposition := epcisproto.CreatePersistentDispositionRequest{}
	persistentDisposition.SetDisp = "completeness_verified"
	persistentDisposition.UnsetDisp = "completeness_inferred"
	persistentDisposition.EventId = uint32(1)
	persistentDisposition.TypeOfEvent = "AssociationEvent"
	persistentDisposition.UserId = "auth0|673ee1a719dd4000cd5a3832"
	persistentDisposition.UserEmail = "sprov300@gmail.com"
	persistentDisposition.RequestId = "bks1m1g91jau4nkks2f0"
	persistentDispositions := []*epcisproto.CreatePersistentDispositionRequest{}
	persistentDispositions = append(persistentDispositions, &persistentDisposition)
	associationEvent.PersistentDispositions = persistentDispositions

	epc := epcisproto.CreateEpcRequest{}
	epc.EpcValue = "urn:epc:id:giai:4000001.12345"
	epc.EventId = uint32(1)
	epc.TypeOfEvent = "AssociationEvent"
	epc.TypeOfEpc = "output"
	epc.UserId = "auth0|673ee1a719dd4000cd5a3832"
	epc.UserEmail = "sprov300@gmail.com"
	epc.RequestId = "bks1m1g91jau4nkks2f0"

	epcs := []*epcisproto.CreateEpcRequest{}
	epcs = append(epcs, &epc)
	associationEvent.ChildEpcs = epcs

	bizTransaction := epcisproto.CreateBizTransactionRequest{}
	bizTransaction.BizTransactionType = "po"
	bizTransaction.BizTransaction = "urn:epc:id:gdti:0614141.00001.1618034"
	bizTransaction.EventId = uint32(1)
	bizTransaction.TypeOfEvent = "AssociationEvent"
	bizTransaction.UserId = "auth0|673ee1a719dd4000cd5a3832"
	bizTransaction.UserEmail = "sprov300@gmail.com"
	bizTransaction.RequestId = "bks1m1g91jau4nkks2f0"

	bizTransactions := []*epcisproto.CreateBizTransactionRequest{}
	bizTransactions = append(bizTransactions, &bizTransaction)
	associationEvent.BizTransactionList = bizTransactions

	quantityElement := epcisproto.CreateQuantityElementRequest{}
	quantityElement.EpcClass = "urn:epc:idpat:sgtin:4012345.098765"
	quantityElement.Quantity = float64(10)
	quantityElement.Uom = "KGM"
	quantityElement.EventId = uint32(1)
	quantityElement.TypeOfEvent = "AssociationEvent"
	quantityElement.TypeOfQuantity = "output"
	quantityElement.UserId = "auth0|673ee1a719dd4000cd5a3832"
	quantityElement.UserEmail = "sprov300@gmail.com"
	quantityElement.RequestId = "bks1m1g91jau4nkks2f0"

	quantityElements := []*epcisproto.CreateQuantityElementRequest{}
	quantityElements = append(quantityElements, &quantityElement)
	associationEvent.ChildQuantityList = quantityElements

	source := epcisproto.CreateSourceRequest{}
	source.SourceType = "location"
	source.Source = "urn:epc:id:sgln:4012345.00225.0"
	source.EventId = uint32(1)
	source.TypeOfEvent = "AssociationEvent"
	source.UserId = "auth0|673ee1a719dd4000cd5a3832"
	source.UserEmail = "sprov300@gmail.com"
	source.RequestId = "bks1m1g91jau4nkks2f0"
	sources := []*epcisproto.CreateSourceRequest{}
	sources = append(sources, &source)
	associationEvent.SourceList = sources

	destination := epcisproto.CreateDestinationRequest{}
	destination.DestType = "location"
	destination.Destination = "urn:epc:id:sgln:0614141.00777.0"
	destination.EventId = uint32(1)
	destination.TypeOfEvent = "AssociationEvent"
	destination.UserId = "auth0|673ee1a719dd4000cd5a3832"
	destination.UserEmail = "sprov300@gmail.com"
	destination.RequestId = "bks1m1g91jau4nkks2f0"

	destinations := []*epcisproto.CreateDestinationRequest{}
	destinations = append(destinations, &destination)
	associationEvent.DestinationList = destinations

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
	sensorElement.TypeOfEvent = "AssociationEvent"
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
	associationEvent.SensorElementList = sensorElements

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateAssociationEventRequest
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
				in:  &associationEvent,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		associationEventResp, err := tt.es.CreateAssociationEvent(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateAssociationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, associationEventResp)
		associationEventResult := associationEventResp.AssociationEvent
		assert.Equal(t, associationEventResult.EpcisEventD.EventId, "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.EventTimeZoneOffset, "+01:00", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizStep, "assembling", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetAssociationEvents(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)
	associationEvents := []*epcisproto.AssociationEvent{}

	associationEvent, err := GetAssociationEvent(uint32(1), []byte{53, 177, 235, 109, 144, 222, 71, 95, 151, 128, 79, 211, 169, 157, 131, 60}, "35b1eb6d-90de-475f-9780-4fd3a99d833c", "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0", "+01:00", "https://accreditation-council.example.org/certificate/ABC12345", "2019-11-01T13:00:00Z", "incorrect_data", "2019-11-07T14:00:00Z", "urn:epc:id:grai:4012345.55555.987", "ADD", "assembling", "in_progress", "urn:epc:id:sgln:4012345.00001.0", "urn:epc:id:sgln:0614141.00888.0", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	associationEvents = append(associationEvents, associationEvent)

	form := epcisproto.GetAssociationEventsRequest{}
	form.Limit = "2"
	form.NextCursor = ""
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"

	nextc := "MA=="
	associationEventsResponse := epcisproto.GetAssociationEventsResponse{AssociationEvents: associationEvents, NextCursor: nextc}

	type args struct {
		ctx context.Context
		in  *epcisproto.GetAssociationEventsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetAssociationEventsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &associationEventsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		associationEventResp, err := tt.es.GetAssociationEvents(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetAssociationEvents() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(associationEventResp, tt.want) {
			t.Errorf("EpcisService.GetAssociationEvents() = %v, want %v", associationEventResp, tt.want)
		}
		assert.NotNil(t, associationEventResp)
		associationEventResult := associationEventResp.AssociationEvents[0]
		assert.Equal(t, associationEventResult.EpcisEventD.EventId, "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.EventTimeZoneOffset, "+01:00", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizStep, "assembling", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func TestEpcisService_GetAssociationEvent(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	associationEvent, err := GetAssociationEvent(uint32(1), []byte{53, 177, 235, 109, 144, 222, 71, 95, 151, 128, 79, 211, 169, 157, 131, 60}, "35b1eb6d-90de-475f-9780-4fd3a99d833c", "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0", "+01:00", "https://accreditation-council.example.org/certificate/ABC12345", "2019-11-01T13:00:00Z", "incorrect_data", "2019-11-07T14:00:00Z", "urn:epc:id:grai:4012345.55555.987", "ADD", "assembling", "in_progress", "urn:epc:id:sgln:4012345.00001.0", "urn:epc:id:sgln:0614141.00888.0", "active", "auth0|673ee1a719dd4000cd5a3832", "auth0|673ee1a719dd4000cd5a3832")
	if err != nil {
		t.Error(err)
		return
	}

	associationEventResponse := epcisproto.GetAssociationEventResponse{}
	associationEventResponse.AssociationEvent = associationEvent

	form := epcisproto.GetAssociationEventRequest{}
	gform := commonproto.GetRequest{}
	gform.Id = "35b1eb6d-90de-475f-9780-4fd3a99d833c"
	gform.UserEmail = "sprov300@gmail.com"
	gform.RequestId = "bks1m1g91jau4nkks2f0"
	form.GetRequest = &gform

	type args struct {
		ctx   context.Context
		inReq *epcisproto.GetAssociationEventRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetAssociationEventResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx:   ctx,
				inReq: &form,
			},
			want:    &associationEventResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		associationEventResp, err := tt.es.GetAssociationEvent(tt.args.ctx, tt.args.inReq)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetAssociationEvent() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(associationEventResp, tt.want) {
			t.Errorf("EpcisService.GetAssociationEvent() = %v, want %v", associationEventResp, tt.want)
		}
		assert.NotNil(t, associationEventResp)
		associationEventResult := associationEventResp.AssociationEvent
		assert.Equal(t, associationEventResult.EpcisEventD.EventId, "ni:///sha-256;3785a2a509892681bb6695cfde36dd75aabdc5d801a028e7b17cded4e0fa320f?ver=CBV2.0", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.EventTimeZoneOffset, "+01:00", "they should be equal")
		assert.Equal(t, associationEventResult.EpcisEventD.Certification, "https://accreditation-council.example.org/certificate/ABC12345", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Action, "ADD", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizStep, "assembling", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.Disposition, "in_progress", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.ReadPoint, "urn:epc:id:sgln:4012345.00001.0", "they should be equal")
		assert.Equal(t, associationEventResult.AssociationEventD.BizLocation, "urn:epc:id:sgln:0614141.00888.0", "they should be equal")
	}
}

func GetAssociationEvent(id uint32, uuid4 []byte, idS string, eventId string, eventTimeZoneOffset string, certification string, eventTime string, reason string, declarationTime string, parentId string, action string, bizStep string, disposition string, readPoint string, bizLocation string, statusCode string, createdByUserId string, updatedByUserId string) (*epcisproto.AssociationEvent, error) {
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

	associationEventD := epcisproto.AssociationEventD{}
	associationEventD.Id = id
	associationEventD.Uuid4 = uuid4
	associationEventD.IdS = idS
	associationEventD.ParentId = parentId
	associationEventD.Action = action
	associationEventD.BizStep = bizStep
	associationEventD.Disposition = disposition
	associationEventD.ReadPoint = readPoint
	associationEventD.BizLocation = bizLocation

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = statusCode
	crUpdUser.CreatedByUserId = createdByUserId
	crUpdUser.UpdatedByUserId = updatedByUserId

	associationEvent := epcisproto.AssociationEvent{EpcisEventD: &epcisEventD, EpcisEventT: &epcisEventT, ErrorDeclarationD: &errorDeclarationD, ErrorDeclarationT: &errorDeclarationT, AssociationEventD: &associationEventD, CrUpdUser: &crUpdUser}

	return &associationEvent, nil
}
