package handler

import (
	"encoding/json"
	"net/http"

	"github.com/adham90/boilerplate/pkg/entity"
	uServ "github.com/adham90/boilerplate/pkg/user/service"
	"github.com/go-chi/chi"
)

// HTTPUserHandler object
type uHandler struct {
	serv uServ.Service
}

func (h uHandler) userIndex() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *entity.User
		var err error

		user, err = h.serv.FindByID("1")

		if err != nil {
			http.Error(w, "User not found", 400)
		}

		jsonUser, err := json.Marshal(user)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	})
}

// MakeUserHandlers create user handler mux
func MakeUserHandlers(r *chi.Mux, us uServ.Service) {
	handler := &uHandler{
		serv: us,
	}

	r.Method("get", "/", handler.userIndex())
}
