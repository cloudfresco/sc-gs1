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

const insertDespatchAdviceLineItemSQL = `insert into despatch_advice_line_items
	  (
	  uuid4,
actual_processed_quantity,
measurement_unit_code,
code_list_version,
country_of_last_processing,
country_of_origin,
despatched_quantity,
dq_measurement_unit_code,
dq_code_list_version,
duty_fee_tax_liability,
extension,
free_goods_quantity,
fgq_measurement_unit_code,
fgq_code_list_version,
handling_instruction_code,
has_item_been_scanned_at_pos,
inventory_status_type,
line_item_number,
parent_line_item_number,
requested_quantity,
rq_measurement_unit_code,
rq_code_list_version,
contract,
coupon_clearing_house,
customer,
customer_document_reference,
customer_reference,
delivery_note,
item_owner,
original_supplier,
product_certification,
promotional_deal,
purchase_conditions,
purchase_order,
referenced_consignment,
requested_item_identification,
specification,
despatch_advice_id,
first_in_first_out_date_time,
pick_up_date_time,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(
 :uuid4, 
 :actual_processed_quantity,
:measurement_unit_code,
:code_list_version,
:country_of_last_processing,
:country_of_origin,
:despatched_quantity,
:dq_measurement_unit_code,
:dq_code_list_version,
:duty_fee_tax_liability,
:extension,
:free_goods_quantity,
:fgq_measurement_unit_code,
:fgq_code_list_version,
:handling_instruction_code,
:has_item_been_scanned_at_pos,
:inventory_status_type,
:line_item_number,
:parent_line_item_number,
:requested_quantity,
:rq_measurement_unit_code,
:rq_code_list_version,
:contract,
:coupon_clearing_house,
:customer,
:customer_document_reference,
:customer_reference,
:delivery_note,
:item_owner,
:original_supplier,
:product_certification,
:promotional_deal,
:purchase_conditions,
:purchase_order,
:referenced_consignment,
:requested_item_identification,
:specification,
:despatch_advice_id,
:first_in_first_out_date_time,
:pick_up_date_time,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

const selectDespatchAdviceLineItemsSQL = `select
  id,
  uuid4,
  actual_processed_quantity,
  measurement_unit_code,
  code_list_version,
  country_of_last_processing,
  country_of_origin,
  despatched_quantity,
  dq_measurement_unit_code,
  dq_code_list_version,
  duty_fee_tax_liability,
  extension,
  free_goods_quantity,
  fgq_measurement_unit_code,
  fgq_code_list_version,
  handling_instruction_code,
  has_item_been_scanned_at_pos,
  inventory_status_type,
  line_item_number,
  parent_line_item_number,
  requested_quantity,
  rq_measurement_unit_code,
  rq_code_list_version,
  contract,
  coupon_clearing_house,
  customer,
  customer_document_reference,
  customer_reference,
  delivery_note,
  item_owner,
  original_supplier,
  product_certification,
  promotional_deal,
  purchase_conditions,
  purchase_order,
  referenced_consignment,
  requested_item_identification,
  specification,
  despatch_advice_id,
  first_in_first_out_date_time,
  pick_up_date_time,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from despatch_advice_line_items`

func (das *DespatchAdviceService) CreateDespatchAdviceLineItem(ctx context.Context, in *logisticsproto.CreateDespatchAdviceLineItemRequest) (*logisticsproto.CreateDespatchAdviceLineItemResponse, error) {
	despatchAdviceLineItem, err := das.ProcessDespatchAdviceLineItemRequest(ctx, in)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = das.insertDespatchAdviceLineItem(ctx, insertDespatchAdviceLineItemSQL, despatchAdviceLineItem, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceLineItemResponse := logisticsproto.CreateDespatchAdviceLineItemResponse{}
	despatchAdviceLineItemResponse.DespatchAdviceLineItem = despatchAdviceLineItem
	return &despatchAdviceLineItemResponse, nil
}

// ProcessDespatchAdviceLineItemRequest - ProcessDespatchAdviceLineItemRequest
func (das *DespatchAdviceService) ProcessDespatchAdviceLineItemRequest(ctx context.Context, in *logisticsproto.CreateDespatchAdviceLineItemRequest) (*logisticsproto.DespatchAdviceLineItem, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, das.UserServiceClient)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	pickUpDateTime, err := time.Parse(common.Layout, in.PickUpDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	firstInFirstOutDateTime, err := time.Parse(common.Layout, in.FirstInFirstOutDateTime)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdviceLineItemD := logisticsproto.DespatchAdviceLineItemD{}
	despatchAdviceLineItemD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	despatchAdviceLineItemD.ActualProcessedQuantity = in.ActualProcessedQuantity
	despatchAdviceLineItemD.MeasurementUnitCode = in.MeasurementUnitCode
	despatchAdviceLineItemD.CodeListVersion = in.CodeListVersion
	despatchAdviceLineItemD.CountryOfLastProcessing = in.CountryOfLastProcessing
	despatchAdviceLineItemD.CountryOfOrigin = in.CountryOfOrigin
	despatchAdviceLineItemD.DespatchedQuantity = in.DespatchedQuantity
	despatchAdviceLineItemD.DqMeasurementUnitCode = in.DqMeasurementUnitCode
	despatchAdviceLineItemD.DqCodeListVersion = in.DqCodeListVersion
	despatchAdviceLineItemD.DutyFeeTaxLiability = in.DutyFeeTaxLiability
	despatchAdviceLineItemD.Extension = in.Extension
	despatchAdviceLineItemD.FreeGoodsQuantity = in.FreeGoodsQuantity
	despatchAdviceLineItemD.FgqMeasurementUnitCode = in.FgqMeasurementUnitCode
	despatchAdviceLineItemD.FgqCodeListVersion = in.FgqCodeListVersion
	despatchAdviceLineItemD.HandlingInstructionCode = in.HandlingInstructionCode
	despatchAdviceLineItemD.HasItemBeenScannedAtPos = in.HasItemBeenScannedAtPos
	despatchAdviceLineItemD.InventoryStatusType = in.InventoryStatusType
	despatchAdviceLineItemD.LineItemNumber = in.LineItemNumber
	despatchAdviceLineItemD.ParentLineItemNumber = in.ParentLineItemNumber
	despatchAdviceLineItemD.RequestedQuantity = in.RequestedQuantity
	despatchAdviceLineItemD.RqMeasurementUnitCode = in.RqMeasurementUnitCode
	despatchAdviceLineItemD.RqCodeListVersion = in.RqCodeListVersion
	despatchAdviceLineItemD.Contract = in.Contract
	despatchAdviceLineItemD.CouponClearingHouse = in.CouponClearingHouse
	despatchAdviceLineItemD.Customer = in.Customer
	despatchAdviceLineItemD.CustomerDocumentReference = in.CustomerDocumentReference
	despatchAdviceLineItemD.CustomerReference = in.CustomerReference
	despatchAdviceLineItemD.DeliveryNote = in.DeliveryNote
	despatchAdviceLineItemD.ItemOwner = in.ItemOwner
	despatchAdviceLineItemD.OriginalSupplier = in.OriginalSupplier
	despatchAdviceLineItemD.ProductCertification = in.ProductCertification
	despatchAdviceLineItemD.PromotionalDeal = in.PromotionalDeal
	despatchAdviceLineItemD.PurchaseConditions = in.PurchaseConditions
	despatchAdviceLineItemD.PurchaseOrder = in.PurchaseOrder
	despatchAdviceLineItemD.ReferencedConsignment = in.ReferencedConsignment
	despatchAdviceLineItemD.RequestedItemIdentification = in.RequestedItemIdentification
	despatchAdviceLineItemD.Specification = in.Specification
	despatchAdviceLineItemD.DespatchAdviceId = in.DespatchAdviceId

	despatchAdviceLineItemT := logisticsproto.DespatchAdviceLineItemT{}
	despatchAdviceLineItemT.FirstInFirstOutDateTime = common.TimeToTimestamp(firstInFirstOutDateTime.UTC().Truncate(time.Second))
	despatchAdviceLineItemT.PickUpDateTime = common.TimeToTimestamp(pickUpDateTime.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	despatchAdviceLineItem := logisticsproto.DespatchAdviceLineItem{DespatchAdviceLineItemD: &despatchAdviceLineItemD, DespatchAdviceLineItemT: &despatchAdviceLineItemT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &despatchAdviceLineItem, nil
}

// insertDespatchAdviceLineItem - Insert DespatchAdviceLineItem into database
func (das *DespatchAdviceService) insertDespatchAdviceLineItem(ctx context.Context, insertDespatchAdviceLineItemSQL string, despatchAdviceLineItem *logisticsproto.DespatchAdviceLineItem, userEmail string, requestID string) error {
	despatchAdviceLineItemTmp, err := das.crDespatchAdviceLineItemStruct(ctx, despatchAdviceLineItem, userEmail, requestID)
	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = das.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertDespatchAdviceLineItemSQL, despatchAdviceLineItemTmp)
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		despatchAdviceLineItem.DespatchAdviceLineItemD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(despatchAdviceLineItem.DespatchAdviceLineItemD.Uuid4)
		if err != nil {
			das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		despatchAdviceLineItem.DespatchAdviceLineItemD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		das.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crDespatchAdviceLineItemStruct - process DespatchAdviceLineItem details
func (das *DespatchAdviceService) crDespatchAdviceLineItemStruct(ctx context.Context, despatchAdviceLineItem *logisticsproto.DespatchAdviceLineItem, userEmail string, requestID string) (*logisticsstruct.DespatchAdviceLineItem, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(despatchAdviceLineItem.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(despatchAdviceLineItem.CrUpdTime.UpdatedAt)

	despatchAdviceLineItemT := new(logisticsstruct.DespatchAdviceLineItemT)
	despatchAdviceLineItemT.FirstInFirstOutDateTime = common.TimestampToTime(despatchAdviceLineItem.DespatchAdviceLineItemT.FirstInFirstOutDateTime)
	despatchAdviceLineItemT.PickUpDateTime = common.TimestampToTime(despatchAdviceLineItem.DespatchAdviceLineItemT.PickUpDateTime)

	despatchAdviceLineItemTmp := logisticsstruct.DespatchAdviceLineItem{DespatchAdviceLineItemD: despatchAdviceLineItem.DespatchAdviceLineItemD, DespatchAdviceLineItemT: despatchAdviceLineItemT, CrUpdUser: despatchAdviceLineItem.CrUpdUser, CrUpdTime: crUpdTime}

	return &despatchAdviceLineItemTmp, nil
}

// GetDespatchAdviceLineItems - GetDespatchAdviceLineItems
func (das *DespatchAdviceService) GetDespatchAdviceLineItems(ctx context.Context, inReq *logisticsproto.GetDespatchAdviceLineItemsRequest) (*logisticsproto.GetDespatchAdviceLineItemsResponse, error) {
	in := inReq.GetRequest
	getRequest := commonproto.GetRequest{}
	getRequest.Id = in.Id
	getRequest.UserEmail = in.UserEmail
	getRequest.RequestId = in.RequestId
	form := logisticsproto.GetDespatchAdviceRequest{}
	form.GetRequest = &getRequest

	despatchAdviceResponse, err := das.GetDespatchAdvice(ctx, &form)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	despatchAdvice := despatchAdviceResponse.DespatchAdvice
	despatchAdviceLineItems := []*logisticsproto.DespatchAdviceLineItem{}

	nselectDespatchAdviceLineItemsSQL := selectDespatchAdviceLineItemsSQL + ` where despatch_advice_id = ?;`
	rows, err := das.DBService.DB.QueryxContext(ctx, nselectDespatchAdviceLineItemsSQL, despatchAdvice.DespatchAdviceD.Id)
	if err != nil {
		das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}
	for rows.Next() {

		despatchAdviceLineItemTmp := logisticsstruct.DespatchAdviceLineItem{}
		err = rows.StructScan(&despatchAdviceLineItemTmp)
		if err != nil {
			das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		despatchAdviceLineItemT := logisticsproto.DespatchAdviceLineItemT{}
		despatchAdviceLineItemT.FirstInFirstOutDateTime = common.TimeToTimestamp(despatchAdviceLineItemTmp.DespatchAdviceLineItemT.FirstInFirstOutDateTime)
		despatchAdviceLineItemT.PickUpDateTime = common.TimeToTimestamp(despatchAdviceLineItemTmp.DespatchAdviceLineItemT.PickUpDateTime)

		crUpdTime := new(commonproto.CrUpdTime)
		crUpdTime.CreatedAt = common.TimeToTimestamp(despatchAdviceLineItemTmp.CrUpdTime.CreatedAt)
		crUpdTime.UpdatedAt = common.TimeToTimestamp(despatchAdviceLineItemTmp.CrUpdTime.UpdatedAt)

		uuid4Str, err := common.UUIDBytesToStr(despatchAdviceLineItemTmp.DespatchAdviceLineItemD.Uuid4)
		if err != nil {
			das.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
			return nil, err
		}

		despatchAdviceLineItemTmp.DespatchAdviceLineItemD.IdS = uuid4Str

		despatchAdviceLineItem := logisticsproto.DespatchAdviceLineItem{DespatchAdviceLineItemD: despatchAdviceLineItemTmp.DespatchAdviceLineItemD, DespatchAdviceLineItemT: &despatchAdviceLineItemT, CrUpdUser: despatchAdviceLineItemTmp.CrUpdUser, CrUpdTime: crUpdTime}

		despatchAdviceLineItems = append(despatchAdviceLineItems, &despatchAdviceLineItem)
	}
	despatchAdviceLineItemsResponse := logisticsproto.GetDespatchAdviceLineItemsResponse{}
	despatchAdviceLineItemsResponse.DespatchAdviceLineItems = despatchAdviceLineItems
	return &despatchAdviceLineItemsResponse, nil
}
