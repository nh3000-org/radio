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

var AC = widget.NewAccordion()

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
	edhour := widget.NewSelect([]string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}, func(string) {})

	lapos := widget.NewLabel("Position on Schedule: ")
	edpos := widget.NewSelect([]string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21"}, func(string) {})

	lacategory := widget.NewLabel("Category to Pick From: ")
	edcategory := widget.NewSelect(config.CategoriesToArray(), func(string) {})

	laspins := widget.NewLabel("Spins to Play From Category: ")
	edspins := widget.NewSelect([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20"}, func(string) {})

	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)
	gridday := container.New(layout.NewGridLayoutWithRows(2), laday, edday)
	gridhour := container.New(layout.NewGridLayoutWithRows(2), lahour, edhour)
	gridpos := container.New(layout.NewGridLayoutWithRows(2), lapos, edpos)
	gridcat := container.New(layout.NewGridLayoutWithRows(2), lacategory, edcategory)
	gridspins := container.New(layout.NewGridLayoutWithRows(2), laspins, edspins)
	saveaddbutton := widget.NewButtonWithIcon("Add Schedule Item", theme.ContentCopyIcon(), func() {
		myspins, _ := strconv.Atoi(edspins.Selected)

		config.ScheduleAdd(edday.Selected, edhour.Selected, edpos.Selected, edcategory.Selected, myspins)
		config.ScheduleGet()
	})

	//List := layout.NewVBoxLayout()

	List := widget.NewList(
		func() int {
			return len(config.ScheduleStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {

			myspins := strconv.Itoa(config.ScheduleStore[id].Spinstoplay)
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText("Day: " + config.ScheduleStore[id].Days + " Hour: " + config.ScheduleStore[id].Hours + " Position: " + config.ScheduleStore[id].Position + " Category: " + config.ScheduleStore[id].Category + " Spins: " + myspins)
			//gridcols := container.New(layout.NewGridLayoutWithColumns(5), widget.NewLabel(config.ScheduleStore[id].Days) , widget.NewLabel(config.ScheduleStore[id].Hours) , widget.NewLabel(config.ScheduleStore[id].Position) , widget.NewLabel(config.ScheduleStore[id].Category) , widget.NewLabel(myspins))
			//item.(*fyne.Container).Objects[0].(*widget.GridWrap(layout.NewGridLayoutWithColumns(5), widget.NewLabel(config.ScheduleStore[id].Days) , widget.NewLabel(config.ScheduleStore[id].Hours) , widget.NewLabel(config.ScheduleStore[id].Position) , widget.NewLabel(config.ScheduleStore[id].Category) , widget.NewLabel(myspins)))
			//item.(*fyne.Container).Objects[0].((*widget.GridWrap).(layout.NewGridLayoutWithColumns(5), widget.NewLabel(config.ScheduleStore[id].Days) , widget.NewLabel(config.ScheduleStore[id].Hours) , widget.NewLabel(config.ScheduleStore[id].Position) , widget.NewLabel(config.ScheduleStore[id].Category) , widget.NewLabel(myspins)))
			//item.(*fyne.Container).Objects[0].(*widget.GridWrap).layout.NewGridLayoutWithColumns((5), widget.NewLabel(config.ScheduleStore[id].Days) , widget.NewLabel(config.ScheduleStore[id].Hours) , widget.NewLabel(config.ScheduleStore[id].Position) , widget.NewLabel(config.ScheduleStore[id].Category) , widget.NewLabel(myspins))
		},
	)

	//config.FyneScheduleList = List

	List.OnSelected = func(id widget.ListItemID) {
		config.SelectedDay = id
		myspins := strconv.Itoa(config.ScheduleStore[id].Spinstoplay)
		Details.SetText(config.ScheduleStore[id].Days + " " + config.ScheduleStore[id].Hours + " " + config.ScheduleStore[id].Position + " " + config.ScheduleStore[id].Category + " " + myspins)
		edrow.SetText(strconv.Itoa(config.ScheduleStore[id].Row))
		edrow.Disable()
		edday.SetSelected(config.ScheduleStore[id].Days)
		edhour.SetSelected(config.ScheduleStore[id].Hours)
		edpos.SetSelected(config.ScheduleStore[id].Position)
		edcategory.SetSelected(config.ScheduleStore[id].Category)
		edspins.SetSelected(strconv.Itoa(config.ScheduleStore[id].Spinstoplay))
		deletebutton := widget.NewButtonWithIcon("Delete Schedule Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.ScheduleDelete(myrow)
			config.ScheduleGet()
		})
		savebutton := widget.NewButtonWithIcon("Save Schedule", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			myspins, _ := strconv.Atoi(edspins.Selected)
			config.ScheduleUpdate(myrow, edday.Selected, edhour.Selected, edpos.Selected, edcategory.Selected, myspins)
			config.ScheduleGet()

		})
		gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)

		databox := container.NewVBox(
			deletebutton,
			gridrow,
			gridday,
			gridhour,
			gridpos,
			gridcat,
			gridspins,
			savebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Schedule")

		//DetailsBottom := container.NewBorder(databox, nil, nil, nil, nil)
		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add New Schedule Item", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridday,
			gridhour,
			gridpos,
			gridcat,
			gridspins,
			saveaddbutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Schedule Item")

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
