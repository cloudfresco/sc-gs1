package epcisservices

import (
	"context"

	epcisproto "github.com/cloudfresco/sc-gs1/internal/protogen/epcis/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertQuantityElementSQL = `insert into quantity_elements
  (epc_class,
  quantity,
  uom,
  event_id,
  type_of_event,
  type_of_quantity)
    values(
  :epc_class,
  :quantity,
  :uom,
  :event_id,
  :type_of_event,
  :type_of_quantity);`

const selectQuantityElementsSQL = `select
  epc_class,
  quantity,
  uom,
  event_id,
  type_of_event,
  type_of_quantity from quantity_elements`

// CreateQuantityElement - Create QuantityElement
func (es *EpcisService) CreateQuantityElement(ctx context.Context, in *epcisproto.CreateQuantityElementRequest) (*epcisproto.CreateQuantityElementResponse, error) {
	quantityElement, err := es.ProcessQuantityElementRequest(ctx, in)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = es.insertQuantityElement(ctx, insertQuantityElementSQL, quantityElement, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	quantityElementResponse := epcisproto.CreateQuantityElementResponse{}
	quantityElementResponse.QuantityElement = quantityElement
	return &quantityElementResponse, nil
}

// ProcessQuantityElementRequest - ProcessQuantityElementRequest
func (es *EpcisService) ProcessQuantityElementRequest(ctx context.Context, in *epcisproto.CreateQuantityElementRequest) (*epcisproto.QuantityElement, error) {
	quantityElement := epcisproto.QuantityElement{}
	quantityElement.EpcClass = in.EpcClass
	quantityElement.Quantity = in.Quantity
	quantityElement.Uom = in.Uom
	quantityElement.EventId = in.EventId
	quantityElement.TypeOfEvent = in.TypeOfEvent
	quantityElement.TypeOfQuantity = in.TypeOfQuantity
	return &quantityElement, nil
}

// insertQuantityElement - Insert QuantityElement details into database
func (es *EpcisService) insertQuantityElement(ctx context.Context, insertQuantityElementSQL string, quantityElement *epcisproto.QuantityElement, userEmail string, requestID string) error {
	err := es.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertQuantityElementSQL, quantityElement)
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

// GetQuantityElements - Get QuantityElements
func (es *EpcisService) GetQuantityElements(ctx context.Context, in *epcisproto.GetQuantityElementsRequest) (*epcisproto.GetQuantityElementsResponse, error) {
	query := "event_id = ? and type_of_event = ? and type_of_quantity = ?"

	quantityElements := []*epcisproto.QuantityElement{}

	nselectQuantityElementsSQL := selectQuantityElementsSQL + ` where ` + query

	rows, err := es.DBService.DB.QueryxContext(ctx, nselectQuantityElementsSQL, in.EventId, in.TypeOfEvent, in.TypeOfQuantity)
	if err != nil {
		es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		quantityElement := epcisproto.QuantityElement{}
		err = rows.StructScan(&quantityElement)
		if err != nil {
			es.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		quantityElements = append(quantityElements, &quantityElement)
	}

	quantityElementsResponse := epcisproto.GetQuantityElementsResponse{QuantityElements: quantityElements}

	return &quantityElementsResponse, nil
}
