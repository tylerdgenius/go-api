package handlers

import (
	"api-template/internal/utils"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, http.StatusOK, "Test endpoint is working fine!", nil)
}
