package handlers

import (
	"api-template/internal/utils"
	"net/http"
)

func (h *BaseHandler) SampleHandler(w http.ResponseWriter, r *http.Request) {
	utils.HandleResponse(w, http.StatusOK, "Test endpoint is working fine!", nil)
}
