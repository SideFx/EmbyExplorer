// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// UI, table panels & utilities
// Using Unison library (c) Richard A. Wilkes
// https://github.com/richardwilkes/unison
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/models"
	"github.com/richardwilkes/toolbox/tid"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/behavior"
	"github.com/richardwilkes/unison/enums/paintstyle"
)

const (
	toolbuttonWidth          = 20
	toolbuttonHeight         = 20
	toolbarFontSize  float32 = 9
	viewsPopupWidth          = 150
	viewsPopupHeight         = 20
	coverMaxWidth            = "300"
	coverMaxHeight           = "300"
)

var viewsPopupMenu *unison.PopupMenu[string]
var prefsBtn *unison.Button
var authBtn *unison.Button
var fetchBtn *unison.Button
var detailsBtn *unison.Button
var exportBtn *unison.Button

var mainContent *unison.Panel
var logoPanel *unison.Panel
var tableScrollArea *unison.ScrollPanel
var collectionType = ""
var canDisplayDetails = false
var titleIcons []*unison.Image

func newSVGButton(svg *unison.SVG) *unison.Button {
	btn := unison.NewButton()
	btn.HideBase = true
	btn.Drawable = &unison.DrawableSVG{
		SVG:  svg,
		Size: unison.NewSize(toolbuttonWidth, toolbuttonHeight),
	}
	btn.Font = unison.LabelFont.Face().Font(toolbarFontSize)
	return btn
}

func createButton(title string, svgcontent string) (*unison.Button, error) {
	svg, err := unison.NewSVGFromContentString(svgcontent)
	if err != nil {
		return nil, err
	}
	btn := newSVGButton(svg)
	btn.SetTitle(title)
	btn.SetLayoutData(align.Middle)
	return btn, nil
}

func createSpacer(width float32, panel *unison.Panel) {
	spacer := &unison.Panel{}
	spacer.Self = spacer
	spacer.SetSizer(func(_ unison.Size) (minSize, prefSize, maxSize unison.Size) {
		minSize.Width = width
		prefSize.Width = width
		maxSize.Width = width
		return
	})
	panel.AddChild(spacer)
}

func setFunctions(prefs bool, auth bool, fetch bool, details bool, export bool) {
	prefsBtn.SetEnabled(prefs)
	authBtn.SetEnabled(auth)
	fetchBtn.SetEnabled(fetch)
	detailsBtn.SetEnabled(details)
	exportBtn.SetEnabled(export)
}

func createToolbarPanel() *unison.Panel {
	var err error
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: 1,
		VSpacing: unison.StdVSpacing,
	})
	prefsBtn, err = createButton(assets.CapPreferences, assets.IconPreferences)
	if err == nil {
		prefsBtn.SetEnabled(true)
		prefsBtn.SetFocusable(false)
		panel.AddChild(prefsBtn)
		prefsBtn.ClickCallback = func() { PreferencesDialog() }
	}
	authBtn, err = createButton(assets.CapAuthenticate, assets.IconLogin)
	if err == nil {
		authBtn.SetEnabled(true)
		authBtn.SetFocusable(false)
		panel.AddChild(authBtn)
		authBtn.ClickCallback = func() { embyAuthenticateUser() }
	}
	createSpacer(25, panel)
	lblItems := unison.NewLabel()
	lblItems.Font = unison.LabelFont.Face().Font(toolbarFontSize)
	lblItems.SetTitle(assets.CapViews)
	lblItems.SetLayoutData(align.Middle)
	panel.AddChild(lblItems)
	createSpacer(5, panel)
	viewsPopupMenu = unison.NewPopupMenu[string]()
	viewsPopupMenu.SetLayoutData(align.Middle)
	viewsPopupMenu.Font = unison.LabelFont.Face().Font(toolbarFontSize)
	viewsPopupSize := unison.NewSize(viewsPopupWidth, viewsPopupHeight)
	viewsPopupMenu.SetSizer(func(_ unison.Size) (minSize, prefSize, maxSize unison.Size) {
		minSize = viewsPopupSize
		prefSize = viewsPopupSize
		maxSize = viewsPopupSize
		return
	})
	viewsPopupMenu.SetFocusable(false)
	panel.AddChild(viewsPopupMenu)
	createSpacer(5, panel)
	fetchBtn, err = createButton(assets.CapFetch, assets.IconFetch)
	if err == nil {
		fetchBtn.SetEnabled(true)
		fetchBtn.SetFocusable(false)
		panel.AddChild(fetchBtn)
		fetchBtn.ClickCallback = func() { embyFetchItemsForUser() }
	}
	detailsBtn, err = createButton(assets.CapDetails, assets.IconDetails)
	if err == nil {
		detailsBtn.SetEnabled(true)
		detailsBtn.SetFocusable(false)
		panel.AddChild(detailsBtn)
		detailsBtn.ClickCallback = func() { embyFetchDetails() }
	}
	exportBtn, err = createButton(assets.CapExport, assets.IconExport)
	if err == nil {
		exportBtn.SetEnabled(true)
		exportBtn.SetFocusable(false)
		panel.AddChild(exportBtn)
		exportBtn.ClickCallback = func() { embyExport() }
	}
	return panel
}

