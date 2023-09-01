package service

import "github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"

type UserRepository interface {
	ChangeSegments(userID int, insertEntities, deleteEntities []entity.Segment) error
	GetSegments(userID int) ([]entity.Segment, error)
}

type UserService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) ChangeSegments(userID int, insertEntities, deleteEntities []entity.Segment) error {
	err := s.repository.ChangeSegments(userID, insertEntities, deleteEntities)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetSegments(userID int) ([]entity.Segment, error) {
	segments, err := s.repository.GetSegments(userID)
	if err != nil {
		return nil, err
	}

	return segments, nil
}
