package web

import (
	"encoding/json"
	"net/http"

	"github.com/flpnascto/clean-architecture-go/internal/entity"
	"github.com/flpnascto/clean-architecture-go/internal/usecase"
	"github.com/flpnascto/clean-architecture-go/pkg/events"
)

type WebOrderHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	OrderRepository   entity.OrderRepositoryInterface
	OrderCreatedEvent events.EventInterface
	ListOrdersEvent   events.EventInterface
}
type OrderCreatedEventInterface interface{ events.EventInterface }
type ListOrdersEventInterface interface{ events.EventInterface }

func NewWebOrderHandler(
	EventDispatcher events.EventDispatcherInterface,
	OrderRepository entity.OrderRepositoryInterface,
	OrderCreatedEvent OrderCreatedEventInterface,
	ListOrdersEvent ListOrdersEventInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		EventDispatcher:   EventDispatcher,
		OrderRepository:   OrderRepository,
		OrderCreatedEvent: OrderCreatedEvent,
		ListOrdersEvent:   ListOrdersEvent,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.OrderInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createOrder := usecase.NewCreateOrderUseCase(h.OrderRepository, h.OrderCreatedEvent, h.EventDispatcher)
	output, err := createOrder.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebOrderHandler) List(w http.ResponseWriter, r *http.Request) {
	listOrders := usecase.ListOrdersUseCase(h.OrderRepository, h.ListOrdersEvent, h.EventDispatcher)
	output, err := listOrders.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
