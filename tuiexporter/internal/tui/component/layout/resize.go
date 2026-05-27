package layout

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type ResizeDirection int

const (
	ResizeDirectionHorizontal ResizeDirection = iota
	ResizeDirectionVertical

	WidenHorizontallyKey  = tcell.KeyCtrlL
	NarrowHorizontallyKey = tcell.KeyCtrlH
	WidenVerticallyKey    = tcell.KeyCtrlJ
	NarrowVerticallyKey   = tcell.KeyCtrlK
)

type ResizeManager struct {
	direction        ResizeDirection
	parent           *tview.Flex
	first, second    tview.Primitive
	firstProportion  int
	secondProportion int
	commands         *tview.TextView
}

func NewResizeManager(direction ResizeDirection) *ResizeManager {
	_ = "STUB: not implemented"
	return nil
}

func (m *ResizeManager) Register(
	parent *tview.Flex,
	first, second tview.Primitive,
	firstProportion, secondProportion int,
	commands *tview.TextView,
) {
	_ = "STUB: not implemented"
	return
}

func (m *ResizeManager) KeyMaps() KeyMaps { _ = "STUB: not implemented"; return *new(KeyMaps) }

// Ctrl-H is often interpreted as backspace by terminals

// Some terminals use KeyBackspace2 for backspace

func (m *ResizeManager) resize() { _ = "STUB: not implemented"; return }
