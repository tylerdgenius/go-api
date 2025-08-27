package repository

import (
	"api-template/internal/data"

	"gorm.io/gorm"
)

type ISampleRepository interface {
	GetData() []data.SampleData
}

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{
		db: db,
	}
}

func (r *SampleRepository) GetData() []data.SampleData {
	return []data.SampleData{
		{ID: 1, Name: "Sample 1"},
		{ID: 2, Name: "Sample 2"},
		{ID: 3, Name: "Sample 3"},
	}
}
