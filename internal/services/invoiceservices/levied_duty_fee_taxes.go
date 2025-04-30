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

const insertLeviedDutyFeeTaxSQL = `insert into levied_duty_fee_taxes
	    (duty_fee_tax_accounting_currency,
duty_fee_tax_agency_name,
duty_fee_tax_amount,
dfta_code_list_version,
dfta_currency_code,
duty_fee_tax_amount_in_accounting_currency,
dftaiac_code_list_version,
dftaiac_currency_code,
duty_fee_tax_basis_amount,
dftba_code_list_version,
dftba_currency_code,
duty_fee_tax_basis_amount_in_accounting_currency,
dftbaiac_code_list_version,
dftbaiac_currency_code,
duty_fee_tax_category_code,
duty_fee_tax_description,
duty_fee_tax_exemption_description,
duty_fee_tax_exemption_reason,
duty_fee_tax_percentage,
duty_fee_tax_type_code,
extension,
order_line_item_id,
invoice_line_item_id,
invoice_id,
duty_fee_tax_point_date)
  values(
:duty_fee_tax_accounting_currency,
:duty_fee_tax_agency_name,
:duty_fee_tax_amount,
:dfta_code_list_version,
:dfta_currency_code,
:duty_fee_tax_amount_in_accounting_currency,
:dftaiac_code_list_version,
:dftaiac_currency_code,
:duty_fee_tax_basis_amount,
:dftba_code_list_version,
:dftba_currency_code,
:duty_fee_tax_basis_amount_in_accounting_currency,
:dftbaiac_code_list_version,
:dftbaiac_currency_code,
:duty_fee_tax_category_code,
:duty_fee_tax_description,
:duty_fee_tax_exemption_description,
:duty_fee_tax_exemption_reason,
:duty_fee_tax_percentage,
:duty_fee_tax_type_code,
:extension,
:order_line_item_id,
:invoice_line_item_id,
:invoice_id,
:duty_fee_tax_point_date);`

/*const selectLeviedDutyFeeTaxesSQL = `select
id,
duty_fee_tax_accounting_currency,
duty_fee_tax_agency_name,
duty_fee_tax_amount,
dfta_code_list_version,
dfta_currency_code,
duty_fee_tax_amount_in_accounting_currency,
dftaiac_code_list_version,
dftaiac_currency_code,
duty_fee_tax_basis_amount,
dftba_code_list_version,
dftba_currency_code,
duty_fee_tax_basis_amount_in_accounting_currency,
dftbaiac_code_list_version,
dftbaiac_currency_code,
duty_fee_tax_category_code,
duty_fee_tax_description,
duty_fee_tax_exemption_description,
duty_fee_tax_exemption_reason,
duty_fee_tax_percentage,
duty_fee_tax_type_code,
extension,
order_line_item_id,
invoice_line_item_id,
invoice_id,
duty_fee_tax_point_date from levied_duty_fee_taxes`*/

