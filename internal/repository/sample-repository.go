package repository

import "gorm.io/gorm"

type ISampleRepository interface {
	GetData() string
}

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{
		db: db,
	}
}

func (r *SampleRepository) GetData() string {
	return "sample data"
}
