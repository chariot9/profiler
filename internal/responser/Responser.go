package responser

import (
	"encoding/json"
	"net/http"
)

func Respond(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Add("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
