package test

import (
	"reflect"
	"testing"

	"github.com/crasico/go-task-tracker-cli/pkg/ui"
)

func TestNewBox(t *testing.T) {
	box := &ui.Box{
		Title:  "Test",
		Width:  1,
		Height: 2,
	}

	result := ui.NewBox("Test", 1, 2)

	if !reflect.DeepEqual(box, result) {
		t.Error("The created box does not equal the expected box")
		t.Fail()
	}
}

func TestRender(t *testing.T) {
	expected := `---Test---
|        |
----------`

	box := ui.NewBox("Test", 7, 3)

	result := box.Render()

	if !reflect.DeepEqual(expected, result) {
		t.Error("The ascii art did not render")
		t.Fail()
	}
}
