package api

import (
	"log"
	"net/http"
)

// bootRouter boots router
func (a *API) bootRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a.router)

	return mux
}

// router for the requests
func (a *API) router(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, " ", r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		switch r.URL.Path {
		case "/init":
			a.handlerInit(w, r)
		default:
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"message":"not found"}`))
		}
	}
}
