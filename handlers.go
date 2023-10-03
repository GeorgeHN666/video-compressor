package main

import (
	"fmt"
	"net/http"
)

func StorePost(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	w.Header().Set("Content-Type", "multipart/form-data")

	var Post Post

	Post.Avatar = r.FormValue("avatar")
	Post.Description = r.FormValue("description")
	Post.Title = r.FormValue("title")

	videos, err := CompressVideos(r)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	Post.Videos = videos
	fmt.Println("About to store videos")
	err = StartDB().InsertPost(&Post)
	if err != nil {
		return
	}

	res.Error = false
	res.Message = "Post inserted"
	WriteJSON(w, r, http.StatusOK, res)

}

func Feed(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool    `json:"error"`
		Message string  `json:"message"`
		Posts   []*Post `json:"posts"`
	}

	posts, err := StartDB().GetPosts()
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Posts = posts

	WriteJSON(w, r, http.StatusOK, res)

}

func Test(w http.ResponseWriter, r *http.Request) {

	var res struct {
		Error   bool   `json:"error"`
		Message string `json:"message"`
	}

	var Test struct {
		Name string `json:"name"`
		Val  string `json:"val"`
	}

	err := ReadJSON(w, r, &Test)
	if err != nil {
		res.Error = true
		res.Message = err.Error()
		WriteJSON(w, r, http.StatusBadRequest, res)
		return
	}

	res.Error = false
	res.Message = "All good"

	WriteJSON(w, r, http.StatusOK, res)

}
