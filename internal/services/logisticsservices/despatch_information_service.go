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

const insertDespatchInformationSQL = `insert into despatch_informations
	  (
	  despatch_advice_id,
	  actual_ship_date_time,
despatch_date_time,
estimated_delivery_date_time,
estimated_delivery_date_time_at_ultimate_consignee,
loading_date_time,
pick_up_date_time,
release_date_time_of_supplier,
estimated_delivery_period_begin,
estimated_delivery_period_end
)
  values(
  :despatch_advice_id,
  :actual_ship_date_time,
  :despatch_date_time,
  :estimated_delivery_date_time,
  :estimated_delivery_date_time_at_ultimate_consignee,
  :loading_date_time,
  :pick_up_date_time,
  :release_date_time_of_supplier,
  :estimated_delivery_period_begin,
  :estimated_delivery_period_end
 );`

/*const selectDespatchInformationsSQL = `select
  id,
  despatch_advice_id,
  actual_ship_date_time,
  despatch_date_time,
  estimated_delivery_date_time,
  estimated_delivery_date_time_at_ultimate_consignee,
  loading_date_time,
  pick_up_date_time,
  release_date_time_of_supplier,
  estimated_delivery_period_begin,
  estimated_delivery_period_end from despatch_informations`*/

func (das *DespatchAdviceService) CreateDespatchInformation(ctx context.Context, in *logisticsproto.CreateDespatchInformationRequest) (*logisticsproto.CreateDespatchInformationResponse, error) {
	despatchInformation, err := das.ProcessDespatchInformationRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchInformation(ctx, insertDespatchInformationSQL, despatchInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchInformationResponse := logisticsproto.CreateDespatchInformationResponse{}
	despatchInformationResponse.DespatchInformation = despatchInformation
	return &despatchInformationResponse, nil
}

// ProcessDespatchInformationRequest - ProcessDespatchInformationRequest
func (das *DespatchAdviceService) ProcessDespatchInformationRequest(ctx context.Context, in *logisticsproto.CreateDespatchInformationRequest) (*logisticsproto.DespatchInformation, error) {
	actualShipDateTime, err := time.Parse(common.Layout, in.ActualShipDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchDateTime, err := time.Parse(common.Layout, in.DespatchDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	estimatedDeliveryDateTime, err := time.Parse(common.Layout, in.EstimatedDeliveryDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	estimatedDeliveryDateTimeAtUltimateConsignee, err := time.Parse(common.Layout, in.EstimatedDeliveryDateTimeAtUltimateConsignee)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	loadingDateTime, err := time.Parse(common.Layout, in.LoadingDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pickUpDateTime, err := time.Parse(common.Layout, in.PickUpDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	releaseDateTimeOfSupplier, err := time.Parse(common.Layout, in.ReleaseDateTimeOfSupplier)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	estimatedDeliveryPeriodBegin, err := time.Parse(common.Layout, in.EstimatedDeliveryPeriodBegin)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	estimatedDeliveryPeriodEnd, err := time.Parse(common.Layout, in.EstimatedDeliveryPeriodEnd)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchInformationD := logisticsproto.DespatchInformationD{}
	despatchInformationD.DespatchAdviceId = in.DespatchAdviceId

	despatchInformationT := logisticsproto.DespatchInformationT{}
	despatchInformationT.ActualShipDateTime = common.TimeToTimestamp(actualShipDateTime.UTC().Truncate(time.Second))
	despatchInformationT.DespatchDateTime = common.TimeToTimestamp(despatchDateTime.UTC().Truncate(time.Second))
	despatchInformationT.EstimatedDeliveryDateTime = common.TimeToTimestamp(estimatedDeliveryDateTime.UTC().Truncate(time.Second))
	despatchInformationT.EstimatedDeliveryDateTimeAtUltimateConsignee = common.TimeToTimestamp(estimatedDeliveryDateTimeAtUltimateConsignee.UTC().Truncate(time.Second))
	despatchInformationT.LoadingDateTime = common.TimeToTimestamp(loadingDateTime.UTC().Truncate(time.Second))
	despatchInformationT.PickUpDateTime = common.TimeToTimestamp(pickUpDateTime.UTC().Truncate(time.Second))
	despatchInformationT.ReleaseDateTimeOfSupplier = common.TimeToTimestamp(releaseDateTimeOfSupplier.UTC().Truncate(time.Second))
	despatchInformationT.EstimatedDeliveryPeriodBegin = common.TimeToTimestamp(estimatedDeliveryPeriodBegin.UTC().Truncate(time.Second))
	despatchInformationT.EstimatedDeliveryPeriodEnd = common.TimeToTimestamp(estimatedDeliveryPeriodEnd.UTC().Truncate(time.Second))

	despatchInformation := logisticsproto.DespatchInformation{DespatchInformationD: &despatchInformationD, DespatchInformationT: &despatchInformationT}

	return &despatchInformation, nil
}

// insertDespatchInformation - Insert DespatchInformation into database
func (das *DespatchAdviceService) insertDespatchInformation(ctx context.Context, insertDespatchInformationSQL string, despatchInformation *logisticsproto.DespatchInformation, userEmail string, requestID string) error {
	despatchInformationTmp, err := das.crDespatchInformationStruct(ctx, despatchInformation, userEmail, requestID)
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err = tx.NamedExecContext(ctx, insertDespatchInformationSQL, despatchInformationTmp)

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

// crDespatchInformationStruct - process DespatchInformation details
func (das *DespatchAdviceService) crDespatchInformationStruct(ctx context.Context, despatchInformation *logisticsproto.DespatchInformation, userEmail string, requestID string) (*logisticsstruct.DespatchInformation, error) {
	despatchInformationT := new(logisticsstruct.DespatchInformationT)
	despatchInformationT.ActualShipDateTime = common.TimestampToTime(despatchInformation.DespatchInformationT.ActualShipDateTime)
	despatchInformationT.DespatchDateTime = common.TimestampToTime(despatchInformation.DespatchInformationT.DespatchDateTime)
	despatchInformationT.EstimatedDeliveryDateTime = common.TimestampToTime(despatchInformation.DespatchInformationT.EstimatedDeliveryDateTime)
	despatchInformationT.EstimatedDeliveryDateTimeAtUltimateConsignee = common.TimestampToTime(despatchInformation.DespatchInformationT.EstimatedDeliveryDateTimeAtUltimateConsignee)
	despatchInformationT.LoadingDateTime = common.TimestampToTime(despatchInformation.DespatchInformationT.LoadingDateTime)
	despatchInformationT.PickUpDateTime = common.TimestampToTime(despatchInformation.DespatchInformationT.PickUpDateTime)
	despatchInformationT.ReleaseDateTimeOfSupplier = common.TimestampToTime(despatchInformation.DespatchInformationT.ReleaseDateTimeOfSupplier)
	despatchInformationT.EstimatedDeliveryPeriodBegin = common.TimestampToTime(despatchInformation.DespatchInformationT.EstimatedDeliveryPeriodBegin)
	despatchInformationT.EstimatedDeliveryPeriodEnd = common.TimestampToTime(despatchInformation.DespatchInformationT.EstimatedDeliveryPeriodEnd)

	despatchInformationTmp := logisticsstruct.DespatchInformation{DespatchInformationD: despatchInformation.DespatchInformationD, DespatchInformationT: despatchInformationT}

	return &despatchInformationTmp, nil
}
