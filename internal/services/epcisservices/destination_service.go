package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertDestinationSQL = `insert into destinations
  (dest_type,
  destination,
  event_id,
  type_of_event)
    values(
  :dest_type,
  :destination,
  :event_id,
  :type_of_event);`

const selectDestinationsSQL = `select
  dest_type,
  destination,
  event_id,
  type_of_event from destinations`

// CreateDestination - Create Destination
func (es *EpcisService) CreateDestination(ctx context.Context, in *epcisproto.CreateDestinationRequest) (*epcisproto.CreateDestinationResponse, error) {
	destination, err := es.ProcessDestinationRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertDestination(ctx, insertDestinationSQL, destination, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	destinationResponse := epcisproto.CreateDestinationResponse{}
	destinationResponse.Destination = destination
	return &destinationResponse, nil
}

// ProcessDestinationRequest - ProcessDestinationRequest
func (es *EpcisService) ProcessDestinationRequest(ctx context.Context, in *epcisproto.CreateDestinationRequest) (*epcisproto.Destination, error) {
	destination := epcisproto.Destination{}
	destination.DestType = in.DestType
	destination.Destination = in.Destination
	destination.EventId = in.EventId
	destination.TypeOfEvent = in.TypeOfEvent
	return &destination, nil
}

// insertDestination - Insert Destination details into database
func (es *EpcisService) insertDestination(ctx context.Context, insertDestinationSQL string, destination *epcisproto.Destination, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertDestinationSQL, destination)
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

// GetDestinations - Get Destinations
func (es *EpcisService) GetDestinations(ctx context.Context, in *epcisproto.GetDestinationsRequest) (*epcisproto.GetDestinationsResponse, error) {
	query := "event_id = ? and type_of_event = ?"

	destinations := []*epcisproto.Destination{}

	nselectDestinationsSQL := selectDestinationsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectDestinationsSQL, in.EventId, in.TypeOfEvent)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		destination := epcisproto.Destination{}
		err = rows.StructScan(&destination)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		destinations = append(destinations, &destination)
	}

	destinationsResponse := epcisproto.GetDestinationsResponse{Destinations: destinations}

	return &destinationsResponse, nil
}
