package logisticscontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	logisticsproto "github.com/cloudfresco/sc-gs1/internal/protogen/logistics/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	logisticsworkflows "github.com/cloudfresco/sc-gs1/internal/workflows/logisticsworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// DespatchAdviceController - Create DespatchAdvice Controller
type DespatchAdviceController struct {
	log                         *zap.Logger
	UserServiceClient           partyproto.UserServiceClient
	DespatchAdviceServiceClient logisticsproto.DespatchAdviceServiceClient
	wfHelper                    common.WfHelper
	workflowClient              client.Client
	ServerOpt                   *config.ServerOptions
}

// NewDespatchAdviceController - Create DespatchAdvice Handler
func NewDespatchAdviceController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, despatchServiceClient logisticsproto.DespatchAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *DespatchAdviceController {
	return &DespatchAdviceController{
		log:                         log,
		UserServiceClient:           userServiceClient,
		DespatchAdviceServiceClient: despatchServiceClient,
		wfHelper:                    wfHelper,
		workflowClient:              workflowClient,
		ServerOpt:                   serverOpt,
	}
}

// CreateDespatchAdvice - Create Despatch Header
func (dc *DespatchAdviceController) CreateDespatchAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"despadv:cud"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        logisticsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := logisticsproto.CreateDespatchAdviceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		dc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := dc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, logisticsworkflows.CreateDespatchAdviceWorkflow, &form, token, user, dc.log)
	workflowClient := dc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var despatchAdvice logisticsproto.CreateDespatchAdviceResponse
	err = workflowRun.Get(ctx, &despatchAdvice)

	if err != nil {
		dc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, despatchAdvice)
}

// Index - list Despatch Headers
func (dc *DespatchAdviceController) Index(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"despadv:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	despatchAdvice, err := dc.DespatchAdviceServiceClient.GetDespatchAdvices(ctx, &logisticsproto.GetDespatchAdvicesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		dc.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, despatchAdvice)
}

// Show - Show Despatch
func (dc *DespatchAdviceController) Show(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"despadv:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	despatchAdvice, err := dc.DespatchAdviceServiceClient.GetDespatchAdvice(ctx, &logisticsproto.GetDespatchAdviceRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		dc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, despatchAdvice)
}

// UpdateDespatchAdvice - Update DespatchAdvice
func (dc *DespatchAdviceController) UpdateDespatchAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"despadv:cud"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        logisticsworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := logisticsproto.UpdateDespatchAdviceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		dc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	form.Id = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := dc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, logisticsworkflows.UpdateDespatchAdviceWorkflow, &form, token, user, dc.log)
	workflowClient := dc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		dc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}

// GetDespatchAdviceLineItems - Get DespatchAdviceLineItems
func (dc *DespatchAdviceController) GetDespatchAdviceLineItems(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"despadv:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	despatchLines, err := dc.DespatchAdviceServiceClient.GetDespatchAdviceLineItems(ctx, &logisticsproto.GetDespatchAdviceLineItemsRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		dc.log.Error("Error",
			zap.String("reqid", user.RequestId),

			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, despatchLines)
}
