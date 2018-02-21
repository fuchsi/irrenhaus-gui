/*
 * irrenhaus-gui, GTK client for irrenhaus.dyndns.dk
 * Copyright (C) 2018  Daniel Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>
 */

package main

import (
	"bufio"
	"fmt"
	"github.com/c2h5oh/datasize"
	"github.com/fuchsi/torrentfile"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"io/ioutil"
	"math"
	"os"
	"path"
	"strings"
	"time"

	. "github.com/fuchsi/irrenhaus-gui/config"
	. "github.com/fuchsi/irrenhaus-gui/maketorrent"

	api "github.com/fuchsi/irrenhaus-api"
	"github.com/fuchsi/irrenhaus-api/Category"
	"github.com/mattn/go-gtk/gtk"
	"github.com/pborman/getopt/v2"
)

// VERSION real value comes from Makefile
var VERSION = ""

// COMMIT real value comes from Makefile
var COMMIT = ""

// BRANCH real value comes from Makefile
var BRANCH = "master"

var configPath = ".irrenhaus-gui/"
var config Configuration
var configFile string

// getopt flags/options
var helpFlag = getopt.BoolLong("help", 'h', "Show this help message and exit")
var versionFlag = getopt.BoolLong("version", 'V', "Print version and quit")
var configOpt = getopt.StringLong("config", 'c', "", "Path to the config directory")

type GUI struct {
	window    *gtk.Window
	statusbar *gtk.Statusbar

	accel *gtk.AccelGroup

	uploadPage     UploadPage
	createMetaPage CreateMetaPage
	settings       SettingsWindow
}

func (g *GUI) Init() {
	g.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	g.window.SetTitle("irrenhaus-gui " + VERSION)
	g.window.Connect("destroy", g.exitApp)
	pb := gdkpixbuf.NewPixbufFromData(logoPng)
	g.window.SetIcon(pb)

	g.accel = gtk.NewAccelGroup()
	g.window.AddAccelGroup(g.accel)

	menubar := g.mkMenuBar()
	vbox := gtk.NewVBox(false, 1)
	vbox.PackStart(menubar, false, false, 0)

	notebook := g.mkNotebook()
	vbox.Add(notebook)

	g.window.Add(vbox)

	g.window.SetSizeRequest(800, 600)

	g.settings = SettingsWindow{}
	g.settings.Init(g)
}

func (g *GUI) mkMenuBar() *gtk.MenuBar {
	menubar := gtk.NewMenuBar()

	fileMenu := gtk.NewMenu()

	settings := gtk.NewMenuItemWithMnemonic("_Settings")
	settings.Connect("activate", g.showSettings)
	fileMenu.Append(settings)

	seperator := gtk.NewSeparatorMenuItem()
	fileMenu.Append(seperator)

	quit := gtk.NewImageMenuItemFromStock(gtk.STOCK_QUIT, g.accel)
	quit.Connect("activate", g.exitApp)
	fileMenu.Append(quit)

	fileMenuItem := gtk.NewMenuItemWithLabel("File")
	fileMenuItem.SetSubmenu(fileMenu)
	menubar.Append(fileMenuItem)

	helpMenu := gtk.NewMenu()

	about := gtk.NewImageMenuItemFromStock(gtk.STOCK_ABOUT, g.accel)
	about.Connect("activate", g.aboutDlg)
	helpMenu.Append(about)

	helpMenuItem := gtk.NewMenuItemWithLabel("Help")
	helpMenuItem.SetSubmenu(helpMenu)
	menubar.Append(helpMenuItem)

	return menubar
}

func (g *GUI) showSettings() {
	g.settings.Show()
}

func (g *GUI) mkNotebook() *gtk.Notebook {
	notebook := gtk.NewNotebook()

	pageUpload := g.mkPageUpload()
	notebook.AppendPage(pageUpload, gtk.NewLabel("Upload Torrent"))

	pageCreateMeta := g.mkPageCreateMeta()
	notebook.AppendPage(pageCreateMeta, gtk.NewLabel("Create Torrent"))

	//pageCreateNfo := gtk.NewFrame("Create NFO")
	//notebook.AppendPage(pageCreateNfo, gtk.NewLabel("Create NFO"))

	return notebook
}

