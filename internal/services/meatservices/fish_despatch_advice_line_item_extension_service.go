package meatservices

import (
	"context"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/services/partyservices"

	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	commonstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/common/v1"
	meatstruct "github.com/cloudfresco/sc-gs1/internal/servicestructs/meat/v1"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

// FishService - For accessing FishDespatchAdviceLineItemExtension services
type FishService struct {
	log               *zap.Logger
	DBService         *common.DBService
	RedisService      *common.RedisService
	UserServiceClient partyproto.UserServiceClient
	meatproto.UnimplementedFishServiceServer
}

// NewFishService - Create FishDespatchAdviceLineItemExtension service
func NewFishService(log *zap.Logger, dbOpt *common.DBService, redisOpt *common.RedisService, userServiceClient partyproto.UserServiceClient) *FishService {
	return &FishService{
		log:               log,
		DBService:         dbOpt,
		RedisService:      redisOpt,
		UserServiceClient: userServiceClient,
	}
}

const insertFishDespatchAdviceLineItemExtensionSQL = `insert into fish_despatch_advice_line_item_extensions
	  (
uuid4,
aquatic_species_code,
aquatic_species_name,
fish_presentation_code,
fp_code_list_agency_name,
fp_code_list_version,
fp_code_list_name,
fish_size_code,
fs_code_list_agency_name,
fs_code_list_version,
fs_code_list_name,
quality_grade_code,
qg_code_list_agency_name,
qg_code_list_version,
qg_code_list_name,
storage_state_code,
aqua_culture_production_unit,
fishing_vessel,
place_of_slaughter,
port_of_landing,
fish_catch_or_production_date_id,
fish_catch_or_production_id,
date_of_landing,
date_of_slaughter,
status_code,
created_by_user_id,
updated_by_user_id,
created_at,
updated_at)
values(
:uuid4,
:aquatic_species_code,
:aquatic_species_name,
:fish_presentation_code,
:fp_code_list_agency_name,
:fp_code_list_version,
:fp_code_list_name,
:fish_size_code,
:fs_code_list_agency_name,
:fs_code_list_version,
:fs_code_list_name,
:quality_grade_code,
:qg_code_list_agency_name,
:qg_code_list_version,
:qg_code_list_name,
:storage_state_code,
:aqua_culture_production_unit,
:fishing_vessel,
:place_of_slaughter,
:port_of_landing,
:fish_catch_or_production_date_id,
:fish_catch_or_production_id,
:date_of_landing,
:date_of_slaughter,
:status_code,
:created_by_user_id,
:updated_by_user_id,
:created_at,
:updated_at);`

/*const selectFishDespatchAdviceLineItemExtensionsSQL = `select
  id,
  uuid4,
  aquatic_species_code,
  aquatic_species_name,
  fish_presentation_code,
  fp_code_list_agency_name,
  fp_code_list_version,
  fp_code_list_name,
  fish_size_code,
  fs_code_list_agency_name,
  fs_code_list_version,
  fs_code_list_name,
  quality_grade_code,
  qg_code_list_agency_name,
  qg_code_list_version,
  qg_code_list_name,
  storage_state_code,
  aqua_culture_production_unit,
  fishing_vessel,
  place_of_slaughter,
  port_of_landing,
  fish_catch_or_production_date_id,
  fish_catch_or_production_id,
  date_of_landing,
  date_of_slaughter,
  status_code,
  created_by_user_id,
  updated_by_user_id,
  created_at,
  updated_at from fish_despatch_advice_line_item_extensions`*/

func (fis *FishService) CreateFishDespatchAdviceLineItemExtension(ctx context.Context, in *meatproto.CreateFishDespatchAdviceLineItemExtensionRequest) (*meatproto.CreateFishDespatchAdviceLineItemExtensionResponse, error) {
	fishDespatchAdviceLineItemExtension, err := fis.ProcessFishDespatchAdviceLineItemExtensionRequest(ctx, in)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	err = fis.insertFishDespatchAdviceLineItemExtension(ctx, insertFishDespatchAdviceLineItemExtensionSQL, fishDespatchAdviceLineItemExtension, in.GetUserEmail(), in.GetRequestId())
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishDespatchAdviceLineItemExtensionResponse := meatproto.CreateFishDespatchAdviceLineItemExtensionResponse{}
	fishDespatchAdviceLineItemExtensionResponse.FishDespatchAdviceLineItemExtension = fishDespatchAdviceLineItemExtension
	return &fishDespatchAdviceLineItemExtensionResponse, nil
}

// ProcessFishDespatchAdviceLineItemExtensionRequest - ProcessFishDespatchAdviceLineItemExtensionRequest
func (fis *FishService) ProcessFishDespatchAdviceLineItemExtensionRequest(ctx context.Context, in *meatproto.CreateFishDespatchAdviceLineItemExtensionRequest) (*meatproto.FishDespatchAdviceLineItemExtension, error) {
	user, err := partyservices.GetUserWithNewContext(ctx, in.UserId, in.UserEmail, in.RequestId, fis.UserServiceClient)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	ttime := common.GetTimeDetails()
	tn := common.TimeToTimestamp(ttime)

	dateOfLanding, err := time.Parse(common.Layout, in.DateOfLanding)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	dateOfSlaughter, err := time.Parse(common.Layout, in.DateOfSlaughter)
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishDespatchAdviceLineItemExtensionD := meatproto.FishDespatchAdviceLineItemExtensionD{}
	fishDespatchAdviceLineItemExtensionD.Uuid4, err = common.GetUUIDBytes()
	if err != nil {
		fis.log.Error("Error", zap.String("user", in.GetUserEmail()), zap.String("reqid", in.GetRequestId()), zap.Error(err))
		return nil, err
	}

	fishDespatchAdviceLineItemExtensionD.AquaticSpeciesCode = in.AquaticSpeciesCode
	fishDespatchAdviceLineItemExtensionD.AquaticSpeciesName = in.AquaticSpeciesName
	fishDespatchAdviceLineItemExtensionD.FishPresentationCode = in.FishPresentationCode
	fishDespatchAdviceLineItemExtensionD.FpCodeListAgencyName = in.FpCodeListAgencyName
	fishDespatchAdviceLineItemExtensionD.FpCodeListVersion = in.FpCodeListVersion
	fishDespatchAdviceLineItemExtensionD.FpCodeListName = in.FpCodeListName
	fishDespatchAdviceLineItemExtensionD.FishSizeCode = in.FishSizeCode
	fishDespatchAdviceLineItemExtensionD.FsCodeListAgencyName = in.FsCodeListAgencyName
	fishDespatchAdviceLineItemExtensionD.FsCodeListVersion = in.FsCodeListVersion
	fishDespatchAdviceLineItemExtensionD.FsCodeListName = in.FsCodeListName
	fishDespatchAdviceLineItemExtensionD.QualityGradeCode = in.QualityGradeCode
	fishDespatchAdviceLineItemExtensionD.QgCodeListAgencyName = in.QgCodeListAgencyName
	fishDespatchAdviceLineItemExtensionD.QgCodeListVersion = in.QgCodeListVersion
	fishDespatchAdviceLineItemExtensionD.QgCodeListName = in.QgCodeListName
	fishDespatchAdviceLineItemExtensionD.StorageStateCode = in.StorageStateCode
	fishDespatchAdviceLineItemExtensionD.AquaCultureProductionUnit = in.AquaCultureProductionUnit
	fishDespatchAdviceLineItemExtensionD.FishingVessel = in.FishingVessel
	fishDespatchAdviceLineItemExtensionD.PlaceOfSlaughter = in.PlaceOfSlaughter
	fishDespatchAdviceLineItemExtensionD.PortOfLanding = in.PortOfLanding
	fishDespatchAdviceLineItemExtensionD.FishCatchOrProductionDateId = in.FishCatchOrProductionDateId
	fishDespatchAdviceLineItemExtensionD.FishCatchOrProductionId = in.FishCatchOrProductionId

	fishDespatchAdviceLineItemExtensionT := meatproto.FishDespatchAdviceLineItemExtensionT{}
	fishDespatchAdviceLineItemExtensionT.DateOfLanding = common.TimeToTimestamp(dateOfLanding.UTC().Truncate(time.Second))
	fishDespatchAdviceLineItemExtensionT.DateOfSlaughter = common.TimeToTimestamp(dateOfSlaughter.UTC().Truncate(time.Second))

	crUpdUser := commonproto.CrUpdUser{}
	crUpdUser.StatusCode = "active"
	crUpdUser.CreatedByUserId = user.Id
	crUpdUser.UpdatedByUserId = user.Id

	crUpdTime := commonproto.CrUpdTime{}
	crUpdTime.CreatedAt = tn
	crUpdTime.UpdatedAt = tn

	fishDespatchAdviceLineItemExtension := meatproto.FishDespatchAdviceLineItemExtension{FishDespatchAdviceLineItemExtensionD: &fishDespatchAdviceLineItemExtensionD, FishDespatchAdviceLineItemExtensionT: &fishDespatchAdviceLineItemExtensionT, CrUpdUser: &crUpdUser, CrUpdTime: &crUpdTime}

	return &fishDespatchAdviceLineItemExtension, nil
}

// insertFishDespatchAdviceLineItemExtension - Insert FishDespatchAdviceLineItemExtension into database
func (fis *FishService) insertFishDespatchAdviceLineItemExtension(ctx context.Context, insertFishDespatchAdviceLineItemExtensionSQL string, fishDespatchAdviceLineItemExtension *meatproto.FishDespatchAdviceLineItemExtension, userEmail string, requestID string) error {
	fishDespatchAdviceLineItemExtensionTmp, err := fis.crFishDespatchAdviceLineItemExtensionStruct(ctx, fishDespatchAdviceLineItemExtension, userEmail, requestID)
	if err != nil {
		fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	err = fis.DBService.InsUpd(ctx, userEmail, requestID, func(tx *sqlx.Tx) error {
		res, err := tx.NamedExecContext(ctx, insertFishDespatchAdviceLineItemExtensionSQL, fishDespatchAdviceLineItemExtensionTmp)
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}

		uID, err := res.LastInsertId()
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionD.Id = uint32(uID)
		uuid4Str, err := common.UUIDBytesToStr(fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionD.Uuid4)
		if err != nil {
			fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
			return err
		}
		fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionD.IdS = uuid4Str
		return nil
	})

	if err != nil {
		fis.log.Error("Error", zap.String("user", userEmail), zap.String("reqid", requestID), zap.Error(err))
		return err
	}
	return nil
}

// crFishDespatchAdviceLineItemExtensionStruct - process FishDespatchAdviceLineItemExtension details
func (fis *FishService) crFishDespatchAdviceLineItemExtensionStruct(ctx context.Context, fishDespatchAdviceLineItemExtension *meatproto.FishDespatchAdviceLineItemExtension, userEmail string, requestID string) (*meatstruct.FishDespatchAdviceLineItemExtension, error) {
	crUpdTime := new(commonstruct.CrUpdTime)
	crUpdTime.CreatedAt = common.TimestampToTime(fishDespatchAdviceLineItemExtension.CrUpdTime.CreatedAt)
	crUpdTime.UpdatedAt = common.TimestampToTime(fishDespatchAdviceLineItemExtension.CrUpdTime.UpdatedAt)

	fishDespatchAdviceLineItemExtensionT := new(meatstruct.FishDespatchAdviceLineItemExtensionT)
	fishDespatchAdviceLineItemExtensionT.DateOfLanding = common.TimestampToTime(fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionT.DateOfLanding)
	fishDespatchAdviceLineItemExtensionT.DateOfSlaughter = common.TimestampToTime(fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionT.DateOfSlaughter)

	fishDespatchAdviceLineItemExtensionTmp := meatstruct.FishDespatchAdviceLineItemExtension{FishDespatchAdviceLineItemExtensionD: fishDespatchAdviceLineItemExtension.FishDespatchAdviceLineItemExtensionD, FishDespatchAdviceLineItemExtensionT: fishDespatchAdviceLineItemExtensionT, CrUpdUser: fishDespatchAdviceLineItemExtension.CrUpdUser, CrUpdTime: crUpdTime}

	return &fishDespatchAdviceLineItemExtensionTmp, nil
}
