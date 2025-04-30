package inventoryservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertTransactionalItemDataSQL = `insert into transactional_item_data
	    (uuid4,
batch_number,
country_of_origin,
item_in_contact_with_food_product,
lot_number,
product_quality_indication,
serial_number,
shelf_life,
trade_item_quantity,
available_for_sale_date,
best_before_date,
item_expiration_date,
packaging_date,
production_date,
sell_by_date,
trade_item_inventory_status_id,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
:batch_number,
:country_of_origin,
:item_in_contact_with_food_product,
:lot_number,
:product_quality_indication,
:serial_number,
:shelf_life,
:trade_item_quantity,
:available_for_sale_date,
:best_before_date,
:item_expiration_date,
:packaging_date,
:production_date,
:sell_by_date,
:trade_item_inventory_status_id,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectTransactionalItemDatasSQL = `select
  id,
  uuid4,
  batch_number,
  country_of_origin,
  item_in_contact_with_food_product,
  lot_number,
  product_quality_indication,
  serial_number,
  shelf_life,
  trade_item_quantity,
  available_for_sale_date,
  best_before_date,
  item_expiration_date,
  packaging_date,
  production_date,
  sell_by_date,
  trade_item_inventory_status_id,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from transactional_item_data`*/

// CreateTransactionalItemData - Create TransactionalItemData
func (invs *InventoryService) CreateTransactionalItemData(ctx context.Context, in *inventoryproto.CreateTransactionalItemDataRequest) (*inventoryproto.CreateTransactionalItemDataResponse, error) {
	transactionalItemData, err := invs.ProcessTransactionalItemDataRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertTransactionalItemData(ctx, insertTransactionalItemDataSQL, transactionalItemData, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemDataResponse := inventoryproto.CreateTransactionalItemDataResponse{}
	transactionalItemDataResponse.TransactionalItemData = transactionalItemData
	return &transactionalItemDataResponse, nil
}

// ProcessTransactionalItemDataRequest - ProcessTransactionalItemDataRequest
func (invs *InventoryService) ProcessTransactionalItemDataRequest(ctx context.Context, in *inventoryproto.CreateTransactionalItemDataRequest) (*inventoryproto.TransactionalItemData, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	availableForSaleDate, err := time.Parse(common.Layout, in.AvailableForSaleDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	bestBeforeDate, err := time.Parse(common.Layout, in.BestBeforeDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	itemExpirationDate, err := time.Parse(common.Layout, in.ItemExpirationDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	packagingDate, err := time.Parse(common.Layout, in.PackagingDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	productionDate, err := time.Parse(common.Layout, in.ProductionDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	sellByDate, err := time.Parse(common.Layout, in.SellByDate)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	transactionalItemDataD := inventoryproto.TransactionalItemDataD{}
	transactionalItemDataD.BatchNumber = in.BatchNumber
	transactionalItemDataD.CountryOfOrigin = in.CountryOfOrigin
	transactionalItemDataD.ItemInContactWithFoodProduct = in.ItemInContactWithFoodProduct
	transactionalItemDataD.LotNumber = in.LotNumber
	transactionalItemDataD.ProductQualityIndication = in.ProductQualityIndication
	transactionalItemDataD.SerialNumber = in.SerialNumber
	transactionalItemDataD.ShelfLife = in.ShelfLife
	transactionalItemDataD.TradeItemQuantity = in.TradeItemQuantity
	transactionalItemDataD.TradeItemInventoryStatusId = in.TradeItemInventoryStatusId

	transactionalItemDataT := inventoryproto.TransactionalItemDataT{}
	transactionalItemDataT.AvailableForSaleDate = common.TimeToTimestamp(availableForSaleDate.UTC().Truncate(time.Second))
	transactionalItemDataT.BestBeforeDate = common.TimeToTimestamp(bestBeforeDate.UTC().Truncate(time.Second))
	transactionalItemDataT.ItemExpirationDate = common.TimeToTimestamp(itemExpirationDate.UTC().Truncate(time.Second))
	transactionalItemDataT.PackagingDate = common.TimeToTimestamp(packagingDate.UTC().Truncate(time.Second))
	transactionalItemDataT.ProductionDate = common.TimeToTimestamp(productionDate.UTC().Truncate(time.Second))
	transactionalItemDataT.SellByDate = common.TimeToTimestamp(sellByDate.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	transactionalItemData := inventoryproto.TransactionalItemData{TransactionalItemDataD: &transactionalItemDataD, TransactionalItemDataT: &transactionalItemDataT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &transactionalItemData, nil
}

// insertTransactionalItemData - Insert TransactionalItemData details into database
func (invs *InventoryService) insertTransactionalItemData(ctx context.Context, insertTransactionalItemDataSQL string, transactionalItemData *inventoryproto.TransactionalItemData, userEmail string, requestID string) error {
	transactionalItemDataTmp, err := invs.crTransactionalItemDataStruct(ctx, transactionalItemData, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertTransactionalItemDataSQL, transactionalItemDataTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemData.TransactionalItemDataD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(transactionalItemData.TransactionalItemDataD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		transactionalItemData.TransactionalItemDataD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crTransactionalItemDataStruct - process TransactionalItemData details
func (invs *InventoryService) crTransactionalItemDataStruct(ctx context.Context, transactionalItemData *inventoryproto.TransactionalItemData, userEmail string, requestID string) (*inventorystruct.TransactionalItemData, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(transactionalItemData.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(transactionalItemData.CrUpdTime.UpdatedAt)

	transactionalItemDataT := new(inventorystruct.TransactionalItemDataT)
	transactionalItemDataT.AvailableForSaleDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.AvailableForSaleDate)
	transactionalItemDataT.BestBeforeDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.BestBeforeDate)
	transactionalItemDataT.ItemExpirationDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.ItemExpirationDate)
	transactionalItemDataT.PackagingDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.PackagingDate)
	transactionalItemDataT.ProductionDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.ProductionDate)
	transactionalItemDataT.SellByDate = common.TimestampToTime(transactionalItemData.TransactionalItemDataT.SellByDate)

	transactionalItemDataTmp := inventorystruct.TransactionalItemData{TransactionalItemDataD: transactionalItemData.TransactionalItemDataD, TransactionalItemDataT: transactionalItemDataT, CrUpdUser: transactionalItemData.CrUpdUser, CrUpdTime: crUpdTime}

	return &transactionalItemDataTmp, nil
}
