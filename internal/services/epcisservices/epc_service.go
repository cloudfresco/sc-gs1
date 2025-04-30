package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertEpcSQL = `insert into epcs
  (epc_value,
  event_id,
  type_of_event,
  type_of_epc)
    values(
  :epc_value,
  :event_id,
  :type_of_event,
  :type_of_epc);`

const selectEpcsSQL = `select
 epc_value,
 event_id,
 type_of_event,
 type_of_epc from epcs`

// CreateEpc - Create Epc
func (es *EpcisService) CreateEpc(ctx context.Context, in *epcisproto.CreateEpcRequest) (*epcisproto.CreateEpcResponse, error) {
	epc, err := es.ProcessEpcRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertEpc(ctx, insertEpcSQL, epc, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	epcResponse := epcisproto.CreateEpcResponse{}
	epcResponse.Epc = epc
	return &epcResponse, nil
}

// ProcessEpcRequest - ProcessEpcRequest
func (es *EpcisService) ProcessEpcRequest(ctx context.Context, in *epcisproto.CreateEpcRequest) (*epcisproto.Epc, error) {
	epc := epcisproto.Epc{}
	epc.EpcValue = in.EpcValue
	epc.EventId = in.EventId
	epc.TypeOfEvent = in.TypeOfEvent
	epc.TypeOfEpc = in.TypeOfEpc
	return &epc, nil
}

// insertEpc - Insert Epc details into database
func (es *EpcisService) insertEpc(ctx context.Context, insertEpcSQL string, epc *epcisproto.Epc, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertEpcSQL, epc)
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

// GetEpcs - Get Epcs
func (es *EpcisService) GetEpcs(ctx context.Context, in *epcisproto.GetEpcsRequest) (*epcisproto.GetEpcsResponse, error) {
	query := "event_id = ? and type_of_event = ? and type_of_epc = ?"

	epcs := []*epcisproto.Epc{}

	nselectEpcsSQL := selectEpcsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectEpcsSQL, in.EventId, in.TypeOfEvent, in.TypeOfEpc)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		epc := epcisproto.Epc{}
		err = rows.StructScan(&epc)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		epcs = append(epcs, &epc)
	}

	epcsResponse := epcisproto.GetEpcsResponse{Epcs: epcs}

	return &epcsResponse, nil
}
