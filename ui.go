package main

import (
	"github.com/andlabs/ui"
)

// The main window
var window *ui.Window
var box *ui.Box
var area *ui.Area
var mainText *ui.AttributedString
var height = 1024
var width = 1024

type areaHandler struct {
}

func (areaHandler) DragBroken(a *ui.Area) {

}
func (areaHandler) MouseCrossed(a *ui.Area, left bool) {

}
func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {

}
func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) bool {
	return false
}
func (areaHandler) Draw(a *ui.Area, dp *ui.AreaDrawParams) {
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String: mainText,
		DefaultFont: &ui.FontDescriptor{
			Family:  ui.TextFamily("Times New Roman"),
			Size:    ui.TextSize(13),
			Stretch: ui.TextStretchNormal,
			Weight:  ui.TextWeightNormal,
			Italic:  ui.TextItalicNormal,
		},
		Width: dp.AreaWidth,
		Align: ui.DrawTextAlign(ui.AlignFill),
	})
	defer tl.Free()
	dp.Context.Text(tl, 0, 0)
}

// setupUI setups the UI
func setupUI() {
	window = ui.NewWindow("Markdown Viewer", 640, 480, true)
	window.SetMargined(true)
	window.OnClosing(func(*ui.Window) bool {
		window.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		window.Destroy()
		return true
	})
	box = ui.NewVerticalBox()
	area = ui.NewArea(new(areaHandler))
	box.Append(area, true)
	mainText = ui.NewAttributedString("")
	window.SetChild(box)
	appendWithAttributes("\r\n   MarkdownViewer", ui.TextSize(50))
	area.QueueRedrawAll()
	realMain()
	window.Show()
}

func appendWithAttributes(str string, attrs ...ui.Attribute) {
	start := len(mainText.String())
	end := start + len(str)
	mainText.AppendUnattributed(str)
	for _, v := range attrs {
		mainText.SetAttribute(v, start, end)
	}
}

func restoreText() {
	mainText = ui.NewAttributedString("")
}
