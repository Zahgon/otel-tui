package log

import (
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/tui/component/layout"
)

type body struct {
	view *tview.TextView
}

func newBody(
	commands *tview.TextView,
	resizeManager *layout.ResizeManager,
) *body {
	_ = "STUB: not implemented"
	return nil
}

func (b *body) flush() { _ = "STUB: not implemented"; return }

func (b *body) update(body string) { _ = "STUB: not implemented"; return }

func (b *body) registerCommands(commands *tview.TextView, resizeManager *layout.ResizeManager) {
	_ = "STUB: not implemented"
	return
}
