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

const insertDespatchAdviceQuantityVarianceSQL = `insert into despatch_advice_quantity_variances
	(
  remaining_quantity_status_code,
  variance_quantity,
  vq_measurement_unit_code,
  vq_code_list_version,
  variance_reason_code,
  despatch_advice_id,
  delivery_date_variance
  )
  values(
  :remaining_quantity_status_code,
  :variance_quantity,
  :vq_measurement_unit_code,
  :vq_code_list_version,
  :variance_reason_code,
  :despatch_advice_id,
  :delivery_date_variance
 );`

/*const selectDespatchAdviceQuantityVariancesSQL = `select
  id,
  remaining_quantity_status_code,
  variance_quantity,
  vq_measurement_unit_code,
  vq_code_list_version,
  variance_reason_code,
  despatch_advice_id,
  delivery_date_variance from despatch_advice_quantity_variances`*/

func (das *DespatchAdviceService) CreateDespatchAdviceQuantityVariance(ctx context.Context, in *logisticsproto.CreateDespatchAdviceQuantityVarianceRequest) (*logisticsproto.CreateDespatchAdviceQuantityVarianceResponse, error) {
	despatchAdviceQuantityVariance, err := das.ProcessDespatchAdviceQuantityVarianceRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceQuantityVariance(ctx, insertDespatchAdviceQuantityVarianceSQL, despatchAdviceQuantityVariance, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceQuantityVarianceResponse := logisticsproto.CreateDespatchAdviceQuantityVarianceResponse{}
	despatchAdviceQuantityVarianceResponse.DespatchAdviceQuantityVariance = despatchAdviceQuantityVariance
	return &despatchAdviceQuantityVarianceResponse, nil
}

// ProcessDespatchAdviceQuantityVarianceRequest - ProcessDespatchAdviceQuantityVarianceRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceQuantityVarianceRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceQuantityVarianceRequest) (*logisticsproto.DespatchAdviceQuantityVariance, error) {
	deliveryDateVariance, err := time.Parse(common.Layout, in.DeliveryDateVariance)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceQuantityVarianceD := logisticsproto.DespatchAdviceQuantityVarianceD{}
	despatchAdviceQuantityVarianceD.RemainingQuantityStatusCode = in.RemainingQuantityStatusCode
	despatchAdviceQuantityVarianceD.VarianceQuantity = in.VarianceQuantity
	despatchAdviceQuantityVarianceD.VqMeasurementUnitCode = in.VqMeasurementUnitCode
	despatchAdviceQuantityVarianceD.VqCodeListVersion = in.VqCodeListVersion
	despatchAdviceQuantityVarianceD.VarianceReasonCode = in.VarianceReasonCode
	despatchAdviceQuantityVarianceD.DespatchAdviceId = in.DespatchAdviceId

	despatchAdviceQuantityVarianceT := logisticsproto.DespatchAdviceQuantityVarianceT{}
	despatchAdviceQuantityVarianceT.DeliveryDateVariance = common.TimeToTimestamp(deliveryDateVariance.UTC().Truncate(time.Second))

	despatchAdviceQuantityVariance := logisticsproto.DespatchAdviceQuantityVariance{DespatchAdviceQuantityVarianceD: &despatchAdviceQuantityVarianceD, DespatchAdviceQuantityVarianceT: &despatchAdviceQuantityVarianceT}

	return &despatchAdviceQuantityVariance, nil
}

// insertDespatchAdviceQuantityVariance - Insert DespatchAdviceQuantityVariance into database
func (das *DespatchAdviceService) insertDespatchAdviceQuantityVariance(ctx context.Context, insertDespatchAdviceQuantityVarianceSQL string, despatchAdviceQuantityVariance *logisticsproto.DespatchAdviceQuantityVariance, userEmail string, requestID string) error {
	despatchAdviceQuantityVarianceTmp, err := das.crDespatchAdviceQuantityVarianceStruct(ctx, despatchAdviceQuantityVariance, userEmail, requestID)
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err = tx.NamedExecContext(ctx, insertDespatchAdviceQuantityVarianceSQL, despatchAdviceQuantityVarianceTmp)

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

// crDespatchAdviceQuantityVarianceStruct - process DespatchAdviceQuantityVariance details
func (das *DespatchAdviceService) crDespatchAdviceQuantityVarianceStruct(ctx context.Context, despatchAdviceQuantityVariance *logisticsproto.DespatchAdviceQuantityVariance, userEmail string, requestID string) (*logisticsstruct.DespatchAdviceQuantityVariance, error) {
	despatchAdviceQuantityVarianceT := new(logisticsstruct.DespatchAdviceQuantityVarianceT)
	despatchAdviceQuantityVarianceT.DeliveryDateVariance = common.TimestampToTime(despatchAdviceQuantityVariance.DespatchAdviceQuantityVarianceT.DeliveryDateVariance)

	despatchAdviceQuantityVarianceTmp := logisticsstruct.DespatchAdviceQuantityVariance{DespatchAdviceQuantityVarianceD: despatchAdviceQuantityVariance.DespatchAdviceQuantityVarianceD, DespatchAdviceQuantityVarianceT: despatchAdviceQuantityVarianceT}

	return &despatchAdviceQuantityVarianceTmp, nil
}
