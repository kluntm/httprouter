package main

import (
	"github.com/julienschmidt/httprouter"
	"httprouter/api"
)

type Route struct {
	Name string
	Method string
	Path string
	HandleFunc httprouter.Handle
}

type Routes []Route

func AllRoutes() Routes{
	routes := Routes{
		Route{"Index", "GET", "/", api.Index},
		Route{"BookIndex", "GET", "/index", api.BookIndex},
		Route{"BookShow", "GET", "/book/:isdn", api.BookShow},
		Route{"BookShow", "POST", "/book", api.BookShow},
	}
	return routes
}


