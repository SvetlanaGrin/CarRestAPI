package service

import "carRestAPI/internal/repository"

type Service struct {
	RequestCarCatalog
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		RequestCarCatalog: NewReqService(repos.RequestCarCatalog),
	}
}
