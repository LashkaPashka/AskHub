package storage

import (
	"errors"
	"log/slog"

	"github.com/LashkaPashka/AskHub/internal/Answer/model"
	"github.com/LashkaPashka/AskHub/pkg/db"
	"gorm.io/gorm"
)

type Storage struct {
	pool   *db.Db
	logger *slog.Logger
}

func New(connStr string, logger *slog.Logger) *Storage {
	return &Storage{
		pool:   db.NewDb(connStr),
		logger: logger,
	}
}

func (s *Storage) Create(answer *model.Answer) (success bool, err error) {
	const op = "askHub.answer.storage.create"

	result := s.pool.DB.Table("answers").Create(answer)
	if err = result.Error; err != nil {
		s.logger.Error("Falied storage create",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)
		return false, err
	}

	s.logger.Info("answer successfully created",
		slog.String("op", op),
		slog.Uint64("question_id", uint64(answer.QuestionID)),
		slog.Bool("success", true),
	)

	return true, err
}

func (s *Storage) GetByID(answerID uint) (answer *model.Answer, err error) {
	const op = "askHub.answer.storage.getbyid"

	result := s.pool.DB.Table("answers").First(&answer, answerID)
	if err = result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("Record not found",
				slog.String("op", op),
			)

			return nil, err
		}

		s.logger.Error("Falied storage create",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return nil, err
	}

	return answer, err
}

func (s *Storage) Delete(answerID uint) (success bool, err error) {
	const op = "askHub.answer.storage.delete"

	result := s.pool.DB.Table("answers").Delete(&model.Answer{}, answerID)
	if err = result.Error; err != nil {
		s.logger.Error("Falied storage delete",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return false, err
	}

	s.logger.Info("answer successfully deleted",
		slog.String("op", op),
		slog.Bool("success", true),
	)

	return true, err
}

func (s *Storage) FoundQuestion(ID uint) (exists bool, err error) {
	const op = "askHub.answer.foundquestion"

	err = s.pool.DB.
		Table("questions").
		Select("COUNT(1) > 0").
		Where("id = ?", ID).
		Scan(&exists).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.logger.Error("record not found",
				slog.String("op", op),
			)

			return false, err
		}

		s.logger.Error("Failed storage getbyid question",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return false, err
	}

	return exists, nil
}
