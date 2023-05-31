package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/core"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/auth"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type PremiumController struct {
	client auth.AdminHandlerClient
}

func NewPremiumController(client auth.AdminHandlerClient) *PremiumController {
	return &PremiumController{client: client}
}

func (c *PremiumController) BuyPremium(w http.ResponseWriter, r *http.Request) {
	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	payload := auth.UpdateUserTypePayload{
		UserId:   int32(id),
		UserType: string(core.Premium),
	}

	jwt, err := c.client.UpdateUserType(context.Background(), &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	err = json.NewEncoder(w).Encode(&struct {
		Jwt string `json:"jwt"`
	}{
		jwt.Jwt,
	})
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
