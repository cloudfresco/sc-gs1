package inventoryservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	inventoryproto "github.com/cloudfresco/sc-gs1/internal/protogen/inventory/v1"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	inventorystruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/inventory/v1"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertInventoryReportSQL = `insert into inventory_reports
	    (uuid4,
	    inventory_report_type_code,
      structure_type_code,
      inventory_report_identification,
      inventory_reporting_party,
      inventory_report_to_party,
      reporting_period_begin,
      reporting_period_end,
      status_code,
      created_by_user_id,
      updated_by_user_id,
      created_at,
      updated_at)
  values(:uuid4,
        :inventory_report_type_code,
        :structure_type_code,
        :inventory_report_identification,
        :inventory_reporting_party,
        :inventory_report_to_party,
        :reporting_period_begin,
        :reporting_period_end,
        :status_code,
        :created_by_user_id,
        :updated_by_user_id,
        :created_at,
        :updated_at);`

/*const selectInventoryReportsSQL = `select
  id,
  uuid4,
  inventory_report_type_code,
  structure_type_code,
  inventory_report_identification,
  inventory_reporting_party,
  inventory_report_to_party,
  reporting_period_begin,
  reporting_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from inventory_reports`*/

// CreateInventoryReport - Create InventoryReport
func (invs *InventoryService) CreateInventoryReport(ctx context.Context, in *inventoryproto.CreateInventoryReportRequest) (*inventoryproto.CreateInventoryReportResponse, error) {
	inventoryReport, err := invs.ProcessInventoryReportRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertInventoryReport(ctx, insertInventoryReportSQL, inventoryReport, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryReportResponse := inventoryproto.CreateInventoryReportResponse{}
	inventoryReportResponse.InventoryReport = inventoryReport
	return &inventoryReportResponse, nil
}

// ProcessInventoryReportRequest - ProcessInventoryReportRequest
func (invs *InventoryService) ProcessInventoryReportRequest(ctx context.Context, in *inventoryproto.CreateInventoryReportRequest) (*inventoryproto.InventoryReport, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, invs.UserServiceClient)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	reportingPeriodBegin, err := time.Parse(common.Layout, in.ReportingPeriodBegin)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	reportingPeriodEnd, err := time.Parse(common.Layout, in.ReportingPeriodEnd)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	inventoryReportD := inventoryproto.InventoryReportD{}
	inventoryReportD.InventoryReportTypeCode = in.InventoryReportTypeCode
	inventoryReportD.StructureTypeCode = in.StructureTypeCode
	inventoryReportD.InventoryReportIdentification = in.InventoryReportIdentification
	inventoryReportD.InventoryReportingParty = in.InventoryReportingParty
	inventoryReportD.InventoryReportToParty = in.InventoryReportToParty

	inventoryReportT := inventoryproto.InventoryReportT{}
	inventoryReportT.ReportingPeriodBegin = common.TimeToTimestamp(reportingPeriodBegin.UTC().Truncate(time.Second))
	inventoryReportT.ReportingPeriodEnd = common.TimeToTimestamp(reportingPeriodEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	inventoryReport := inventoryproto.InventoryReport{InventoryReportD: &inventoryReportD, InventoryReportT: &inventoryReportT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &inventoryReport, nil
}

// insertInventoryReport - Insert InventoryReport details into database
func (invs *InventoryService) insertInventoryReport(ctx context.Context, insertInventoryReportSQL string, inventoryReport *inventoryproto.InventoryReport, userEmail string, requestID string) error {
	inventoryReportTmp, err := invs.crInventoryReportStruct(ctx, inventoryReport, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertInventoryReportSQL, inventoryReportTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryReport.InventoryReportD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(inventoryReport.InventoryReportD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		inventoryReport.InventoryReportD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crInventoryReportStruct - process InventoryReport details
func (invs *InventoryService) crInventoryReportStruct(ctx context.Context, inventoryReport *inventoryproto.InventoryReport, userEmail string, requestID string) (*inventorystruct.InventoryReport, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(inventoryReport.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(inventoryReport.CrUpdTime.UpdatedAt)

	inventoryReportT := new(inventorystruct.InventoryReportT)
	inventoryReportT.ReportingPeriodBegin = common.TimestampToTime(inventoryReport.InventoryReportT.ReportingPeriodBegin)
	inventoryReportT.ReportingPeriodEnd = common.TimestampToTime(inventoryReport.InventoryReportT.ReportingPeriodEnd)

	inventoryReportTmp := inventorystruct.InventoryReport{InventoryReportD: inventoryReport.InventoryReportD, InventoryReportT: inventoryReportT, CrUpdUser: inventoryReport.CrUpdUser, CrUpdTime: crUpdTime}

	return &inventoryReportTmp, nil
}