func createTablePanel() *unison.Panel {
	mainContent = unison.NewPanel()
	mainContent.SetLayout(&unison.FlexLayout{
		Columns:  1,
		HSpacing: 1,
		VSpacing: 1,
	})
	mainContent.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	mainContent.SetBorder(unison.NewDefaultFieldBorder(false))
	logoPanel = createEmbyLogoPanel()
	if logoPanel != nil {
		mainContent.AddChild(logoPanel)
	}
	return mainContent.AsPanel()
}

func createEmbyLogoPanel() *unison.Panel {
	svg, err := unison.NewSVGFromContentString(assets.EmbyLogo)
	if err != nil {
		return nil
	}
	panel := unison.NewPanel()
	panel.SetLayoutData(&unison.FlexLayoutData{
		MinSize: unison.NewSize(50, 50),
		HSpan:   1,
		VSpan:   1,
		HAlign:  align.Fill,
		VAlign:  align.Fill,
		HGrab:   true,
		VGrab:   true,
	})
	panel.DrawCallback = func(gc *unison.Canvas, dirty unison.Rect) {
		gc.DrawRect(dirty, unison.ThemeSurface.Light.Paint(gc, dirty, paintstyle.Fill))
		svg.DrawInRectPreservingAspectRatio(gc, panel.ContentRect(false), nil, nil)
	}
	return panel
}

func setLogoPanel() {
	if logoPanel != nil {
		mainContent.RemoveAllChildren()
		mainContent.AddChild(logoPanel)
	}
}

func switchView() {
	collectionType = api.AllowedCollectionTypes[viewsPopupMenu.SelectedIndex()]
	setFunctions(false, false, true, false, false)
	setLogoPanel()
}

func newMovieTable(content *unison.Panel, movieData []models.MovieData) {
	models.MovieTable = unison.NewTable[*models.MovieRow](&unison.SimpleTableModel[*models.MovieRow]{})
	models.MovieTable.Columns = make([]unison.ColumnInfo, models.MovieTableDescription.NoOfColumns)
	for i := range models.MovieTable.Columns {
		models.MovieTable.Columns[i].ID = i
		models.MovieTable.Columns[i].Minimum = 20
		models.MovieTable.Columns[i].Maximum = 10000
	}
	rows := make([]*models.MovieRow, 0)
	for _, m := range movieData {
		r := models.NewMovieRow(tid.MustNewTID('a'), m)
		rows = append(rows, r)
	}
	models.MovieTable.SetRootRows(rows)
	models.MovieTable.SizeColumnsToFit(true)
	header := unison.NewTableHeader[*models.MovieRow](models.MovieTable,
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[0].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[1].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[2].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[3].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[4].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[5].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[6].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[7].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[8].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[9].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[10].Caption, ""),
		unison.NewTableColumnHeader[*models.MovieRow](models.MovieTableDescription.Columns[11].Caption, ""),
	)
	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
	})
	tableScrollArea = unison.NewScrollPanel()
	tableScrollArea.SetContent(models.MovieTable, behavior.Fill, behavior.Fill)
	tableScrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	tableScrollArea.SetColumnHeader(header)
	models.MovieTable.SelectionChangedCallback = func() {
		if canDisplayDetails {
			detailsWindowDisplay()
		}
	}
	content.AddChild(tableScrollArea)
}

