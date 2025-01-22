package panes

import (
	"log"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

func ReportsScreen(win fyne.Window) fyne.CanvasObject {
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

	laalbum := widget.NewLabel("Campaign (Album): ")
	selalbum := widget.NewSelect(config.AlbumToArray(), func(string) {})

	trafficreport := widget.NewButton("Traffic", func() {
		donebutton := widget.NewButtonWithIcon("Done", theme.ContentCopyIcon(), func() {
			config.TrafficStart = TrafStart.Text
			config.TrafficEnd = TrafEnd.Text
			config.TrafficAlbum = selalbum.Selected
			config.ToPDF("TrafficReport", "ADMIN")

		})
		databox := container.NewGridWithRows((7),
			widget.NewLabel("Start"),
			TrafStart,
			widget.NewLabel("End"),
			TrafEnd,
			laalbum,
			selalbum,
			donebutton,
		)

		TrafStart.SetMinRowsVisible(5)

		TrafEnd.SetMinRowsVisible(5)

		dlg := fyne.CurrentApp().NewWindow("Select Traffic Range")

		dlg.SetContent(databox)
		dlg.Resize(fyne.NewSize(240, 180))
		dlg.Show()

	})

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
	RAC.Refresh()
	return container.NewBorder(
		RAC,
		nil,
		nil,
		nil,
		nil,
	)

}
