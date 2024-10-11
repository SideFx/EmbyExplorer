// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// UI MainWindow, using Unison library (c) Richard A. Wilkes
// https://github.com/richardwilkes/unison
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/settings"
	"github.com/richardwilkes/unison"
)

const (
	wndMinWidth  float32 = 768
	wndMinHeight float32 = 480
)

var mainWindow *unison.Window

func NewMainWindow() error {
	var err error
	mainWindow, err = unison.NewWindow("")
	if err != nil {
		return err
	}
	mainWindow.SetTitle(assets.AppName + " " + assets.AppVersion)
	content := mainWindow.Content()
	content.SetBorder(unison.NewEmptyBorder(unison.NewUniformInsets(5)))
	content.SetLayout(&unison.FlexLayout{
		Columns:  1,
		HSpacing: 1,
		VSpacing: 5,
	})
	content.AddChild(createToolbarPanel())
	content.AddChild(createTablePanel())
	installDefaultMenus(mainWindow)
	installCallbacks()
	_ = LoadPreferences()
	prefs := settings.GetPreferences()
	rect := prefs.WindowRect
	if rect.Width < wndMinWidth {
		rect.Width = wndMinWidth
	}
	if rect.Height < wndMinHeight {
		rect.Height = wndMinHeight
	}
	dispRect := unison.PrimaryDisplay().Usable
	if rect.X == 0 || rect.X > dispRect.Width-rect.Width {
		if dispRect.Width > rect.Width {
			rect.X = (dispRect.Width - rect.Width) / 2
		}
	}
	if rect.Y == 0 || rect.Y > dispRect.Height-rect.Height {
		if dispRect.Height > rect.Height {
			rect.Y = (dispRect.Height - rect.Height) / 2
		}
	}
	mainWindow.SetFrameRect(rect)
	v := settings.Valid()
	if v {
		api.InitApiPreferences(prefs.EmbySecure, prefs.EmbyServer, prefs.EmbyPort, prefs.EmbyUser, string(prefs.EmbyPassword))
	}
	setFunctions(true, v, false, false, false)
	prepareTitleIcon()
	if len(titleIcons) > 0 {
		mainWindow.SetTitleIcons(titleIcons)
	}
	mainWindow.ToFront()
	return nil
}

func installDefaultMenus(wnd *unison.Window) {
	unison.DefaultMenuFactory().BarForWindow(wnd, func(m unison.Menu) {
		unison.InsertStdMenus(m, AboutDialog, PreferencesDialogFromMenu, nil)
	})
}

func installCallbacks() {
	viewsPopupMenu.SelectionChangedCallback = func(popup *unison.PopupMenu[string]) {
		switchView()
	}
	mainWindow.MinMaxContentSizeCallback = func() (minSize, maxSize unison.Size) {
		return windowMinMaxResizeCallback()
	}
	mainWindow.WillCloseCallback = func() {
		mainWindowWillClose()
	}
}

func windowMinMaxResizeCallback() (minSize, maxSize unison.Size) {
	var _min, _max unison.Size
	_min = unison.NewSize(wndMinWidth, wndMinHeight)
	disp := unison.PrimaryDisplay()
	_max.Width = disp.Usable.Width
	_max.Height = disp.Usable.Height
	return _min, _max
}

func mainWindowWillClose() {
	rect := mainWindow.FrameRect()
	prefs := settings.GetPreferences()
	prefs.WindowRect = rect
	settings.SetPreferences(prefs)
	_ = SavePreferences()
}

func AllowQuitCallback() bool {
	mainWindow.AttemptClose()
	return true
}
