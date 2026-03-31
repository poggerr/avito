package app

import (
	"encoding/json"
	"io"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func decodeJSONBody(req *http.Request, dst any) error {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, dst)
}

func writeJSON(res http.ResponseWriter, status int, payload any) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	if payload == nil {
		return
	}

	if err := json.NewEncoder(res).Encode(payload); err != nil {
		http.Error(res, "failed to encode response", http.StatusInternalServerError)
	}
}

func writeError(res http.ResponseWriter, status int, msg string) {
	writeJSON(res, status, errorResponse{Error: msg})
}
