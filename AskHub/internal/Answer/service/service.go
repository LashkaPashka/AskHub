package service

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/LashkaPashka/AskHub/internal/Answer/model"
)

type Storage interface {
	Create(answer *model.Answer) (success bool, err error)
	GetByID(answerID uint) (answer *model.Answer, err error)
	Delete(answerID uint) (success bool, err error)
	FoundQuestion(ID uint) (exists bool, err error)
}

type Service struct {
	logger  *slog.Logger
	storage Storage
}

func New(storage Storage, logger *slog.Logger) *Service {
	return &Service{
		storage: storage,
		logger:  logger,
	}
}

func (s *Service) Create(answer *model.Answer) (bool, error) {
	const op = "askHub.answer.service.create"

	exists, err := s.storage.FoundQuestion(answer.QuestionID)
	if err != nil {
		s.logger.Error("failed to check question existence",
			slog.String("op", op),
			slog.Uint64("question_id", uint64(answer.QuestionID)),
			slog.String("error", err.Error()),
		)
		return false, fmt.Errorf("%s: check question existence: %w", op, err)
	}

	if !exists {
		s.logger.Warn("attempt to create answer for non-existent question",
			slog.String("op", op),
			slog.Uint64("question_id", uint64(answer.QuestionID)),
		)
		return false, errors.New("cannot create answer: question does not exist")
	}

	ok, err := s.storage.Create(answer)
	if err != nil {
		s.logger.Error("failed to create answer",
			slog.String("op", op),
			slog.Uint64("question_id", uint64(answer.QuestionID)),
			slog.String("error", err.Error()),
		)
		return false, fmt.Errorf("%s: create answer: %w", op, err)
	}

	return ok, nil
}

func (s *Service) GetByID(ID uint) (*model.Answer, error) {
	const op = "askHub.answer.service.create"

	answer, err := s.storage.GetByID(ID)
	if err != nil {
		s.logger.Error("failed to get answer",
			slog.String("op", op),
			slog.Uint64("question_id", uint64(ID)),
			slog.String("error", err.Error()),
		)
		return nil, fmt.Errorf("%s: get answer: %w", op, err)
	}

	return answer, nil
}

func (s *Service) Delete(ID uint) (success bool, err error) {
	const op = "askHub.answer.service.create"

	success, err = s.storage.Delete(ID)
	if err != nil {
		s.logger.Error("failed to delete answer",
			slog.String("op", op),
			slog.Uint64("question_id", uint64(ID)),
			slog.String("error", err.Error()),
		)
		return false, fmt.Errorf("%s: delete answer: %w", op, err)
	}

	return success, err
}
