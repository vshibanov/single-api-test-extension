package controller

import (
	"encoding/json"
	"net/http"
)

type TestController struct {
}

func (c *TestController) IndexAction(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode("Test extension controller")
}
