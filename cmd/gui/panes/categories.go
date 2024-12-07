package panes

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

func CategoriesScreen(win fyne.Window) fyne.CanvasObject {

	//config.HoursGet() moved to logon

	Details := widget.NewLabel("")

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laid := widget.NewLabel("Category: ")
	edid :=  widget.NewEntry()

	ladesc := widget.NewLabel("Description: ")
	eddesc := widget.NewEntry()

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laid, edid)
	griddesc := container.New(layout.NewGridLayoutWithRows(2), ladesc, eddesc)

	saveaddbutton := widget.NewButtonWithIcon("Add Hour Part", theme.ContentCopyIcon(), func() {

		config.CategoriesAdd(edid.Text, eddesc.Text)
		config.CategoriesGet()
	})
	List := widget.NewList(
		func() int {
			return len(config.CategoriesStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {

			mymessage = config.CategoriesStore[id].Desc

			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(config.CategoriesStore[id].Id + " " + config.CategoriesStore[id].Desc)
		},
	)
	config.FyneDaysList = List
	List.OnSelected = func(id widget.ListItemID) {
		config.SelectedCategory = id

		Details.SetText(config.CategoriesStore[id].Id)

		edrow.SetText(strconv.Itoa(config.CategoriesStore[id].Row))
		edrow.Disable()

		edid.SetText(config.CategoriesStore[id].Id)

		eddesc.SetText(config.CategoriesStore[id].Desc)

		deletebutton := widget.NewButtonWithIcon("Delete Inventory Category", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.CategoriesDelete(myrow)
			config.CategoriesGet()
		})
		savebutton := widget.NewButtonWithIcon("Save Inventory Category", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)

			config.CategoriesUpdate(myrow, edid.Text, eddesc.Text)
			config.CategoriesGet()

		})
		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			griddesc,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Category")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Inventory Category", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			griddesc,

			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Category")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)dlg.Show()
	})
	topbox := container.NewBorder(addbutton, nil, nil, nil)

	bottombox := container.NewBorder(
		nil,
		Errors,
		nil,
		nil,
		nil,
	)

	return container.NewBorder(
		topbox,
		bottombox,
		nil,
		nil,
		List,
	)

}
