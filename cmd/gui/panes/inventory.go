package panes

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
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

// var Song []byte
// var Intro []byte
// var Outro []byte
var imcategory string
var imartist string
var imsong string
var imalbum string
var imimportdir string
var sp fyne.URI
var sp1 string
var startpath string
var walkstuberr error
var removepath string

var videofull string
var rmcat string
var songfull string
var songunparsed string
var result []string
var da time.Time
var added = "YYYY-MM-DD 00:00:00"
var m string
var d string
var rowreturned int
var songbytes []byte
var songerr error

func InventoryScreen(win fyne.Window) fyne.CanvasObject {

	//config.HoursGet() moved to logon
	config.FyneInventoryWin = win
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

	lastartson := widget.NewLabel("Starts On: ")
	edstartson := widget.NewEntry()
	edstartson.SetText(time.Now().String())
	gridstartson := container.New(layout.NewGridLayoutWithRows(2), lastartson, edstartson)

	lalastplayed := widget.NewLabel("Last Played: ")
	edlastplayed := widget.NewEntry()
	edlastplayed.SetText("2000-01-01 00:00:00")
	gridlastplayed := container.New(layout.NewGridLayoutWithRows(2), lalastplayed, edlastplayed)

	ladateadded := widget.NewLabel("Date Added: ")
	eddateadded := widget.NewEntry()
	eddateadded.Disable()

	da := time.Now()
	added := "YYYY-MM-DD 00:00:00"
	added = strings.Replace(added, "YYYY", strconv.Itoa(da.Year()), 1)
	m := strconv.Itoa(int(da.Month()))
	if len(m) == 1 {
		m = "0" + m
	}
	added = strings.Replace(added, "MM", m, 1)
	d := strconv.Itoa(int(da.Day()))
	if len(d) == 1 {
		d = "0" + d
	}
	added = strings.Replace(added, "DD", d, 1)
	eddateadded.SetText(added)
	gridedateadded := container.New(layout.NewGridLayoutWithRows(2), ladateadded, eddateadded)

	laspinstoday := widget.NewLabel("Spins Today: ")
	edspinstoday := widget.NewEntry()
	edspinstoday.Disable()
	edspinstoday.SetText("0")
	gridspinstoday := container.New(layout.NewGridLayoutWithRows(2), laspinstoday, edspinstoday)

	laspinsweek := widget.NewLabel("Spins Weekly: ")
	edspinsweek := widget.NewEntry()
	edspinsweek.Disable()
	edspinsweek.SetText("0")
	gridspinsweek := container.New(layout.NewGridLayoutWithRows(2), laspinsweek, edspinsweek)

	lalinks := widget.NewLabel("Source Links: ")
	edlinks := widget.NewEntry()
	edlinks.SetPlaceHolder("Enter Website: ")
	gridlinks := container.New(layout.NewGridLayoutWithRows(2), lalinks, edlinks)

	laspinstotal := widget.NewLabel("Spins Total: ")
	edspinstotal := widget.NewEntry()
	edspinstotal.Disable()
	edspinstotal.SetText("0")

	edsongsz := widget.NewLabel("0")
	edintrosz := widget.NewLabel("0")
	edoutrosz := widget.NewLabel("0")

	gridspinstotal := container.New(layout.NewGridLayoutWithRows(2), laspinstotal, edspinstotal)
	importbutton := widget.NewButtonWithIcon("Import Inventory From Stub", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("openInventory err ", err)
				//dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			log.Println("Import Inventory Walk path", reader.URI())

			sp = reader.URI()

			sp1 = strings.Replace(sp.Path(), "file//", "", 1)
			startpath = strings.Replace(sp1, "/README.txt", "", 1)
			os.Chdir(startpath)
			// get category
			log.Println("Start path", startpath)
			walkstuberr = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
				if walkstuberr != nil {
					config.Send("messages.IMPORT", "Inventory Walk Err FileInfo "+err.Error(), "onair")
					log.Println("Import Inventory Walk Err FileInfo", err)
					return walkstuberr
				}
				log.Println("walk path", path, "size", info.Size())
				// strip out last part of path for category
				removepath = startpath + "/"
				cat := strings.Replace(path, removepath, "", 1)
				imimportdir = startpath + "/" + cat
				log.Println("import directory ", imimportdir)

				log.Println("cat", cat)
				if info.IsDir() {
					imcategory = cat
					//log.Println("category", imcategory)
				}
				if strings.HasSuffix(cat, "mp4") {
					videofull = strings.ReplaceAll(path, imcategory+"/", "")
					log.Println("import base video ", videofull)

				}
				//var artist = "na"
				//var song = "na"
				//var album = "na"
				if strings.HasSuffix(cat, "mp3") {
					rmcat = imcategory + "/"
					songfull = strings.ReplaceAll(path, rmcat, "")
					//log.Println("import base song ", songfull, "path", path, "category", imcategory)
					songunparsed = strings.ReplaceAll(songfull, ".mp3", "")
					result = strings.Split(songunparsed, "-")
					fmt.Println("Result:", result, len(result))
					if len(result) == 0 {
						log.Printf("Song not parsed")
					}
					if len(result) == 3 {
						imartist = result[0]
						imsong = result[1]
						imalbum = result[2]
					}
					if len(result) == 2 {
						imartist = result[0]
						imsong = result[1]
						imalbum = "Digital"
					}
					if len(result) == 1 {
						imartist = result[0]
						imsong = result[0]
						imalbum = "Digital"
					}
					log.Println("Song parsed imartist:", imartist, ":song:", imsong, ":album:", imalbum, ":category:", imcategory, ":")
					win.SetTitle("Importing:" + imartist + " - " + imsong)
					var length, _ = strconv.Atoi("0")
					var today, _ = strconv.Atoi("0")
					var week, _ = strconv.Atoi("0")
					var total, _ = strconv.Atoi("0")

					da = time.Now()
					added = strings.Replace(added, "YYYY", strconv.Itoa(da.Year()), 1)
					m = strconv.Itoa(int(da.Month()))
					if len(m) == 1 {
						m = "0" + m
					}
					added = strings.Replace(added, "MM", m, 1)
					d = strconv.Itoa(int(da.Day()))
					if len(d) == 1 {
						d = "0" + d
					}
					added = strings.Replace(added, "DD", d, 1)
					rowreturned = config.InventoryAdd(imcategory, imartist, imsong, imalbum, length, "000000", "2023-12-31 00:00:00", "9999-12-31 00:00:00", "1999-01-01 00:00:00", added, today, week, total, "Stub")
					row := strconv.Itoa(rowreturned)
					if row != "0" {
						log.Println("put bucket song ", imsong)
						songbytes, songerr = os.ReadFile(imimportdir)
						if songerr != nil {
							log.Println("put bucket song ", songerr)
						}
						if songerr != nil {
							log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))
						}
						config.PutBucket("mp3", row, songbytes)
						log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))

						if strings.HasSuffix(cat, "INTRO.mp3") {
							log.Println("import base song intro ", path)
							log.Println("put bucket song ", imsong)
							songbytes, songerr = os.ReadFile(imimportdir)
							if songerr != nil {
								log.Println("put bucket song ", songerr)
							}
							if songerr != nil {
								log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))
							}
							config.PutBucket("mp3", row, songbytes)
							log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))

						}
						if strings.HasSuffix(cat, "OUTRO.mp3") {
							log.Println("import base song outro ", path)
							log.Println("put bucket song ", imsong)
							songbytes, songerr = os.ReadFile(imimportdir)
							if songerr != nil {
								log.Println("put bucket song ", songerr)
							}
							if songerr != nil {
								log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))
							}
							config.PutBucket("mp3", row, songbytes)
							log.Println("PutBucket song ", "item", row, "song size", strconv.Itoa(len(songbytes)))

						}
					}
				}

				return nil
			})
			if walkstuberr != nil {
				log.Println("Import Inventory Walk Error", walkstuberr)
			}

		}, win)

		fd.Show()
		win.SetTitle("Importing Complete")

	})
	openSong := widget.NewButtonWithIcon("Load Song ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("openSong err ", err)
				//dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}
			var song = reader
			//log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text, songbytes)
			edsongsz.SetText(strconv.Itoa(len(songbytes)))

		}, win)

		fd.Show()

	})

	openSongIntro := widget.NewButtonWithIcon("Load Song Intro ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("openSongIntro err ", err)
				//dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			var song = reader
			//log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text+"INTRO", songbytes)
			edintrosz.SetText(strconv.Itoa(len(songbytes)))

		}, win)

		fd.Show()
	})

	openSongOutro := widget.NewButtonWithIcon("Load Song Outro ", theme.UploadIcon(), func() {
		fd := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil {
				log.Println("openSongOutro err ", err)
				//dialog.ShowError(err, win)
				return
			}
			if reader == nil {
				log.Println("Cancelled")
				return
			}

			var song = reader
			//log.Println(os.Stat(strings.Replace(song.URI().String(), "file://", "", -1)))
			songbytes, songerr := os.ReadFile(strings.Replace(song.URI().String(), "file://", "", -1))
			if songerr != nil {
				log.Println("put bucket song ", songerr)
			}

			//inv := 			//var layout = "1999-01-02 15:16:17"
			//var starts, _ = time.Parse(layout, edstartson.Text)
			//var expires, _ = time.Parse(time.RFC3339, edexpires.Text)
			//var expcvtboolires, _ = time.Parse(layout, edexpires.Text)
			//var lastplayed, _ = time.Parse(layout, edlastplayed.Text)
			//var dateadded, _ = time.Parse(layout, eddateadded.Text)strconv.Itoa(edrow)
			if songerr != nil {
				log.Println("PutBucket song ", "item", edrow.Text, "song size", strconv.Itoa(len(songbytes)))
			}
			config.PutBucket("mp3", edrow.Text+"OUTRO", songbytes)
			edoutrosz.SetText(strconv.Itoa(len(songbytes)))
		}, win)

		fd.Show()
	})
	gridfile := container.New(layout.NewGridLayoutWithColumns(3), openSong, openSongIntro, openSongOutro)
	openSongsz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text)))
	edsongsz.SetText("Song Size:" + openSongsz)
	openSongIntrosz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text+"INTRO")))
	edintrosz.SetText("Intro Size:" + openSongIntrosz)
	openSongOutrosz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text+"OUTRO")))
	edoutrosz.SetText("Intro Size:" + openSongOutrosz)
	gridfilesz := container.New(layout.NewGridLayoutWithColumns(3), edsongsz, edintrosz, edoutrosz)

	saveaddbutton := widget.NewButtonWithIcon("Add Inventory Item", theme.ContentCopyIcon(), func() {
		var length, _ = strconv.Atoi(edlength.Text)
		var today, _ = strconv.Atoi(edspinstoday.Text)
		var week, _ = strconv.Atoi(edspinsweek.Text)
		var total, _ = strconv.Atoi(edspinstotal.Text)
		rowreturned := config.InventoryAdd(edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, edstartson.Text, edexpires.Text, edlastplayed.Text, eddateadded.Text, today, week, total, edlinks.Text)
		row := strconv.Itoa(rowreturned)
		edrow.SetText(row)
		openSong.Enable()
		openSongIntro.Enable()
		openSongOutro.Enable()
		config.InventoryGet()
		config.FyneInventoryList.Refresh()

	})
	List := widget.NewList(
		func() int {
			return len(config.InventoryStore)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(widget.NewLabel("Template Object"))
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*fyne.Container).Objects[0].(*widget.Label).SetText("[" + config.InventoryStore[id].Category + "] " + config.InventoryStore[id].Artist + " - " + config.InventoryStore[id].Song + " (" + strconv.Itoa(config.InventoryStore[id].Row) + ")")
		},
	)
	config.FyneInventoryList = List
	List.OnSelected = func(id widget.ListItemID) {

		config.SelectedInventory = id

		Details.SetText("[" + config.InventoryStore[id].Category + "] " + config.InventoryStore[id].Artist + " - " + config.InventoryStore[id].Song)

		edrow.SetText(strconv.Itoa(config.InventoryStore[id].Row))
		edrow.Disable()
		edcategory.SetSelected(config.InventoryStore[id].Category)
		edartist.SetText(config.InventoryStore[id].Artist)
		edsong.SetText(config.InventoryStore[id].Song)
		edalbum.SetText(config.InventoryStore[id].Album)
		edlength.SetText(strconv.Itoa(config.InventoryStore[id].Songlength))
		edlength.Disable()
		edorder.SetText(config.InventoryStore[id].Rndorder)
		edorder.Disable()

		edstartson.SetText(config.InventoryStore[id].Startson)
		edexpires.SetText(config.InventoryStore[id].Expireson)
		eddateadded.SetText(config.InventoryStore[id].Dateadded)
		edlastplayed.SetText(config.InventoryStore[id].Lastplayed)
		edlinks.SetText(config.InventoryStore[id].Sourcelink)
		openSongsz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text)))
		edsongsz.SetText("Song Size: " + openSongsz)
		openSongIntrosz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text+"INTRO")))
		edintrosz.SetText("Intro Size: " + openSongIntrosz)
		openSongOutrosz := strconv.Itoa(int(config.GetBucketSize("mp3", edrow.Text+"OUTRO")))
		edoutrosz.SetText("Outro Size: " + openSongOutrosz)
		edspinstoday.SetText(strconv.Itoa(config.InventoryStore[id].Spinstoday))
		edspinsweek.SetText(strconv.Itoa(config.InventoryStore[id].Spinsweek))
		edspinstotal.SetText(strconv.Itoa(config.InventoryStore[id].Spinstotal))

		deletebutton := widget.NewButtonWithIcon("Delete Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			config.InventoryDelete(myrow)
			config.InventoryGet()
			// publish to nats
		})
		savebutton := widget.NewButtonWithIcon("Save Inventory Item", theme.ContentCopyIcon(), func() {
			myrow, _ := strconv.Atoi(edrow.Text)
			var length, _ = strconv.Atoi(edlength.Text)
			var today, _ = strconv.Atoi(edspinstoday.Text)
			var week, _ = strconv.Atoi(edspinsweek.Text)
			var total, _ = strconv.Atoi(edspinstotal.Text)

			config.InventoryUpdate(myrow, edcategory.Selected, edartist.Text, edsong.Text, edalbum.Text, length, edorder.Text, edstartson.Text, edexpires.Text, edlastplayed.Text, eddateadded.Text, today, week, total, edlinks.Text)
			config.InventoryGet()

			config.FyneInventoryList.Refresh()

		})
		databox := container.NewVBox(

			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridfilesz,
			gridlength,
			gridorder,
			gridstartson,
			gridexpires,
			gridlastplayed,
			gridedateadded,
			gridlinks,
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
	//	importbutton := widget.NewButtonWithIcon("Import Inventory From Stub ", theme.ContentCopyIcon(), func() {
	//	})

	addbutton := widget.NewButtonWithIcon("Add New Inventory Item ", theme.ContentCopyIcon(), func() {

		edrow.SetText("")
		edartist.SetText("")
		edsong.SetText("")
		edalbum.SetText("")
		edalbum.SetText("")
		edsongsz.SetText("0")
		edintrosz.SetText("0")
		edoutrosz.SetText("0")

		edstartson.SetText("2023-12-31 00:00:00")
		edexpires.SetText("9999-12-31 00:00:00")

		edlastplayed.SetText("1999-01-01 00:00:00")

		var da = time.Now()
		//if daerr != nil {
		//	log.Println("Date  daerr ", daerr)
		//}

		//log.Println("Date Added da ", da)
		added := "YYYY-MM-DD 00:00:00"
		added = strings.Replace(added, "YYYY", strconv.Itoa(da.Year()), 1)
		m := strconv.Itoa(int(da.Month()))
		if len(m) == 1 {
			m = "0" + m
		}
		added = strings.Replace(added, "MM", m, 1)
		d := strconv.Itoa(int(da.Day()))
		if len(d) == 1 {
			d = "0" + d
		}
		added = strings.Replace(added, "DD", d, 1)
		//log.Println("Date Added", added)
		eddateadded.SetText(added)

		edspinstoday.SetText("0")
		edspinsweek.SetText("0")
		edspinstotal.SetText("0")

		databox := container.NewVBox(

			gridrow,
			gridcategory,
			gridartist,
			gridsong,
			gridalbum,
			gridfile,
			gridfilesz,
			gridlength,
			gridorder,
			gridstartson,
			gridexpires,
			gridlastplayed,
			gridedateadded,
			gridlinks,
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
	topbox := container.NewBorder(addbutton, importbutton, nil, nil)

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
