package navigation

import "github.com/rivo/tview"

var (
	focusFn func(tview.Primitive)
	showMFn func(tview.Primitive, string) *tview.TextView
	hideMFn func(tview.Primitive)
)

func Init(
	setFocusFn func(tview.Primitive),
	showModalFn func(tview.Primitive, string) *tview.TextView,
	hideModalFn func(tview.Primitive),
) {
	_ = "STUB: not implemented"
	return
}

func Focus(primitive tview.Primitive) { _ = "STUB: not implemented"; return }

func ShowModal(primitive tview.Primitive, title string) *tview.TextView {
	_ = "STUB: not implemented"
	return nil
}

func HideModal(primitive tview.Primitive) { _ = "STUB: not implemented"; return }
