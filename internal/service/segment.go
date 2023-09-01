package service

import (
	"fmt"
	"github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"
)

type SegmentRepository interface {
	Create(segment *entity.Segment) error
	Delete(segment *entity.Segment) error
	CreateCSV(fileName string, year, month int) error
}

type SegmentService struct {
	repository SegmentRepository
}

func NewSegmentService(repository SegmentRepository) *SegmentService {
	return &SegmentService{
		repository: repository,
	}
}

func (s *SegmentService) Create(segment *entity.Segment) error {
	err := s.repository.Create(segment)
	if err != nil {
		return err
	}

	return nil
}

func (s *SegmentService) Delete(segment *entity.Segment) error {
	err := s.repository.Delete(segment)
	if err != nil {
		return err
	}

	return nil
}

func (s *SegmentService) CreateCSV(year, month int) (string, error) {
	file := fmt.Sprintf("%d%d.csv", year, month)
	err := s.repository.CreateCSV("csv/"+file, year, month)
	if err != nil {
		return "", nil
	}

	return file, nil
}
