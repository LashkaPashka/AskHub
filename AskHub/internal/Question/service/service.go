package service

import (
	"github.com/LashkaPashka/AskHub/internal/Question/model"
)

type Storage interface {
	Create(question *model.Question) (success bool, err error)
	GetByID(ID uint) (question *model.Question, err error)
	GetAll() (questions []model.Question, err error)
	Delete(ID uint) (success bool, err error)
}

type Service struct {
	storage Storage
}

func New(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) Create(question *model.Question) (bool, error) {
	return s.storage.Create(question)
}

func (s *Service) GetByID(ID uint) (*model.Question, error) {
	return s.storage.GetByID(ID)
}

func (s *Service) GetAll() ([]model.Question, error) {
	return s.storage.GetAll()
}

func (s *Service) Delete(ID uint) (bool, error) {
	return s.storage.Delete(ID)
}
