package router

import "net/http"

type Router interface {
	GET(uri string, f func(rw http.ResponseWriter, r *http.Request))
	POST(uri string, f func(rw http.ResponseWriter, r *http.Request))
	PUT(uri string, f func(rw http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(rw http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
