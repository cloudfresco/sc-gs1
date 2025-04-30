package invoiceservices

import (
	"context"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDebitCreditAdviceLineItemDetailSQL = `insert into debit_credit_advice_line_item_details
	  (
  aligned_price,
  ap_code_list_version,
  ap_currency_code,
  invoiced_price,
  ip_code_list_version,
  ip_currency_code,
  quantity,
  q_measurement_unit_code,
  q_code_list_version,
  debit_credit_advice_id,
  debit_credit_advice_line_item_id)
      values(
    :aligned_price,
    :ap_code_list_version,
    :ap_currency_code,
    :invoiced_price,
    :ip_code_list_version,
    :ip_currency_code,
    :quantity,
    :q_measurement_unit_code,
    :q_code_list_version,
    :debit_credit_advice_id,
    :debit_credit_advice_line_item_id);`

/*const selectDebitCreditAdviceLineItemDetailsSQL = `select
  id,
  aligned_price,
  ap_code_list_version,
  ap_currency_code,
  invoiced_price,
  ip_code_list_version,
  ip_currency_code,
  quantity,
  q_measurement_unit_code,
  q_code_list_version,
  debit_credit_advice_id,
  debit_credit_advice_line_item_id from debit_credit_advice_line_item_details`*/

// CreateDebitCreditAdviceLineItemDetail - Create DebitCreditAdviceLineItemDetail
func (ds *DebitCreditAdviceService) CreateDebitCreditAdviceLineItemDetail(ctx context.Context, in *invoiceproto.CreateDebitCreditAdviceLineItemDetailRequest) (*invoiceproto.CreateDebitCreditAdviceLineItemDetailResponse, error) {
	debitCreditAdviceLineItemDetail, err := ds.ProcessDebitCreditAdviceLineItemDetailRequest(ctx, in)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ds.insertDebitCreditAdviceLineItemDetail(ctx, insertDebitCreditAdviceLineItemDetailSQL, debitCreditAdviceLineItemDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceLineItemDetailResponse := invoiceproto.CreateDebitCreditAdviceLineItemDetailResponse{}
	debitCreditAdviceLineItemDetailResponse.DebitCreditAdviceLineItemDetail = debitCreditAdviceLineItemDetail
	return &debitCreditAdviceLineItemDetailResponse, nil
}

// ProcessDebitCreditAdviceLineItemDetailRequest - ProcessDebitCreditAdviceLineItemDetailRequest
func (ds *DebitCreditAdviceService) ProcessDebitCreditAdviceLineItemDetailRequest(ctx context.Context, in *invoiceproto.CreateDebitCreditAdviceLineItemDetailRequest) (*invoiceproto.DebitCreditAdviceLineItemDetail, error) {
	debitCreditAdviceLineItemDetail := invoiceproto.DebitCreditAdviceLineItemDetail{}
	debitCreditAdviceLineItemDetail.AlignedPrice = in.AlignedPrice
	debitCreditAdviceLineItemDetail.ApCodeListVersion = in.ApCodeListVersion
	debitCreditAdviceLineItemDetail.ApCurrencyCode = in.ApCurrencyCode
	debitCreditAdviceLineItemDetail.InvoicedPrice = in.InvoicedPrice
	debitCreditAdviceLineItemDetail.IpCodeListVersion = in.IpCodeListVersion
	debitCreditAdviceLineItemDetail.IpCurrencyCode = in.IpCurrencyCode
	debitCreditAdviceLineItemDetail.Quantity = in.Quantity
	debitCreditAdviceLineItemDetail.QMeasurementUnitCode = in.QMeasurementUnitCode
	debitCreditAdviceLineItemDetail.QCodeListVersion = in.QCodeListVersion
	debitCreditAdviceLineItemDetail.DebitCreditAdviceId = in.DebitCreditAdviceId
	debitCreditAdviceLineItemDetail.DebitCreditAdviceLineItemId = in.DebitCreditAdviceLineItemId
	return &debitCreditAdviceLineItemDetail, nil
}

// insertDebitCreditAdviceLineItemDetail - Insert DebitCreditAdviceLineItemDetail details into database
func (ds *DebitCreditAdviceService) insertDebitCreditAdviceLineItemDetail(ctx context.Context, insertDebitCreditAdviceLineItemDetailSQL string, debitCreditAdviceLineItemDetail *invoiceproto.DebitCreditAdviceLineItemDetail, userEmail string, requestID string) error {
	err := ds.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDebitCreditAdviceLineItemDetailSQL, debitCreditAdviceLineItemDetail)
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
