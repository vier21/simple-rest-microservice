package middlewares

import (
	"encoding/json"
	"net/http"
)

type WebErrorResp struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func NotfoundHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	resp := WebErrorResp{
		Code:   http.StatusNotFound,
		Status: "404 status not found",
		Data:   "Data not found",
	}

	json.NewEncoder(w).Encode(resp)

}
