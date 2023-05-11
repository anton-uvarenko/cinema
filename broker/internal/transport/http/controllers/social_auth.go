package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/auth"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type SocialAuthController struct {
	client auth.SocialAuthClient
}

func NewSocialAuthController(client auth.SocialAuthClient) *SocialAuthController {
	return &SocialAuthController{
		client: client,
	}
}

func (c *SocialAuthController) GoogleAuth(w http.ResponseWriter, r *http.Request) {
	payload := &auth.SocialAuthPayload{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, "error decoding payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := c.client.GoogleAuth(ctx, payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}

func (c *SocialAuthController) FacebookAuth(w http.ResponseWriter, r *http.Request) {
	payload := &auth.SocialAuthPayload{}
	err := json.NewDecoder(r.Body).Decode(payload)
	if err != nil {
		http.Error(w, "error decoding payload", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := c.client.FacebookAuth(ctx, payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
