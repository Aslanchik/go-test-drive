package controllers

import (
	"encoding/json"
	"fmt"
	"go-test-drive/mvc/services"
	"go-test-drive/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		apiErr := &utils.AppplicationError{
			Message:    fmt.Errorf("user_id must be a number"),
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(apiErr.StatusCode)

		res.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write([]byte(jsonValue))
		return
	}
	jsonValue, _ := json.Marshal(user)

	res.Write(jsonValue)
}