func (g *GUI) mkPageUpload() gtk.IWidget {
	g.uploadPage = UploadPage{gui: g}

	vbox := gtk.NewVBox(false, 5)

	var frame *gtk.Frame
	var hbox *gtk.HBox
	var entry *gtk.Entry
	var filechooser *gtk.FileChooserButton
	var textview *gtk.TextView
	var checkbutton *gtk.CheckButton
	var combobox *gtk.ComboBoxText
	var button *gtk.Button
	var filter *gtk.FileFilter

	hbox = gtk.NewHBox(true, 2)
	// Meta file
	frame = gtk.NewFrame("Meta file:")
	filechooser = gtk.NewFileChooserButton("Meta File", gtk.FILE_CHOOSER_ACTION_OPEN)
	filter = gtk.NewFileFilter()
	filter.AddMimeType("application/x-bittorrent")
	filter.SetName(".torrent files")
	filechooser.SetFilter(filter)
	g.uploadPage.metaFile = filechooser
	g.uploadPage.metaFile.Connect("file-set", g.uploadPage.selectMetafile)
	frame.Add(filechooser)
	hbox.Add(frame)
	// Name
	frame = gtk.NewFrame("Name:")
	entry = gtk.NewEntry()
	g.uploadPage.name = entry
	frame.Add(entry)
	hbox.Add(frame)
	vbox.PackStart(hbox, false, false, 2)

	hbox = gtk.NewHBox(true, 2)
	// NFO
	frame = gtk.NewFrame("NFO file:")
	filechooser = gtk.NewFileChooserButton("NFO File", gtk.FILE_CHOOSER_ACTION_OPEN)
	filter = gtk.NewFileFilter()
	filter.AddPattern("*.nfo")
	filter.SetName(".nfo files")
	filechooser.SetFilter(filter)
	g.uploadPage.nfoFile = filechooser
	frame.Add(filechooser)
	hbox.Add(frame)
	// Category
	frame = gtk.NewFrame("Category:")
	combobox = gtk.NewComboBoxText()
	for _, name := range Category.GetCategories() {
		combobox.AppendText(name)
	}
	g.uploadPage.category = combobox
	frame.Add(combobox)
	hbox.Add(frame)
	vbox.PackStart(hbox, false, false, 2)

	// Images
	frame = gtk.NewFrame("Images:")
	hbox = gtk.NewHBox(true, 2)
	filechooser = gtk.NewFileChooserButton("Image 1", gtk.FILE_CHOOSER_ACTION_OPEN)
	filter = gtk.NewFileFilter()
	filter.AddMimeType("image/jpeg")
	filter.AddMimeType("image/png")
	filter.AddMimeType("image/gif")
	filter.SetName("Image files")
	filechooser.SetFilter(filter)
	g.uploadPage.imageFile1 = filechooser
	hbox.Add(filechooser)
	filechooser = gtk.NewFileChooserButton("Image 2", gtk.FILE_CHOOSER_ACTION_OPEN)
	filechooser.SetFilter(filter)
	g.uploadPage.imageFile2 = filechooser
	hbox.Add(filechooser)
	frame.Add(hbox)
	vbox.PackStart(frame, false, false, 2)

	// Description
	frame = gtk.NewFrame("Description:")
	vbox2 := gtk.NewVBox(false, 2)
	textview = gtk.NewTextView()
	textview.SetEditable(true)
	textview.SetCursorVisible(true)
	textview.SetSizeRequest(280, 280)
	g.uploadPage.description = textview
	vbox2.PackStart(textview, true, true, 2)
	checkbutton = gtk.NewCheckButtonWithLabel("Use NFO as description")
	g.uploadPage.useNfo = checkbutton
	g.uploadPage.useNfo.Connect("toggled", g.uploadPage.nfoSwitch)
	vbox2.PackEnd(checkbutton, false, false, 2)
	frame.Add(vbox2)
	vbox.PackStart(frame, false, false, 2)

	// Buttons
	hbox = gtk.NewHBox(false, 0)
	button = gtk.NewButtonWithLabel("Upload")
	button.Clicked(g.uploadPage.upload)
	hbox.Add(button)
	button = gtk.NewButtonWithLabel("Clear")
	button.Clicked(g.uploadPage.clear)
	hbox.Add(button)
	vbox.PackEnd(hbox, false, false, 2)

	return vbox
}

