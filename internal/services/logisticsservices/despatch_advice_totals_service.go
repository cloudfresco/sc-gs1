package logisticsservices

import (
	"context"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDespatchAdviceTotalSQL = `insert into despatch_advice_totals
	  (
measurement_type,
measurement_value,
package_total,
despatch_advice_id
)
  values(
  :measurement_type,
  :measurement_value,
  :package_total,
  :despatch_advice_id
 );`

/*const selectDespatchAdviceTotalsSQL = `select
  id,
  measurement_type,
  measurement_value,
  package_total,
  despatch_advice_id from despatch_advice_totals`*/

func (das *DespatchAdviceService) CreateDespatchAdviceTotal(ctx context.Context, in *logisticsproto.CreateDespatchAdviceTotalRequest) (*logisticsproto.CreateDespatchAdviceTotalResponse, error) {
	despatchAdviceTotal, err := das.ProcessDespatchAdviceTotalRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceTotal(ctx, insertDespatchAdviceTotalSQL, despatchAdviceTotal, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceTotalResponse := logisticsproto.CreateDespatchAdviceTotalResponse{}
	despatchAdviceTotalResponse.DespatchAdviceTotal = despatchAdviceTotal
	return &despatchAdviceTotalResponse, nil
}

// ProcessDespatchAdviceTotalRequest - ProcessDespatchAdviceTotalRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceTotalRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceTotalRequest) (*logisticsproto.DespatchAdviceTotal, error) {
	despatchAdviceTotal := logisticsproto.DespatchAdviceTotal{}
	despatchAdviceTotal.MeasurementType = in.MeasurementType
	despatchAdviceTotal.MeasurementValue = in.MeasurementValue
	despatchAdviceTotal.PackageTotal = in.PackageTotal
	despatchAdviceTotal.DespatchAdviceId = in.DespatchAdviceId
	return &despatchAdviceTotal, nil
}

// insertDespatchAdviceTotal - Insert DespatchAdviceTotal into database
func (das *DespatchAdviceService) insertDespatchAdviceTotal(ctx context.Context, insertDespatchAdviceTotalSQL string, despatchAdviceTotal *logisticsproto.DespatchAdviceTotal, userEmail string, requestID string) error {
	err := das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDespatchAdviceTotalSQL, despatchAdviceTotal)
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
