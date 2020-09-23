package main

import (
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/zarulzakuan/avab011/docs"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes List of all routes
type Routes []Route

// NewRouter New router, duhh
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}

	router.PathPrefix("/").Handler(httpSwagger.WrapHandler)

	return router
}

var routes = Routes{
	Route{
		"POST",
		"/order",
		createOrder,
	},
	Route{
		"POST",
		"/order/hotel",
		searchOrderByHotel,
	},
	Route{
		"POST",
		"/order/customer",
		searchOrderByCustomer,
	},
	Route{
		"GET",
		"/paymentstatus/{orderid}",
		getPaymentStatus,
	},
	Route{
		"GET",
		"/makepayment/{paymentid}",
		makePayment,
	},
}
