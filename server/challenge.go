package server

import (
	"context"
	"github.com/marcsj/ocaptchas/challenge"
)

type challengeServer struct {

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
	return nil, nil
}

func (s challengeServer) GetQuestionsChallenge(
	ctx context.Context, req *challenge.GetQuestionsRequest) (*challenge.GetQuestionsResponse, error) {
	return nil, nil
}