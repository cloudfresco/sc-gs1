package logisticsservices

import (
	"context"

	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDespatchAdviceItemTotalSQL = `insert into despatch_advice_item_totals
	  (additional_trade_item_identification,
     additional_trade_item_identification_type_code,
     code_list_version,
     gtin,
     trade_item_identification,
     despatch_advice_id,
     despatch_advice_line_item_id
)
  values(
  :additional_trade_item_identification,
  :additional_trade_item_identification_type_code,
  :code_list_version,
  :gtin,
  :trade_item_identification,
  :despatch_advice_id,
  :despatch_advice_line_item_id
 );`

/*const selectDespatchAdviceItemTotalsSQL = `select
  id,
  additional_trade_item_identification,
  additional_trade_item_identification_type_code,
  code_list_version,
  gtin,
  trade_item_identification,
  despatch_advice_id,
  despatch_advice_line_item_id
  from despatch_advice_item_totals`*/

func (das *DespatchAdviceService) CreateDespatchAdviceItemTotal(ctx context.Context, in *logisticsproto.CreateDespatchAdviceItemTotalRequest) (*logisticsproto.CreateDespatchAdviceItemTotalResponse, error) {
	despatchAdviceItemTotal, err := das.ProcessDespatchAdviceItemTotalRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceItemTotal(ctx, insertDespatchAdviceItemTotalSQL, despatchAdviceItemTotal, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceItemTotalResponse := logisticsproto.CreateDespatchAdviceItemTotalResponse{}
	despatchAdviceItemTotalResponse.DespatchAdviceItemTotal = despatchAdviceItemTotal
	return &despatchAdviceItemTotalResponse, nil
}

// ProcessDespatchAdviceItemTotalRequest - ProcessDespatchAdviceItemTotalRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceItemTotalRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceItemTotalRequest) (*logisticsproto.DespatchAdviceItemTotal, error) {
	despatchAdviceItemTotal := logisticsproto.DespatchAdviceItemTotal{}
	despatchAdviceItemTotal.AdditionalTradeItemIdentification = in.AdditionalTradeItemIdentification
	despatchAdviceItemTotal.AdditionalTradeItemIdentificationTypeCode = in.AdditionalTradeItemIdentificationTypeCode
	despatchAdviceItemTotal.CodeListVersion = in.CodeListVersion
	despatchAdviceItemTotal.Gtin = in.Gtin
	despatchAdviceItemTotal.TradeItemIdentification = in.TradeItemIdentification
	despatchAdviceItemTotal.DespatchAdviceId = in.DespatchAdviceId
	despatchAdviceItemTotal.DespatchAdviceLineItemId = in.DespatchAdviceLineItemId
	return &despatchAdviceItemTotal, nil
}

// insertDespatchAdviceItemTotal - Insert DespatchAdviceItemTotal into database
func (das *DespatchAdviceService) insertDespatchAdviceItemTotal(ctx context.Context, insertDespatchAdviceItemTotalSQL string, despatchAdviceItemTotal *logisticsproto.DespatchAdviceItemTotal, userEmail string, requestID string) error {
	err := das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDespatchAdviceItemTotalSQL, despatchAdviceItemTotal)
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
