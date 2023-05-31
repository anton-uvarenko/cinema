package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/auth"
	general "github.com/anton-uvarenko/cinema/broker-service/protobufs/general"
	"net/http"
	"strconv"
	"time"
)

type AdminController struct {
	client auth.AdminHandlerClient
}

func NewAdminController(client auth.AdminHandlerClient) *AdminController {
	return &AdminController{
		client: client,
	}
}

func (c *AdminController) UpdateUserType(w http.ResponseWriter, r *http.Request) {
	payload := auth.UpdateUserTypePayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "error unmarshalling data", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err = c.client.UpdateUserType(ctx, &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}
}

func (c *AdminController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	payload := auth.DeleteUserPayload{}
	id, err := strconv.Atoi(r.URL.Query().Get("user_id"))
	if err != nil {
		http.Error(w, "bad user_id", http.StatusBadRequest)
		return
	}
	payload.UserId = int32(id)

	_, err = c.client.DelteUser(context.Background(), &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}
}

func (c *AdminController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := c.client.GetAllUsers(context.Background(), &general.Empty{})
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
