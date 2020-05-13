package captcha

import (
	"image"
	"testing"
)

func TestUserCaptchaFail(t *testing.T) {
	got := ChallengeUser(stubChallenger("52"), stubPrompter("43"))
	if got {
		t.Fatal("expected ChallengeUser to return false")
	}
}

func TestUserCaptchaSuccess(t *testing.T) {
	got := ChallengeUser(stubChallenger("52"), stubPrompter("52"))
	if !got {
		t.Fatal("expected ChallengeUser to return true")
	}
}

type stubChallenger string

func (s stubChallenger) Challenge() (image.Image, string) {
	return image.NewRGBA(image.Rect(0, 0, 100, 100)), string(s)
}

type stubPrompter string

func (p stubPrompter) Prompt(img image.Image) string {
	return string(p)
}
