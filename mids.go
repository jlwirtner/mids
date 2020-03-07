package mids

import (
	"fmt"
	"net/http"
)

type methodHandlers struct {
	get http.HandlerFunc
	put http.HandlerFunc
	post http.HandlerFunc
	delete http.HandlerFunc
}

var endpoints = make(map[string]methodHandlers)

func Get(pattern string, handler http.HandlerFunc) {
	tmp := endpoints[pattern]
	tmp.get = handler
	endpoints[pattern] = tmp
}

func Put(pattern string, handler http.HandlerFunc) {
	tmp := endpoints[pattern]
	tmp.put = handler
	endpoints[pattern] = tmp
}

func Post(pattern string, handler http.HandlerFunc) {
	tmp := endpoints[pattern]
	tmp.post = handler
	endpoints[pattern] = tmp
}

func Delete(pattern string, handler http.HandlerFunc) {
	tmp := endpoints[pattern]
	tmp.delete = handler
	endpoints[pattern] = tmp
}

func execIfNotNil(fn http.HandlerFunc, w http.ResponseWriter, r *http.Request) {
	if fn == nil {
		fmt.Fprintf(w, "404 page not found")
		return
	}
	fn(w,r)
}

func ListenAndServe(port string) {
	for pattern, handlers := range endpoints {
		get := handlers.get
		put := handlers.put
		post := handlers.post
		delete := handlers.delete
		fn := func(w http.ResponseWriter,r *http.Request) {
			switch r.Method {
			case http.MethodGet:
				execIfNotNil(get,w,r)
			case http.MethodPut:
				execIfNotNil(put,w,r)
			case http.MethodPost:
				execIfNotNil(post,w,r)
			case http.MethodDelete:
				execIfNotNil(delete,w,r)
			default:
				fmt.Fprintf(w,"404 page not found")
			}

		}
		http.HandleFunc(pattern, fn)
	}
	http.ListenAndServe(port, nil)
}