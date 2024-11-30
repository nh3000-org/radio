package panes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/nh3000-org/radio/config"
)

func RadioScreen(_ fyne.Window) fyne.CanvasObject {
	header := widget.NewLabel("")
	daysbutton := widget.NewButton(config.GetLangs("ra-days"), func() {
		header.SetText("Manage Days of Week")
	})
	categoriesbutton := widget.NewButton(config.GetLangs("ra-cats"), func() {
		header.SetText("Manage Schedule Categories")

	})
	schedulebutton := widget.NewButton(config.GetLangs("ra-sched"), func() {
		header.SetText("Manage Schedule")

	})
	inventorybutton := widget.NewButton(config.GetLangs("ra-inv"), func() {
		header.SetText("Manage Inventory")

	})

	var DetailsBorder = container.NewBorder(header, nil, nil, nil, nil)

	DetailsVW := container.NewScroll(DetailsBorder)
	bgrid := container.NewGridWithColumns(4, daysbutton, categoriesbutton, schedulebutton, inventorybutton)

	return container.NewBorder(bgrid, DetailsVW, nil, nil)

}