// CreateLeviedDutyFeeTax - Create LeviedDutyFeeTax
func (invs *InvoiceService) CreateLeviedDutyFeeTax(ctx context.Context, in *invoiceproto.CreateLeviedDutyFeeTaxRequest) (*invoiceproto.CreateLeviedDutyFeeTaxResponse, error) {
	leviedDutyFeeTax, err := invs.ProcessLeviedDutyFeeTaxRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLeviedDutyFeeTax(ctx, insertLeviedDutyFeeTaxSQL, leviedDutyFeeTax, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	leviedDutyFeeTaxResponse := invoiceproto.CreateLeviedDutyFeeTaxResponse{}
	leviedDutyFeeTaxResponse.LeviedDutyFeeTax = leviedDutyFeeTax
	return &leviedDutyFeeTaxResponse, nil
}

// ProcessLeviedDutyFeeTaxRequest - ProcessLeviedDutyFeeTaxRequest
func (invs *InvoiceService) ProcessLeviedDutyFeeTaxRequest(ctx context.Context, in *invoiceproto.CreateLeviedDutyFeeTaxRequest) (*invoiceproto.LeviedDutyFeeTax, error) {
	dutyFeeTaxPointDate, err := time.Parse(common.Layout, in.DutyFeeTaxPointDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	leviedDutyFeeTaxD := invoiceproto.LeviedDutyFeeTaxD{}
	leviedDutyFeeTaxD.DutyFeeTaxAccountingCurrency = in.DutyFeeTaxAccountingCurrency
	leviedDutyFeeTaxD.DutyFeeTaxAgencyName = in.DutyFeeTaxAgencyName
	leviedDutyFeeTaxD.DutyFeeTaxAmount = in.DutyFeeTaxAmount
	leviedDutyFeeTaxD.DftaCodeListVersion = in.DftaCodeListVersion
	leviedDutyFeeTaxD.DftaCurrencyCode = in.DftaCurrencyCode
	leviedDutyFeeTaxD.DutyFeeTaxAmountInAccountingCurrency = in.DutyFeeTaxAmountInAccountingCurrency
	leviedDutyFeeTaxD.DftaiacCodeListVersion = in.DftaiacCodeListVersion
	leviedDutyFeeTaxD.DftaiacCurrencyCode = in.DftaiacCurrencyCode
	leviedDutyFeeTaxD.DutyFeeTaxBasisAmount = in.DutyFeeTaxBasisAmount
	leviedDutyFeeTaxD.DftbaCodeListVersion = in.DftbaCodeListVersion
	leviedDutyFeeTaxD.DftbaCurrencyCode = in.DftbaCurrencyCode
	leviedDutyFeeTaxD.DutyFeeTaxBasisAmountInAccountingCurrency = in.DutyFeeTaxBasisAmountInAccountingCurrency
	leviedDutyFeeTaxD.DftbaiacCodeListVersion = in.DftbaiacCodeListVersion
	leviedDutyFeeTaxD.DftbaiacCurrencyCode = in.DftbaiacCurrencyCode
	leviedDutyFeeTaxD.DutyFeeTaxCategoryCode = in.DutyFeeTaxCategoryCode
	leviedDutyFeeTaxD.DutyFeeTaxDescription = in.DutyFeeTaxDescription
	leviedDutyFeeTaxD.DutyFeeTaxExemptionDescription = in.DutyFeeTaxExemptionDescription
	leviedDutyFeeTaxD.DutyFeeTaxExemptionReason = in.DutyFeeTaxExemptionReason
	leviedDutyFeeTaxD.DutyFeeTaxPercentage = in.DutyFeeTaxPercentage
	leviedDutyFeeTaxD.DutyFeeTaxTypeCode = in.DutyFeeTaxTypeCode
	leviedDutyFeeTaxD.Extension = in.Extension
	leviedDutyFeeTaxD.OrderLineItemId = in.OrderLineItemId
	leviedDutyFeeTaxD.InvoiceLineItemId = in.InvoiceLineItemId
	leviedDutyFeeTaxD.InvoiceId = in.InvoiceId

	leviedDutyFeeTaxT := invoiceproto.LeviedDutyFeeTaxT{}
	leviedDutyFeeTaxT.DutyFeeTaxPointDate = common.TimeToTimestamp(dutyFeeTaxPointDate.UTC().Truncate(time.Second))

	leviedDutyFeeTax := invoiceproto.LeviedDutyFeeTax{LeviedDutyFeeTaxD: &leviedDutyFeeTaxD, LeviedDutyFeeTaxT: &leviedDutyFeeTaxT}

	return &leviedDutyFeeTax, nil
}

// insertLeviedDutyFeeTax - Insert LeviedDutyFeeTax details into database
func (invs *InvoiceService) insertLeviedDutyFeeTax(ctx context.Context, insertLeviedDutyFeeTaxSQL string, leviedDutyFeeTax *invoiceproto.LeviedDutyFeeTax, userEmail string, requestID string) error {
	leviedDutyFeeTaxTmp, err := invs.crLeviedDutyFeeTaxStruct(ctx, leviedDutyFeeTax, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLeviedDutyFeeTaxSQL, leviedDutyFeeTaxTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		leviedDutyFeeTax.LeviedDutyFeeTaxD.Id = uint32(uID)

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crLeviedDutyFeeTaxStruct - process LeviedDutyFeeTax details
func (invs *InvoiceService) crLeviedDutyFeeTaxStruct(ctx context.Context, leviedDutyFeeTax *invoiceproto.LeviedDutyFeeTax, userEmail string, requestID string) (*invoicestruct.LeviedDutyFeeTax, error) {
	leviedDutyFeeTaxT := new(invoicestruct.LeviedDutyFeeTaxT)
	leviedDutyFeeTaxT.DutyFeeTaxPointDate = common.TimestampToTime(leviedDutyFeeTax.LeviedDutyFeeTaxT.DutyFeeTaxPointDate)

	leviedDutyFeeTaxTmp := invoicestruct.LeviedDutyFeeTax{LeviedDutyFeeTaxD: leviedDutyFeeTax.LeviedDutyFeeTaxD, LeviedDutyFeeTaxT: leviedDutyFeeTaxT}

	return &leviedDutyFeeTaxTmp, nil
}
