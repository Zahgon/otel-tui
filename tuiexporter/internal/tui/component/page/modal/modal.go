package modal

import (
	"github.com/rivo/tview"
)

const ModalTitle = "Scroll (Ctrl+J, Ctrl+K)"

type ModalPage struct {
	view     *tview.Flex
	textView *tview.TextView
}

func NewModalPage() *ModalPage { _ = "STUB: not implemented"; return nil }

func (m *ModalPage) SetText(text string) { _ = "STUB: not implemented"; return }

func (m *ModalPage) GetPrimitive() tview.Primitive {
	_ = "STUB: not implemented"
	return *new(tview.Primitive)
}

func (m *ModalPage) ShowModalFunc(showModalPageFn func()) func(current tview.Primitive, text string) *tview.TextView {
	_ = "STUB: not implemented"
	return nil
}

func (m *ModalPage) HideModalFunc(hideModalPageFn func()) func(current tview.Primitive) {
	_ = "STUB: not implemented"
	return nil
}
