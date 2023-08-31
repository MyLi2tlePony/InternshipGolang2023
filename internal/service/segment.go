package service

import "github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"

type SegmentRepository interface {
	Create(segment *entity.Segment) error
	Delete(segment *entity.Segment) error
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
