package bot

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shahinrahimi/telelancerbot/types"
)

type Handler func(*tgbotapi.Update, context.Context)
type ErrorHandler func(*tgbotapi.Update, context.Context) error
type Middleware func(Handler) Handler

type Router struct {
	middlewares []Middleware
	handlers    map[types.CommandType]Handler
	routes      map[string]*Route
}

type Route struct {
	middlewares []Middleware
	handlers    map[types.CommandType]Handler
}

func (r *Router) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) NewRoute(name string) *Route {
	if _, exist := r.routes[name]; exist {
		log.Panicf("Route %s already exists", name)
	}
	newRoute := &Route{
		middlewares: make([]Middleware, 0),
		handlers:    make(map[types.CommandType]Handler),
	}
	r.routes[name] = newRoute
	return newRoute
}

func (r *Route) Use(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *Router) HandleCommand(c types.CommandType, h Handler) {
	r.handlers[c] = h
}

func (r *Route) HandleCommand(c types.CommandType, h Handler) {
	r.handlers[c] = h
}
