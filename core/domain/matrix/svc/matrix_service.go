package svc

import (
	"image"

	"github.com/KonstantinGasser/davinci/core/domain/matrix"
)

type Service interface {
	Print(img image.Image) error
	Animate(frames []*image.Paletted) error
}

type service struct {
	matrix matrix.Matrix
}

func New(m matrix.Matrix) Service {
	return &service{
		matrix: m,
	}
}

func (s service) Print(img image.Image) error {
	return s.matrix.Print(img)
}

func (s service) Animate(frames []*image.Paletted) error {
	return s.matrix.Animate(frames)
}
