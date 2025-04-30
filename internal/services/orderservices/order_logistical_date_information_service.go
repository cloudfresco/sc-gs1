package orderservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	orderstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderLogisticalDateInformationSQL = `insert into order_logistical_date_informations
	    (requested_delivery_date_range_begin,
      requested_delivery_date_range_end,
      requested_delivery_date_range_at_ultimate_consignee_begin,
      requested_delivery_date_range_at_ultimate_consignee_end,
      requested_delivery_date_time,
      requested_delivery_date_time_at_ultimate_consignee,
      requested_pick_up_date_time,
      requested_ship_date_range_begin,
      requested_ship_date_range_end,
      requested_ship_date_time,
      order_response_id)
        values(
      :requested_delivery_date_range_begin,
      :requested_delivery_date_range_end,
      :requested_delivery_date_range_at_ultimate_consignee_begin,
      :requested_delivery_date_range_at_ultimate_consignee_end,
      :requested_delivery_date_time,
      :requested_delivery_date_time_at_ultimate_consignee,
      :requested_pick_up_date_time,
      :requested_ship_date_range_begin,
      :requested_ship_date_range_end,
      :requested_ship_date_time,
      :order_response_id);`

/*const selectOrderLogisticalDateInformationsSQL = `select
  id,
  requested_delivery_date_range_begin,
  requested_delivery_date_range_end,
  requested_delivery_date_range_at_ultimate_consignee_begin,
  requested_delivery_date_range_at_ultimate_consignee_end,
  requested_delivery_date_time,
  requested_delivery_date_time_at_ultimate_consignee,
  requested_pick_up_date_time,
  requested_ship_date_range_begin,
  requested_ship_date_range_end,
  requested_ship_date_time,
  order_response_id from order_logistical_date_informations`*/

