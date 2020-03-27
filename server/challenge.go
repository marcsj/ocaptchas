package server

import (
	"context"
	"github.com/marcsj/ocaptchas/challenge"
	"github.com/marcsj/ocaptchas/controller"
)

type challengeServer struct {
	controller controller.ChallengeController
}

func NewChallengeServer() challenge.ChallengeServer {
	return challengeServer{}
}

func (s challengeServer) GetImagesChallenge(
	ctx context.Context, req *challenge.GetImagesRequest) (*challenge.GetImagesResponse, error) {
	return nil, nil
}

func (s challengeServer) GetAlphanumericChallenge(
	ctx context.Context, req *challenge.GetAlphanumericRequest) (*challenge.GetAlphanumericResponse, error) {
	img, prompt, err := s.controller.GetAlphanumericChallenge(int(req.GetLength()), int(req.GetSize()))
	if err != nil {
		return nil, err
	}
	return &challenge.GetAlphanumericResponse{
		Image: img, Prompt: prompt}, nil
}

func (s challengeServer) GetQuestionsChallenge(
	ctx context.Context, req *challenge.GetQuestionsRequest) (*challenge.GetQuestionsResponse, error) {
	return nil, nil
}