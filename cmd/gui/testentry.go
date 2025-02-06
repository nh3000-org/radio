package main

import (
	"log"
	"strconv"
	"time"

	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type PaneTest struct {
	Title, Intro string
	Icon         fyne.Resource
	View         func(w fyne.Window) fyne.CanvasObject
	SupportWeb   bool
}

func mainxxx() {

	var a = app.NewWithID("org.nh3000.nh3000")
	var wi = a.NewWindow("NH3000")
	var PanesTest = map[string]PaneTest{
		"logon":     {"Logon", "", theme.LoginIcon(), widgettest, true},
		"reports":   {"Reports", "", theme.ListIcon(), widgettest, true},
		"inventory": {"Inventory", "", theme.ListIcon(), widgettest, true},
	}

	wi.SetContent(container.NewAppTabs(
		container.NewTabItemWithIcon(PanesTest["logon"].Title, PanesTest["logon"].Icon, widget.NewLabel("log")),
		container.NewTabItemWithIcon(PanesTest["reports"].Title, PanesTest["reports"].Icon, widgettest(wi)),
		container.NewTabItemWithIcon(PanesTest["inventory"].Title, PanesTest["inventory"].Icon, widget.NewLabel("inv")),
	))

	wi.Resize(fyne.NewSize(640, 480))
	wi.ShowAndRun()

}
func widgettest(win fyne.Window) fyne.CanvasObject {
	t := time.Now()

	ts := "YYYY-MM-DD 00:00:00"
	ts = strings.Replace(ts, "YYYY", strconv.Itoa(t.Year()), 1)
	ms := strconv.Itoa(int(t.Month()))
	if len(ms) == 1 {
		ms = "0" + ms
	}
	ts = strings.Replace(ts, "MM", ms, 1)
	ds := strconv.Itoa(int(t.Day()))
	if len(ds) == 1 {
		ds = "0" + ds
	}
	ts = strings.Replace(ts, "DD", ds, 1)
	TrafStart := widget.NewEntry()
	TrafStart.SetPlaceHolder(ts)
	TrafStart.SetText(ts)

	te := "YYYY-MM-DD 23:59:59"
	log.Println("TE", te)
	te = strings.Replace(te, "YYYY", strconv.Itoa(t.Year()), 1)
	me := strconv.Itoa(int(t.Month()))
	if len(me) == 1 {
		me = "0" + me
	}
	te = strings.Replace(te, "MM", me, 1)
	de := strconv.Itoa(int(t.Day()))
	if len(de) == 1 {
		de = "0" + de
	}
	te = strings.Replace(te, "DD", de, 1)
	TrafEnd := widget.NewEntry()
	//TrafEnd.SetPlaceHolder(te)
	TrafEnd.SetText(te)
	trafgrid := container.New(layout.NewGridLayout(3), TrafStart, TrafEnd)
	RAC := widget.NewAccordion(
		widget.NewAccordionItem("TrafficReport", widget.NewLabel("Traffic Top")),
		widget.NewAccordionItem("TrafficReport", trafgrid),
		widget.NewAccordionItem("TrafficReport", widget.NewLabel("Traffic Bottom")),
	)
	return RAC
}
