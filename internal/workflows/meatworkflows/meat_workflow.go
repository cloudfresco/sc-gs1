package meatworkflows

import (
	"time"

	meatproto "github.com/cloudfresco/sc-gs1/internal/protogen/meat/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"

	"go.uber.org/cadence/workflow"
	"go.uber.org/zap"
)

const (
	// ApplicationName is the task list
	ApplicationName = "gs1"
)

// CreateMeatActivityHistoryWorkflow - Create MeatActivityHistory workflow
func CreateMeatActivityHistoryWorkflow(ctx workflow.Context, form *meatproto.CreateMeatActivityHistoryRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatActivityHistoryResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatActivityHistory meatproto.CreateMeatActivityHistoryResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatActivityHistoryActivity, form, tokenString, user, log).Get(ctx, &meatActivityHistory)
	if err != nil {
		logger.Error("Failed to CreateMeatActivityHistoryWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatActivityHistory, nil
}

// CreateMeatAcidityWorkflow - Create MeatAcidity workflow
func CreateMeatAcidityWorkflow(ctx workflow.Context, form *meatproto.CreateMeatAcidityRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatAcidityResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatAcidity meatproto.CreateMeatAcidityResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatAcidityActivity, form, tokenString, user, log).Get(ctx, &meatAcidity)
	if err != nil {
		logger.Error("Failed to CreateMeatAcidityWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatAcidity, nil
}

// CreateMeatTestWorkflow - Create MeatTest workflow
func CreateMeatTestWorkflow(ctx workflow.Context, form *meatproto.CreateMeatTestRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatTestResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatTest meatproto.CreateMeatTestResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatTestActivity, form, tokenString, user, log).Get(ctx, &meatTest)
	if err != nil {
		logger.Error("Failed to CreateMeatTestWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatTest, nil
}

// CreateMeatBreedingDetailWorkflow - Create MeatBreedingDetail workflow
func CreateMeatBreedingDetailWorkflow(ctx workflow.Context, form *meatproto.CreateMeatBreedingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatBreedingDetailResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatBreedingDetail meatproto.CreateMeatBreedingDetailResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatBreedingDetailActivity, form, tokenString, user, log).Get(ctx, &meatBreedingDetail)
	if err != nil {
		logger.Error("Failed to CreateMeatBreedingDetailWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatBreedingDetail, nil
}

// CreateMeatCuttingDetailWorkflow - Create MeatCuttingDetail workflow
func CreateMeatCuttingDetailWorkflow(ctx workflow.Context, form *meatproto.CreateMeatCuttingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatCuttingDetailResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatCuttingDetail meatproto.CreateMeatCuttingDetailResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatCuttingDetailActivity, form, tokenString, user, log).Get(ctx, &meatCuttingDetail)
	if err != nil {
		logger.Error("Failed to CreateMeatCuttingDetailWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatCuttingDetail, nil
}

// CreateMeatFatteningDetailWorkflow - Create MeatFatteningDetail workflow
func CreateMeatFatteningDetailWorkflow(ctx workflow.Context, form *meatproto.CreateMeatFatteningDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatFatteningDetailResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatFatteningDetail meatproto.CreateMeatFatteningDetailResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatFatteningDetailActivity, form, tokenString, user, log).Get(ctx, &meatFatteningDetail)
	if err != nil {
		logger.Error("Failed to CreateMeatFatteningDetailWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatFatteningDetail, nil
}

// CreateMeatMincingDetailWorkflow - Create MeatMincingDetail workflow
func CreateMeatMincingDetailWorkflow(ctx workflow.Context, form *meatproto.CreateMeatMincingDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatMincingDetailResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatMincingDetail meatproto.CreateMeatMincingDetailResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatMincingDetailActivity, form, tokenString, user, log).Get(ctx, &meatMincingDetail)
	if err != nil {
		logger.Error("Failed to CreateMeatMincingDetailWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatMincingDetail, nil
}

// CreateMeatProcessingPartyWorkflow - Create MeatProcessingParty workflow
func CreateMeatProcessingPartyWorkflow(ctx workflow.Context, form *meatproto.CreateMeatProcessingPartyRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatProcessingPartyResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatProcessingParty meatproto.CreateMeatProcessingPartyResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatProcessingPartyActivity, form, tokenString, user, log).Get(ctx, &meatProcessingParty)
	if err != nil {
		logger.Error("Failed to CreateMeatProcessingPartyWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatProcessingParty, nil
}

// CreateMeatSlaughteringDetailWorkflow - Create MeatSlaughteringDetail workflow
func CreateMeatSlaughteringDetailWorkflow(ctx workflow.Context, form *meatproto.CreateMeatSlaughteringDetailRequest, tokenString string, user *partyproto.GetAuthUserDetailsResponse, log *zap.Logger) (*meatproto.CreateMeatSlaughteringDetailResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToStartTimeout: time.Minute,
		StartToCloseTimeout:    time.Minute,
		HeartbeatTimeout:       time.Second * 20,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)
	logger := workflow.GetLogger(ctx)
	var ma *MeatActivities
	var meatSlaughteringDetail meatproto.CreateMeatSlaughteringDetailResponse
	err := workflow.ExecuteActivity(ctx, ma.CreateMeatSlaughteringDetailActivity, form, tokenString, user, log).Get(ctx, &meatSlaughteringDetail)
	if err != nil {
		logger.Error("Failed to CreateMeatSlaughteringDetailWorkflow", zap.Error(err))
		log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		return nil, err
	}
	return &meatSlaughteringDetail, nil
}
