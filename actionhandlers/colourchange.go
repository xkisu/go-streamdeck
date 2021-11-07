package actionhandlers

import (
	"image/color"

	streamdeck "github.com/xkisu/go-streamdeck"
	"github.com/xkisu/go-streamdeck/buttons"
)

type ColourChangeAction struct {
	NewColour color.Color
}

func (action *ColourChangeAction) Pressed(btn streamdeck.Button) {
	mybtn := btn.(*buttons.ColourButton)
	mybtn.SetColour(action.NewColour)
}

func NewColourChangeAction(newColour color.Color) *ColourChangeAction {
	return &ColourChangeAction{NewColour: newColour}
}
