package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertSourceSQL = `insert into sources
  (source_type,
  source,
  event_id,
  type_of_event)
    values(
  :source_type,
  :source,
  :event_id,
  :type_of_event);`

const selectSourcesSQL = `select
  source_type,
  source,
  event_id,
  type_of_event from sources`

// CreateSource - Create Source
func (es *EpcisService) CreateSource(ctx context.Context, in *epcisproto.CreateSourceRequest) (*epcisproto.CreateSourceResponse, error) {
	source, err := es.ProcessSourceRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertSource(ctx, insertSourceSQL, source, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sourceResponse := epcisproto.CreateSourceResponse{}
	sourceResponse.Source = source
	return &sourceResponse, nil
}

// ProcessSourceRequest - ProcessSourceRequest
func (es *EpcisService) ProcessSourceRequest(ctx context.Context, in *epcisproto.CreateSourceRequest) (*epcisproto.Source, error) {
	source := epcisproto.Source{}
	source.SourceType = in.SourceType
	source.Source = in.Source
	source.EventId = in.EventId
	source.TypeOfEvent = in.TypeOfEvent
	return &source, nil
}

// insertSource - Insert Source details into database
func (es *EpcisService) insertSource(ctx context.Context, insertSourceSQL string, source *epcisproto.Source, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertSourceSQL, source)
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

// GetSources - Get Sources
func (es *EpcisService) GetSources(ctx context.Context, in *epcisproto.GetSourcesRequest) (*epcisproto.GetSourcesResponse, error) {
	query := "event_id = ? and type_of_event = ?"

	sources := []*epcisproto.Source{}

	nselectSourcesSQL := selectSourcesSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectSourcesSQL, in.EventId, in.TypeOfEvent)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		source := epcisproto.Source{}
		err = rows.StructScan(&source)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		sources = append(sources, &source)
	}

	sourcesResponse := epcisproto.GetSourcesResponse{Sources: sources}

	return &sourcesResponse, nil
}
