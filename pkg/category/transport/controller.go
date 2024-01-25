package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"gorest/pkg/category/middlewares"
	"gorest/pkg/category/repository"
	"gorest/pkg/category/service"

	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type controller struct {
	service service.CategoryServices
}

func InitHttpHandler(mux *mux.Router, svc service.CategoryServices) {
	handler := &controller{
		service: svc,
	}
	mux.NotFoundHandler = http.HandlerFunc(middlewares.NotfoundHandler)

	s := mux.PathPrefix("/category").Subrouter()
	s.HandleFunc("/", handler.SaveCtgHandler).Methods(http.MethodPost)
	s.HandleFunc("/{id}", handler.UpdateCtgHandler).Methods(http.MethodPut)
	s.HandleFunc("/{id}", handler.DeleteCtgHandler).Methods(http.MethodDelete)
	s.HandleFunc("/", handler.FindAllCtgHandler).Methods(http.MethodGet)
	s.HandleFunc("/{id}", handler.FIndByIdCtgHandler).Methods(http.MethodGet)
}

type Response struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

var (
	ErrBadRequestMsg     = "invalid request"
	ErrInternalServerMsg = "internal server error"
)

func encodeJson(w http.ResponseWriter, code int, status string, data any) error {
	resp := Response{
		Code:   code,
		Status: status,
		Data:   data,
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		return err
	}

	return nil
}

func decodeJson(r *http.Request, data any) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}

	return nil
}

func writeHeader(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

}

func (c *controller) SaveCtgHandler(w http.ResponseWriter, r *http.Request) {
	var ctg repository.Category
	writeHeader(w)

	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, ErrBadRequestMsg, nil)
	}

	if err := decodeJson(r, &ctg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, ErrBadRequestMsg, nil)

		return
	}

	save, err := c.service.SaveCategory(context.Background(), ctg)

	if err != nil {
		if err.Error() == validator.ValidationErrors.Error(err.(validator.ValidationErrors)) {
			w.WriteHeader(http.StatusBadRequest)
			encodeJson(w, http.StatusBadRequest, ErrBadRequestMsg, err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		encodeJson(w, http.StatusInternalServerError, ErrInternalServerMsg, err.Error())
		return
	}

	if err := encodeJson(w, http.StatusOK, "success save category", save); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeJson(w, http.StatusInternalServerError, ErrInternalServerMsg, err.Error())
		return
	}

}

func (c *controller) UpdateCtgHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var ctg service.CategoryUpdate
	writeHeader(w)
	if err := decodeJson(r, &ctg); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	category, err := c.service.UpdateCategory(context.Background(), id, ctg)

	if err != nil {
		if err.Error() == validator.ValidationErrors.Error(err.(validator.ValidationErrors)) {
			w.WriteHeader(http.StatusBadRequest)
			encodeJson(w, http.StatusBadRequest, ErrBadRequestMsg, err.Error())
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		encodeJson(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := encodeJson(w, http.StatusOK, "success update category", category); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		encodeJson(w, http.StatusInternalServerError, ErrInternalServerMsg, err.Error())
		return
	}
}

func (c *controller) DeleteCtgHandler(w http.ResponseWriter, r *http.Request) {
	writeHeader(w)
	id := mux.Vars(r)["id"]
	err := c.service.DeleteCategory(context.Background(), id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	encodeJson(w, http.StatusOK, "Success Deleted Item", fmt.Sprintf("item deleted with id %s", id))
}

func (c *controller) FIndByIdCtgHandler(w http.ResponseWriter, r *http.Request) {
	writeHeader(w)
	id := mux.Vars(r)["id"]
	ctgs, err := c.service.GetById(context.Background(), id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	encodeJson(w, http.StatusOK, "Success", ctgs)
}

func (c *controller) FindAllCtgHandler(w http.ResponseWriter, r *http.Request) {
	writeHeader(w)

	ctgs, err := c.service.GetAllCategory(context.Background())

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encodeJson(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	encodeJson(w, http.StatusOK, "Success", ctgs)
}
