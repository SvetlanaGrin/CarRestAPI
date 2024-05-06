package service

import (
	"carRestAPI/internal/models"
	"carRestAPI/internal/repository"
)

type RequestCarCatalog interface {
	Create(input []models.Car) error
	Delete(regNum string) error
	GetAll(params models.Params) ([]models.Car, error)
	Update(regNum string, input models.UpdateCar) error
}

type CarCatalogService struct {
	repos repository.RequestCarCatalog
}

func NewReqService(repos repository.RequestCarCatalog) *CarCatalogService {
	return &CarCatalogService{repos: repos}
}

func (cp *CarCatalogService) Create(input []models.Car) error {
	return cp.repos.Create(input)
}
func (cp *CarCatalogService) Delete(regNum string) error {
	return cp.repos.Delete(regNum)
}
func (cp *CarCatalogService) GetAll(params models.Params) ([]models.Car, error) {
	return cp.repos.GetAll(params)
}

func (cp *CarCatalogService) Update(regNum string, input models.UpdateCar) error {
	return cp.repos.Update(regNum, input)
}
