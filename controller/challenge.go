package controller

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/marcsj/ocaptchas/repo"
	"github.com/marcsj/ocaptchas/util"
	"image/jpeg"
	"strings"
)

type ChallengeController interface {
	GetAlphanumericChallenge(length int, size int) (string, []byte, string, error)
	GetQuestionsChallenge(number int, label string) (string, []string, error)
}

type challengeController struct {
	sessionRepo repo.SessionRepo
	questionsRepo repo.ChallengeQuestionsRepo
}

func NewChallengeController(
	sessionRepo repo.SessionRepo, questionsRepo repo.ChallengeQuestionsRepo) ChallengeController {
	return &challengeController{
		sessionRepo: sessionRepo,
		questionsRepo: questionsRepo,
	}
}

func (c challengeController) GetImagesChallenge() {
}

func (c challengeController) GetAlphanumericChallenge(
	length int, size int) (string, []byte, string, error) {
	answer := util.RandStringRunes(length)
	img, err := util.CreateTextImage(answer, size)
	if err != nil {
		return "", nil, "", err
	}
	answer = strings.ReplaceAll(answer, " ", "")
	answer = strings.ToLower(answer)

	sessionID := getUUID()
	err = c.sessionRepo.CreateSession(sessionID, repo.SessionType_Alphanumeric, answer)
	if err != nil {
		return "", nil, "", err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return "", nil, "", err
	}
	return sessionID, buf.Bytes(), "Enter all characters shown in order left to right.", nil
}

func (c challengeController) GetQuestionsChallenge(
	number int, label string) (string, []string, error) {
	challenges, err := c.questionsRepo.GetChallengeQuestions(number, label)
	if err != nil {
		return "", nil, err
	}

	answers := make([]string, 0)
	questions := make([]string, 0)
	for _, challenge := range challenges {
		answers = append(answers, challenge.Answer)
		questions = append(questions, challenge.Question)
	}
	sessionID := getUUID()
	err = c.sessionRepo.CreateSession(
		sessionID, repo.SessionType_Questions, strings.Join(answers, ","))
	if err != nil {
		return "", nil, err
	}
	return sessionID, questions, nil
}

func getUUID() string {
	return uuid.New().String()
}