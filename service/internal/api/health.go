package api

import (
	"encoding/json"
	"net/http"
)

var health = &Health{
	Status: "ok",
}

type Health struct {
	Status string `json:"status"`
}

func GetHealthHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	content, _ := json.Marshal(health)
	_, _ = w.Write(content)
}
