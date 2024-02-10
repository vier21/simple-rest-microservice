package handler

import (
	"encoding/json"
	"gorest/internal/utils/errors"
	"gorest/pkg/user/domain/web"
	"gorest/pkg/user/service"
	"net/http"

	"github.com/gorilla/mux"
)

type httpHandler struct {
	svc service.UserService
}

func UserHttpHandler(s service.UserService, mux *mux.Router) {
	h := &httpHandler{
		svc: s,
	}

	mux.HandleFunc("/register", h.RegisterUserHandler).Methods(http.MethodPost)
	mux.HandleFunc("/user/{username}", h.GetUserByUsernameHandler).Methods(http.MethodGet)
}

func (h *httpHandler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var payload web.RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := web.ResponsePayload{
			Code:   http.StatusBadRequest,
			Status: http.StatusText(http.StatusBadRequest),
			Data:   nil,
		}
		json.NewEncoder(w).Encode(resp)
		return
	}

	regist := h.svc.RegisterUser(r.Context(), payload)

	if regist.Error != nil {
		err := regist.Error.(*errors.ErrorString)
		w.WriteHeader(err.Code())
		resp := web.ResponsePayload{
			Code:   err.Code(),
			Status: err.Message(),
			Data:   nil,
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	json.NewEncoder(w).Encode(web.ResponsePayload{
		Code:   http.StatusCreated,
		Status: "success",
		Data:   regist.Data,
	})

}

func (h *httpHandler) GetUserByUsernameHandler(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	username := mux.Vars(r)["username"]

	user := h.svc.GetUserByUsername(r.Context(), username)

	if user.Error != nil {
		err := user.Error.(*errors.ErrorString)
		w.WriteHeader(err.Code())
		json.NewEncoder(w).Encode(web.ResponsePayload{
			Code:   err.Code(),
			Status: "failed",
			Data:   err.Message(),
		})
		return
	}

	json.NewEncoder(w).Encode(web.ResponsePayload{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   user,
	})
}

func setHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
}