// CreateOrderLogisticalDateInformation - Create OrderLogisticalDateInformation
func (o *OrderService) CreateOrderLogisticalDateInformation(ctx context.Context, in *orderproto.CreateOrderLogisticalDateInformationRequest) (*orderproto.CreateOrderLogisticalDateInformationResponse, error) {
	orderLogisticalDateInformation, err := o.ProcessOrderLogisticalDateInformationRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderLogisticalDateInformation(ctx, insertOrderLogisticalDateInformationSQL, orderLogisticalDateInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLogisticalDateInformationResponse := orderproto.CreateOrderLogisticalDateInformationResponse{}
	orderLogisticalDateInformationResponse.OrderLogisticalDateInformation = orderLogisticalDateInformation
	return &orderLogisticalDateInformationResponse, nil
}

// ProcessOrderLogisticalDateInformationRequest - ProcessOrderLogisticalDateInformationRequest
func (o *OrderService) ProcessOrderLogisticalDateInformationRequest(ctx context.Context, in *orderproto.CreateOrderLogisticalDateInformationRequest) (*orderproto.OrderLogisticalDateInformation, error) {
	requestedDeliveryDateRangeBegin, err := time.Parse(common.Layout, in.RequestedDeliveryDateRangeBegin)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedDeliveryDateRangeEnd, err := time.Parse(common.Layout, in.RequestedDeliveryDateRangeEnd)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedDeliveryDateRangeAtUltimateConsigneeBegin, err := time.Parse(common.Layout, in.RequestedDeliveryDateRangeAtUltimateConsigneeBegin)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedDeliveryDateRangeAtUltimateConsigneeEnd, err := time.Parse(common.Layout, in.RequestedDeliveryDateRangeAtUltimateConsigneeEnd)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedDeliveryDateTime, err := time.Parse(common.Layout, in.RequestedDeliveryDateTime)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedDeliveryDateTimeAtUltimateConsignee, err := time.Parse(common.Layout, in.RequestedDeliveryDateTimeAtUltimateConsignee)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedPickUpDateTime, err := time.Parse(common.Layout, in.RequestedPickUpDateTime)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedShipDateRangeBegin, err := time.Parse(common.Layout, in.RequestedShipDateRangeBegin)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedShipDateRangeEnd, err := time.Parse(common.Layout, in.RequestedShipDateRangeEnd)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	requestedShipDateTime, err := time.Parse(common.Layout, in.RequestedShipDateTime)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLogisticalDateInformationT := orderproto.OrderLogisticalDateInformationT{}
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeBegin = common.TimeToTimestamp(requestedDeliveryDateRangeBegin.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeEnd = common.TimeToTimestamp(requestedDeliveryDateRangeEnd.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeBegin = common.TimeToTimestamp(requestedDeliveryDateRangeAtUltimateConsigneeBegin.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeEnd = common.TimeToTimestamp(requestedDeliveryDateRangeAtUltimateConsigneeEnd.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedDeliveryDateTime = common.TimeToTimestamp(requestedDeliveryDateTime.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedDeliveryDateTimeAtUltimateConsignee = common.TimeToTimestamp(requestedDeliveryDateTimeAtUltimateConsignee.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedPickUpDateTime = common.TimeToTimestamp(requestedPickUpDateTime.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedShipDateRangeBegin = common.TimeToTimestamp(requestedShipDateRangeBegin.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedShipDateRangeEnd = common.TimeToTimestamp(requestedShipDateRangeEnd.UTC().Truncate(time.Second))
	orderLogisticalDateInformationT.RequestedShipDateTime = common.TimeToTimestamp(requestedShipDateTime.UTC().Truncate(time.Second))

	orderLogisticalDateInformationD := orderproto.OrderLogisticalDateInformationD{}
	orderLogisticalDateInformationD.OrderResponseId = in.OrderResponseId

	orderLogisticalDateInformation := orderproto.OrderLogisticalDateInformation{OrderLogisticalDateInformationD: &orderLogisticalDateInformationD, OrderLogisticalDateInformationT: &orderLogisticalDateInformationT}

	return &orderLogisticalDateInformation, nil
}

// insertOrderLogisticalDateInformation - Insert OrderLogisticalDateInformation details into database
func (o *OrderService) insertOrderLogisticalDateInformation(ctx context.Context, insertOrderLogisticalDateInformationSQL string, orderLogisticalDateInformation *orderproto.OrderLogisticalDateInformation, userEmail string, requestID string) error {
	orderLogisticalDateInformationTmp, err := o.crOrderLogisticalDateInformationStruct(ctx, orderLogisticalDateInformation, userEmail, requestID)
	if err != nil {
		o.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err = tx.NamedExecContext(ctx, insertOrderLogisticalDateInformationSQL, orderLogisticalDateInformationTmp)

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

// crOrderLogisticalDateInformationStruct - process OrderLogisticalDateInformation details
func (o *OrderService) crOrderLogisticalDateInformationStruct(ctx context.Context, orderLogisticalDateInformation *orderproto.OrderLogisticalDateInformation, userEmail string, requestID string) (*orderstruct.OrderLogisticalDateInformation, error) {
	orderLogisticalDateInformationT := new(orderstruct.OrderLogisticalDateInformationT)
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeBegin = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateRangeBegin)
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeEnd = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateRangeEnd)
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeBegin = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeBegin)
	orderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeEnd = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateRangeAtUltimateConsigneeEnd)
	orderLogisticalDateInformationT.RequestedDeliveryDateTime = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateTime)
	orderLogisticalDateInformationT.RequestedDeliveryDateTimeAtUltimateConsignee = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedDeliveryDateTimeAtUltimateConsignee)
	orderLogisticalDateInformationT.RequestedPickUpDateTime = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedPickUpDateTime)
	orderLogisticalDateInformationT.RequestedShipDateRangeBegin = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedShipDateRangeBegin)
	orderLogisticalDateInformationT.RequestedShipDateRangeEnd = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedShipDateRangeEnd)
	orderLogisticalDateInformationT.RequestedShipDateTime = common.TimestampToTime(orderLogisticalDateInformation.OrderLogisticalDateInformationT.RequestedShipDateTime)

	orderLogisticalDateInformationTmp := orderstruct.OrderLogisticalDateInformation{OrderLogisticalDateInformationD: orderLogisticalDateInformation.OrderLogisticalDateInformationD, OrderLogisticalDateInformationT: orderLogisticalDateInformationT}

	return &orderLogisticalDateInformationTmp, nil
}
