package layout

import (
	"regexp"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var keyMapRegex = regexp.MustCompile(`Rune|\[|\]`)

type KeyMap struct {
	Key         *tcell.EventKey
	Arrow       bool
	Hidden      bool
	Description string
	Handler     func(event *tcell.EventKey) *tcell.EventKey
}

type KeyMaps []*KeyMap

func (m *KeyMaps) Merge(m2 KeyMaps) { _ = "STUB: not implemented"; return }

func (m KeyMaps) keyTexts() string { _ = "STUB: not implemented"; return "" }

func getInt32Key(key *tcell.EventKey) int32 { _ = "STUB: not implemented"; return 0 }

type FocusableBox interface {
	SetFocusFunc(func()) *tview.Box
	SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey) *tview.Box
}

func NewCommandList() *tview.TextView { _ = "STUB: not implemented"; return nil }

func AttachCommandList(commands *tview.TextView, p tview.Primitive) *tview.Flex {
	_ = "STUB: not implemented"
	return nil
}

func RegisterCommandList(commands *tview.TextView, c FocusableBox, origFocusFn func(), keys KeyMaps) {
	_ = "STUB: not implemented"
	return
}
