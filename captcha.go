package captcha

import (
	"crypto/subtle"
	"image"
)

type Challenger interface {
	Challenge() (image.Image, string)
}

type Prompter interface {
	Prompt(image.Image) string
}

func ChallengeUser(c Challenger, p Prompter) bool {
	img, expectedAns := c.Challenge()
	userAnswer := p.Prompt(img)
	if subtle.ConstantTimeEq(int32(len(expectedAns)), int32(len(userAnswer))) == 0 {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(expectedAns), []byte(userAnswer)) == 1
}
