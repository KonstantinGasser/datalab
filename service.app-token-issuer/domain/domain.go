package domain

import "github.com/KonstantinGasser/datalab/service.app-token-issuer/repo"

type AppTokenIssuer interface{}

type apptokenissuer struct {
	repo repo.Repo
}
