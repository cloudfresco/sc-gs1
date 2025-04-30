package meatservices

import (
	"context"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

const insertFishCatchOrProductionSQL = `insert into fish_catch_or_productions
	  (
catch_area,
fishing_gear_type_code,
production_method_for_fish_and_seafood_code
)
values(
:catch_area,
:fishing_gear_type_code,
:production_method_for_fish_and_seafood_code);`

/*const selectFishCatchOrProductionsSQL = `select
  id,
  catch_area,
  fishing_gear_type_code,
  production_method_for_fish_and_seafood_code fish_catch_or_productions`*/

func (fis *FishService) CreateFishCatchOrProduction(ctx context.Context, in *meatproto.CreateFishCatchOrProductionRequest) (*meatproto.CreateFishCatchOrProductionResponse, error) {
	fishCatchOrProduction, err := fis.ProcessFishCatchOrProductionRequest(ctx, in)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = fis.insertFishCatchOrProduction(ctx, insertFishCatchOrProductionSQL, fishCatchOrProduction, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishCatchOrProductionResponse := meatproto.CreateFishCatchOrProductionResponse{}
	fishCatchOrProductionResponse.FishCatchOrProduction = fishCatchOrProduction
	return &fishCatchOrProductionResponse, nil
}

// ProcessFishCatchOrProductionRequest - ProcessFishCatchOrProductionRequest
func (fis *FishService) ProcessFishCatchOrProductionRequest(ctx context.Context, in *meatproto.CreateFishCatchOrProductionRequest) (*meatproto.FishCatchOrProduction, error) {
	fishCatchOrProduction := meatproto.FishCatchOrProduction{}
	fishCatchOrProduction.CatchArea = in.CatchArea
	fishCatchOrProduction.FishingGearTypeCode = in.FishingGearTypeCode
	fishCatchOrProduction.ProductionMethodForFishAndSeafoodCode = in.ProductionMethodForFishAndSeafoodCode

	return &fishCatchOrProduction, nil
}

// insertFishCatchOrProduction - Insert FishCatchOrProduction into database
func (fis *FishService) insertFishCatchOrProduction(ctx context.Context, insertFishCatchOrProductionSQL string, fishCatchOrProduction *meatproto.FishCatchOrProduction, userEmail string, requestID string) error {
	err := fis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertFishCatchOrProductionSQL, fishCatchOrProduction)
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		fishCatchOrProduction.Id = uint32(uID)
		return nil
	})
	if err != nil {
		fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}
