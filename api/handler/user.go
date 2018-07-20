package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/adham90/boilerplate/pkg/entity"
	uServ "github.com/adham90/boilerplate/pkg/user/service"
	"github.com/adham90/boilerplate/pkg/utils"
	"github.com/go-chi/chi"
	"github.com/google/jsonapi"
)

const (
	headerAccept      = "Accept"
	headerContentType = "Content-Type"
	contextKeyUser    = utils.ContextKey("user")
)

// HTTPUserHandler object
type uHandler struct {
	serv uServ.Service
}

func (h uHandler) getUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, ok := ctx.Value(contextKeyUser).(*entity.User)

	if !ok {
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Not Found Error",
			Detail: "can not find user",
			Status: "404",
		}})
		return
	}

	w.Header().Set("Content-Type", jsonapi.MediaType)
	w.WriteHeader(201)
	if err := jsonapi.MarshalPayload(w, user); err != nil {
		jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
			Title:  "Not Found Error",
			Detail: err.Error(),
			Status: "404",
		}})
		return
	}
}

func (h uHandler) createUser(w http.ResponseWriter, r *http.Request) {
	u := entity.User{}

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	}

	result, err := h.serv.CreateUser(u)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	jsonUser, err := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonUser)
}

func (h uHandler) UserCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "userID")
		userID, _ := strconv.Atoi(id)
		user, err := h.serv.FindByID(userID)
		if err != nil {
			jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
				Title:  "Not Found Error",
				Detail: "can not find user",
				Status: "404",
			}})
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyUser, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// MakeUserHandlers create user handler mux
func MakeUserHandlers(r *chi.Mux, us uServ.Service) {
	h := &uHandler{
		serv: us,
	}

	r.Route("/users", func(g chi.Router) {
		g.Route("/{userID}", func(sg chi.Router) {
			sg.Use(h.UserCtx)
			sg.Get("/", h.getUser)
			sg.Post("/", h.createUser)
		})
	})
}
