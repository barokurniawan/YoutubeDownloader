package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type RouteRegister struct {
	Path    string
	Handler http.Handler
}

type RouteServiceProvide struct {
	Router   *mux.Router
	Register []RouteRegister
}

func (routeServiceProvider *RouteServiceProvide) SetRouter(router *mux.Router) {
	routeServiceProvider.Router = router
}

// GetRouterInstance get instance of gorila mux route
func (routeServiceProvider RouteServiceProvide) GetRouterInstance() *mux.Router {
	return routeServiceProvider.Router
}

func (routeServiceProvider *RouteServiceProvide) RegisterRoute(path string, handler http.Handler) {
	routeServiceProvider.Register = append(routeServiceProvider.Register, RouteRegister{
		Path:    path,
		Handler: handler,
	})
}

func (routeServiceProvider RouteServiceProvide) ProvideRoute() {
	for _, element := range routeServiceProvider.Register {
		var route = routeServiceProvider.GetRouterInstance()
		route.Handle(element.Path, element.Handler)

		fmt.Println("init route: " + element.Path)
	}
}
