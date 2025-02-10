package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/uuid"
	"github.com/recovery-flow/comtools/cifractx"
	"github.com/recovery-flow/comtools/httpkit"
	"github.com/recovery-flow/comtools/httpkit/problems"
	"github.com/recovery-flow/roles"
	"github.com/recovery-flow/sso-oauth/internal/config"
	"github.com/recovery-flow/sso-oauth/internal/service/events"
	"github.com/recovery-flow/tokens"
	"github.com/sirupsen/logrus"
)

func AdminRoleUpdate(w http.ResponseWriter, r *http.Request) {
	server, err := cifractx.GetValue[*config.Server](r.Context(), config.SERVER)
	if err != nil {
		logrus.Errorf("Failed to retrieve service configuration %s", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	log := server.Logger

	initiatorID, ok := r.Context().Value(tokens.UserIDKey).(uuid.UUID)
	if !ok {
		log.Warn("UserID not found in context")
		httpkit.RenderErr(w, problems.Unauthorized("User not authenticated"))
		return
	}

	InitiatorRoleStr, ok := r.Context().Value(tokens.RoleKey).(string)
	if !ok {
		log.Warn("Role not found in context")
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	InitiatorRole, err := roles.StringToRoleUser(InitiatorRoleStr)
	if err != nil {
		log.Errorf("Failed to parse Initiator updatedRole: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	updatedUserID, err := uuid.Parse(chi.URLParam(r, "user_id"))
	if err != nil {
		log.Errorf("Failed to parse user_id: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"user_id": validation.NewError("user_id", "invalid UUID"),
		})...)
		return
	}

	updatedRole, err := roles.StringToRoleUser(chi.URLParam(r, "updatedRole"))
	if err != nil {
		log.Errorf("Failed to parse role: %v", err)
		httpkit.RenderErr(w, problems.BadRequest(validation.Errors{
			"role": validation.NewError("role", "invalid role"),
		})...)
		return
	}

	if roles.CompareRolesUser(InitiatorRole, updatedRole) != 1 {
		log.Warn("User can't update role to higher level than his own")
		httpkit.RenderErr(w, problems.Forbidden("User can't update role to higher level"))
		return
	}

	user, err := server.SqlDB.Accounts.GetById(r, updatedUserID)
	if err != nil {
		log.Errorf("Failed to get user: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	userRole, err := roles.StringToRoleUser(user.Role)
	if err != nil {
		log.Errorf("Failed to parse user role: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	if roles.CompareRolesUser(InitiatorRole, userRole) == -1 {
		log.Warn("User can't update role of user with higher role than his own")
		httpkit.RenderErr(w, problems.Forbidden("User can't update role of user with higher role"))
		return
	}

	res, err := server.SqlDB.Accounts.UpdateRole(r, updatedUserID, updatedRole)
	if err != nil {
		log.Errorf("Failed to update role: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	event := events.RoleUpdated{
		Event:     "RoleUpdated",
		UserID:    res.ID.String(),
		Role:      res.Role,
		Timestamp: time.Now().UTC(),
	}

	body, err := json.Marshal(event)
	if err != nil {
		log.Errorf("error serializing event: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}
	err = server.Broker.Publish(
		server.Config.Rabbit.Exchange,
		"account",
		"account.role_updated",
		body)
	if err != nil {
		log.Errorf("error publishing event: %v", err)
		httpkit.RenderErr(w, problems.InternalError())
		return
	}

	log.Infof("Role updated for user %s to %s by user %s", updatedUserID, updatedRole, initiatorID)
	httpkit.Render(w, http.StatusOK)
}
