package controllers

import (
	"encoding/json"
	"net/http"
	"profiler/internal/message"
	"profiler/internal/responser"
	u "profiler/internal/user"
)

var CreateUser = func(w http.ResponseWriter, r *http.Request) {
	user := &u.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		responser.Respond(w, message.Message(false, "NG0001", "Invalid request"))
		return
	}

	res := user.Create()
	responser.Respond(w, res)
}