func newTVShowTable(content *unison.Panel, tvshowData []models.TVShowData) {
	models.TVShowTable = unison.NewTable[*models.TVShowRow](&unison.SimpleTableModel[*models.TVShowRow]{})
	models.TVShowTable.Columns = make([]unison.ColumnInfo, models.TVShowTableDescription.NoOfColumns)
	for i := range models.TVShowTable.Columns {
		models.TVShowTable.Columns[i].ID = i
		models.TVShowTable.Columns[i].Minimum = 20
		models.TVShowTable.Columns[i].Maximum = 10000
	}
	rows := make([]*models.TVShowRow, 0)
	for _, m := range tvshowData {
		r := models.NewTVShowRow(tid.MustNewTID('a'), m)
		rows = append(rows, r)
	}
	models.TVShowTable.SetRootRows(rows)
	models.TVShowTable.SizeColumnsToFit(true)
	header := unison.NewTableHeader[*models.TVShowRow](models.TVShowTable,
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[0].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[1].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[2].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[3].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[4].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[5].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[6].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[7].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[8].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[9].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[10].Caption, ""),
		unison.NewTableColumnHeader[*models.TVShowRow](models.TVShowTableDescription.Columns[11].Caption, ""),
	)
	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
	})
	tableScrollArea = unison.NewScrollPanel()
	tableScrollArea.SetContent(models.TVShowTable, behavior.Fill, behavior.Fill)
	tableScrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	tableScrollArea.SetColumnHeader(header)
	models.TVShowTable.SelectionChangedCallback = func() {
		if canDisplayDetails {
			detailsWindowDisplay()
		}
	}
	content.AddChild(tableScrollArea)
}

func newHomeVideoTable(content *unison.Panel, homevideoData []models.HomeVideoData) {
	models.HomeVideoTable = unison.NewTable[*models.HomeVideoRow](&unison.SimpleTableModel[*models.HomeVideoRow]{})
	models.HomeVideoTable.Columns = make([]unison.ColumnInfo, models.HomeVideoTableDescription.NoOfColumns)
	for i := range models.HomeVideoTable.Columns {
		models.HomeVideoTable.Columns[i].ID = i
		models.HomeVideoTable.Columns[i].Minimum = 20
		models.HomeVideoTable.Columns[i].Maximum = 10000
	}
	rows := make([]*models.HomeVideoRow, 0)
	for _, m := range homevideoData {
		r := models.NewHomeVideoRow(tid.MustNewTID('a'), m)
		rows = append(rows, r)
	}
	models.HomeVideoTable.SetRootRows(rows)
	models.HomeVideoTable.SizeColumnsToFit(true)
	header := unison.NewTableHeader[*models.HomeVideoRow](models.HomeVideoTable,
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[0].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[1].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[2].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[3].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[4].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[5].Caption, ""),
		unison.NewTableColumnHeader[*models.HomeVideoRow](models.HomeVideoTableDescription.Columns[6].Caption, ""),
	)
	header.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
	})
	tableScrollArea = unison.NewScrollPanel()
	tableScrollArea.SetContent(models.HomeVideoTable, behavior.Fill, behavior.Fill)
	tableScrollArea.SetLayoutData(&unison.FlexLayoutData{
		HAlign: align.Fill,
		VAlign: align.Fill,
		HGrab:  true,
		VGrab:  true,
	})
	tableScrollArea.SetColumnHeader(header)
	content.AddChild(tableScrollArea)
}

func prepareTitleIcon() {
	newImage, err := unison.NewImageFromBytes(assets.Emby, 1)
	if err == nil {
		titleIcons = append(titleIcons, newImage)
	}
}

func newImageFromBytes(itemid string) (*unison.Image, error) {
	var newImage *unison.Image
	image, err := api.GetPrimaryImageForItemInt(itemid, api.ImageFormatPng, coverMaxWidth, coverMaxHeight)
	if err != nil {
		return nil, err
	}
	newImage, err = unison.NewImageFromBytes(image, 1)
	return newImage, nil
}
