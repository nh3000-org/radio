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

func ScheduleScreen(win fyne.Window) fyne.CanvasObject {

	//config.DaysGet() moved to logon

	Details := widget.NewLabel("")
	//var DetailsBorder = container.NewBorder(Details, nil, nil, nil, nil)
	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")

	laday := widget.NewLabel("Day: ")
	edday := widget.NewRadioGroup([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}, func(string) {})
	edday.Horizontal = true

	lahour := widget.NewLabel("Hour: ")
	edhour := widget.NewSelect([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	lapos := widget.NewLabel("Position on Schedule: ")
	edpos := widget.NewSelect([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}, func(string) {})

	lacategory := widget.NewLabel("Category to Pick From: ")
	edcategory := widget.NewSelect(config.CategoriesToArray(), func(string) {})

	laspins := widget.NewLabel("Spins to Play From Category: ")
	edspins := widget.NewSelect([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}, func(string) {})

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laday, edday)
	gridhour := container.New(layout.NewGridLayoutWithRows(2), lanour, edhour)
	griddow := container.New(layout.NewGridLayoutWithRows(2), ladow, eddow)
	saveaddbutton := widget.NewButtonWithIcon("Add Day of Week", theme.ContentCopyIcon(), func() {
		mydow, _ := strconv.Atoi(eddow.Selected)

		config.DaysAdd(edday.Selected, eddesc.Text, mydow)
		config.DaysGet()
	})
	List := widget.NewList(
		func() int {
			return len(config.DaysStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {

			mymessage = config.DaysStore[id].Desc

			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(config.DaysStore[id].Day + " " + config.DaysStore[id].Desc)
		},
	)
	config.FyneDaysList = List
	List.OnSelected = func(id widget.ListItemID) {
		config.SelectedDay = id

		Details.SetText(config.DaysStore[id].Day)

		edrow.SetText(strconv.Itoa(config.DaysStore[id].Row))
		edrow.Disable()

		edday.SetSelected(config.DaysStore[id].Day)

		eddesc.SetText(config.DaysStore[id].Desc)

		eddow.SetSelected(strconv.Itoa(config.DaysStore[id].Dow))

		deletebutton := widget.NewButtonWithIcon("Delete Day of Week", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.DaysDelete(myrow)
			config.DaysGet()
		})
		savebutton := widget.NewButtonWithIcon("Save Day of Week", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			mydow, _ := strconv.Atoi(eddow.Selected)
			config.DaysUpdate(myrow, edday.Selected, eddesc.Text, mydow)
			config.DaysGet()

		})
		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			griddesc,
			griddow,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Days")

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Day of Week", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			griddesc,
			griddow,
			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Days")

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
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