func (g *GUI) mkPageCreateMeta() gtk.IWidget {
	g.createMetaPage = CreateMetaPage{gui: g}

	vbox := gtk.NewVBox(false, 5)

	var frame *gtk.Frame
	var hbox *gtk.HBox
	var entry *gtk.Entry
	var filechooser *gtk.FileChooserButton
	var combobox *gtk.ComboBoxText
	var button *gtk.Button
	var textview *gtk.TextView
	var checkbutton *gtk.CheckButton
	var progressbar *gtk.ProgressBar

	hbox = gtk.NewHBox(true, 2)
	// Source
	frame = gtk.NewFrame("Source:")
	filechooser = gtk.NewFileChooserButton("Source", gtk.FILE_CHOOSER_ACTION_OPEN|gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER)
	g.createMetaPage.source = filechooser
	g.createMetaPage.source.Connect("file-set", g.createMetaPage.selectSource)
	frame.Add(filechooser)
	hbox.Add(frame)
	// Name
	frame = gtk.NewFrame("Name:")
	entry = gtk.NewEntry()
	g.createMetaPage.name = entry
	frame.Add(entry)
	hbox.Add(frame)
	vbox.PackStart(hbox, false, false, 2)

	hbox = gtk.NewHBox(true, 2)
	// Comment
	frame = gtk.NewFrame("Comment:")
	entry = gtk.NewEntry()
	g.createMetaPage.comment = entry
	frame.Add(entry)
	hbox.Add(frame)
	// Piece length
	frame = gtk.NewFrame("Piece Length:")
	combobox = gtk.NewComboBoxText()
	combobox.AppendText("256 KB")
	combobox.AppendText("512 KB")
	combobox.AppendText("1 MB")
	combobox.AppendText("2 MB")
	combobox.AppendText("4 MB")
	combobox.AppendText("8 MB")
	g.createMetaPage.pieceLength = combobox
	frame.Add(combobox)
	hbox.Add(frame)
	vbox.PackStart(hbox, false, false, 2)

	// Announce URLs
	frame = gtk.NewFrame("Announce URLs:")
	vbox2 := gtk.NewVBox(false, 2)
	textview = gtk.NewTextView()
	textview.SetEditable(true)
	textview.SetCursorVisible(true)
	textview.SetSizeRequest(280, 250)
	g.createMetaPage.announce = textview
	vbox2.PackStart(textview, true, true, 2)
	checkbutton = gtk.NewCheckButtonWithLabel("Private Torrent")
	g.createMetaPage.private = checkbutton
	vbox2.PackEnd(checkbutton, false, false, 2)
	frame.Add(vbox2)
	vbox.PackStart(frame, false, false, 2)

	// Progress
	frame = gtk.NewFrame("Progress:")
	progressbar = gtk.NewProgressBar()
	g.createMetaPage.progress = progressbar
	progressbar.SetText("idle")
	frame.Add(progressbar)
	vbox.PackStart(frame, false, false, 2)

	// Buttons
	hbox = gtk.NewHBox(false, 0)
	button = gtk.NewButtonWithLabel("Create")
	button.Clicked(g.createMetaPage.create)
	hbox.Add(button)
	button = gtk.NewButtonWithLabel("Clear")
	button.Clicked(g.createMetaPage.clear)
	hbox.Add(button)
	vbox.PackEnd(hbox, false, false, 2)

	return vbox
}

func (g *GUI) aboutDlg() {
	dlg := gtk.NewAboutDialog()
	dlg.SetProgramName("irrenhaus-gui")
	dlg.SetComments("A GTK Client for the irrenhaus tracker")
	dlg.SetVersion(VERSION)
	dlg.SetCopyright("© 2018 by Daniel Müller")
	dlg.SetAuthors([]string{"Daniel Müller <perlfuchsi@gmail.com>"})
	dlg.SetLicense(`This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
`)
	dlg.Run()
	dlg.Destroy()
}

