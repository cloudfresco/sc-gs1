package epcisservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	epcisstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertSensorElementSQL = `insert into sensor_elements
  (device_id,
  device_metadata,
  raw_data,
  data_processing_method,
  biz_rules,
  sensor_time,
  start_time,
  end_time,
  event_id,
  type_of_event)
    values(
  :device_id,
  :device_metadata,
  :raw_data,
  :data_processing_method,
  :biz_rules,
  :sensor_time,
  :start_time,
  :end_time,
  :event_id,
  :type_of_event);`

const selectSensorElementsSQL = `select
  id,
  device_id,
  device_metadata,
  raw_data,
  data_processing_method,
  biz_rules,
  sensor_time,
  start_time,
  end_time,
  event_id,
  type_of_event from sensor_elements`

const insertSensorReportSQL = `insert into sensor_reports
  (sensor_report_type,
  device_id,
  raw_data,
  data_processing_method,
  microorganism,
  chemical_substance,
  sensor_value,
  component,
  string_value,
  boolean_value,
  hex_binary_value,
  uri_value,
  min_value,
  max_value,
  mean_value,
  perc_rank,
  perc_value,
  uom,
  s_dev,
  device_metadata,
  sensor_element_id,
  sensor_report_time)
    values(
  :sensor_report_type,
  :device_id,
  :raw_data,
  :data_processing_method,
  :microorganism,
  :chemical_substance,
  :sensor_value,
  :component,
  :string_value,
  :boolean_value,
  :hex_binary_value,
  :uri_value,
  :min_value,
  :max_value,
  :mean_value,
  :perc_rank,
  :perc_value,
  :uom,
  :s_dev,
  :device_metadata,
  :sensor_element_id,
  :sensor_report_time);`

const selectSensorReportsSQL = `select
  sensor_report_type,
  device_id,
  raw_data,
  data_processing_method,
  microorganism,
  chemical_substance,
  sensor_value,
  component,
  string_value,
  boolean_value,
  hex_binary_value,
  uri_value,
  min_value,
  max_value,
  mean_value,
  perc_rank,
  perc_value,
  uom,
  s_dev,
  device_metadata,
  sensor_element_id,
  sensor_report_time from sensor_reports`

