package server

import (
	"context"
	"github.com/marcsj/ocaptchas/challenge"
	"github.com/marcsj/ocaptchas/controller"
)

type challengeServer struct {
	controller controller.ChallengeController
}

func NewChallengeServer(controller controller.ChallengeController) challenge.ChallengeServer {
	return &challengeServer{
		controller: controller,
	}
}

func (s challengeServer) GetImagesChallenge(
	ctx context.Context, req *challenge.GetImagesRequest) (*challenge.GetImagesResponse, error) {
	sessionID, images, prompt, err := s.controller.GetImagesChallenge(int(req.GetNumber()), req.GetLabel())
	if err != nil {
		return nil, err
	}
	return &challenge.GetImagesResponse{
		SessionId: sessionID,
		Images: images,
		Prompt: prompt,
	}, nil
}

func (s challengeServer) GetAlphanumericChallenge(
	ctx context.Context, req *challenge.GetAlphanumericRequest) (*challenge.GetAlphanumericResponse, error) {
	session, img, prompt, err := s.controller.
		GetAlphanumericChallenge(int(req.GetLength()), int(req.GetSize()))
	if err != nil {
		return nil, err
	}
	return &challenge.GetAlphanumericResponse{
		SessionId: session,
		Image:     img,
		Prompt:    prompt,
	}, nil
}

func (s challengeServer) GetQuestionsChallenge(
	ctx context.Context, req *challenge.GetQuestionsRequest) (*challenge.GetQuestionsResponse, error) {
	session, questions, err := s.controller.GetQuestionsChallenge(int(req.GetNumber()), req.GetLabel())
	if err != nil {
		return nil, err
	}
	return &challenge.GetQuestionsResponse{
		SessionId: session,
		Questions: questions,
	}, nil
}

func (s challengeServer) SolveSession(
	ctx context.Context, req *challenge.SessionSolution) (*challenge.SolutionResponse, error) {
	return nil, s.controller.SolveChallenge(req.GetUuid(), req.GetAnswer())
}