func (g *GUI) betaWarning() {
	dlg := gtk.NewMessageDialog(g.window, gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, "This software is currently in beta.\nAlthough everything seems to work, you should not expect it to work at all.")
	dlg.Run()
	dlg.Destroy()
}

func (g *GUI) Show() {
	g.window.SetPosition(gtk.WIN_POS_CENTER)
	g.window.ShowAll()
	g.betaWarning()

	if err := getConnection().Login(); err != nil {
		g.error("Login failed: " + err.Error())
	}
}

func (g *GUI) exitApp() {
	gtk.MainQuit()
}

func (g *GUI) info(message string) {
	dlg := gtk.NewMessageDialog(g.window, gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_INFO, gtk.BUTTONS_OK, message)
	dlg.Run()
	dlg.Destroy()
}

func (g *GUI) error(message string) {
	dlg := gtk.NewMessageDialog(g.window, gtk.DIALOG_MODAL|gtk.DIALOG_DESTROY_WITH_PARENT, gtk.MESSAGE_ERROR, gtk.BUTTONS_OK, message)
	dlg.Run()
	dlg.Destroy()
}

type UploadPage struct {
	gui *GUI

	metaFile    *gtk.FileChooserButton
	nfoFile     *gtk.FileChooserButton
	imageFile1  *gtk.FileChooserButton
	imageFile2  *gtk.FileChooserButton
	name        *gtk.Entry
	category    *gtk.ComboBoxText
	description *gtk.TextView
	useNfo      *gtk.CheckButton
}

func (u *UploadPage) nfoSwitch() {
	if u.useNfo.GetActive() {
		u.description.SetEditable(false)
	} else {
		u.description.SetEditable(true)
	}
}

func (u *UploadPage) selectMetafile() {
	filename := u.metaFile.GetFilename()
	filename = path.Base(filename)
	filename = strings.Replace(filename, path.Ext(filename), "", 1)
	u.name.SetText(filename)
}

func (u *UploadPage) clear() {
	u.metaFile.UnselectAll()
	u.nfoFile.UnselectAll()
	u.imageFile1.UnselectAll()
	u.imageFile2.UnselectAll()
	u.name.SetText("")
	u.category.SetActive(0)
	u.description.GetBuffer().SetText("")
	u.useNfo.SetActive(false)
}

func (u *UploadPage) upload() {
	var description string

	if u.useNfo.GetActive() {
		file, err := ioutil.ReadFile(u.nfoFile.GetFilename())
		if err != nil {
			u.gui.error("Could not open NFO file: " + err.Error())
			return
		}
		description = string(file)
	} else {
		var start, end gtk.TextIter
		buffer := u.description.GetBuffer()
		buffer.GetStartIter(&start)
		buffer.GetEndIter(&end)
		description = buffer.GetText(&start, &end, true)
	}

	category, err := Category.ToInt(u.category.GetActiveText())
	if err != nil {
		u.gui.error("Invalid category: " + err.Error())
		return
	}

	id, err := upload(u.metaFile.GetFilename(), u.nfoFile.GetFilename(), u.imageFile1.GetFilename(), u.imageFile2.GetFilename(), u.name.GetText(), description, category)

	if err != nil {
		u.gui.error("Upload failed: " + err.Error())
		return
	}

	u.gui.info(fmt.Sprintf("Upload successful: %s/details.php?id=%d\n", config.URL, id))
	u.clear()
}

type CreateMetaPage struct {
	gui *GUI

	source      *gtk.FileChooserButton
	name        *gtk.Entry
	comment     *gtk.Entry
	pieceLength *gtk.ComboBoxText
	announce    *gtk.TextView
	private     *gtk.CheckButton
	progress    *gtk.ProgressBar
}

func (c *CreateMetaPage) selectSource() {
	filename := c.source.GetFilename()
	filename = path.Base(filename)
	c.name.SetText(filename)
}

