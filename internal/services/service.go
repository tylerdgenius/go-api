package services

import (
	"api-template/internal/data"
	"api-template/internal/repository"
)

type ISampleService interface {
	// Define the methods that the SampleService should implement
}

type SampleService struct {
	sampleRepo repository.ISampleRepository
}

func NewSampleService(sampleRepo repository.ISampleRepository) *SampleService {
	return &SampleService{
		sampleRepo: sampleRepo,
	}
}

func (s *SampleService) SomeServiceMethod() []data.SampleData {
	return s.sampleRepo.GetData()
}
