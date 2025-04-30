package invoiceservices

import (
	"context"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInvoiceLineItemInformationAfterTaxesQL = `insert into invoice_line_item_information_after_taxes
	    (amount_exclusive_allowances_charges,
aeac_code_list_version,
aeac_currency_code,
amount_inclusive_allowances_charges,
aiac_code_list_version,
aiac_currency_code,
invoice_id,
invoice_line_item_id)
  values(
:amount_exclusive_allowances_charges,
:aeac_code_list_version,
:aeac_currency_code,
:amount_inclusive_allowances_charges,
:aiac_code_list_version,
:aiac_currency_code,
:invoice_id,
:invoice_line_item_id);`

/*const selectInvoiceLineItemInformationAfterTaxesSQL = `select
    id,
amount_exclusive_allowances_charges,
aeac_code_list_version,
aeac_currency_code,
amount_inclusive_allowances_charges,
aiac_code_list_version,
aiac_currency_code,
invoice_id,
invoice_line_item_id from invoice_line_item_information_after_taxes`*/

// CreateInvoiceLineItemInformationAfterTaxes - Create InvoiceLineItemInformationAfterTaxes
func (invs *InvoiceService) CreateInvoiceLineItemInformationAfterTaxes(ctx context.Context, in *invoiceproto.CreateInvoiceLineItemInformationAfterTaxesRequest) (*invoiceproto.CreateInvoiceLineItemInformationAfterTaxesResponse, error) {
	invoiceLineItemInformationAfterTax, err := invs.ProcessInvoiceLineItemInformationAfterTaxesRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInvoiceLineItemInformationAfterTaxes(ctx, insertInvoiceLineItemInformationAfterTaxesQL, invoiceLineItemInformationAfterTax, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	invoiceLineItemInformationAfterTaxResponse := invoiceproto.CreateInvoiceLineItemInformationAfterTaxesResponse{}
	invoiceLineItemInformationAfterTaxResponse.InvoiceLineItemInformationAfterTaxes = invoiceLineItemInformationAfterTax
	return &invoiceLineItemInformationAfterTaxResponse, nil
}

// ProcessInvoiceLineItemInformationAfterTaxesRequest - ProcessInvoiceLineItemInformationAfterTaxesRequest
func (invs *InvoiceService) ProcessInvoiceLineItemInformationAfterTaxesRequest(ctx context.Context, in *invoiceproto.CreateInvoiceLineItemInformationAfterTaxesRequest) (*invoiceproto.InvoiceLineItemInformationAfterTaxes, error) {
	invoiceLineItemInformationAfterTax := invoiceproto.InvoiceLineItemInformationAfterTaxes{}

	invoiceLineItemInformationAfterTax.AmountExclusiveAllowancesCharges = in.AmountExclusiveAllowancesCharges
	invoiceLineItemInformationAfterTax.AeacCodeListVersion = in.AeacCodeListVersion
	invoiceLineItemInformationAfterTax.AeacCurrencyCode = in.AeacCurrencyCode
	invoiceLineItemInformationAfterTax.AmountInclusiveAllowancesCharges = in.AmountInclusiveAllowancesCharges
	invoiceLineItemInformationAfterTax.AiacCodeListVersion = in.AiacCodeListVersion
	invoiceLineItemInformationAfterTax.AiacCurrencyCode = in.AiacCurrencyCode
	invoiceLineItemInformationAfterTax.InvoiceId = in.InvoiceId
	invoiceLineItemInformationAfterTax.InvoiceLineItemId = in.InvoiceLineItemId

	return &invoiceLineItemInformationAfterTax, nil
}

// insertInvoiceLineItemInformationAfterTaxes - Insert InvoiceLineItemInformationAfterTaxes details into database
func (invs *InvoiceService) insertInvoiceLineItemInformationAfterTaxes(ctx context.Context, insertInvoiceLineItemInformationAfterTaxesQL string, invoiceLineItemInformationAfterTax *invoiceproto.InvoiceLineItemInformationAfterTaxes, userEmail string, requestID string) error {
	err := invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInvoiceLineItemInformationAfterTaxesQL, invoiceLineItemInformationAfterTax)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoiceLineItemInformationAfterTax.Id = uint32(uID)

		return nil
	})
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
