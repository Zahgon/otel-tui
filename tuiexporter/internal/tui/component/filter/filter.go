package filter

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/ymtdzzz/otel-tui/tuiexporter/internal/telemetry"
)

type onInputEnterFn func(inputConfirmed string, sortType telemetry.SortType)
type onInputDoneFn func()
type onInputChangedFn func(text string)
type onSortTypeChangedFn func(inputConfirmed string, sortType telemetry.SortType)

type Filter struct {
	view                  *tview.InputField
	sortType              telemetry.SortType
	input, inputConfirmed string
	onInputEnterFn        onInputEnterFn
	onInputDoneFn         onInputDoneFn
	onInputChangedFn      onInputChangedFn
	onSortTypeChangedFn   onSortTypeChangedFn
}

func NewFilter(
	commands *tview.TextView,
	label string,
	onInputEnterFn onInputEnterFn,
	onInputDoneFn onInputDoneFn,
	onInputChangedFn onInputChangedFn,
	onSortTypeChangedFn onSortTypeChangedFn,
) *Filter {
	_ = "STUB: not implemented"
	return nil
}

func (f *Filter) onChangedFunc() func(text string) { _ = "STUB: not implemented"; return nil }

func (f *Filter) onDoneFunc() func(key tcell.Key) { _ = "STUB: not implemented"; return nil }

func (f *Filter) registerCommands(commands *tview.TextView) { _ = "STUB: not implemented"; return }

func (f *Filter) RotateSortType() { _ = "STUB: not implemented"; return }

func (f *Filter) InputConfirmed() string { _ = "STUB: not implemented"; return "" }

func (f *Filter) SortType() *telemetry.SortType { _ = "STUB: not implemented"; return nil }

func (f *Filter) View() *tview.InputField { _ = "STUB: not implemented"; return nil }
