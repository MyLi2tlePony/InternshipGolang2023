package service

import "github.com/MyLi2tlePony/AvitoInternshipGolang2023/internal/entity"

type UserRepository interface {
	ChangeSegments(userID int, insert, delete []string) error
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
	insertSegments := make([]string, 0, len(insertEntities))
	deleteSegments := make([]string, 0, len(deleteEntities))

	for i := range insertEntities {
		insertSegments = append(insertSegments, insertEntities[i].Name)
	}

	for i := range deleteEntities {
		deleteSegments = append(deleteSegments, deleteEntities[i].Name)
	}

	err := s.repository.ChangeSegments(userID, insertSegments, deleteSegments)
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
