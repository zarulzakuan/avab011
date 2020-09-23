package main

import (
	"net/http"

	"github.com/gorilla/mux"
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

	return router
}

var routes = Routes{
	Route{
		"GET",
		"/",
		index,
	},
	Route{
		"GET",
		"/todos",
		todoIndex,
	},
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
		"/paymentstatus",
		getPaymentStatus,
	},
	Route{
		"GET",
		"/payment/{paymentid}",
		makePayment,
	},
}
