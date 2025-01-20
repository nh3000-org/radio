package panes

import (
	"log"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

func ReportsScreen(win fyne.Window) fyne.CanvasObject {
	t := time.Now()

	ts := "YYYY-MM-DD 00:00:00"
	ts = strings.Replace(added, "YYYY", strconv.Itoa(t.Year()), 1)
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
	TrafEnd.SetPlaceHolder(te)
	TrafEnd.SetText(te)
	log.Println("TE1", te)
	trafficreport := widget.NewButton("Traffic", func() {
		config.TrafficStart = TrafStart.Text
		config.TrafficEnd = TrafEnd.Text
		config.ToPDF("TrafficReport", "ADMIN")
	})
	//trafgrid := container.New(layout.NewGridLayout(3), TrafStart, TrafEnd, trafficreport)

	daysreport := widget.NewButton("Days", func() {
		config.ToPDF("DaysList", "ADMIN")
	})
	hoursreport := widget.NewButton("Hours", func() {
		config.ToPDF("HoursList", "ADMIN")
	})
	categoriesreport := widget.NewButton("Categories", func() {
		config.ToPDF("CategoryList", "ADMIN")
	})
	schedulereport := widget.NewButton("Schedule", func() {
		config.ToPDF("ScheduleList", "ADMIN")
	})
	inventoryreport := widget.NewButton("Inventory", func() {
		log.Println("InventoryByCategoryFULL")
		config.ToPDF("InventoryByCategoryFULL", "ADMIN")
	})
	spinsperday := widget.NewButton("SpinsPerDay", func() {
		config.ToPDF("SpinsPerDay", "ADMIN")
	})
	spinsperweek := widget.NewButton("SpinsPerWeek", func() {
		config.ToPDF("SpinsPerWeek", "ADMIN")
	})
	spinstotal := widget.NewButton("SpinsTotal", func() {
		config.ToPDF("SpinsTotal", "ADMIN")
	})

	RAC := widget.NewAccordion(
		widget.NewAccordionItem("TrafficReport", trafficreport),
		widget.NewAccordionItem("Days Report", daysreport),
		widget.NewAccordionItem("Hours Report", hoursreport),
		widget.NewAccordionItem("Categories Report", categoriesreport),
		widget.NewAccordionItem("Schedule Report", schedulereport),
		widget.NewAccordionItem("Inventory Report", inventoryreport),
		widget.NewAccordionItem("Spins per Day", spinsperday),
		widget.NewAccordionItem("Spins per Week", spinsperweek),
		widget.NewAccordionItem("Spins Total", spinstotal),
	)

	return container.NewBorder(
		RAC,
		nil,
		nil,
		nil,
		nil,
	)

}
