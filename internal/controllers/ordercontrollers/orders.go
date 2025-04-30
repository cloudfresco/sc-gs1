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

// OrderController - Create Order Controller
type OrderController struct {
	log                *zap.Logger
	UserServiceClient  partyproto.UserServiceClient
	OrderServiceClient orderproto.OrderServiceClient
	wfHelper           common.WfHelper
	workflowClient     client.Client
	ServerOpt          *config.ServerOptions
}

// NewOrderController - Create Order Handler
func NewOrderController(log *zap.Logger, userServiceClient partyproto.UserServiceClient, orderServiceClient orderproto.OrderServiceClient, wfHelper common.WfHelper, workflowClient client.Client, serverOpt *config.ServerOptions) *OrderController {
	return &OrderController{
		log:                log,
		UserServiceClient:  userServiceClient,
		OrderServiceClient: orderServiceClient,
		wfHelper:           wfHelper,
		workflowClient:     workflowClient,
		ServerOpt:          serverOpt,
	}
}

// CreateOrder - Create Order
func (oc *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"order:cud"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
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

	form := orderproto.CreateOrderRequest{}
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
	result := wHelper.StartWorkflow(workflowOptions, orderworkflows.CreateOrderWorkflow, &form, token, user, oc.log)
	workflowClient := oc.workflowClient
	workflowRun := workflowClient.GetWorkflow(ctx, result.ID, result.RunID)
	var order orderproto.CreateOrderResponse
	err = workflowRun.Get(ctx, &order)
	if err != nil {
		oc.log.Error("Error", zap.String("user", user.Email), zap.String("reqid", user.RequestId), zap.Error(err))
		common.RenderErrorJSON(w, "4002", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, order)
}

// Index - list CrediNotes
func (oc *OrderController) Index(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"order:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	cursor := r.URL.Query().Get("cursor")
	limit := r.URL.Query().Get("limit")
	orders, err := oc.OrderServiceClient.GetOrders(ctx, &orderproto.GetOrdersRequest{Limit: limit, NextCursor: cursor, UserEmail: user.Email, RequestId: user.RequestId})
	if err != nil {
		oc.log.Error("Error",
			zap.String("user", user.Email),
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1301", err.Error(), 402, user.RequestId)
		return
	}

	common.RenderJSON(w, orders)
}

// Show - Show OrderOrder
func (oc *OrderController) Show(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"order:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	order, err := oc.OrderServiceClient.GetOrder(ctx, &orderproto.GetOrderRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		oc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, order)
}

// UpdateOrder - Update Order
func (oc *OrderController) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	ctx, user, token, err := common.GetContextAuthUser(w, r, []string{"order:cud"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
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

	form := orderproto.UpdateOrderRequest{}
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
	result := wHelper.StartWorkflow(workflowOptions, orderworkflows.UpdateOrderWorkflow, &form, token, user, oc.log)
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

// GetOrderLineItems - Get OrderOrder Lines
func (oc *OrderController) GetOrderLineItems(w http.ResponseWriter, r *http.Request) {
	ctx, user, _, err := common.GetContextAuthUser(w, r, []string{"order:read"}, oc.ServerOpt.Auth0Audience, oc.ServerOpt.Auth0Domain, oc.UserServiceClient)
	if err != nil {
		common.RenderErrorJSON(w, "1001", err.Error(), 401, user.RequestId)
		return
	}

	id := r.PathValue("id")

	orderLineItems, err := oc.OrderServiceClient.GetOrderLineItems(ctx, &orderproto.GetOrderLineItemsRequest{GetRequest: &commonproto.GetRequest{Id: id, UserEmail: user.Email, RequestId: user.RequestId}})
	if err != nil {
		oc.log.Error("Error",
			zap.String("reqid", user.RequestId),
			zap.Error(err))
		common.RenderErrorJSON(w, "1103", err.Error(), 402, user.RequestId)
		return
	}
	common.RenderJSON(w, orderLineItems)
}