func (c *CreateMetaPage) create() {
	tf := torrentfile.TorrentFile{}

	var start, end gtk.TextIter
	var announceList []string
	buffer := c.announce.GetBuffer()
	buffer.GetStartIter(&start)
	buffer.GetEndIter(&end)
	announceStr := buffer.GetText(&start, &end, true)
	announceList = strings.Split(announceStr, "\n")

	if len(announceList) == 0 {
		c.gui.error("Empty announce list")
		return
	}

	tf.AnnounceUrl = announceList[0]
	if len(announceList) > 1 {
		tf.AnnounceList = announceList
	}

	tf.Name = c.name.GetText()
	tf.Comment = c.comment.GetText()
	tf.Private = c.private.GetActive()
	tf.PieceLength = uint64(math.Pow(2, float64(c.pieceLength.GetActive()+18)))
	tf.CreatedBy = "irrenhaus-gui " + VERSION
	tf.CreationDate = time.Now()
	tf.Encoding = "UTF-8"

	finfo, err := os.Stat(c.source.GetFilename())
	if err != nil {
		c.gui.error("Could not open source: " + err.Error())
		return
	}

	cProgress := make(chan CreateProgress)
	go func() {
		var progress CreateProgress
		var totalSize datasize.ByteSize
		var hashedSize datasize.ByteSize

		for {
			progress = <-cProgress
			fraction := float64(progress.HashedPieces) / float64(progress.NumPieces)
			if progress.HashedPieces == 0 {
				fraction = 0
			}

			totalSize = datasize.ByteSize(progress.NumPieces * tf.PieceLength)
			hashedSize = datasize.ByteSize(progress.HashedPieces * tf.PieceLength)

			gdk.ThreadsEnter()
			c.progress.SetFraction(fraction)
			c.progress.SetText(fmt.Sprintf("%.2f%% %s / %s hashed", (fraction * 100), hashedSize.HumanReadable(), totalSize.HumanReadable()))
			gdk.ThreadsLeave()

			if progress.Finished {
				return
			}
		}
	}()

	go func() {
		if finfo.IsDir() { // Dir mode
			files, pieces := CreateFromDirectory(c.source.GetFilename(), tf.PieceLength, cProgress)
			tf.Files = files
			tf.Pieces = pieces
		} else { // Single file mode
			tf.Files = make([]torrentfile.File, 1)
			file, pieces := CreateFromSingleFile(c.source.GetFilename(), tf.PieceLength, cProgress)
			tf.Files[0] = file
			tf.Pieces = pieces
		}

		gdk.ThreadsEnter()
		filechooserdialog := gtk.NewFileChooserDialog(
			"Save meta file",
			c.gui.window,
			gtk.FILE_CHOOSER_ACTION_SAVE,
			gtk.STOCK_SAVE,
			gtk.RESPONSE_ACCEPT)
		filter := gtk.NewFileFilter()
		filter.AddPattern("*.torrent")
		filter.SetName("Torrent files")
		filechooserdialog.AddFilter(filter)
		filechooserdialog.SetCurrentName(tf.Name + ".torrent")

		filechooserdialog.Response(func() {
			fp, err := os.Create(filechooserdialog.GetFilename())
			if err != nil {
				c.gui.error("Could not save meta file: " + err.Error())
				return
			}
			defer fp.Close()

			writer := bufio.NewWriter(fp)
			writer.Write(tf.Encode())

			filechooserdialog.Destroy()
			c.clear()
		})
		filechooserdialog.Run()
		gdk.ThreadsLeave()
	}()

}

func (c *CreateMetaPage) clear() {
	c.source.UnselectAll()
	c.name.SetText("")
	c.comment.SetText("")
	c.announce.GetBuffer().SetText("")
	c.pieceLength.SetActive(0)
	c.private.SetActive(false)
}

type SettingsWindow struct {
	gui    *GUI
	window *gtk.Window

	username *gtk.Entry
	password *gtk.Entry
	pin      *gtk.Entry
	url      *gtk.Entry
}

