package invoiceservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	invoicestruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/invoice/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDebitCreditAdviceLineItemSQL = `insert into debit_credit_advice_line_items
	  (
	  uuid4,
    adjustment_amount,
    adjustment_amount_currency,
    aa_code_list_version,
    debit_credit_indicator_code,
    financial_adjustment_reason_code,
    line_item_number,
    parent_line_item_number,
    debit_credit_advice_id,
    status_code,
    created_by_user_id,
    updated_by_user_id,
    created_at,
    updated_at)
      values(
    :uuid4,
    :adjustment_amount,
    :adjustment_amount_currency,
    :aa_code_list_version,
    :debit_credit_indicator_code,
    :financial_adjustment_reason_code,
    :line_item_number,
    :parent_line_item_number,
    :debit_credit_advice_id,
    :status_code,
    :created_by_user_id,
    :updated_by_user_id,
    :created_at,
    :updated_at);`

const selectDebitCreditAdviceLineItemsSQL = `select
  id,
  uuid4,
  adjustment_amount,
  adjustment_amount_currency,
  aa_code_list_version, 
  debit_credit_indicator_code,
  financial_adjustment_reason_code,
  line_item_number,
  parent_line_item_number,
  debit_credit_advice_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from debit_credit_advice_line_items`

// CreateDebitCreditAdviceLineItem - Create DebitCreditAdviceLineItem
func (ds *DebitCreditAdviceService) CreateDebitCreditAdviceLineItem(ctx context.Context, in *invoiceproto.CreateDebitCreditAdviceLineItemRequest) (*invoiceproto.CreateDebitCreditAdviceLineItemResponse, error) {
	debitCreditAdviceLineItem, err := ds.ProcessDebitCreditAdviceLineItemRequest(ctx, in)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ds.insertDebitCreditAdviceLineItem(ctx, insertDebitCreditAdviceLineItemSQL, debitCreditAdviceLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceLineItemResponse := invoiceproto.CreateDebitCreditAdviceLineItemResponse{}
	debitCreditAdviceLineItemResponse.DebitCreditAdviceLineItem = debitCreditAdviceLineItem
	return &debitCreditAdviceLineItemResponse, nil
}

// ProcessDebitCreditAdviceLineItemRequest - ProcessDebitCreditAdviceLineItemRequest
func (ds *DebitCreditAdviceService) ProcessDebitCreditAdviceLineItemRequest(ctx context.Context, in *invoiceproto.CreateDebitCreditAdviceLineItemRequest) (*invoiceproto.DebitCreditAdviceLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ds.UserServiceClient)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	debitCreditAdviceLineItemD := invoiceproto.DebitCreditAdviceLineItemD{}
	debitCreditAdviceLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	debitCreditAdviceLineItemD.AdjustmentAmount = in.AdjustmentAmount
	debitCreditAdviceLineItemD.AaCodeListVersion = in.AaCodeListVersion
	debitCreditAdviceLineItemD.DebitCreditIndicatorCode = in.DebitCreditIndicatorCode
	debitCreditAdviceLineItemD.FinancialAdjustmentReasonCode = in.FinancialAdjustmentReasonCode
	debitCreditAdviceLineItemD.LineItemNumber = in.LineItemNumber
	debitCreditAdviceLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	debitCreditAdviceLineItemD.DebitCreditAdviceId = in.DebitCreditAdviceId

	adjustmentAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, in.AdjustmentAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	adjustmentAmountMinor, err := common.ParseAmountString(in.AdjustmentAmount, adjustmentAmountCurrency)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdviceLineItemD.AdjustmentAmountCurrency = adjustmentAmountCurrency.Code
	debitCreditAdviceLineItemD.AdjustmentAmount = adjustmentAmountMinor
	debitCreditAdviceLineItemD.AdjustmentAmountString = common.FormatAmountString(adjustmentAmountMinor, adjustmentAmountCurrency)

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	debitCreditAdviceLineItem := invoiceproto.DebitCreditAdviceLineItem{DebitCreditAdviceLineItemD: &debitCreditAdviceLineItemD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &debitCreditAdviceLineItem, nil
}

// insertDebitCreditAdviceLineItem - Insert DebitCreditAdviceLineItem details into database
func (ds *DebitCreditAdviceService) insertDebitCreditAdviceLineItem(ctx context.Context, insertDebitCreditAdviceLineItemSQL string, debitCreditAdviceLineItem *invoiceproto.DebitCreditAdviceLineItem, userEmail string, requestID string) error {
	debitCreditAdviceLineItemTmp, err := ds.crDebitCreditAdviceLineItemStruct(ctx, debitCreditAdviceLineItem, userEmail, requestID)
	if err != nil {
		ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ds.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertDebitCreditAdviceLineItemSQL, debitCreditAdviceLineItemTmp)
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.Uuid4)
		if err != nil {
			ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		ds.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDebitCreditAdviceLineItemStruct - process DebitCreditAdviceLineItem details
func (ds *DebitCreditAdviceService) crDebitCreditAdviceLineItemStruct(ctx context.Context, debitCreditAdviceLineItem *invoiceproto.DebitCreditAdviceLineItem, userEmail string, requestID string) (*invoicestruct.DebitCreditAdviceLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(debitCreditAdviceLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(debitCreditAdviceLineItem.CrUpdTime.UpdatedAt)

	debitCreditAdviceLineItemTmp := invoicestruct.DebitCreditAdviceLineItem{DebitCreditAdviceLineItemD: debitCreditAdviceLineItem.DebitCreditAdviceLineItemD, CrUpdUser: debitCreditAdviceLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &debitCreditAdviceLineItemTmp, nil
}

// GetDebitCreditAdviceLineItems - GetDebitCreditAdviceLineItems
func (ds *DebitCreditAdviceService) GetDebitCreditAdviceLineItems(ctx context.Context, inReq *invoiceproto.GetDebitCreditAdviceLineItemsRequest) (*invoiceproto.GetDebitCreditAdviceLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := invoiceproto.GetDebitCreditAdviceRequest{}
	form.GetRequest = &getRequest

	debitCreditAdviceResponse, err := ds.GetDebitCreditAdvice(ctx, &form)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	debitCreditAdvice := debitCreditAdviceResponse.DebitCreditAdvice

	debitCreditAdviceLineItems := []*invoiceproto.DebitCreditAdviceLineItem{}

	nselectDebitCreditAdviceLineItemsSQL := selectDebitCreditAdviceLineItemsSQL + ` where debit_credit_advice_id = ?;`
	rows, err := ds.DBService.DB.QueryxContext(ctx, nselectDebitCreditAdviceLineItemsSQL, debitCreditAdvice.DebitCreditAdviceD.Id)
	if err != nil {
		ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	for rows.Next() {

		debitCreditAdviceLineItemTmp := invoicestruct.DebitCreditAdviceLineItem{}
		err = rows.StructScan(&debitCreditAdviceLineItemTmp)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		uuid4Str, err := common.UUIDBytesToStr(debitCreditAdviceLineItemTmp.DebitCreditAdviceLineItemD.Uuid4)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		debitCreditAdviceLineItemTmp.DebitCreditAdviceLineItemD.IdS = uuid4Str

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(debitCreditAdviceLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(debitCreditAdviceLineItemTmp.CrUpdTime.UpdatedAt)

		debitCreditAdviceLineItem := invoiceproto.DebitCreditAdviceLineItem{DebitCreditAdviceLineItemD: debitCreditAdviceLineItemTmp.DebitCreditAdviceLineItemD, CrUpdUser: debitCreditAdviceLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		adjustmentAmountCurrency, err := ds.CurrencyService.GetCurrency(ctx, debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.AdjustmentAmountCurrency)
		if err != nil {
			ds.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.AdjustmentAmountString = common.FormatAmountString(debitCreditAdviceLineItem.DebitCreditAdviceLineItemD.AdjustmentAmount, adjustmentAmountCurrency)

		debitCreditAdviceLineItems = append(debitCreditAdviceLineItems, &debitCreditAdviceLineItem)
	}
	debitCreditAdviceLineItemsResponse := invoiceproto.GetDebitCreditAdviceLineItemsResponse{}
	debitCreditAdviceLineItemsResponse.DebitCreditAdviceLineItems = debitCreditAdviceLineItems

	return &debitCreditAdviceLineItemsResponse, nil
}
