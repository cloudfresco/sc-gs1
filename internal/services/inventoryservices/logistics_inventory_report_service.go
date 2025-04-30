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

const insertLogisticsInventoryReportSQL = `insert into logistics_inventory_reports
	    (uuid4,
	    structure_type_code,
type_of_service_transaction,
inventory_reporting_party,
inventory_report_to_party,
logistics_inventory_report_identification,
logistics_inventory_report_request,
reporting_period_begin,
reporting_period_end,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
  values(:uuid4,
 :structure_type_code,
:type_of_service_transaction,
:inventory_reporting_party,
:inventory_report_to_party,
:logistics_inventory_report_identification,
:logistics_inventory_report_request,
:reporting_period_begin,
:reporting_period_end,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectLogisticsInventoryReportsSQL = `select
  id,
  uuid4,
  structure_type_code,
  type_of_service_transaction,
  inventory_reporting_party,
  inventory_report_to_party,
  logistics_inventory_report_identification,
  logistics_inventory_report_request,
  reporting_period_begin,
  reporting_period_end,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from logistics_inventory_reports`*/

// CreateLogisticsInventoryReport - Create LogisticsInventoryReport
func (invs *InventoryService) CreateLogisticsInventoryReport(ctx context.Context, in *inventoryproto.CreateLogisticsInventoryReportRequest) (*inventoryproto.CreateLogisticsInventoryReportResponse, error) {
	logisticsInventoryReport, err := invs.ProcessLogisticsInventoryReportRequest(ctx, in)
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = invs.insertLogisticsInventoryReport(ctx, insertLogisticsInventoryReportSQL, logisticsInventoryReport, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticsInventoryReportResponse := inventoryproto.CreateLogisticsInventoryReportResponse{}
	logisticsInventoryReportResponse.LogisticsInventoryReport = logisticsInventoryReport
	return &logisticsInventoryReportResponse, nil
}

// ProcessLogisticsInventoryReportRequest - ProcessLogisticsInventoryReportRequest
func (invs *InventoryService) ProcessLogisticsInventoryReportRequest(ctx context.Context, in *inventoryproto.CreateLogisticsInventoryReportRequest) (*inventoryproto.LogisticsInventoryReport, error) {
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

	logisticsInventoryReportD := inventoryproto.LogisticsInventoryReportD{}
	logisticsInventoryReportD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		invs.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	logisticsInventoryReportD.StructureTypeCode = in.StructureTypeCode
	logisticsInventoryReportD.TypeOfServiceTransaction = in.TypeOfServiceTransaction
	logisticsInventoryReportD.InventoryReportingParty = in.InventoryReportingParty
	logisticsInventoryReportD.InventoryReportToParty = in.InventoryReportToParty
	logisticsInventoryReportD.LogisticsInventoryReportIdentification = in.LogisticsInventoryReportIdentification
	logisticsInventoryReportD.LogisticsInventoryReportRequest = in.LogisticsInventoryReportRequest

	logisticsInventoryReportT := inventoryproto.LogisticsInventoryReportT{}
	logisticsInventoryReportT.ReportingPeriodBegin = common.TimeToTimestamp(reportingPeriodBegin.UTC().Truncate(time.Second))
	logisticsInventoryReportT.ReportingPeriodEnd = common.TimeToTimestamp(reportingPeriodEnd.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	logisticsInventoryReport := inventoryproto.LogisticsInventoryReport{LogisticsInventoryReportD: &logisticsInventoryReportD, LogisticsInventoryReportT: &logisticsInventoryReportT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &logisticsInventoryReport, nil
}

// insertLogisticsInventoryReport - Insert LogisticsInventoryReport details into database
func (invs *InventoryService) insertLogisticsInventoryReport(ctx context.Context, insertLogisticsInventoryReportSQL string, logisticsInventoryReport *inventoryproto.LogisticsInventoryReport, userEmail string, requestID string) error {
	logisticsInventoryReportTmp, err := invs.crLogisticsInventoryReportStruct(ctx, logisticsInventoryReport, userEmail, requestID)
	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}

	err = invs.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertLogisticsInventoryReportSQL, logisticsInventoryReportTmp)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticsInventoryReport.LogisticsInventoryReportD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(logisticsInventoryReport.LogisticsInventoryReportD.Uuid4)
		if err != nil {
			invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		logisticsInventoryReport.LogisticsInventoryReportD.IdS = uuid4Str

		return nil
	})

	if err != nil {
		invs.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crLogisticsInventoryReportStruct - process LogisticsInventoryReport details
func (invs *InventoryService) crLogisticsInventoryReportStruct(ctx context.Context, logisticsInventoryReport *inventoryproto.LogisticsInventoryReport, userEmail string, requestID string) (*inventorystruct.LogisticsInventoryReport, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(logisticsInventoryReport.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(logisticsInventoryReport.CrUpdTime.UpdatedAt)

	logisticsInventoryReportT := new(inventorystruct.LogisticsInventoryReportT)
	logisticsInventoryReportT.ReportingPeriodBegin = common.TimestampToTime(logisticsInventoryReport.LogisticsInventoryReportT.ReportingPeriodBegin)
	logisticsInventoryReportT.ReportingPeriodEnd = common.TimestampToTime(logisticsInventoryReport.LogisticsInventoryReportT.ReportingPeriodEnd)

	logisticsInventoryReportTmp := inventorystruct.LogisticsInventoryReport{LogisticsInventoryReportD: logisticsInventoryReport.LogisticsInventoryReportD, LogisticsInventoryReportT: logisticsInventoryReportT, CrUpdUser: logisticsInventoryReport.CrUpdUser, CrUpdTime: crUpdTime}

	return &logisticsInventoryReportTmp, nil
}
