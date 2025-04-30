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

func TestEpcisService_CreateSensorElement(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

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
	sensorElement.TypeOfEvent = "ObjectEvent"
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

	type args struct {
		ctx context.Context
		in  *epcisproto.CreateSensorElementRequest
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
				in:  &sensorElement,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sensorElementResp, err := tt.es.CreateSensorElement(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.CreateSensorElement() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		assert.NotNil(t, sensorElementResp)
		sensorElementResult := sensorElementResp.SensorElement
		assert.Equal(t, sensorElementResult.SensorMetadataD.DeviceId, "urn:epc:id:giai:4000001.111", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.DeviceMetadata, "https://id.gs1.org/8004/4000001111", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.RawData, "https://example.org/8004/401234599999", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.DataProcessingMethod, "https://example.com/253/4012345000054987", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.BizRules, "https://example.com/253/4012345000054987", "they should be equal")

		sensorReportResult := sensorElementResult.SensorReports[0]
		assert.Equal(t, sensorReportResult.SensorReportD.DeviceId, "urn:epc:id:giai:4000001.111", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.RawData, "https://example.org/8004/401234599999", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.DataProcessingMethod, "https://example.com/253/4012345000054987", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Microorganism, "https://www.ncbi.nlm.nih.gov/taxonomy/1126011", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.ChemicalSubstance, "https://identifiers.org/inchikey:CZMRCDWAGMRECN-UGDNZRGBSA-N", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.SensorValue, float64(26), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Component, "example:x", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.UriValue, "https://id.gs1.org/8004/4000001111", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MinValue, float64(26), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MaxValue, float64(26.2), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MeanValue, float64(13.2), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.PercRank, float64(50), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.PercValue, float64(12.7), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Uom, "CEL", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.SDev, float64(0.1), "they should be equal")
	}
}

func TestEpcisService_GetSensorElements(t *testing.T) {
	err := test.LoadSQL(logUser, dbService)
	if err != nil {
		t.Error(err)
		return
	}

	ctx := LoginUser()

	epcisService := NewEpcisService(log, dbService, redisService, userServiceClient)

	sensorElement1, err := GetSensorElement("urn:epc:id:giai:4000001.111", "https://id.gs1.org/8004/4000001111", "https://example.org/8004/401234599999", "https://example.com/253/4012345000054987", "https://example.com/253/4012345000054987", "2019-04-02T13:05:00Z", "2019-04-02T12:55:01Z", "2019-04-02T13:55:00Z", uint32(1), uint32(1), "ObjectEvent")
	if err != nil {
		t.Error(err)
	}

	sensorReport1, err := GetSensorReport("Temperature", "urn:epc:id:giai:4000001.111", "https://example.org/8004/401234599999", "https://example.com/253/4012345000054987", "https://www.ncbi.nlm.nih.gov/taxonomy/1126011", "https://identifiers.org/inchikey:CZMRCDWAGMRECN-UGDNZRGBSA-N", float64(26), "example:x", "SomeString", true, "f0f0f0", "https://id.gs1.org/8004/4000001111", float64(26), float64(26.2), float64(13.2), float64(50), float64(12.7), "CEL", float64(0.1), "https://id.gs1.org/8004/4000001111", uint32(1), "2019-07-19T13:00:00Z")
	if err != nil {
		t.Error(err)
	}

	sensorElement1.SensorReports = append(sensorElement1.SensorReports, sensorReport1)

	sensorElements := []*epcisproto.SensorElement{}
	sensorElements = append(sensorElements, sensorElement1)

	sensorElementsResponse := epcisproto.GetSensorElementsResponse{}
	sensorElementsResponse.SensorElements = sensorElements

	form := epcisproto.GetSensorElementsRequest{}
	form.EventId = uint32(1)
	form.TypeOfEvent = "ObjectEvent"
	form.UserEmail = "sprov300@gmail.com"
	form.RequestId = "bks1m1g91jau4nkks2f0"
	type args struct {
		ctx context.Context
		in  *epcisproto.GetSensorElementsRequest
	}
	tests := []struct {
		es      *EpcisService
		args    args
		want    *epcisproto.GetSensorElementsResponse
		wantErr bool
	}{
		{
			es: epcisService,
			args: args{
				ctx: ctx,
				in:  &form,
			},
			want:    &sensorElementsResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		sensorElementResponse, err := tt.es.GetSensorElements(tt.args.ctx, tt.args.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("EpcisService.GetSensorElements() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		if !reflect.DeepEqual(sensorElementResponse, tt.want) {
			t.Errorf("EpcisService.GetSensorElements() = %v, want %v", sensorElementResponse, tt.want)
		}
		sensorElementResult := sensorElementResponse.SensorElements[0]
		assert.NotNil(t, sensorElementResult)
		assert.Equal(t, sensorElementResult.SensorMetadataD.DeviceId, "urn:epc:id:giai:4000001.111", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.DeviceMetadata, "https://id.gs1.org/8004/4000001111", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.RawData, "https://example.org/8004/401234599999", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.DataProcessingMethod, "https://example.com/253/4012345000054987", "they should be equal")
		assert.Equal(t, sensorElementResult.SensorMetadataD.BizRules, "https://example.com/253/4012345000054987", "they should be equal")

		sensorReportResult := sensorElementResult.SensorReports[0]
		assert.Equal(t, sensorReportResult.SensorReportD.DeviceId, "urn:epc:id:giai:4000001.111", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.RawData, "https://example.org/8004/401234599999", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.DataProcessingMethod, "https://example.com/253/4012345000054987", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Microorganism, "https://www.ncbi.nlm.nih.gov/taxonomy/1126011", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.ChemicalSubstance, "https://identifiers.org/inchikey:CZMRCDWAGMRECN-UGDNZRGBSA-N", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.SensorValue, float64(26), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Component, "example:x", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.UriValue, "https://id.gs1.org/8004/4000001111", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MinValue, float64(26), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MaxValue, float64(26.2), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.MeanValue, float64(13.2), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.PercRank, float64(50), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.PercValue, float64(12.7), "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.Uom, "CEL", "they should be equal")
		assert.Equal(t, sensorReportResult.SensorReportD.SDev, float64(0.1), "they should be equal")
	}
}

func GetSensorElement(deviceId string, deviceMetadata string, rawData string, dataProcessingMethod string, bizRules string, sensorTime string, startTime string, endTime string, id uint32, eventId uint32, typeOfEvent string) (*epcisproto.SensorElement, error) {
	sensorTime1, err := common.ConvertTimeToTimestamp(Layout, sensorTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	startTime1, err := common.ConvertTimeToTimestamp(Layout, startTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	endTime1, err := common.ConvertTimeToTimestamp(Layout, endTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}

	sensorMetadataD := commonproto.SensorMetadataD{}
	sensorMetadataD.DeviceId = deviceId
	sensorMetadataD.DeviceMetadata = deviceMetadata
	sensorMetadataD.RawData = rawData
	sensorMetadataD.DataProcessingMethod = dataProcessingMethod
	sensorMetadataD.BizRules = bizRules

	sensorMetadataT := commonproto.SensorMetadataT{}
	sensorMetadataT.SensorTime = sensorTime1
	sensorMetadataT.StartTime = startTime1
	sensorMetadataT.EndTime = endTime1

	sensorElementD := epcisproto.SensorElementD{}
	sensorElementD.Id = id
	sensorElementD.EventId = eventId
	sensorElementD.TypeOfEvent = typeOfEvent

	sensorElement := epcisproto.SensorElement{SensorMetadataD: &sensorMetadataD, SensorMetadataT: &sensorMetadataT, SensorElementD: &sensorElementD}

	return &sensorElement, nil
}

func GetSensorReport(sensorReportType string, deviceId string, rawData string, dataProcessingMethod string, microorganism string, chemicalSubstance string, sensorValue float64, component string, stringValue string, booleanValue bool, hexBinaryValue string, uriValue string, minValue float64, maxValue float64, meanValue float64, percRank float64, percValue float64, uom string, sDev float64, deviceMetadata string, sensorElementId uint32, sensorReportTime string) (*epcisproto.SensorReport, error) {
	sensorReportTime1, err := common.ConvertTimeToTimestamp(Layout, sensorReportTime)
	if err != nil {
		log.Error("Error", zap.Error(err))
	}
	sensorReportD := epcisproto.SensorReportD{}
	sensorReportD.SensorReportType = sensorReportType
	sensorReportD.DeviceId = deviceId
	sensorReportD.RawData = rawData
	sensorReportD.DataProcessingMethod = dataProcessingMethod
	sensorReportD.Microorganism = microorganism
	sensorReportD.ChemicalSubstance = chemicalSubstance
	sensorReportD.SensorValue = sensorValue
	sensorReportD.Component = component
	sensorReportD.StringValue = stringValue
	sensorReportD.BooleanValue = booleanValue
	sensorReportD.HexBinaryValue = hexBinaryValue
	sensorReportD.UriValue = uriValue
	sensorReportD.MinValue = minValue
	sensorReportD.MaxValue = maxValue
	sensorReportD.MeanValue = meanValue
	sensorReportD.PercRank = percRank
	sensorReportD.PercValue = percValue
	sensorReportD.Uom = uom
	sensorReportD.SDev = sDev
	sensorReportD.DeviceMetadata = deviceMetadata
	sensorReportD.SensorElementId = sensorElementId

	sensorReportT := epcisproto.SensorReportT{}
	sensorReportT.SensorReportTime = sensorReportTime1

	sensorReport := epcisproto.SensorReport{SensorReportD: &sensorReportD, SensorReportT: &sensorReportT}
	return &sensorReport, nil
}
