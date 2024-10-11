// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// Preferences dialog, using Unison library (c) Richard A. Wilkes
// https://github.com/richardwilkes/unison
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/settings"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/check"
)

const inpTextSizeMax = 200
const inpTextSizeMin = 40
const obscureRune = 0x2a

var okButton *unison.Button
var inpServer *unison.Field
var inpPort *unison.Field
var inpUser *unison.Field
var inpPassword *unison.Field
var chkSecure *unison.CheckBox

func PreferencesDialogFromMenu(_ unison.MenuItem) {
	PreferencesDialog()
}

func PreferencesDialog() {
	dialog, err := unison.NewDialog(nil, nil, newPreferencesPanel(),
		[]*unison.DialogButtonInfo{unison.NewOKButtonInfo(), unison.NewCancelButtonInfo()},
		unison.NotResizableWindowOption())
	if err == nil {
		wnd := dialog.Window()
		wnd.SetTitle(assets.CapPreferences)
		prepareTitleIcon()
		if len(titleIcons) > 0 {
			wnd.SetTitleIcons(titleIcons)
		}
		okButton = dialog.Button(unison.ModalResponseOK)
		okButton.ClickCallback = func() {
			saveSettings()
			dialog.StopModal(unison.ModalResponseOK)
		}
		_ = dialog.Button(unison.ModalResponseCancel)
		s := settings.GetPreferences()
		chkSecure.State = check.Off
		if s.EmbySecure {
			chkSecure.State = check.On
		}
		inpServer.SetText(s.EmbyServer)
		inpPort.SetText(s.EmbyPort)
		inpUser.SetText(s.EmbyUser)
		inpPassword.SetText(string(s.EmbyPassword))
		okButton.SetEnabled(checkOk())
		dialog.RunModal()
	}
}

func newPreferencesPanel() *unison.Panel {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlexLayout{
		Columns:  2,
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	lblSecure := unison.NewLabel()
	lblSecure.Font = unison.LabelFont
	lblSecure.SetTitle(assets.CapSecure)
	chkSecure = unison.NewCheckBox()
	lblServer := unison.NewLabel()
	lblServer.Font = unison.LabelFont
	lblServer.SetTitle(assets.CapServer)
	inpServer = unison.NewField()
	inpServer.Font = unison.FieldFont
	inpServer.MinimumTextWidth = inpTextSizeMax
	lblPort := unison.NewLabel()
	lblPort.Font = unison.LabelFont
	lblPort.SetTitle(assets.CapPort)
	inpPort = unison.NewField()
	inpPort.Font = unison.FieldFont
	inpPort.MinimumTextWidth = inpTextSizeMin
	lblUser := unison.NewLabel()
	lblUser.Font = unison.LabelFont
	lblUser.SetTitle(assets.CapUser)
	inpUser = unison.NewField()
	inpUser.Font = unison.FieldFont
	inpUser.MinimumTextWidth = inpTextSizeMax
	lblPassword := unison.NewLabel()
	lblPassword.Font = unison.LabelFont
	lblPassword.SetTitle(assets.CapPassword)
	inpPassword = unison.NewField()
	inpPassword.Font = unison.FieldFont
	inpPassword.ObscurementRune = obscureRune
	inpPassword.MinimumTextWidth = inpTextSizeMax
	inpServer.ModifiedCallback = func(before, after *unison.FieldState) {
		inpModifiedCallback(before, after)
	}
	inpPort.ModifiedCallback = func(before, after *unison.FieldState) {
		inpModifiedCallback(before, after)
	}
	inpUser.ModifiedCallback = func(before, after *unison.FieldState) {
		inpModifiedCallback(before, after)
	}
	inpPassword.ModifiedCallback = func(before, after *unison.FieldState) {
		inpModifiedCallback(before, after)
	}
	panel.SetLayoutData(&unison.FlexLayoutData{
		MinSize: unison.Size{Width: 300},
		HSpan:   1,
		VSpan:   5,
		VAlign:  align.Middle,
	})
	panel.AddChild(lblSecure)
	panel.AddChild(chkSecure)
	panel.AddChild(lblServer)
	panel.AddChild(inpServer)
	panel.AddChild(lblPort)
	panel.AddChild(inpPort)
	panel.AddChild(lblUser)
	panel.AddChild(inpUser)
	panel.AddChild(lblPassword)
	panel.AddChild(inpPassword)
	panel.Pack()
	return panel
}

func saveSettings() {
	settings.SetPreferencesDetail(mainWindow.FrameRect(), chkSecure.State == check.On, inpServer.Text(),
		inpPort.Text(), inpUser.Text(), inpPassword.Text())
	_ = SavePreferences()
	if settings.Valid() {
		// must update emby session parameters for REST api
		s := settings.GetPreferences()
		api.InitApiPreferences(s.EmbySecure, s.EmbyServer, s.EmbyPort, s.EmbyUser, string(s.EmbyPassword))
		api.CheckEmby(s.EmbyServer)
		authBtn.SetEnabled(true) // enable button for authorization
	}
}

func inpModifiedCallback(_, _ *unison.FieldState) {
	okButton.SetEnabled(checkOk())
}

// authorization data complete?
func checkOk() bool {
	return inpServer.Text() != "" && inpPort.Text() != "" && inpUser.Text() != "" && inpPassword.Text() != ""
}
