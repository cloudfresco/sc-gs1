package inventoryservices

import (
	"context"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertConsumptionReportSQL = `insert into consumption_reports
	    (uuid4,
	    buyer,
      consumption_report_identification,
      seller,
      status_code,
      created_by_user_id,
      updated_by_user_id,
      created_at,
      updated_at)
  values(
  :uuid4,
  :buyer,
  :consumption_report_identification,
  :seller,
  :status_code,
  :created_by_user_id,
  :updated_by_user_id,
  :created_at,
  :updated_at);`

/*const selectConsumptionReportsSQL = `select
  id,
  uuid4,
  buyer,
  consumption_report_identification,
  seller,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from consumption_reports`*/

// CreateConsumptionReport - Create ConsumptionReport
func (invs *InventoryService) CreateConsumptionReport(ctx context.Context, in *inventoryproto.CreateConsumptionReportRequest) (*inventoryproto.CreateConsumptionReportResponse, error) {
	consumptionReport, err := invs.ProcessConsumptionReportRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertConsumptionReport(ctx, insertConsumptionReportSQL, consumptionReport, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportResponse := inventoryproto.CreateConsumptionReportResponse{}
	consumptionReportResponse.ConsumptionReport = consumptionReport
	return &consumptionReportResponse, nil
}

// ProcessConsumptionReportRequest - ProcessConsumptionReportRequest
func (invs *InventoryService) ProcessConsumptionReportRequest(ctx context.Context, in *inventoryproto.CreateConsumptionReportRequest) (*inventoryproto.ConsumptionReport, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	consumptionReportD := inventoryproto.ConsumptionReportD{}
	consumptionReportD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	consumptionReportD.Buyer = in.Buyer
	consumptionReportD.ConsumptionReportIdentification = in.ConsumptionReportIdentification
	consumptionReportD.Seller = in.Seller

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	consumptionReport := inventoryproto.ConsumptionReport{ConsumptionReportD: &consumptionReportD, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}
	return &consumptionReport, nil
}

// insertConsumptionReport - Insert ConsumptionReport details into database
func (invs *InventoryService) insertConsumptionReport(ctx context.Context, insertConsumptionReportSQL string, consumptionReport *inventoryproto.ConsumptionReport, userEmail string, requestID string) error {
	consumptionReportTmp, err := invs.crConsumptionReportStruct(ctx, consumptionReport, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertConsumptionReportSQL, consumptionReportTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consumptionReport.ConsumptionReportD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(consumptionReport.ConsumptionReportD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		consumptionReport.ConsumptionReportD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crConsumptionReportStruct - process ConsumptionReport details
func (invs *InventoryService) crConsumptionReportStruct(ctx context.Context, consumptionReport *inventoryproto.ConsumptionReport, userEmail string, requestID string) (*inventorystruct.ConsumptionReport, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(consumptionReport.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(consumptionReport.CrUpdTime.UpdatedAt)

	consumptionReportTmp := inventorystruct.ConsumptionReport{ConsumptionReportD: consumptionReport.ConsumptionReportD, CrUpdUser: consumptionReport.CrUpdUser, CrUpdTime: crUpdTime}

	return &consumptionReportTmp, nil
}
