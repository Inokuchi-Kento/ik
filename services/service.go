package services

import "blrpc/ent"

type AppService struct {
	client *ent.Client
}

func NewAppService(client *ent.Client) *AppService {
	return &AppService{client: client}
}