func (s *SettingsWindow) Init(gui *GUI) {
	s.gui = gui

	s.window = gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	s.window.SetTitle("Settings")
	s.window.SetModal(true)
	s.window.SetPosition(gtk.WIN_POS_CENTER_ON_PARENT)
	pb := gdkpixbuf.NewPixbufFromData(logoPng)
	s.window.SetIcon(pb)

	vbox := gtk.NewVBox(false, 1)

	s.username = gtk.NewEntry()
	s.password = gtk.NewEntry()
	s.password.SetInvisibleChar('*')
	s.password.SetVisibility(false)
	s.pin = gtk.NewEntry()
	s.url = gtk.NewEntry()

	var frame *gtk.Frame
	var hbox *gtk.HBox
	var button *gtk.Button

	frame = gtk.NewFrame("Username:")
	frame.Add(s.username)
	vbox.PackStart(frame, false, false, 2)

	frame = gtk.NewFrame("Password:")
	frame.Add(s.password)
	vbox.PackStart(frame, false, false, 2)

	frame = gtk.NewFrame("Pin:")
	frame.Add(s.pin)
	vbox.PackStart(frame, false, false, 2)

	frame = gtk.NewFrame("URL:")
	frame.Add(s.url)
	vbox.PackStart(frame, false, false, 2)

	hbox = gtk.NewHBox(false, 2)
	hbox.Add(gtk.NewLabel(""))

	button = gtk.NewButtonFromStock(gtk.STOCK_OK)
	button.Clicked(s.save)
	hbox.PackEnd(button, false, false, 2)

	button = gtk.NewButtonFromStock(gtk.STOCK_CANCEL)
	button.Clicked(func() {
		s.window.Destroy()
	})
	hbox.PackEnd(button, false, false, 2)

	vbox.PackEnd(hbox, false, false, 0)

	s.window.Add(vbox)

	s.window.SetSizeRequest(400, 300)
}

func (s *SettingsWindow) save() {
	config.Username = s.username.GetText()
	config.Password = s.password.GetText()
	config.Pin = s.pin.GetText()
	config.URL = s.url.GetText()

	DumpConfig(config, configFile)

	s.window.Destroy()
}

func (s *SettingsWindow) Show() {
	s.username.SetText(config.Username)
	s.password.SetText(config.Password)
	s.pin.SetText(config.Pin)
	s.url.SetText(config.URL)

	s.window.ShowAll()
}

func main() {
	getopt.SetParameters("command args")

	// Parse the program arguments
	getopt.Parse()

	if *versionFlag {
		fmt.Printf("irrenhaus-gui %s (%s@%s)\n", VERSION, COMMIT, BRANCH)
		return
	}
	if *helpFlag {
		getopt.Usage()
		return
	}

	if *configOpt == "" {
		configPath = os.Getenv("HOME") + configPath
		if _, err := os.Stat(configPath); err != nil {
			os.Mkdir(configPath, 0772)
		}
	} else {
		configPath = *configOpt
	}

	configFile = configPath + "config.json"

	var err error
	config, err = LoadConfig(configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config file: %s\n", err.Error())
	}

	newConnection()

	// Main window
	glib.ThreadInit(nil)
	gdk.ThreadsInit()
	gdk.ThreadsEnter()
	gtk.Init(&os.Args)

	gui := GUI{}
	gui.Init()

	gui.Show()

	gtk.Main()
	gdk.ThreadsLeave()

	DumpCookies(configPath, getConnection().GetCookies())
}

var connection api.Connection

func getConnection() *api.Connection {
	return &connection
}

func newConnection() *api.Connection {
	connection = api.NewConnection(config.URL, config.Username, config.Password, config.Pin)
	connection.SetUserAgent("irrenhaus-gui " + VERSION)
	cookies, err := LoadCookies(configPath)
	if err == nil {
		connection.SetCookies(cookies)
	}

	return &connection
}

func upload(meta string, nfo string, image1 string, image2 string, name string, description string, category int) (int64, error) {
	metard, err := os.Open(meta)
	if err != nil {
		return 0, err
	}
	defer metard.Close()

	nford, err := os.Open(nfo)
	if err != nil {
		return 0, err
	}
	defer nford.Close()

	imagerd, err := os.Open(image1)
	if err != nil {
		return 0, err
	}
	defer imagerd.Close()

	t, err := api.NewUpload(getConnection(), metard, nford, imagerd, name, category, description)
	if err != nil {
		return 0, err
	}
	if image2 != "" {
		image2rd, err := os.Open(image2)
		if err != nil {
			return 0, err
		}
		defer image2rd.Close()
		t.Image2 = image2rd
	}

	if err := t.Upload(); err != nil {
		return 0, err
	}

	return t.Id, nil
}
