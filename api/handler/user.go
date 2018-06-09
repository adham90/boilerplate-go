package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adham90/boilerplate/pkg/entity"
	pkg "github.com/adham90/boilerplate/pkg/user"
	"github.com/go-chi/chi"
)

func userIndex(s pkg.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *entity.User
		var err error

		user, err = s.Find("5")

		if err != nil {
			panic(err)
		}

		jsonUser, err := json.Marshal(user)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	})
}

func MakeUserHandlers(r *chi.Mux, service pkg.Service) {
	r.Method("get", "/", userIndex(service))
}
