package orderservices

import (
	"context"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderLineItemDetailSQL = `insert into order_line_item_details
	    (requested_quantity,
       rq_measurement_unit_code,
       rq_code_list_version,
       order_line_item_id)
  values(
  :requested_quantity,
  :rq_measurement_unit_code,
  :rq_code_list_version,
  :order_line_item_id);`

/*const selectOrderLineItemDetailsSQL = `select
  id,
  requested_quantity,
  rq_measurement_unit_code,
  rq_code_list_version,
  order_line_item_id from order_line_item_details`*/

// CreateOrderLineItemDetail - Create OrderLineItemDetail
func (o *OrderService) CreateOrderLineItemDetail(ctx context.Context, in *orderproto.CreateOrderLineItemDetailRequest) (*orderproto.CreateOrderLineItemDetailResponse, error) {
	orderLineItemDetail, err := o.ProcessOrderLineItemDetailRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderLineItemDetail(ctx, insertOrderLineItemDetailSQL, orderLineItemDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLineItemDetailResponse := orderproto.CreateOrderLineItemDetailResponse{}
	orderLineItemDetailResponse.OrderLineItemDetail = orderLineItemDetail
	return &orderLineItemDetailResponse, nil
}

// ProcessOrderLineItemDetailRequest - ProcessOrderLineItemDetailRequest
func (o *OrderService) ProcessOrderLineItemDetailRequest(ctx context.Context, in *orderproto.CreateOrderLineItemDetailRequest) (*orderproto.OrderLineItemDetail, error) {
	orderLineItemDetail := orderproto.OrderLineItemDetail{}

	orderLineItemDetail.RequestedQuantity = in.RequestedQuantity
	orderLineItemDetail.RqMeasurementUnitCode = in.RqMeasurementUnitCode
	orderLineItemDetail.RqCodeListVersion = in.RqCodeListVersion
	orderLineItemDetail.OrderLineItemId = in.OrderLineItemId

	return &orderLineItemDetail, nil
}

// insertOrderLineItemDetail - Insert OrderLineItemDetail details into database
func (o *OrderService) insertOrderLineItemDetail(ctx context.Context, insertOrderLineItemDetailSQL string, orderLineItemDetail *orderproto.OrderLineItemDetail, userEmail string, requestID string) error {
	err := o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertOrderLineItemDetailSQL, orderLineItemDetail)
		if err != nil {
			o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