// CreateSensorElement - Create SensorElement
func (es *EpcisService) CreateSensorElement(ctx context.Context, in *epcisproto.CreateSensorElementRequest) (*epcisproto.CreateSensorElementResponse, error) {
	sensorElement, err := es.ProcessSensorElementRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertSensorElement(ctx, insertSensorElementSQL, insertSensorReportSQL, sensorElement, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sensorElementResponse := epcisproto.CreateSensorElementResponse{}
	sensorElementResponse.SensorElement = sensorElement
	return &sensorElementResponse, nil
}

// ProcessSensorElementRequest - ProcessSensorElementRequest
func (es *EpcisService) ProcessSensorElementRequest(ctx context.Context, in *epcisproto.CreateSensorElementRequest) (*epcisproto.SensorElement, error) {
	sensorTime, err := time.Parse(common.Layout, in.SensorTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	startTime, err := time.Parse(common.Layout, in.StartTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	endTime, err := time.Parse(common.Layout, in.EndTime)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sensorMetadataD := commonproto.SensorMetadataD{}
	sensorMetadataD.DeviceId = in.DeviceId
	sensorMetadataD.DeviceMetadata = in.DeviceMetadata
	sensorMetadataD.RawData = in.RawData
	sensorMetadataD.DataProcessingMethod = in.DataProcessingMethod
	sensorMetadataD.BizRules = in.BizRules

	sensorMetadataT := commonproto.SensorMetadataT{}
	sensorMetadataT.SensorTime = common.TimeToTimestamp(sensorTime.UTC().Truncate(time.Second))
	sensorMetadataT.StartTime = common.TimeToTimestamp(startTime.UTC().Truncate(time.Second))
	sensorMetadataT.EndTime = common.TimeToTimestamp(endTime.UTC().Truncate(time.Second))

	sensorElementD := epcisproto.SensorElementD{}
	sensorElementD.EventId = in.EventId
	sensorElementD.TypeOfEvent = in.TypeOfEvent

	sensorElement := epcisproto.SensorElement{SensorMetadataD: &sensorMetadataD, SensorMetadataT: &sensorMetadataT, SensorElementD: &sensorElementD}

	sensorReports := []*epcisproto.SensorReport{}
	for _, sReport := range in.SensorReports {
		sensorReportTime, err := time.Parse(common.Layout, sReport.SensorReportTime)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		sensorReportD := epcisproto.SensorReportD{}
		sensorReportD.SensorReportType = sReport.SensorReportType
		sensorReportD.DeviceId = sReport.DeviceId
		sensorReportD.RawData = sReport.RawData
		sensorReportD.DataProcessingMethod = sReport.DataProcessingMethod
		sensorReportD.Microorganism = sReport.Microorganism
		sensorReportD.ChemicalSubstance = sReport.ChemicalSubstance
		sensorReportD.SensorValue = sReport.SensorValue
		sensorReportD.Component = sReport.Component
		sensorReportD.StringValue = sReport.StringValue
		sensorReportD.BooleanValue = sReport.BooleanValue
		sensorReportD.HexBinaryValue = sReport.HexBinaryValue
		sensorReportD.UriValue = sReport.UriValue
		sensorReportD.MinValue = sReport.MinValue
		sensorReportD.MaxValue = sReport.MaxValue
		sensorReportD.MeanValue = sReport.MeanValue
		sensorReportD.PercRank = sReport.PercRank
		sensorReportD.PercValue = sReport.PercValue
		sensorReportD.Uom = sReport.Uom
		sensorReportD.SDev = sReport.SDev
		sensorReportD.DeviceMetadata = sReport.DeviceMetadata
		sensorReportD.SensorElementId = sReport.SensorElementId

		sensorReportT := epcisproto.SensorReportT{}
		sensorReportT.SensorReportTime = common.TimeToTimestamp(sensorReportTime.UTC().Truncate(time.Second))

		sensorReport := epcisproto.SensorReport{SensorReportD: &sensorReportD, SensorReportT: &sensorReportT}
		sensorReports = append(sensorReports, &sensorReport)
	}
	sensorElement.SensorReports = sensorReports

	return &sensorElement, nil
}

// insertSensorElement - Insert SensorElement details into database
func (es *EpcisService) insertSensorElement(ctx context.Context, insertSensorElementSQL string, insertSensorReportSQL string, sensorElement *epcisproto.SensorElement, userEmail string, requestID string) error {
	sensorElementTmp, err := es.crSensorElementStruct(ctx, sensorElement, userEmail, requestID)
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertSensorElementSQL, sensorElementTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		sensorElement.SensorElementD.Id = uint32(uID)

		for _, sensorReport := range sensorElement.SensorReports {
			sensorReport.SensorReportD.SensorElementId = sensorElement.SensorElementD.Id

			sensorReportTmp, err := es.crSensorReportStruct(ctx, sensorReport, userEmail, requestID)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
			_, err = tx.NamedExecContext(ctx, insertSensorReportSQL, sensorReportTmp)
			if err != nil {
				es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
				return err
			}
		}
		return nil
	})

	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crSensorElementStruct - process Sensor Element
func (es *EpcisService) crSensorElementStruct(ctx context.Context, sensorElement *epcisproto.SensorElement, userEmail string, requestID string) (*epcisstruct.SensorElement, error) {
	sensorMetadataT := new(commonstruct.SensorMetadataT)
	sensorMetadataT.SensorTime = common.TimestampToTime(sensorElement.SensorMetadataT.SensorTime)
	sensorMetadataT.StartTime = common.TimestampToTime(sensorElement.SensorMetadataT.StartTime)
	sensorMetadataT.EndTime = common.TimestampToTime(sensorElement.SensorMetadataT.EndTime)

	sensorElementTmp := epcisstruct.SensorElement{SensorMetadataD: sensorElement.SensorMetadataD, SensorMetadataT: sensorMetadataT, SensorElementD: sensorElement.SensorElementD}

	return &sensorElementTmp, nil
}

// crSensorReportStruct - process Sensor Report
func (es *EpcisService) crSensorReportStruct(ctx context.Context, sensorReport *epcisproto.SensorReport, userEmail string, requestID string) (*epcisstruct.SensorReport, error) {
	sensorReportT := new(epcisstruct.SensorReportT)
	sensorReportT.SensorReportTime = common.TimestampToTime(sensorReport.SensorReportT.SensorReportTime)

	sensorReportTmp := epcisstruct.SensorReport{SensorReportD: sensorReport.SensorReportD, SensorReportT: sensorReportT}

	return &sensorReportTmp, nil
}

// GetSensorElements - Get SensorElements
func (es *EpcisService) GetSensorElements(ctx context.Context, in *epcisproto.GetSensorElementsRequest) (*epcisproto.GetSensorElementsResponse, error) {
	query := "event_id = ? and type_of_event = ?"

	sensorElements := []*epcisproto.SensorElement{}

	nselectSensorElementsSQL := selectSensorElementsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectSensorElementsSQL, in.EventId, in.TypeOfEvent)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		sensorElementTmp := epcisstruct.SensorElement{}
		err = rows.StructScan(&sensorElementTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		getRequest := commonproto.GetRequest{}
		getRequest.UserEmail = in.UserEmail
		getRequest.RequestId = in.RequestId
		sensorElement, err := es.getSensorElementStruct(ctx, &getRequest, sensorElementTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}
		sensorReports, err := es.getSensorReports(ctx, &getRequest, sensorElement.SensorElementD.Id)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		sensorElement.SensorReports = sensorReports

		sensorElements = append(sensorElements, sensorElement)
	}

	sensorElementsResponse := epcisproto.GetSensorElementsResponse{SensorElements: sensorElements}

	return &sensorElementsResponse, nil
}

// getSensorElementStruct - Get sensorElement
func (es *EpcisService) getSensorElementStruct(ctx context.Context, in *commonproto.GetRequest, sensorElementTmp epcisstruct.SensorElement) (*epcisproto.SensorElement, error) {
	sensorMetadataT := new(commonproto.SensorMetadataT)
	sensorMetadataT.SensorTime = common.TimeToTimestamp(sensorElementTmp.SensorMetadataT.SensorTime)
	sensorMetadataT.StartTime = common.TimeToTimestamp(sensorElementTmp.SensorMetadataT.StartTime)
	sensorMetadataT.EndTime = common.TimeToTimestamp(sensorElementTmp.SensorMetadataT.EndTime)

	sensorElement := epcisproto.SensorElement{SensorMetadataD: sensorElementTmp.SensorMetadataD, SensorMetadataT: sensorMetadataT, SensorElementD: sensorElementTmp.SensorElementD}

	return &sensorElement, nil
}

// getSensorReports - Get SensorReports
func (es *EpcisService) getSensorReports(ctx context.Context, in *commonproto.GetRequest, sensorElementId uint32) ([]*epcisproto.SensorReport, error) {
	query := "sensor_element_id = ?"

	sensorReports := []*epcisproto.SensorReport{}

	nselectSensorReportsSQL := selectSensorReportsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectSensorReportsSQL, sensorElementId)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		sensorReportTmp := epcisstruct.SensorReport{}
		err = rows.StructScan(&sensorReportTmp)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		sensorReportT := new(epcisproto.SensorReportT)
		sensorReportT.SensorReportTime = common.TimeToTimestamp(sensorReportTmp.SensorReportT.SensorReportTime)

		sensorReport := epcisproto.SensorReport{SensorReportD: sensorReportTmp.SensorReportD, SensorReportT: sensorReportT}

		sensorReports = append(sensorReports, &sensorReport)
	}

	return sensorReports, nil
}
