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

// ReceivingAdviceController - Create ReceivingAdvice Controller
type ReceivingAdviceController struct {
	log                          *zap.Logger
	UserServiceClient            partyproto.UserServiceClient
	ReceivingAdviceServiceClient logisticsproto.ReceivingAdviceServiceClient
	wfHelper                     common.WfHelper
	workflowClient               client.Client
	ServerOpt                    *config.ServerOptions
}

// NewReceivingAdviceController - Create ReceivingAdvice Handler
func NewReceivingAdviceController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, receivingAdviceServiceClient logisticsproto.ReceivingAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *ReceivingAdviceController {
	return &ReceivingAdviceController{
		log:                          log,
		UserServiceClient:            userServiceClient,
		ReceivingAdviceServiceClient: receivingAdviceServiceClient,
		wfHelper:                     wfHelper,
		workflowClient:               workflowClient,
		ServerOpt:                    serverOpt,
	}
}

// CreateReceivingAdvice - Create ReceivingAdvice
func (rc *ReceivingAdviceController) CreateReceivingAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"rcvadv:cud"}, rc.ServerOpt.Auth0Audience, rc.ServerOpt.Auth0Domain, rc.UserServiceClient)
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

	form := logisticsproto.CreateReceivingAdviceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		rc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := rc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, logisticsworkflows.CreateReceivingAdviceWorkflow, &form, token, user, rc.log)
	workflowClient := rc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var receivingAdvice logisticsproto.CreateReceivingAdviceResponse
	err = workflowRun.Get(ctx, &receivingAdvice)

	if err != nil {
		rc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, receivingAdvice)
}

// Index - list ReceivingAdvices
func (rc *ReceivingAdviceController) Index(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"rcvadv:read"}, rc.ServerOpt.Auth0Audience, rc.ServerOpt.Auth0Domain, rc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	receivingAdvice, err := rc.ReceivingAdviceServiceClient.GetReceivingAdvices(ctx, &logisticsproto.GetReceivingAdvicesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		rc.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, receivingAdvice)
}

// Show - Show ReceiptAdvice
func (rc *ReceivingAdviceController) Show(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"rcvadv:read"}, rc.ServerOpt.Auth0Audience, rc.ServerOpt.Auth0Domain, rc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	receivingAdvice, err := rc.ReceivingAdviceServiceClient.GetReceivingAdvice(ctx, &logisticsproto.GetReceivingAdviceRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		rc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, receivingAdvice)
}

// UpdateReceivingAdvice - Update ReceivingAdvice
func (rc *ReceivingAdviceController) UpdateReceivingAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"rcvadv:cud"}, rc.ServerOpt.Auth0Audience, rc.ServerOpt.Auth0Domain, rc.UserServiceClient)
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

	form := logisticsproto.UpdateReceivingAdviceRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		rc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	form.Id = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := rc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, logisticsworkflows.UpdateReceivingAdviceWorkflow, &form, token, user, rc.log)
	workflowClient := rc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		rc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}

// GetReceivingAdviceLineItems - Get ReceiptAdvice Lines
func (rc *ReceivingAdviceController) GetReceivingAdviceLineItems(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"rcvadv:read"}, rc.ServerOpt.Auth0Audience, rc.ServerOpt.Auth0Domain, rc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	receivingAdviceLineItems, err := rc.ReceivingAdviceServiceClient.GetReceivingAdviceLineItems(ctx, &logisticsproto.GetReceivingAdviceLineItemsRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		rc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, receivingAdviceLineItems)
}
