package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"GetFarmasiList",
		"GET",
		"/farmasi/{negeri}",
		GetFarmasi,
	},
	Route{
		"GetHospitalList",
		"GET",
		"/hospital/{negeri}",
		GetHospital,
	},
	Route{
		"GetUserInfo",
		"GET",
		"/user/{userid}",
		GetUserInfo,
	},
	Route{
		"UpdateUserInfo",
		"PUT",
		"/user/{userid}",
		UpdateUserInfo,
	},
	Route{
		"CreateUserInfo",
		"POST",
		"/user/{userid}",
		CreateUserInfo,
	},
}
