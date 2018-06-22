package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/adham90/boilerplate/pkg/entity"
	us "github.com/adham90/boilerplate/pkg/user/service"
	"github.com/go-chi/chi"
)

func userIndex(s us.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var user *entity.User
		var err error

		user, err = s.Find("1")

		if err != nil {
			panic(err)
		}

		jsonUser, err := json.Marshal(user)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonUser)
	})
}

func MakeUserHandlers(r *chi.Mux, db *sql.DB) {
	us := us.New(db)

	r.Method("get", "/", userIndex(us))
}
