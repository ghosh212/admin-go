package base

import (
	"admin-go/model"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//as we are uding json we needto verofy al the request and response ar in json format
//takes input as hepp handler
//gives output for http handler
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func NewHTTPServer(ctx context.Context, endpoints Endpoint, basepath string) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

}

func encoderesponse(ctx context.Context, r http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(r).Encode(response)
}

func decoderequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req model.SearchAdmin
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
