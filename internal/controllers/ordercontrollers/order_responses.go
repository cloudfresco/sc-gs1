package ordercontrollers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/cloudfresco/sc-gs1/internal/common"
	"github.com/cloudfresco/sc-gs1/internal/config"
	commonproto "github.com/cloudfresco/sc-gs1/internal/protogen/common/v1"
	orderproto "github.com/cloudfresco/sc-gs1/internal/protogen/order/v1"
	partyproto "github.com/cloudfresco/sc-gs1/internal/protogen/party/v1"
	"github.com/cloudfresco/sc-gs1/internal/workflows/orderworkflows"
	"github.com/pborman/uuid"
	"go.uber.org/cadence/client"
	"go.uber.org/zap"
)

// OrderResponseController - Create OrderResponse Controller
type OrderResponseController struct {
	log                        *zap.Logger
	UserServiceClient          partyproto.UserServiceClient
	OrderResponseServiceClient orderproto.OrderResponseServiceClient
	wfHelper                   common.WfHelper
	workflowClient             client.Client
	ServerOpt                  *config.ServerOptions
}

// NewOrderResponseController - Create OrderResponse Handler
func NewOrderResponseController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, orderResponseServiceClient orderproto.OrderResponseServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *OrderResponseController {
	return &OrderResponseController{
		log:                        log,
		UserServiceClient:          userServiceClient,
		OrderResponseServiceClient: orderResponseServiceClient,
		wfHelper:                   wfHelper,
		workflowClient:             workflowClient,
		ServerOpt:                  serverOpt,
	}
}

// CreateOrderResponse - Create OrderResponse
func (oc *OrderResponseController) CreateOrderResponse(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"orderresp:cud"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        orderworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := orderproto.CreateOrderResponseRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		oc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := oc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, orderworkflows.CreateOrderResponseWorkflow, &form, token, user, oc.log)
	workflowClient := oc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var orderResponse orderproto.CreateOrderResponseResponse
	err = workflowRun.Get(ctx, &orderResponse)
	if err != nil {
		oc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, orderResponse)
}

// Index - list CrediNotes
func (oc *OrderResponseController) Index(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"orderresp:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")
	orderResponses, err := oc.OrderResponseServiceClient.GetOrderResponses(ctx, &orderproto.GetOrderResponsesRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		oc.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, orderResponses)
}

// Show - Show OrderResponseOrderResponse
func (oc *OrderResponseController) Show(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"orderresp:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	orderResponse, err := oc.OrderResponseServiceClient.GetOrderResponse(ctx, &orderproto.GetOrderResponseRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		oc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, orderResponse)
}

// UpdateOrderResponse - Update OrderResponse
func (oc *OrderResponseController) UpdateOrderResponse(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"orderresp:cud"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	workflowOptions := client.StartWorkflowOptions{
		ID:                              "gs1_" + uuid.New(),
		TaskList:                        orderworkflows.ApplicationName,
		ExecutionStartToCloseTimeout:    time.Minute,
		DecisionTaskStartToCloseTimeout: time.Minute,
	}

	form := orderproto.UpdateOrderResponseRequest{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&form)
	if err != nil {
		oc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	form.Id = id
	form.UserId = user.UserId
	form.UserEmail = user.Email
	form.RequestId = user.RequestId

	wHelper := oc.wfHelper
	result := wHelper.StartWorkflow(workflowOptions, orderworkflows.UpdateOrderResponseWorkflow, &form, token, user, oc.log)
	workflowClient := oc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var response string
	err = workflowRun.Get(ctx, &response)
	if err != nil {
		oc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4009", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, response)
}

// GetOrderResponseLineItems - Get OrderResponseOrderResponse Lines
func (oc *OrderResponseController) GetOrderResponseLineItems(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"orderresp:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	orderResponseLineItems, err := oc.OrderResponseServiceClient.GetOrderResponseLineItems(ctx, &orderproto.GetOrderResponseLineItemsRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		oc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, orderResponseLineItems)
}
