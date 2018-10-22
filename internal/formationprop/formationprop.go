package formationprop

type FormationProp interface {
	Style() string
	Width() int
	FileSpacing() float64
	RankSpacing() float64
}

type Pprop struct {
	Pstyle       string  `json:"style"`
	Pwidth       int     `json:"width"`
	PfileSpacing float64 `json:"fileSpacing"`
	PrankSpacing float64 `json:"rankSpacing"`
}

func NewFormationProp(Pstyle string, Pwidth int, PfileSpacing float64, PrankSpacing float64) *Pprop {
	return &Pprop{Pstyle, Pwidth, PfileSpacing, PrankSpacing}
}

func (p *Pprop) Style() string {
	return p.Pstyle
}

func (p *Pprop) Width() int {
	return p.Pwidth
}

func (p *Pprop) FileSpacing() float64 {
	return p.PfileSpacing
}

func (p *Pprop) RankSpacing() float64 {
	return p.PrankSpacing
}
