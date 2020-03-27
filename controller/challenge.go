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
	GetAlphanumericChallenge(length int, size int) ([]byte, string, error)
}

type challengeController struct {
	sessionRepo repo.SessionRepo
}

func NewChallengeController(sessionRepo repo.SessionRepo) ChallengeController {
	return &challengeController{
		sessionRepo: sessionRepo,
	}
}

func (c challengeController) GetImagesChallenge() {
}

func (c challengeController) GetAlphanumericChallenge(length int, size int) ([]byte, string, error) {
	answer := util.RandStringRunes(length)
	img, err := util.CreateTextImage(answer, size)
	if err != nil {
		return nil, "", err
	}
	answer = strings.ReplaceAll(answer, " ", "")
	answer = strings.ToLower(answer)

	sessionID := uuid.New().String()
	err = c.sessionRepo.CreateSession(sessionID, repo.SessionType_Alphanumeric, answer)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), "Enter all characters shown in order left to right.", nil
}

func (c challengeController) GetQuestionsChallenge() {
}