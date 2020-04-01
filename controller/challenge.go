package controller

import (
	"bytes"
	"errors"
	"github.com/marcsj/ocaptchas/repo"
	"github.com/marcsj/ocaptchas/util"
	"image"
	"image/jpeg"
	"strings"
)

type ChallengeController interface {
	GetImagesChallenge(number int, label string) (string, []image.Image, string, error)
	GetAlphanumericChallenge(length int, size int) (string, []byte, string, error)
	GetQuestionsChallenge(number int, label string) (string, []string, error)
	SolveChallenge(sessionID string, answer []string) error
}

type challengeController struct {
	sessionRepo repo.SessionRepo
	questionsRepo repo.ChallengeQuestionsRepo
	imagesRepo repo.ChallengeImagesRepo
}

func NewChallengeController(
	sessionRepo repo.SessionRepo, questionsRepo repo.ChallengeQuestionsRepo) ChallengeController {
	return &challengeController{
		sessionRepo: sessionRepo,
		questionsRepo: questionsRepo,
	}
}

func (c challengeController) GetImagesChallenge(number int, label string) (
	session string, images []image.Image, prompt string, err error) {
	images, answer, err := c.imagesRepo.GetChallengeImages(number, label)
	if err != nil {
		return
	}
	session = util.GetUUID()
	err = c.sessionRepo.CreateSession(session, repo.SessionType_Images, answer)
	if err != nil {
		return
	}
	prompt = "images with"
	return
}

func (c challengeController) GetAlphanumericChallenge(
	length int, size int) (
	session string, img []byte, prompt string, err error) {
	img, answer, err := c.createAlphanumericImage(length, size)
	if err != nil {
		return
	}
	session = util.GetUUID()
	err = c.sessionRepo.CreateSession(session, repo.SessionType_Alphanumeric, answer)
	if err != nil {
		return
	}
	prompt = "all shown characters"
	return
}

func (c challengeController) GetQuestionsChallenge(
	number int, label string) (
	session string, questions []string, err error) {
	challenges, err := c.questionsRepo.GetChallengeQuestions(number, label)
	if err != nil {
		return
	}
	answers := make([]string, 0)
	questions = make([]string, 0)
	for _, challenge := range challenges {
		answers = append(answers, challenge.Answer)
		questions = append(questions, challenge.Question)
	}
	session = util.GetUUID()
	err = c.sessionRepo.CreateSession(
		session, repo.SessionType_Questions, strings.Join(answers, ","))
	if err != nil {
		return
	}
	return
}

func (c challengeController) SolveChallenge(sessionID string, answer []string) error {
	session, err := c.sessionRepo.GetSession(sessionID)
	if err != nil {
		return err
	}
	if session.Answer != strings.ToLower(strings.Join(answer, ",")) {
		return errors.New("answer incorrect")
	}
	return nil
}

func (c challengeController) createAlphanumericImage(length int, size int) (
	b []byte, answer string, err error) {
	answer = util.RandStringRunes(length)
	img, err := util.CreateTextImage(answer, size)
	if err != nil {
		return
	}
	answer = strings.ReplaceAll(answer, " ", "")
	answer = strings.ToLower(answer)
	buf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		return
	}
	b = buf.Bytes()
	return
}
