package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/sso-oauth/internal/service/api/responses"
)

func AdminAccountGet(w http.ResponseWriter, r *http.Request) {
	accountID, err := uuid.Parse(chi.URLParam(r, "account_id"))
	if err != nil {
		Log(r).WithError(err).Error("Failed to parse account_id")
		httpkit.RenderErr(w, problems.BadRequest(fmt.Errorf("invalid account_id"))...)
		return
	}

	account, err := Domain(r).AccountGet(r.Context(), accountID)
	if err != nil {
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	httpkit.Render(w, responses.Account(*account))
}
