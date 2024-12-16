package sorter

import "go.uber.org/zap"

type Stack string

const (
	Standard Stack = "STANDARD"
	Special  Stack = "SPECIAL"
	Rejected Stack = "REJECTED"
)

func (s Stack) String() string {
	return string(s)
}

type Package struct {
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Length int    `json:"length"`
	Mass   int    `json:"mass"`
}

func (p *Package) Volume() int {
	return p.Width * p.Height * p.Length
}

type Sorter struct {
	MaxSingleDimension int
	MaxVolume          int
	MaxMass            int
}

func NewSorter(maxSingleDimension, maxVolume, maxMass int) *Sorter {
	return &Sorter{
		MaxSingleDimension: maxSingleDimension,
		MaxVolume:          maxVolume,
		MaxMass:            maxMass,
	}
}

func (s *Sorter) IsBulky(p *Package) bool {
	return p.Volume() >= s.MaxVolume || p.Width >= s.MaxSingleDimension || p.Height >= s.MaxSingleDimension || p.Length >= s.MaxSingleDimension
}

func (s *Sorter) IsHeavy(p *Package) bool {
	return p.Mass >= s.MaxMass
}

func (s *Sorter) Sort(width, length, height, mass int) string {
	p := &Package{
		Width:  width,
		Height: height,
		Length: length,
		Mass:   mass,
	}

	bulky := s.IsBulky(p)
	heavy := s.IsHeavy(p)

	if bulky && heavy {
		zap.L().Debug("rejected package", zap.Any("package", p), zap.Bool("bulky", bulky), zap.Bool("heavy", heavy))
		return Rejected.String()
	}

	if bulky || heavy {
		zap.L().Debug("rejected package", zap.Any("package", p), zap.Bool("bulky", bulky), zap.Bool("heavy", heavy))
		return Special.String()
	}

	zap.L().Debug("rejected package", zap.Any("package", p), zap.Bool("bulky", bulky), zap.Bool("heavy", heavy))
	return Standard.String()
}
