package invoicecontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	invoiceproto "github.com/cloudfresco/sc-gs1/internal/protogen/invoice/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	invoiceworkflows "github.com/cloudfresco/sc-gs1/internal/workflows/invoiceworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// DebitCreditAdviceController - Create DebitCreditAdvice Controller
type DebitCreditAdviceController struct {
	log                            *zap.Logger
	UserServiceClient              partyproto.UserServiceClient
	DebitCreditAdviceServiceClient invoiceproto.DebitCreditAdviceServiceClient
	wfHelper                       common.WfHelper
	workflowClient                 client.Client
	ServerOpt                      *config.ServerOptions
}

// NewDebitCreditAdviceController - Create DebitCreditAdvice Handler
func NewDebitCreditAdviceController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, debitCreditAdviceServiceClient invoiceproto.DebitCreditAdviceServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *DebitCreditAdviceController {
	return &DebitCreditAdviceController{
		log:                            log,
		UserServiceClient:              userServiceClient,
		DebitCreditAdviceServiceClient: debitCreditAdviceServiceClient,
		wfHelper:                       wfHelper,
		workflowClient:                 workflowClient,
		ServerOpt:                      serverOpt,
	}
}

// CreateDebitCreditAdvice - Create DebitCreditAdvice
func (dc *DebitCreditAdviceController) CreateDebitCreditAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"dca:cud"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        invoiceworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := invoiceproto.CreateDebitCreditAdviceRequest{}
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
	result := wHelper.StartWorkflow(workflowOptions, invoiceworkflows.CreateDebitCreditAdviceWorkflow, &form, token, user, dc.log)
	workflowClient := dc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var debitCreditAdvice invoiceproto.CreateDebitCreditAdviceResponse
	err = workflowRun.Get(ctx, &debitCreditAdvice)

	if err != nil {
		dc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, debitCreditAdvice)
}

// Index - list DebitCreditAdvice
func (dc *DebitCreditAdviceController) Index(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"dca:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")

	debitCreditAdvice, err := dc.DebitCreditAdviceServiceClient.GetDebitCreditAdvices(ctx, &invoiceproto.GetDebitCreditAdvicesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		dc.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, debitCreditAdvice)
}

// Show - Show DebitCreditAdvice
func (dc *DebitCreditAdviceController) Show(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"dca:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	debitCreditAdvice, err := dc.DebitCreditAdviceServiceClient.GetDebitCreditAdvice(ctx, &invoiceproto.GetDebitCreditAdviceRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		dc.log.Error("Error",
			zap.String("reqid", user.RequestId),

			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, debitCreditAdvice)
}

// UpdateDebitCreditAdvice - Update DebitCreditAdvice
func (dc *DebitCreditAdviceController) UpdateDebitCreditAdvice(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"dca:cud"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        invoiceworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := invoiceproto.UpdateDebitCreditAdviceRequest{}
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
	result := wHelper.StartWorkflow(workflowOptions, invoiceworkflows.UpdateDebitCreditAdviceWorkflow, &form, token, user, dc.log)
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

// GetDebitCreditAdviceLineItems - Get DebitCreditAdvice LineItems
func (dc *DebitCreditAdviceController) GetDebitCreditAdviceLineItems(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"dca:read"}, dc.ServerOpt.Auth0Audience, dc.ServerOpt.Auth0Domain, dc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	debitCreditAdviceLines, err := dc.DebitCreditAdviceServiceClient.GetDebitCreditAdviceLineItems(ctx, &invoiceproto.GetDebitCreditAdviceLineItemsRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		dc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, debitCreditAdviceLines)
}
