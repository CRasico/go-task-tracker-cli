package ui

type Box struct {
	Title   string
	Padding int64
}

func NewBox(title string, padding int64) *Box {
	return &Box{
		Title:   title,
		Padding: padding,
	}
}
