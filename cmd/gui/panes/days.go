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

func DaysScreen(win fyne.Window) fyne.CanvasObject {
	//header := widget.NewLabel("Manage Days of Week")
	config.GetDays()
	//ac2 := widget.NewAccordionItem("NEW1", widget.NewLabel("NEW1"))
	//widget.NewAccordionItem("NEW2", widget.NewLabel("NEW2"))

	/* 	for _, d := range config.DaysStore {
			//labrow := widget.NewLabel(strconv.Itoa(d.Row))
			//dowrow := widget.NewLabel(strconv.Itoa(d.Dow))
			//dayrow := widget.NewLabel(d.Day)
			//descrow := widget.NewLabel(d.Desc)

			type DaysStruct struct {
		Row int     // rowid
		Day  string // message id
		Desc string // alias
		Dow  int    // hostname
	}

	var DaysStore = make(map[int]DaysStruct)

		} */
	Details := widget.NewLabel("")
	//var DetailsBorder = container.NewBorder(Details, nil, nil, nil, nil)

	List := widget.NewList(
		func() int {
			return len(config.NatsMessages)
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
		/* 			Row int     // rowid
		Day  string // message id
		Desc string // alias
		Dow  int    // hostname */
		larow := widget.NewLabel("Row: ")
		edrow := widget.NewEntry()
		edrow.SetPlaceHolder("Automatically Assigned")
		edrow.SetText(strconv.Itoa(config.DaysStore[id].Row))
		gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)

		laday := widget.NewLabel("Day: ")
		edday := widget.NewRadioGroup([]string{"MON", "TUE", "WED", "THU", "FRI", "SAT", "SUN"}, func(string) {})
		edday.Horizontal = true
		edday.SetSelected(config.DaysStore[id].Day)

		gridday := container.New(layout.NewGridLayoutWithRows(2), laday, edday)

		ladesc := widget.NewLabel("Description: ")
		eddesc := widget.NewEntry()

		eddesc.SetText(config.DaysStore[id].Day)
		griddesc := container.New(layout.NewGridLayoutWithRows(2), ladesc, eddesc)

		ladow := widget.NewLabel("Day of Week: ")
		eddow := widget.NewRadioGroup([]string{"1", "2", "3", "4", "5", "6", "7"}, func(string) {})
		eddow.Horizontal = true
		eddow.SetSelected(strconv.Itoa(config.DaysStore[id].Dow))
		griddow := container.New(layout.NewGridLayoutWithRows(2), ladow, eddow)

		deletebutton := widget.NewButtonWithIcon("Delete Day of Week", theme.ContentCopyIcon(), func() {

		})
		savebutton := widget.NewButtonWithIcon("save Day of Week", theme.ContentCopyIcon(), func() {

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
