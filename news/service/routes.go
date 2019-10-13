package service

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{
		"GetNews",
		"GET",
		"/news/get",
		GetNews,
	},
	{
		"AddNews",
		"GET",
		"/news/add",
		AddNews,
	},
}
