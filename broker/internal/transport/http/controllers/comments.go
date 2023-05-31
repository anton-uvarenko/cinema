package controllers

import (
	"context"
	"encoding/json"
	"github.com/anton-uvarenko/cinema/broker-service/internal/pkg"
	"github.com/anton-uvarenko/cinema/broker-service/protobufs/users"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
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

	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	payload.UserId = int32(id)

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

func (c *CommentController) GetPublicComments(w http.ResponseWriter, r *http.Request) {
	filmId, err := strconv.Atoi(r.URL.Query().Get("filmId"))
	if err != nil {
		http.Error(w, "invalid film id", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "invalid film id", http.StatusBadRequest)
		return
	}
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		http.Error(w, "invalid film id", http.StatusBadRequest)
		return
	}
	respAmount, _ := strconv.Atoi(r.URL.Query().Get("resp_amount"))
	userId, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])

	payload := users.GetPublicCommentsPayload{
		FilmId:          int32(filmId),
		Page:            int32(page),
		Amount:          int32(amount),
		ResponsesAmount: int32(respAmount),
		UserId:          int32(userId),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	resp, err := c.client.GetPublicComments(ctx, &payload)
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

func (c *CommentController) GetPrivateComments(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	filmId, err := strconv.Atoi(r.URL.Query().Get("filmId"))
	if err != nil {
		http.Error(w, "invalid film id", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "invalid page", http.StatusBadRequest)
		return
	}
	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		http.Error(w, "invalid amount", http.StatusBadRequest)
		return
	}

	payload := users.GetPrivateCommentsPayload{
		UserId: int32(id),
		FilmId: int32(filmId),
		Page:   int32(page),
		Amount: int32(amount),
	}

	resp, err := c.client.GetPrivateComments(ctx, &payload)
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

func (c *CommentController) LikeComment(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	id, _ := pkg.ParseWithId(strings.Split(r.Header.Get("Authorization"), " ")[1])
	commentId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "wrong comment_id", http.StatusBadRequest)
		return
	}

	payload := users.LikeCommentPayload{
		UserId:    int32(id),
		CommentId: int32(commentId),
	}

	resp, err := c.client.LikeComment(ctx, &payload)
	if err != nil {
		fail := pkg.CustToPkgError(err.Error())
		http.Error(w, fail.Error(), fail.Code())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		logrus.Error(err)
		http.Error(w, "error encoding response", http.StatusInternalServerError)
		return
	}
}
