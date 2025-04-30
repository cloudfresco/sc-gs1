package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertBizTransactionSQL = `insert into biz_transactions
	(biz_transaction_type,
  biz_transaction,
  event_id,
  type_of_event)
    values(
  :biz_transaction_type,
  :biz_transaction,
  :event_id,
  :type_of_event);`

const selectBizTransactionsSQL = `select
  biz_transaction_type,
  biz_transaction,
  event_id,
  type_of_event from biz_transactions`

// CreateBizTransaction - Create BizTransaction
func (es *EpcisService) CreateBizTransaction(ctx context.Context, in *epcisproto.CreateBizTransactionRequest) (*epcisproto.CreateBizTransactionResponse, error) {
	bizTransaction, err := es.ProcessBizTransactionRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertBizTransaction(ctx, insertBizTransactionSQL, bizTransaction, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bizTransactionResponse := epcisproto.CreateBizTransactionResponse{}
	bizTransactionResponse.BizTransaction = bizTransaction
	return &bizTransactionResponse, nil
}

// ProcessBizTransactionRequest - ProcessBizTransactionRequest
func (es *EpcisService) ProcessBizTransactionRequest(ctx context.Context, in *epcisproto.CreateBizTransactionRequest) (*epcisproto.BizTransaction, error) {
	bizTransaction := epcisproto.BizTransaction{}
	bizTransaction.BizTransactionType = in.BizTransactionType
	bizTransaction.BizTransaction = in.BizTransaction
	bizTransaction.EventId = in.EventId
	bizTransaction.TypeOfEvent = in.TypeOfEvent
	return &bizTransaction, nil
}

// insertBizTransaction - Insert BizTransaction details into database
func (es *EpcisService) insertBizTransaction(ctx context.Context, insertBizTransactionSQL string, bizTransaction *epcisproto.BizTransaction, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertBizTransactionSQL, bizTransaction)
		if err != nil {
			es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		return nil
	})
	if err != nil {
		es.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// GetBizTransactions - Get BizTransactions
func (es *EpcisService) GetBizTransactions(ctx context.Context, in *epcisproto.GetBizTransactionsRequest) (*epcisproto.GetBizTransactionsResponse, error) {
	query := "event_id = ? and type_of_event = ?"

	bizTransactions := []*epcisproto.BizTransaction{}

	nselectBizTransactionsSQL := selectBizTransactionsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectBizTransactionsSQL, in.EventId, in.TypeOfEvent)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		bizTransaction := epcisproto.BizTransaction{}
		err = rows.StructScan(&bizTransaction)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		bizTransactions = append(bizTransactions, &bizTransaction)
	}

	bizTransactionsResponse := epcisproto.GetBizTransactionsResponse{BizTransactions: bizTransactions}

	return &bizTransactionsResponse, nil
}
