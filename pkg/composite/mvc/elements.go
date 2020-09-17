package mvc

import "fyne.io/fyne/widget"

// Элементы окна визуального представления.
type viewElms struct {
	progress *widget.Label
	label    *widget.Label
}

// Конструктор.
func newViewElms() viewElms {
	v := viewElms{}
	v.progress = widget.NewLabel("")
	v.label = widget.NewLabel("offline")

	return v
}
