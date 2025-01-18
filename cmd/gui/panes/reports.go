package panes

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

func ReportsScreen(win fyne.Window) fyne.CanvasObject {

	//config.DaysGet() moved to logon

	//var DetailsBorder = container.NewBorder(Details, nil, nil, nil, nil)
	//Header := widget.NewLabel("Reports: ")
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
