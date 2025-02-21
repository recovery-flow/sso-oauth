package responses

import (
	"github.com/recovery-flow/sso-oauth/internal/service/domain/models"
	"github.com/recovery-flow/sso-oauth/resources"
)

func Session(session models.Session) resources.Session {
	return resources.Session{
		Data: resources.SessionData{
			Type: resources.AccountSessionType,
			Id:   session.ID.String(),
			Attributes: resources.SessionAttributes{
				AccountId: session.AccountID.String(),
				Client:    session.Client,
				Ip:        session.IP,
				CreatedAt: session.CreatedAt,
				LastUsed:  session.LastUsed,
			},
		},
	}
}
