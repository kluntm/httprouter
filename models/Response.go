package models

import (
	"net/http"
	"encoding/json"
)

type JsonResponse struct {
	Meta interface{} `json:"meta"`
	Data interface{} `json:data`
}

type JsonErrResponse struct {
	Error *ApiErrResponse `json:"error"`
}

type ApiErrResponse struct {
	Status int `json:"status"`
	Message string `json:"Message"`
}


func SendJsonResponse(w http.ResponseWriter, data interface{}) {
	//设置header头
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//设置http状态码
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(&JsonResponse{Data: data}); err != nil {
		SendJsonErrResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}

}

func SendJsonErrResponse(w http.ResponseWriter, errCode int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errCode)

	json.NewEncoder(w).Encode(&JsonErrResponse{
		Error: &ApiErrResponse{
			Status: errCode,
			Message:message,
		},
	})
}