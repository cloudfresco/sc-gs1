package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertPersistentDispositionSQL = `insert into persistent_dispositions
  (set_disp,
  unset_disp,
  event_id,
  type_of_event)
    values(:set_disp,
  :unset_disp,
  :event_id,
  :type_of_event);`

const selectPersistentDispositionsSQL = `select
  set_disp,
  unset_disp,
  event_id,
  type_of_event from persistent_dispositions`

// CreatePersistentDisposition - Create PersistentDisposition
func (es *EpcisService) CreatePersistentDisposition(ctx context.Context, in *epcisproto.CreatePersistentDispositionRequest) (*epcisproto.CreatePersistentDispositionResponse, error) {
	persistentDisposition, err := es.ProcessPersistentDispositionRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertPersistentDisposition(ctx, insertPersistentDispositionSQL, persistentDisposition, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	persistentDispositionResponse := epcisproto.CreatePersistentDispositionResponse{}
	persistentDispositionResponse.PersistentDisposition = persistentDisposition
	return &persistentDispositionResponse, nil
}

// ProcessPersistentDispositionRequest - ProcessPersistentDispositionRequest
func (es *EpcisService) ProcessPersistentDispositionRequest(ctx context.Context, in *epcisproto.CreatePersistentDispositionRequest) (*epcisproto.PersistentDisposition, error) {
	persistentDisposition := epcisproto.PersistentDisposition{}
	persistentDisposition.SetDisp = in.SetDisp
	persistentDisposition.UnsetDisp = in.UnsetDisp
	persistentDisposition.EventId = in.EventId
	persistentDisposition.TypeOfEvent = in.TypeOfEvent
	return &persistentDisposition, nil
}

// insertPersistentDisposition - Insert PersistentDisposition details into database
func (es *EpcisService) insertPersistentDisposition(ctx context.Context, insertPersistentDispositionSQL string, persistentDisposition *epcisproto.PersistentDisposition, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertPersistentDispositionSQL, persistentDisposition)
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

// GetPersistentDispositions - Get PersistentDispositions
func (es *EpcisService) GetPersistentDispositions(ctx context.Context, in *epcisproto.GetPersistentDispositionsRequest) (*epcisproto.GetPersistentDispositionsResponse, error) {
	query := "event_id = ? and type_of_event = ?"

	persistentDispositions := []*epcisproto.PersistentDisposition{}

	nselectPersistentDispositionsSQL := selectPersistentDispositionsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectPersistentDispositionsSQL, in.EventId, in.TypeOfEvent)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		persistentDisposition := epcisproto.PersistentDisposition{}
		err = rows.StructScan(&persistentDisposition)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		persistentDispositions = append(persistentDispositions, &persistentDisposition)
	}

	persistentDispositionsResponse := epcisproto.GetPersistentDispositionsResponse{PersistentDispositions: persistentDispositions}

	return &persistentDispositionsResponse, nil
}
