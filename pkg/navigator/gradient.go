package navigator

import "errors"

type Gradient struct {
	Colors   []string `json:"colors" yaml:"colors"`
	MinValue int      `json:"minValue" yaml:"minValue"`
	MaxValue int      `json:"maxValue" yaml:"maxValue"`
}

func NewGradient(colors []string, minValue, maxValue int) (*Gradient, error) {
	if minValue == maxValue {
		return nil, errors.New("minValue and maxValue cannot be equal")
	}
	if minValue > maxValue {
		return nil, errors.New("minValue cannot be greater than maxValue")
	}
	if len(colors) < 2 {
		return nil, errors.New("colors must contain at least two colors")
	}
	return &Gradient{
		Colors:   colors,
		MinValue: minValue,
		MaxValue: maxValue,
	}, nil
}
