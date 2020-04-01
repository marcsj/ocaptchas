package repo

import (
	"github.com/jinzhu/gorm"
)

type ChallengeQuestionsRepo interface {
	GetChallengeQuestions(number int, label string) ([]*QuestionChallenge, error)
	CreateChallenge(label string, question string, answer string) (*QuestionChallenge, error)
	DeleteChallenge(id uint) error
}

type challengeQuestionsRepo struct {
	db *gorm.DB
}

func NewChallengeQuestionsRepo(db *gorm.DB) (ChallengeQuestionsRepo, error) {
	err := db.AutoMigrate(QuestionChallenge{}).Error
	if err != nil {
		return nil, err
	}
	return &challengeQuestionsRepo{
		db: db,
	}, nil
}

type QuestionChallenge struct {
	gorm.Model
	Label    string
	Question string
	Answer   string
}

func (r challengeQuestionsRepo) GetChallengeQuestions(
	number int, label string) ([]*QuestionChallenge, error) {
	challenges := make([]*QuestionChallenge, number)
	err := r.db.
		Find(challenges, "label = ?", label).
		Order(gorm.Expr("random()")).
		Error
	if err != nil {
		return nil, err
	}
	return challenges, nil
}

func (r challengeQuestionsRepo) CreateChallenge(
	label string, question string, answer string) (*QuestionChallenge, error) {
	newChallenge := &QuestionChallenge{}
	db := r.db.Create(
		&QuestionChallenge{
			Label:    label,
			Question: question,
			Answer:   answer,
		})
	if db.Error != nil {
		return nil, db.Error
	}
	err := db.Scan(newChallenge).Error
	if err != nil {
		return nil, err
	}
	return newChallenge, nil
}

func (r challengeQuestionsRepo) DeleteChallenge(id uint) error {
	return r.db.Delete(&QuestionChallenge{}, "id = ?", id).Error
}
