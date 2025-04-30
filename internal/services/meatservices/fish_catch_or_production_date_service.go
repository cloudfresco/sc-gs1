package meatservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	meatstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertFishCatchOrProductionDateSQL = `insert into fish_catch_or_production_dates
	  (
catch_end_date,
catch_start_date,
first_freeze_date,
catch_date_time
)
values(
:catch_end_date,
:catch_start_date,
:first_freeze_date,
:catch_date_time);`

/*const selectFishCatchOrProductionDatesSQL = `select
  id,
  catch_end_date,
  catch_start_date,
  first_freeze_date,
  catch_date_time from fish_catch_or_production_dates`*/

func (fis *FishService) CreateFishCatchOrProductionDate(ctx context.Context, in *meatproto.CreateFishCatchOrProductionDateRequest) (*meatproto.CreateFishCatchOrProductionDateResponse, error) {
	fishCatchOrProductionDate, err := fis.ProcessFishCatchOrProductionDateRequest(ctx, in)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = fis.insertFishCatchOrProductionDate(ctx, insertFishCatchOrProductionDateSQL, fishCatchOrProductionDate, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishCatchOrProductionDateResponse := meatproto.CreateFishCatchOrProductionDateResponse{}
	fishCatchOrProductionDateResponse.FishCatchOrProductionDate = fishCatchOrProductionDate
	return &fishCatchOrProductionDateResponse, nil
}

// ProcessFishCatchOrProductionDateRequest - ProcessFishCatchOrProductionDateRequest
func (fis *FishService) ProcessFishCatchOrProductionDateRequest(ctx context.Context, in *meatproto.CreateFishCatchOrProductionDateRequest) (*meatproto.FishCatchOrProductionDate, error) {
	var err error

	catchEndDate, err := time.Parse(common.Layout, in.CatchEndDate)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	catchStartDate, err := time.Parse(common.Layout, in.CatchStartDate)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	firstFreezeDate, err := time.Parse(common.Layout, in.FirstFreezeDate)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	catchDateTime, err := time.Parse(common.Layout, in.CatchDateTime)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishCatchOrProductionDateD := meatproto.FishCatchOrProductionDateD{}

	fishCatchOrProductionDateT := meatproto.FishCatchOrProductionDateT{}
	fishCatchOrProductionDateT.CatchEndDate = common.TimeToTimestamp(catchEndDate.UTC().Truncate(time.Second))
	fishCatchOrProductionDateT.CatchStartDate = common.TimeToTimestamp(catchStartDate.UTC().Truncate(time.Second))
	fishCatchOrProductionDateT.FirstFreezeDate = common.TimeToTimestamp(firstFreezeDate.UTC().Truncate(time.Second))
	fishCatchOrProductionDateT.CatchDateTime = common.TimeToTimestamp(catchDateTime.UTC().Truncate(time.Second))

	fishCatchOrProductionDate := meatproto.FishCatchOrProductionDate{FishCatchOrProductionDateD: &fishCatchOrProductionDateD, FishCatchOrProductionDateT: &fishCatchOrProductionDateT}

	return &fishCatchOrProductionDate, nil
}

// insertFishCatchOrProductionDate - Insert FishCatchOrProductionDate into database
func (fis *FishService) insertFishCatchOrProductionDate(ctx context.Context, insertFishCatchOrProductionDateSQL string, fishCatchOrProductionDate *meatproto.FishCatchOrProductionDate, userEmail string, requestID string) error {
	fishCatchOrProductionDateTmp, err := fis.crFishCatchOrProductionDateStruct(ctx, fishCatchOrProductionDate, userEmail, requestID)
	if err != nil {
		fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = fis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertFishCatchOrProductionDateSQL, fishCatchOrProductionDateTmp)
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		fishCatchOrProductionDate.FishCatchOrProductionDateD.Id = uint32(uID)
		return nil
	})

	if err != nil {
		fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crFishCatchOrProductionDateStruct - process FishCatchOrProductionDate details
func (fis *FishService) crFishCatchOrProductionDateStruct(ctx context.Context, fishCatchOrProductionDate *meatproto.FishCatchOrProductionDate, userEmail string, requestID string) (*meatstruct.FishCatchOrProductionDate, error) {
	fishCatchOrProductionDateT := new(meatstruct.FishCatchOrProductionDateT)
	fishCatchOrProductionDateT.CatchEndDate = common.TimestampToTime(fishCatchOrProductionDate.FishCatchOrProductionDateT.CatchEndDate)
	fishCatchOrProductionDateT.CatchStartDate = common.TimestampToTime(fishCatchOrProductionDate.FishCatchOrProductionDateT.CatchStartDate)
	fishCatchOrProductionDateT.FirstFreezeDate = common.TimestampToTime(fishCatchOrProductionDate.FishCatchOrProductionDateT.FirstFreezeDate)
	fishCatchOrProductionDateT.CatchDateTime = common.TimestampToTime(fishCatchOrProductionDate.FishCatchOrProductionDateT.CatchDateTime)

	fishCatchOrProductionDateTmp := meatstruct.FishCatchOrProductionDate{FishCatchOrProductionDateD: fishCatchOrProductionDate.FishCatchOrProductionDateD, FishCatchOrProductionDateT: fishCatchOrProductionDateT}

	return &fishCatchOrProductionDateTmp, nil
}
