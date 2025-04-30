package logisticsservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	logisticsstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDespatchAdviceLogisticUnitSQL = `insert into despatch_advice_logistic_units
	  (
    additional_logisitic_unit_identification,
    additional_logistic_unit_identification_type_code,
    code_list_version,
    sscc,
    ultimate_consignee,
    despatch_advice_id,
    estimated_delivery_date_time_at_ultimate_consignee)
  values(
  :additional_logisitic_unit_identification,
  :additional_logistic_unit_identification_type_code,
  :code_list_version,
  :sscc,
  :ultimate_consignee,
  :despatch_advice_id,
  :estimated_delivery_date_time_at_ultimate_consignee);`

/*const selectDespatchAdviceLogisticUnitsSQL = `select
  id,
  additional_logisitic_unit_identification,
  additional_logistic_unit_identification_type_code,
  code_list_version,
  sscc,
  ultimate_consignee,
  despatch_advice_id,
  estimated_delivery_date_time_at_ultimate_consignee
  from despatch_advice_logistic_units`*/

func (das *DespatchAdviceService) CreateDespatchAdviceLogisticUnit(ctx context.Context, in *logisticsproto.CreateDespatchAdviceLogisticUnitRequest) (*logisticsproto.CreateDespatchAdviceLogisticUnitResponse, error) {
	despatchAdviceLogisticUnit, err := das.ProcessDespatchAdviceLogisticUnitRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceLogisticUnit(ctx, insertDespatchAdviceLogisticUnitSQL, despatchAdviceLogisticUnit, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceLogisticUnitResponse := logisticsproto.CreateDespatchAdviceLogisticUnitResponse{}
	despatchAdviceLogisticUnitResponse.DespatchAdviceLogisticUnit = despatchAdviceLogisticUnit
	return &despatchAdviceLogisticUnitResponse, nil
}

// ProcessDespatchAdviceLogisticUnitRequest - ProcessDespatchAdviceLogisticUnitRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceLogisticUnitRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceLogisticUnitRequest) (*logisticsproto.DespatchAdviceLogisticUnit, error) {
	estimatedDeliveryDateTimeAtUltimateConsignee, err := time.Parse(common.Layout, in.EstimatedDeliveryDateTimeAtUltimateConsignee)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceLogisticUnitD := logisticsproto.DespatchAdviceLogisticUnitD{}
	despatchAdviceLogisticUnitD.AdditionalLogisiticUnitIdentification = in.AdditionalLogisiticUnitIdentification
	despatchAdviceLogisticUnitD.AdditionalLogisticUnitIdentificationTypeCode = in.AdditionalLogisticUnitIdentificationTypeCode
	despatchAdviceLogisticUnitD.CodeListVersion = in.CodeListVersion
	despatchAdviceLogisticUnitD.Sscc = in.Sscc
	despatchAdviceLogisticUnitD.UltimateConsignee = in.UltimateConsignee
	despatchAdviceLogisticUnitD.DespatchAdviceId = in.DespatchAdviceId

	despatchAdviceLogisticUnitT := logisticsproto.DespatchAdviceLogisticUnitT{}
	despatchAdviceLogisticUnitT.EstimatedDeliveryDateTimeAtUltimateConsignee = common.TimeToTimestamp(estimatedDeliveryDateTimeAtUltimateConsignee.UTC().Truncate(time.Second))

	despatchAdviceLogisticUnit := logisticsproto.DespatchAdviceLogisticUnit{DespatchAdviceLogisticUnitD: &despatchAdviceLogisticUnitD, DespatchAdviceLogisticUnitT: &despatchAdviceLogisticUnitT}

	return &despatchAdviceLogisticUnit, nil
}

// insertDespatchAdviceLogisticUnit - Insert DespatchAdviceLogisticUnit into database
func (das *DespatchAdviceService) insertDespatchAdviceLogisticUnit(ctx context.Context, insertDespatchAdviceLogisticUnitSQL string, despatchAdviceLogisticUnit *logisticsproto.DespatchAdviceLogisticUnit, userEmail string, requestID string) error {
	despatchAdviceLogisticUnitTmp, err := das.crDespatchAdviceLogisticUnitStruct(ctx, despatchAdviceLogisticUnit, userEmail, requestID)
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err = tx.NamedExecContext(ctx, insertDespatchAdviceLogisticUnitSQL, despatchAdviceLogisticUnitTmp)

		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})

	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDespatchAdviceLogisticUnitStruct - process DespatchAdviceLogisticUnitC details
func (das *DespatchAdviceService) crDespatchAdviceLogisticUnitStruct(ctx context.Context, despatchAdviceLogisticUnit *logisticsproto.DespatchAdviceLogisticUnit, userEmail string, requestID string) (*logisticsstruct.DespatchAdviceLogisticUnit, error) {
	despatchAdviceLogisticUnitT := new(logisticsstruct.DespatchAdviceLogisticUnitT)
	despatchAdviceLogisticUnitT.EstimatedDeliveryDateTimeAtUltimateConsignee = common.TimestampToTime(despatchAdviceLogisticUnit.DespatchAdviceLogisticUnitT.EstimatedDeliveryDateTimeAtUltimateConsignee)

	despatchAdviceLogisticUnitTmp := logisticsstruct.DespatchAdviceLogisticUnit{DespatchAdviceLogisticUnitD: despatchAdviceLogisticUnit.DespatchAdviceLogisticUnitD, DespatchAdviceLogisticUnitT: despatchAdviceLogisticUnitT}

	return &despatchAdviceLogisticUnitTmp, nil
}
