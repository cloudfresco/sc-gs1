package orderservices

import (
	"context"

	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertOrderLogisticalInformationSQL = `insert into order_logistical_informations
	    ( commodity_type_code,
    shipment_split_method_code,
    intermediate_delivery_party,
    inventory_location,
    ship_from,
    ship_to,
    ultimate_consignee,
    order_id)
        values(
      :commodity_type_code,
      :shipment_split_method_code,
      :intermediate_delivery_party,
      :inventory_location,
      :ship_from,
      :ship_to,
      :ultimate_consignee,
      :order_id);`

/*const selectOrderLogisticalInformationsSQL = `select
  id,
  commodity_type_code,
  shipment_split_method_code,
  intermediate_delivery_party,
  inventory_location,
  ship_from,
  ship_to,
  ultimate_consignee,
  order_id from order_logistical_informations`*/

// CreateOrderLogisticalInformation - Create OrderLogisticalInformation
func (o *OrderService) CreateOrderLogisticalInformation(ctx context.Context, in *orderproto.CreateOrderLogisticalInformationRequest) (*orderproto.CreateOrderLogisticalInformationResponse, error) {
	orderLogisticalInformation, err := o.ProcessOrderLogisticalInformationRequest(ctx, in)
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = o.insertOrderLogisticalInformation(ctx, insertOrderLogisticalInformationSQL, orderLogisticalInformation, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		o.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	orderLogisticalInformationResponse := orderproto.CreateOrderLogisticalInformationResponse{}
	orderLogisticalInformationResponse.OrderLogisticalInformation = orderLogisticalInformation
	return &orderLogisticalInformationResponse, nil
}

// ProcessOrderLogisticalInformationRequest - ProcessOrderLogisticalInformationRequest
func (o *OrderService) ProcessOrderLogisticalInformationRequest(ctx context.Context, in *orderproto.CreateOrderLogisticalInformationRequest) (*orderproto.OrderLogisticalInformation, error) {
	orderLogisticalInformation := orderproto.OrderLogisticalInformation{}
	orderLogisticalInformation.CommodityTypeCode = in.CommodityTypeCode
	orderLogisticalInformation.ShipmentSplitMethodCode = in.ShipmentSplitMethodCode
	orderLogisticalInformation.IntermediateDeliveryParty = in.IntermediateDeliveryParty
	orderLogisticalInformation.InventoryLocation = in.InventoryLocation
	orderLogisticalInformation.ShipFrom = in.ShipFrom
	orderLogisticalInformation.ShipTo = in.ShipTo
	orderLogisticalInformation.UltimateConsignee = in.UltimateConsignee
	orderLogisticalInformation.OrderId = in.OrderId

	return &orderLogisticalInformation, nil
}

// insertOrderLogisticalInformation - Insert OrderLogisticalInformation details into database
func (o *OrderService) insertOrderLogisticalInformation(ctx context.Context, insertOrderLogisticalInformationSQL string, orderLogisticalInformation *orderproto.OrderLogisticalInformation, userEmail string, requestID string) error {
	err := o.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		_, err := tx.NamedExecContext(ctx, insertOrderLogisticalInformationSQL, orderLogisticalInformation)
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
