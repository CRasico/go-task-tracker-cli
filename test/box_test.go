package test

import (
	"reflect"
	"testing"

	"github.com/crasico/go-task-tracker-cli/pkg/ui"
)

func TestNewBox(t *testing.T) {
	want := &ui.Box{
		Title:   "Test",
		Padding: 1,
	}

	got := ui.NewBox("Test", 1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %q box but expected %q`, got, want)
		t.Fail()
	}
}
