package logisticsservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	logisticsstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/logistics/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertReceivingAdviceLineItemSQL = `insert into receiving_advice_line_items
	  (
uuid4,	  
line_item_number,
parent_line_item_number,
quantity_accepted,
qa_measurement_unit_code,
qa_code_list_version,
quantity_despatched,
qd_measurement_unit_code,
qd_code_list_version,
quantity_received,
qr_measurement_unit_code,
qr_code_list_version,
transactional_trade_item,
ecom_consignment_identification,
contract,
customer_reference,
delivery_note,
despatch_advice,
product_certification,
promotional_deal,
purchase_conditions,
purchase_order,
requested_item_identification,
specification,
receiving_advice_id,
pick_up_date_time_begin,
pick_up_date_time_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
:uuid4,  
:line_item_number,
:parent_line_item_number,
:quantity_accepted,
:qa_measurement_unit_code,
:qa_code_list_version,
:quantity_despatched,
:qd_measurement_unit_code,
:qd_code_list_version,
:quantity_received,
:qr_measurement_unit_code,
:qr_code_list_version,
:transactional_trade_item,
:ecom_consignment_identification,
:contract,
:customer_reference,
:delivery_note,
:despatch_advice,
:product_certification,
:promotional_deal,
:purchase_conditions,
:purchase_order,
:requested_item_identification,
:specification,
:receiving_advice_id,
:pick_up_date_time_begin,
:pick_up_date_time_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectReceivingAdviceLineItemsSQL = `select
  id,
  uuid4,
  line_item_number,
  parent_line_item_number,
  quantity_accepted,
  qa_measurement_unit_code,
  qa_code_list_version,
  quantity_despatched,
  qd_measurement_unit_code,
  qd_code_list_version,
  quantity_received,
  qr_measurement_unit_code,
  qr_code_list_version,
  transactional_trade_item,
  ecom_consignment_identification,
  contract,
  customer_reference,
  delivery_note,
  despatch_advice,
  product_certification,
  promotional_deal,
  purchase_conditions,
  purchase_order,
  requested_item_identification,
  specification,
  receiving_advice_id,
  pick_up_date_time_begin,
  pick_up_date_time_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from receiving_advice_line_items`

func (ras *ReceivingAdviceService) CreateReceivingAdviceLineItem(ctx context.Context, in *logisticsproto.CreateReceivingAdviceLineItemRequest) (*logisticsproto.CreateReceivingAdviceLineItemResponse, error) {
	receivingAdviceLineItem, err := ras.ProcessReceivingAdviceLineItemRequest(ctx, in)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = ras.insertReceivingAdviceLineItem(ctx, insertReceivingAdviceLineItemSQL, receivingAdviceLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceLineItemResponse := logisticsproto.CreateReceivingAdviceLineItemResponse{}
	receivingAdviceLineItemResponse.ReceivingAdviceLineItem = receivingAdviceLineItem
	return &receivingAdviceLineItemResponse, nil
}

// ProcessReceivingAdviceLineItemRequest - ProcessReceivingAdviceLineItemRequest
func (ras *ReceivingAdviceService) ProcessReceivingAdviceLineItemRequest(ctx context.Context, in *logisticsproto.CreateReceivingAdviceLineItemRequest) (*logisticsproto.ReceivingAdviceLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, ras.UserServiceClient)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	pickUpDateTimeBegin, err := time.Parse(common.Layout, in.PickUpDateTimeBegin)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	pickUpDateTimeEnd, err := time.Parse(common.Layout, in.PickUpDateTimeEnd)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceLineItemD := logisticsproto.ReceivingAdviceLineItemD{}
	receivingAdviceLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdviceLineItemD.LineItemNumber = in.LineItemNumber
	receivingAdviceLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	receivingAdviceLineItemD.QuantityAccepted = in.QuantityAccepted
	receivingAdviceLineItemD.QaMeasurementUnitCode = in.QaMeasurementUnitCode
	receivingAdviceLineItemD.QaCodeListVersion = in.QaCodeListVersion
	receivingAdviceLineItemD.QuantityDespatched = in.QuantityDespatched
	receivingAdviceLineItemD.QdMeasurementUnitCode = in.QdMeasurementUnitCode
	receivingAdviceLineItemD.QdCodeListVersion = in.QdCodeListVersion
	receivingAdviceLineItemD.QuantityReceived = in.QuantityReceived
	receivingAdviceLineItemD.QrMeasurementUnitCode = in.QrMeasurementUnitCode
	receivingAdviceLineItemD.QrCodeListVersion = in.QrCodeListVersion
	receivingAdviceLineItemD.TransactionalTradeItem = in.TransactionalTradeItem
	receivingAdviceLineItemD.EcomConsignmentIdentification = in.EcomConsignmentIdentification
	receivingAdviceLineItemD.Contract = in.Contract
	receivingAdviceLineItemD.CustomerReference = in.CustomerReference
	receivingAdviceLineItemD.DeliveryNote = in.DeliveryNote
	receivingAdviceLineItemD.DespatchAdvice = in.DespatchAdvice
	receivingAdviceLineItemD.ProductCertification = in.ProductCertification
	receivingAdviceLineItemD.PromotionalDeal = in.PromotionalDeal
	receivingAdviceLineItemD.PurchaseConditions = in.PurchaseConditions
	receivingAdviceLineItemD.PurchaseOrder = in.PurchaseOrder
	receivingAdviceLineItemD.RequestedItemIdentification = in.RequestedItemIdentification
	receivingAdviceLineItemD.Specification = in.Specification

	receivingAdviceLineItemT := logisticsproto.ReceivingAdviceLineItemT{}
	receivingAdviceLineItemT.PickUpDateTimeBegin = common.TimeToTimestamp(pickUpDateTimeBegin.UTC().Truncate(time.Second))
	receivingAdviceLineItemT.PickUpDateTimeEnd = common.TimeToTimestamp(pickUpDateTimeEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	receivingAdviceLineItem := logisticsproto.ReceivingAdviceLineItem{ReceivingAdviceLineItemD: &receivingAdviceLineItemD, ReceivingAdviceLineItemT: &receivingAdviceLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &receivingAdviceLineItem, nil
}

// insertReceivingAdviceLineItem - Insert ReceivingAdviceLineItem into database
func (ras *ReceivingAdviceService) insertReceivingAdviceLineItem(ctx context.Context, insertReceivingAdviceLineItemSQL string, receivingAdviceLineItem *logisticsproto.ReceivingAdviceLineItem, userEmail string, requestID string) error {
	receivingAdviceLineItemTmp, err := ras.crReceivingAdviceLineItemStruct(ctx, receivingAdviceLineItem, userEmail, requestID)
	if err != nil {
		ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = ras.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertReceivingAdviceLineItemSQL, receivingAdviceLineItemTmp)
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		receivingAdviceLineItem.ReceivingAdviceLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(receivingAdviceLineItem.ReceivingAdviceLineItemD.Uuid4)
		if err != nil {
			ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		receivingAdviceLineItem.ReceivingAdviceLineItemD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		ras.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crReceivingAdviceLineItemStruct - process ReceivingAdviceLineItem details
func (ras *ReceivingAdviceService) crReceivingAdviceLineItemStruct(ctx context.Context, receivingAdviceLineItem *logisticsproto.ReceivingAdviceLineItem, userEmail string, requestID string) (*logisticsstruct.ReceivingAdviceLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(receivingAdviceLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(receivingAdviceLineItem.CrUpdTime.UpdatedAt)

	receivingAdviceLineItemT := new(logisticsstruct.ReceivingAdviceLineItemT)
	receivingAdviceLineItemT.PickUpDateTimeBegin = common.TimestampToTime(receivingAdviceLineItem.ReceivingAdviceLineItemT.PickUpDateTimeBegin)
	receivingAdviceLineItemT.PickUpDateTimeEnd = common.TimestampToTime(receivingAdviceLineItem.ReceivingAdviceLineItemT.PickUpDateTimeEnd)

	receivingAdviceLineItemTmp := logisticsstruct.ReceivingAdviceLineItem{ReceivingAdviceLineItemD: receivingAdviceLineItem.ReceivingAdviceLineItemD, ReceivingAdviceLineItemT: receivingAdviceLineItemT, CrUpdUser: receivingAdviceLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &receivingAdviceLineItemTmp, nil
}

// GetReceivingAdviceLineItems - GetReceivingAdviceLineItems
func (ras *ReceivingAdviceService) GetReceivingAdviceLineItems(ctx context.Context, inReq *logisticsproto.GetReceivingAdviceLineItemsRequest) (*logisticsproto.GetReceivingAdviceLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := logisticsproto.GetReceivingAdviceRequest{}
	form.GetRequest = &getRequest

	receivingAdviceResponse, err := ras.GetReceivingAdvice(ctx, &form)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	receivingAdvice := receivingAdviceResponse.ReceivingAdvice

	receivingAdviceLineItems := []*logisticsproto.ReceivingAdviceLineItem{}

	nselectReceivingAdviceLineItemsSQL := selectReceivingAdviceLineItemsSQL + ` where receiving_advice_id = ?;`
	rows, err := ras.DBService.DB.QueryxContext(ctx, nselectReceivingAdviceLineItemsSQL, receivingAdvice.ReceivingAdviceD.Id)
	if err != nil {
		ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		receivingAdviceLineItemTmp := logisticsstruct.ReceivingAdviceLineItem{}
		err = rows.StructScan(&receivingAdviceLineItemTmp)
		if err != nil {
			ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		receivingAdviceLineItemT := logisticsproto.ReceivingAdviceLineItemT{}
		receivingAdviceLineItemT.PickUpDateTimeBegin = common.TimeToTimestamp(receivingAdviceLineItemTmp.ReceivingAdviceLineItemT.PickUpDateTimeBegin)
		receivingAdviceLineItemT.PickUpDateTimeEnd = common.TimeToTimestamp(receivingAdviceLineItemTmp.ReceivingAdviceLineItemT.PickUpDateTimeEnd)

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(receivingAdviceLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(receivingAdviceLineItemTmp.CrUpdTime.UpdatedAt)

		uuid4Str, err := common.UUIDBytesToStr(receivingAdviceLineItemTmp.ReceivingAdviceLineItemD.Uuid4)
		if err != nil {
			ras.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		receivingAdviceLineItemTmp.ReceivingAdviceLineItemD.IdS = uuid4Str

		receivingAdviceLineItem := logisticsproto.ReceivingAdviceLineItem{ReceivingAdviceLineItemD: receivingAdviceLineItemTmp.ReceivingAdviceLineItemD, ReceivingAdviceLineItemT: &receivingAdviceLineItemT, CrUpdUser: receivingAdviceLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		receivingAdviceLineItems = append(receivingAdviceLineItems, &receivingAdviceLineItem)
	}
	receivingAdviceLineItemsResponse := logisticsproto.GetReceivingAdviceLineItemsResponse{}
	receivingAdviceLineItemsResponse.ReceivingAdviceLineItems = receivingAdviceLineItems
	return &receivingAdviceLineItemsResponse, nil
}
