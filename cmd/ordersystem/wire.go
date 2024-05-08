//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/flpnascto/clean-architecture-go/internal/entity"
	"github.com/flpnascto/clean-architecture-go/internal/event"
	"github.com/flpnascto/clean-architecture-go/internal/infra/database"
	"github.com/flpnascto/clean-architecture-go/internal/infra/web"
	"github.com/flpnascto/clean-architecture-go/internal/usecase"
	"github.com/flpnascto/clean-architecture-go/pkg/events"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	event.NewListOrders,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventInterface), new(*event.ListOrders)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

var setListOrdersEvent = wire.NewSet(
	event.NewListOrders,
	wire.Bind(new(events.EventInterface), new(*event.ListOrders)),
)

func ListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCaseType {
	wire.Build(
		setOrderRepositoryDependency,
		setListOrdersEvent,
		usecase.ListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCaseType{}
}

var setWebOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(web.OrderCreatedEventInterface), new(*event.OrderCreated)),
)

var setWebListOrdersEvent = wire.NewSet(
	event.NewListOrders,
	wire.Bind(new(web.ListOrdersEventInterface), new(*event.ListOrders)),
)

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setWebOrderCreatedEvent,
		setWebListOrdersEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
