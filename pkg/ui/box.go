package ui

type Box struct {
	Title  string
	Height int64
	Width  int64
}

func NewBox(title string, height int64, width int64) *Box {
	return &Box{
		Title:  title,
		Height: height,
		Width:  width,
	}
}

func (box Box) Render() string {
	return ""
}
