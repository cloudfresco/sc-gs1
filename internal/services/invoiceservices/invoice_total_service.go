package invoiceservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"

	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	invoicestruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInvoiceTotalSQL = `insert into invoice_totals
	    (base_amount,
ba_code_list_version,
ba_currency_code,
prepaid_amount,
pa_code_list_version,
pa_currency_code,
tax_accounting_currency,
total_amount_invoice_allowances_charges,
taiac_code_list_version,
taiac_currency_code,
total_amount_line_allowances_charges,
talac_code_list_version,
talac_currency_code,
total_economic_value,
tev_code_list_version,
tev_currency_code,
total_goods_value,
tgv_code_list_version,
tgv_currency_code,
total_invoice_amount,
tia_code_list_version,
tia_currency_code,
total_invoice_amount_payable,
tiap_code_list_version,
tiap_currency_code,
total_line_amount_exclusive_allowances_charges,
tlaeac_code_list_version,
tlaeac_currency_code,
total_line_amount_inclusive_allowances_charges,
tlaiac_code_list_version,
tlaiac_currency_code,
total_payment_discount_basis_amount,
tpdba_code_list_version,
tpdba_currency_code,
total_retail_value,
trv_code_list_version,
trv_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
total_tax_basis_amount,
ttba_code_list_version,
ttba_currency_code,
total_vat_amount,
tva_code_list_version,
tva_currency_code,
invoice_line_item_id,
invoice_id,
prepaid_amount_date)
  values(
:base_amount,
:ba_code_list_version,
:ba_currency_code,
:prepaid_amount,
:pa_code_list_version,
:pa_currency_code,
:tax_accounting_currency,
:total_amount_invoice_allowances_charges,
:taiac_code_list_version,
:taiac_currency_code,
:total_amount_line_allowances_charges,
:talac_code_list_version,
:talac_currency_code,
:total_economic_value,
:tev_code_list_version,
:tev_currency_code,
:total_goods_value,
:tgv_code_list_version,
:tgv_currency_code,
:total_invoice_amount,
:tia_code_list_version,
:tia_currency_code,
:total_invoice_amount_payable,
:tiap_code_list_version,
:tiap_currency_code,
:total_line_amount_exclusive_allowances_charges,
:tlaeac_code_list_version,
:tlaeac_currency_code,
:total_line_amount_inclusive_allowances_charges,
:tlaiac_code_list_version,
:tlaiac_currency_code,
:total_payment_discount_basis_amount,
:tpdba_code_list_version,
:tpdba_currency_code,
:total_retail_value,
:trv_code_list_version,
:trv_currency_code,
:total_tax_amount,
:tta_code_list_version,
:tta_currency_code,
:total_tax_basis_amount,
:ttba_code_list_version,
:ttba_currency_code,
:total_vat_amount,
:tva_code_list_version,
:tva_currency_code,
:invoice_line_item_id,
:invoice_id,
:prepaid_amount_date);`

/*const selectInvoiceTotalsSQL = `select
id,
base_amount,
ba_code_list_version,
ba_currency_code,
prepaid_amount,
pa_code_list_version,
pa_currency_code,
tax_accounting_currency,
total_amount_invoice_allowances_charges,
taiac_code_list_version,
taiac_currency_code,
total_amount_line_allowances_charges,
talac_code_list_version,
talac_currency_code,
total_economic_value,
tev_code_list_version,
tev_currency_code,
total_goods_value,
tgv_code_list_version,
tgv_currency_code,
total_invoice_amount,
tia_code_list_version,
tia_currency_code,
total_invoice_amount_payable,
tiap_code_list_version,
tiap_currency_code,
total_line_amount_exclusive_allowances_charges,
tlaeac_code_list_version,
tlaeac_currency_code,
total_line_amount_inclusive_allowances_charges,
code_list_version,
currency_code,
total_payment_discount_basis_amount,
tpdba_code_list_version,
tpdba_currency_code,
total_retail_value,
trv_code_list_version,
trv_currency_code,
total_tax_amount,
tta_code_list_version,
tta_currency_code,
total_tax_basis_amount,
ttba_code_list_version,
ttba_currency_code,
total_vat_amount,
tva_code_list_version,
tva_currency_code,
invoice_line_item_id,
invoice_id,
prepaid_amount_date from invoice_totals`*/

