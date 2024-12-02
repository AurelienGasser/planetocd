package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/aureliengasser/planetocd/server/likes"
	"github.com/gorilla/mux"
)

type LikeArticleResponse struct {
	LikeID       int   `json:"likeId"`
	RandomNumber int32 `json:"randomNumber"`
}

func handleLikeArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	likeID, randomNumber, err := likes.Save(id, r.RemoteAddr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	toMarshall := LikeArticleResponse{
		LikeID:       likeID,
		RandomNumber: randomNumber,
	}
	jsonBytes, err := json.Marshal(toMarshall)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	w.Write(jsonBytes)
}

type UpdateArticleLikeRequest struct {
	LikeID       int    `json:"likeId"`
	RandomNumber int32  `json:"randomNumber"`
	UserName     string `json:"userName"`
}

func handleUpdateArticleLike(w http.ResponseWriter, r *http.Request) {
	var req UpdateArticleLikeRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	err = likes.Update(req.LikeID, req.RandomNumber, req.UserName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
