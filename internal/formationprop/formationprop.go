package formationprop

type FormationProp interface {
	Style() string
	Width() int
	FileSpacing() float64
	RankSpacing() float64
}

type prop struct {
	style       string
	width       int
	fileSpacing float64
	rankSpacing float64
}

func NewFormationProp(style string, width int, fileSpacing float64, rankSpacing float64) FormationProp {
	return &prop{style, width, fileSpacing, rankSpacing}
}

func (p *prop) Style() string {
	return p.style
}

func (p *prop) Width() int {
	return p.width
}

func (p *prop) FileSpacing() float64 {
	return p.fileSpacing
}

func (p *prop) RankSpacing() float64 {
	return p.rankSpacing
}
