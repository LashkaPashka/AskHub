package postgresql

import (
	"errors"
	"log/slog"

	"github.com/LashkaPashka/AskHub/internal/Question/model"
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

func (s *Storage) Create(question *model.Question) (success bool, err error) {
	const op = "Question.storage.create"

	result := s.pool.DB.Table("questions").Create(question)
	if result.Error != nil {
		s.logger.Error("failed to create question",
			slog.String("op", op),
			slog.String("error", result.Error.Error()),
			slog.Any("question", question),
			slog.Int64("rows_affected", result.RowsAffected),
		)

		return false, result.Error
	}

	s.logger.Info("question created successfully",
		slog.String("op", op),
		slog.Any("question", question),
		slog.Int64("rows_affected", result.RowsAffected),
	)

	return true, nil

}

func (s *Storage) GetByID(ID uint) (*model.Question, error) {
	const op = "Question.storage.getbyiD"

	var question model.Question

	result := s.pool.DB.Table("questions").First(&question, ID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			s.logger.Warn("question not found",
				slog.String("op", op),
				slog.Uint64("id", uint64(ID)),
			)
			return nil, result.Error
		}

		s.logger.Error("failed to get question by id",
			slog.String("op", op),
			slog.Uint64("id", uint64(ID)),
			slog.String("error", result.Error.Error()),
		)
		return nil, result.Error
	}

	s.logger.Info("question fetched successfully",
		slog.String("op", op),
		slog.Uint64("id", uint64(ID)),
	)

	return &question, nil
}

func (s *Storage) GetAll() ([]model.Question, error) {
	const op = "Question.storage.getall"

	var questions []model.Question

	result := s.pool.DB.Table("questions").Order("id ASC").Scan(&questions)
	if result.Error != nil {
		s.logger.Error("failed to fetch all questions",
			slog.String("op", op),
			slog.String("error", result.Error.Error()),
		)
		return nil, result.Error
	}

	s.logger.Info("questions fetched successfully",
		slog.String("op", op),
		slog.Int("count", len(questions)),
	)

	return questions, nil
}


func (s *Storage) Delete(ID uint) (bool, error) {
	const op = "Question.storage.delete"

	result := s.pool.DB.Table("questions").Delete(&model.Question{}, ID)
	if result.Error != nil {
		s.logger.Error("failed to delete question",
			slog.String("op", op),
			slog.Uint64("id", uint64(ID)),
			slog.String("error", result.Error.Error()),
		)
		return false, result.Error
	}

	if result.RowsAffected == 0 {
		s.logger.Warn("no question to delete",
			slog.String("op", op),
			slog.Uint64("id", uint64(ID)),
		)
		return false, nil
	}

	s.logger.Info("question deleted successfully",
		slog.String("op", op),
		slog.Uint64("id", uint64(ID)),
		slog.Int64("rows_affected", result.RowsAffected),
	)

	return true, nil
}