// CreateInvoiceTotal - Create InvoiceTotal
func (invs *InvoiceService) CreateInvoiceTotal(ctx context.Context, in *invoiceproto.CreateInvoiceTotalRequest) (*invoiceproto.CreateInvoiceTotalResponse, error) {
	invoiceTotal, err := invs.ProcessInvoiceTotalRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInvoiceTotal(ctx, insertInvoiceTotalSQL, invoiceTotal, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceTotalResponse := invoiceproto.CreateInvoiceTotalResponse{}
	invoiceTotalResponse.InvoiceTotal = invoiceTotal
	return &invoiceTotalResponse, nil
}

// ProcessInvoiceTotalRequest - ProcessInvoiceTotalRequest
func (invs *InvoiceService) ProcessInvoiceTotalRequest(ctx context.Context, in *invoiceproto.CreateInvoiceTotalRequest) (*invoiceproto.InvoiceTotal, error) {
	prepaidAmountDate, err := time.Parse(common.Layout, in.PrepaidAmountDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	invoiceTotalD := invoiceproto.InvoiceTotalD{}
	invoiceTotalD.BaseAmount = in.BaseAmount
	invoiceTotalD.BaCodeListVersion = in.BaCodeListVersion
	invoiceTotalD.BaCurrencyCode = in.BaCurrencyCode
	invoiceTotalD.PrepaidAmount = in.PrepaidAmount
	invoiceTotalD.PaCodeListVersion = in.PaCodeListVersion
	invoiceTotalD.PaCurrencyCode = in.PaCurrencyCode
	invoiceTotalD.TaxAccountingCurrency = in.TaxAccountingCurrency
	invoiceTotalD.TotalAmountInvoiceAllowancesCharges = in.TotalAmountInvoiceAllowancesCharges
	invoiceTotalD.TaiacCodeListVersion = in.TaiacCodeListVersion
	invoiceTotalD.TaiacCurrencyCode = in.TaiacCurrencyCode
	invoiceTotalD.TotalAmountLineAllowancesCharges = in.TotalAmountLineAllowancesCharges
	invoiceTotalD.TalacCodeListVersion = in.TalacCodeListVersion
	invoiceTotalD.TalacCurrencyCode = in.TalacCurrencyCode
	invoiceTotalD.TotalEconomicValue = in.TotalEconomicValue
	invoiceTotalD.TevCodeListVersion = in.TevCodeListVersion
	invoiceTotalD.TevCurrencyCode = in.TevCurrencyCode
	invoiceTotalD.TotalGoodsValue = in.TotalGoodsValue
	invoiceTotalD.TgvCodeListVersion = in.TgvCodeListVersion
	invoiceTotalD.TgvCurrencyCode = in.TgvCurrencyCode
	invoiceTotalD.TotalInvoiceAmount = in.TotalInvoiceAmount
	invoiceTotalD.TiaCodeListVersion = in.TiaCodeListVersion
	invoiceTotalD.TiaCurrencyCode = in.TiaCurrencyCode
	invoiceTotalD.TotalInvoiceAmountPayable = in.TotalInvoiceAmountPayable
	invoiceTotalD.TiapCodeListVersion = in.TiapCodeListVersion
	invoiceTotalD.TiapCurrencyCode = in.TiapCurrencyCode
	invoiceTotalD.TotalLineAmountExclusiveAllowancesCharges = in.TotalLineAmountExclusiveAllowancesCharges
	invoiceTotalD.TlaeacCodeListVersion = in.TlaeacCodeListVersion
	invoiceTotalD.TlaeacCurrencyCode = in.TlaeacCurrencyCode
	invoiceTotalD.TotalLineAmountInclusiveAllowancesCharges = in.TotalLineAmountInclusiveAllowancesCharges
	invoiceTotalD.TlaiacCodeListVersion = in.TlaiacCodeListVersion
	invoiceTotalD.TlaiacCurrencyCode = in.TlaiacCurrencyCode
	invoiceTotalD.TotalPaymentDiscountBasisAmount = in.TotalPaymentDiscountBasisAmount
	invoiceTotalD.TpdbaCodeListVersion = in.TpdbaCodeListVersion
	invoiceTotalD.TpdbaCurrencyCode = in.TpdbaCurrencyCode
	invoiceTotalD.TotalRetailValue = in.TotalRetailValue
	invoiceTotalD.TrvCodeListVersion = in.TrvCodeListVersion
	invoiceTotalD.TrvCurrencyCode = in.TrvCurrencyCode
	invoiceTotalD.TotalTaxAmount = in.TotalTaxAmount
	invoiceTotalD.TtaCodeListVersion = in.TtaCodeListVersion
	invoiceTotalD.TtaCurrencyCode = in.TtaCurrencyCode
	invoiceTotalD.TotalTaxBasisAmount = in.TotalTaxBasisAmount
	invoiceTotalD.TtbaCodeListVersion = in.TtbaCodeListVersion
	invoiceTotalD.TtbaCurrencyCode = in.TtbaCurrencyCode
	invoiceTotalD.TotalVATAmount = in.TotalVATAmount
	invoiceTotalD.TvaCodeListVersion = in.TvaCodeListVersion
	invoiceTotalD.TvaCurrencyCode = in.TvaCurrencyCode
	invoiceTotalD.InvoiceLineItemId = in.InvoiceLineItemId
	invoiceTotalD.InvoiceId = in.InvoiceId

	invoiceTotalT := invoiceproto.InvoiceTotalT{}
	invoiceTotalT.PrepaidAmountDate = common.TimeToTimestamp(prepaidAmountDate.UTC().Truncate(time.Second))

	invoiceTotal := invoiceproto.InvoiceTotal{InvoiceTotalD: &invoiceTotalD, InvoiceTotalT: &invoiceTotalT}

	return &invoiceTotal, nil
}

// insertInvoiceTotal - Insert InvoiceTotal details into database
func (invs *InvoiceService) insertInvoiceTotal(ctx context.Context, insertInvoiceTotalSQL string, invoiceTotal *invoiceproto.InvoiceTotal, userEmail string, requestID string) error {
	invoiceTotalTmp, err := invs.crInvoiceTotalStruct(ctx, invoiceTotal, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInvoiceTotalSQL, invoiceTotalTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		invoiceTotal.InvoiceTotalD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInvoiceTotalStruct - process InvoiceTotal details
func (invs *InvoiceService) crInvoiceTotalStruct(ctx context.Context, invoiceTotal *invoiceproto.InvoiceTotal, userEmail string, requestID string) (*invoicestruct.InvoiceTotal, error) {
	invoiceTotalT := new(invoicestruct.InvoiceTotalT)
	invoiceTotalT.PrepaidAmountDate = common.TimestampToTime(invoiceTotal.InvoiceTotalT.PrepaidAmountDate)

	invoiceTotalTmp := invoicestruct.InvoiceTotal{InvoiceTotalD: invoiceTotal.InvoiceTotalD, InvoiceTotalT: invoiceTotalT}

	return &invoiceTotalTmp, nil
}
