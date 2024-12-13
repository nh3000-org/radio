package panes

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
)

//var Song []byte
//var Intro []byte
//var Outro []byte

func InventoryScreen(win fyne.Window) fyne.CanvasObject {

	//config.HoursGet() moved to logon

	Details := widget.NewLabel("")

	larow := widget.NewLabel("Row: ")
	edrow := widget.NewEntry()
	edrow.SetPlaceHolder("Automatically Assigned")
	edrow.Disable()
	gridrow := container.New(layout.NewGridLayoutWithRows(2), larow, edrow)

	lacategory := widget.NewLabel("Category: ")
	edcategory := widget.NewSelect(config.CategoriesToArray(), func(string) {})
	gridcategory := container.New(layout.NewGridLayoutWithRows(2), lacategory, edcategory)

	laartist := widget.NewLabel("Artist: ")
	edartist := widget.NewEntry()
	edartist.SetPlaceHolder("Enter Artist Name")
	gridartist := container.New(layout.NewGridLayoutWithRows(2), laartist, edartist)

	lasong := widget.NewLabel("Song: ")
	edsong := widget.NewEntry()
	edsong.SetPlaceHolder("Enter Song Name")
	gridsong := container.New(layout.NewGridLayoutWithRows(2), lasong, edsong)

	laalbum := widget.NewLabel("Album: ")
	edalbum := widget.NewEntry()
	edalbum.SetPlaceHolder("Enter Album Name")
	gridalbum := container.New(layout.NewGridLayoutWithRows(2), laalbum, edalbum)

	lalength := widget.NewLabel("Length: ")
	edlength := widget.NewEntry()
	edlength.SetPlaceHolder("Enter Song Length")
	edlength.SetText("0")
	edlength.Disable()
	gridlength := container.New(layout.NewGridLayoutWithRows(2), lalength, edlength)

	laorder := widget.NewLabel("Play Order: ")
	edorder := widget.NewEntry()
	edorder.SetText("000000")
	edorder.Disable()
	gridorder := container.New(layout.NewGridLayoutWithRows(2), laorder, edorder)

	laexpires := widget.NewLabel("Expires On: ")
	edexpires := widget.NewEntry()
	edexpires.SetText("9999-01-01 00:00:00")
	gridexpires := container.New(layout.NewGridLayoutWithRows(2), laexpires, edexpires)

	lalastplayed := widget.NewLabel("Last Played: ")
	edlastplayed := widget.NewEntry()
	edlastplayed.SetText("2000-01-01 00:00:00")
	gridlastplayed := container.New(layout.NewGridLayoutWithRows(2), lalastplayed, edlastplayed)

	ladateadded := widget.NewLabel("Date Added: ")
	eddateadded := widget.NewEntry()
	eddateadded.Disable()
	da := time.Now()
	eddateadded.SetText(da.String())
	gridedateadded := container.New(layout.NewGridLayoutWithRows(2), ladateadded, eddateadded)

	laspinstoday := widget.NewLabel("Spins Today: ")
	edspinstoday := widget.NewEntry()
	edspinstoday.Disable()
	edspinstoday.SetText("0")
	gridspinstoday := container.New(layout.NewGridLayoutWithRows(2), laspinstoday, edspinstoday)

	laspinsweek := widget.NewLabel("Spins Today: ")
	edspinsweek := widget.NewEntry()
	edspinsweek.Disable()
	edspinsweek.SetText("0")
	gridspinsweek := container.New(layout.NewGridLayoutWithRows(2), laspinsweek, edspinsweek)

	laspinstotal := widget.NewLabel("Spins Today: ")
	edspinstotal := widget.NewEntry()
	edspinstotal.Disable()
	edspinstotal.SetText("0")
	gridspinstotal := container.New(layout.NewGridLayoutWithRows(2), laspinstotal, edspinstotal)

	openSong := widget.NewButtonWithIcon("Load Song", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			var song = reader
			log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text, songbytes)

		}, win)

		fd.Show()
	})
	openSongIntro := widget.NewButtonWithIcon("Load Song Intro", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			var song = reader
			log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text+"INTRO", songbytes)

		}, win)

		fd.Show()
	})
	openSongOutro := widget.NewButtonWithIcon("Load Song Outro", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			var song = reader
			log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text+"OUTRO", songbytes)

		}, win)

		fd.Show()
	})
	gridfile := container.New(layout.NewGridLayoutWithColumns(3), openSong, openSongIntro, openSongOutro)
	var timelayout = "2000-01-01 00:00:00"
	saveaddbutton := widget.NewButtonWithIcon("Add Inventory Item", theme.ContentCopyIcon(), func() {
		var length, _ = strconv.Atoi(edlength.Text)
		var expires, _ = time.Parse(edexpires.Text, timelayout)
		var lastplayed, _ = time.Parse(edlastplayed.Text, timelayout)
		var dateadded, _ = time.Parse(eddateadded.Text, timelayout)
		var today, _ = strconv.Atoi(edspinstoday.Text)
		var week, _ = strconv.Atoi(edspinsweek.Text)
		var total, _ = strconv.Atoi(edspinstotal.Text)
		rowreturned := config.InventoryAdd(edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, expires, lastplayed, dateadded, today, week, total)
		row := strconv.Itoa(rowreturned)
		edrow.SetText(row)
		//config.InventoryGet()

		//Song = nil
		//Intro = nil
		//Outro = nil
		// copy file into upload 3 posible
		// publish to nats

	})
	List := widget.NewList(
		func() int {
			return len(config.HoursStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {

			mymessage = config.HoursStore[id].Desc

			item.(*fyne.Container).Objects[0].(*widget.Label).SetText(config.InventoryStore[id].Artist + " " + config.InventoryStore[id].Album + " " + config.InventoryStore[id].Song)
		},
	)
	config.FyneInventoryList = List
	List.OnSelected = func(id widget.ListItemID) {
		config.SelectedInventory = id

		Details.SetText(config.InventoryStore[id].Artist + "-" + config.InventoryStore[id].Album + "-" + config.InventoryStore[id].Song)

		edrow.SetText(strconv.Itoa(config.InventoryStore[id].Row))
		edrow.Disable()

		deletebutton := widget.NewButtonWithIcon("Delete Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.InventoryDelete(myrow)
			//config.InventoryGet()
			// publish to nats
		})
		savebutton := widget.NewButtonWithIcon("Save Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			var length, _ = strconv.Atoi(edlength.Text)
			var expires, _ = time.Parse(edexpires.Text, timelayout)
			var lastplayed, _ = time.Parse(edlastplayed.Text, timelayout)
			var dateadded, _ = time.Parse(eddateadded.Text, timelayout)
			var today, _ = strconv.Atoi(edspinstoday.Text)
			var week, _ = strconv.Atoi(edspinsweek.Text)
			var total, _ = strconv.Atoi(edspinstotal.Text)

			config.InventoryUpdate(myrow, edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, expires, lastplayed, dateadded, today, week, total)
			//config.InventoryGet()

		})
		databox := container.NewVBox(

			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridlength,
			gridorder,
			gridexpires,
			gridlastplayed,
			gridedateadded,
			gridspinstoday,
			gridspinsweek,
			gridspinstotal,
			savebutton,
			deletebutton,
		)
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))
		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Items")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
		List.Unselect(id)
	}
	addbutton := widget.NewButtonWithIcon("Add Inventory Item", theme.ContentCopyIcon(), func() {
		databox := container.NewVBox(

			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridlength,
			gridorder,
			gridexpires,
			gridlastplayed,
			gridedateadded,
			gridspinstoday,
			gridspinsweek,
			gridspinstotal,
			saveaddbutton,
		)
		openSong.Disable()
		openSongIntro.Disable()
		openSongOutro.Disable()
		DetailsVW := container.NewScroll(databox)
		DetailsVW.SetMinSize(fyne.NewSize(640, 480))

		dlg := fyne.CurrentApp().NewWindow("Manage Inventory Item")

		dlg.SetContent(container.NewBorder(DetailsVW, nil, nil, nil, nil))
		dlg.Show()
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
