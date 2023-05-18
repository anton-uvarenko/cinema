package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/users"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type CommentController struct {
	client users.CommentsClient
}

func NewCommentController(client users.CommentsClient) *CommentController {
	return &CommentController{
		client: client,
	}
}

func (c *CommentController) AddComment(w http.ResponseWriter, r *http.Request) {
	payload := users.CommentPayload{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	resp, err := c.client.AddComment(ctx, &payload)
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
