package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"IntroductionCreate",
		"POST",
		"/intro",
		IntroCreate,
	},
	Route{
		"IntroductionResult",
		"GET",
		"/intro/{introId}",
		IntroResult,
	},
	Route{
		"IntroductionDelete",
		"DELETE",
		"/intro/{introId}",
		IntroDelete,
	},
	Route{
		"ContactCreate",
		"POST",
		"/contacts/{contactId}",
		ContactCreate,
	},
	Route{
		"ContactDelete",
		"DELETE",
		"/contacts/{contactId}",
		ContactDelete,
	},
	Route{
		"ContactShow",
		"GET",
		"/contacts/{contactId}",
		ContactShow,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}
