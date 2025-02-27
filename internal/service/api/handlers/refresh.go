package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/sso-oauth/internal/service/api/requests"
	"github.com/recovery-flow/sso-oauth/internal/service/api/responses"
	"github.com/recovery-flow/tokens"
)

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	req, err := requests.NewRefresh(r)
	if err != nil {
		h.Log.WithError(err).Warn("Error parsing request")
		httpkit.RenderErr(w, problems.BadRequest(err)...)
		return
	}
	curToken := req.Data.Attributes.RefreshToken

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		h.Log.Warn("Missing Authorization header")
		httpkit.RenderErr(w, problems.Unauthorized("Missing Authorization header"))
		return
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		h.Log.Warn("Invalid Authorization header format")
		httpkit.RenderErr(w, problems.Unauthorized("Invalid Authorization header format"))
		return
	}

	tokenString := parts[1]

	accountData, err := tokens.VerifyJWT(r.Context(), tokenString, h.Config.JWT.AccessToken.SecretKey)
	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		h.Log.WithError(err).Warn("Error validating token")
		httpkit.RenderErr(w, problems.Unauthorized())
		return
	}
	if accountData == nil {
		h.Log.Warn("Token validation failed, no account data")
		httpkit.RenderErr(w, problems.Unauthorized("Token validation failed"))
		return
	}

	sessionID := accountData.SessionID
	accountRole := accountData.Identity

	//-------------------------------------------------------------------------//

	session, err := h.Domain.SessionGetForAccount(r.Context(), *sessionID, *accountData.AccountID)
	if err != nil {
		h.Log.WithError(err).Error("Failed to get session")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	tokenAccess, tokenRefresh, err := h.Domain.SessionRefresh(r.Context(), *session, accountRole, r.RemoteAddr, r.UserAgent(), curToken)
	if err != nil {
		h.Log.WithError(err).Error("Failed to refresh session")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	httpkit.Render(w, responses.TokensPair(*tokenAccess, *tokenRefresh))
}
