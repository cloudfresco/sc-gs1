package orderservices

import (
	"context"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderResponseLineItemDetailSQL = `insert into order_response_line_item_details
	    (confirmed_quantity,
cq_measurement_unit_code,
cq_code_list_version,
return_reason_code,
order_response_line_item_id)
  values(
  :onfirmed_quantity,
:cq_measurement_unit_code,
:cq_code_list_version,
:return_reason_code,
:order_response_line_item_id);`

// CreateOrderResponseLineItemDetail - Create OrderResponseLineItemDetail
func (o *OrderResponseService) CreateOrderResponseLineItemDetail(ctx context.Context, in *orderproto.CreateOrderResponseLineItemDetailRequest) (*orderproto.CreateOrderResponseLineItemDetailResponse, error) {
	orderResponseLineItemDetail, err := o.ProcessOrderResponseLineItemDetailRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderResponseLineItemDetail(ctx, insertOrderResponseLineItemDetailSQL, orderResponseLineItemDetail, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderResponseLineItemDetailResponse := orderproto.CreateOrderResponseLineItemDetailResponse{}
	orderResponseLineItemDetailResponse.OrderResponseLineItemDetail = orderResponseLineItemDetail
	return &orderResponseLineItemDetailResponse, nil
}

// ProcessOrderResponseLineItemDetailRequest - ProcessOrderResponseLineItemDetailRequest
func (o *OrderResponseService) ProcessOrderResponseLineItemDetailRequest(ctx context.Context, in *orderproto.CreateOrderResponseLineItemDetailRequest) (*orderproto.OrderResponseLineItemDetail, error) {
	orderResponseLineItemDetail := orderproto.OrderResponseLineItemDetail{}

	orderResponseLineItemDetail.ConfirmedQuantity = in.ConfirmedQuantity
	orderResponseLineItemDetail.CqMeasurementUnitCode = in.CqMeasurementUnitCode
	orderResponseLineItemDetail.CqCodeListVersion = in.CqCodeListVersion
	orderResponseLineItemDetail.ReturnReasonCode = in.ReturnReasonCode
	orderResponseLineItemDetail.OrderResponseLineItemId = in.OrderResponseLineItemId

	return &orderResponseLineItemDetail, nil
}

// insertOrderResponseLineItemDetail - Insert OrderResponseLineItemDetail details into database
func (o *OrderResponseService) insertOrderResponseLineItemDetail(ctx context.Context, insertOrderResponseLineItemDetailSQL string, orderResponseLineItemDetail *orderproto.OrderResponseLineItemDetail, userEmail string, requestID string) error {
	err := o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertOrderResponseLineItemDetailSQL, orderResponseLineItemDetail)
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
