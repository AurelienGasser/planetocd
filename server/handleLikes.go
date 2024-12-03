package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/aureliengasser/planetocd/server/likes"
	"github.com/gorilla/mux"
	"github.com/x-way/crawlerdetect"
)

type LikeArticleResponse struct {
	LikeID       int   `json:"likeId"`
	RandomNumber int32 `json:"randomNumber"`
}

func handleLikeArticle(w http.ResponseWriter, r *http.Request) {
	if crawlerdetect.IsCrawler(strings.Join(r.Header[http.CanonicalHeaderKey("User-Agent")], "")) {
		http.NotFound(w, r)
		return
	}
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	ip := strings.Join(r.Header[http.CanonicalHeaderKey("CF-Connecting-IP")], ",")
	if ip == "" {
		ip = r.RemoteAddr
	}
	likeID, randomNumber, err := likes.Save(id, ip)

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
	if crawlerdetect.IsCrawler(strings.Join(r.Header[http.CanonicalHeaderKey("User-Agent")], "")) {
		http.NotFound(w, r)
		return
	}
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
