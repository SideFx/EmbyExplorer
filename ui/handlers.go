// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// Event handlers
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/export"
	"Emby_Explorer/models"
	"Emby_Explorer/settings"
	"github.com/richardwilkes/unison"
	"os"
	"path"
	"strconv"
	time2 "time"
)

var userViews []api.UserView

func embyAuthenticateUser() {
	userViews = nil
	err := api.AuthenticateUserInt()
	if err != nil {
		DialogToDisplaySystemError(assets.ErrAuthFailed, err)
		return
	} else {
		userViews, err = api.UserGetViewsInt()
		if err != nil {
			DialogToDisplaySystemError(assets.ErrFetchViewsFailed, err)
			return
		}
		viewsPopupMenu.RemoveAllItems()
		for i, v := range userViews {
			viewsPopupMenu.AddItem(v.Name)
			if i == 0 {
				viewsPopupMenu.SelectIndex(i)
				setFunctions(false, false, true, false, false)
			}
		}
	}
}

func embyFetchItemsForUser() {
	detailsBtn.SetEnabled(false)
	exportBtn.SetEnabled(false)
	index := viewsPopupMenu.SelectedIndex()
	view := userViews[index]
	dto, err := api.UserGetItenmsInt(view.Id, view.CollectionType)
	if err != nil {
		DialogToDisplaySystemError(assets.ErrFetchItemsFailed, err)
		return
	}
	mainContent.RemoveAllChildren()
	switch view.CollectionType {
	case api.CollectionMovies:
		models.MovieDataTable = api.GetMovieDisplayData(dto)
		newMovieTable(mainContent, models.MovieDataTable)
		if len(models.MovieDataTable) > 0 {
			models.MovieTable.SelectByIndex(0)
			detailsBtn.SetEnabled(true)
			exportBtn.SetEnabled(true)
		}
		break
	case api.CollectionTVShows:
		models.TVShowDataTable = api.GetTVShowDisplayData(dto)
		newTVShowTable(mainContent, models.TVShowDataTable)
		if len(models.TVShowDataTable) > 0 {
			models.TVShowTable.SelectByIndex(0)
			detailsBtn.SetEnabled(true)
			exportBtn.SetEnabled(true)
		}
		break
	case api.CollectionHomeVideos:
		models.HomeVideoDataTable = api.GetHomeVideoDisplayData(dto)
		newHomeVideoTable(mainContent, models.HomeVideoDataTable)
		if len(models.HomeVideoDataTable) > 0 {
			models.HomeVideoTable.SelectByIndex(0)
			exportBtn.SetEnabled(true)
		}
		break
	default:
	}
}

func embyFetchDetails() {
	canDisplayDetails = true
	detailsWindowDisplay()
}

func embyExport() {
	index := viewsPopupMenu.SelectedIndex()
	view := userViews[index]
	buildAndExport(view.CollectionType)
}

func buildAndExport(collection string) {
	var i, j int
	var sheet string
	var exp = make([]export.Payload, 0)
	var hdr = make([]export.HeaderData, 0)
	var e export.Payload
	var c export.HeaderData
	j = 1 // xlsx start row
	switch collection {
	case api.CollectionMovies:
		for i = 0; i < models.MovieTableDescription.NoOfColumns; i++ {
			c.XLSCell = models.MovieTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
			c.Name = models.MovieTableDescription.Columns[i].Caption
			c.Column = models.MovieTableDescription.Columns[i].XLSColumn
			c.Width = models.MovieTableDescription.Columns[i].XLSColumnWidth
			hdr = append(hdr, c)
		}
		for _, m := range models.MovieDataTable {
			j++
			for i = 0; i < models.MovieTableDescription.NoOfColumns; i++ {
				e.XLSCell = models.MovieTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
				e.Data = models.GetMovieDataField(i, m)
				exp = append(exp, e)
			}
		}
		sheet = assets.CapMovies
		break
	case api.CollectionTVShows:
		for i = 0; i < models.TVShowTableDescription.NoOfColumns; i++ {
			c.XLSCell = models.TVShowTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
			c.Name = models.TVShowTableDescription.Columns[i].Caption
			c.Column = models.TVShowTableDescription.Columns[i].XLSColumn
			c.Width = models.TVShowTableDescription.Columns[i].XLSColumnWidth
			hdr = append(hdr, c)
		}
		for _, t := range models.TVShowDataTable {
			j++
			for i = 0; i < models.TVShowTableDescription.NoOfColumns; i++ {
				e.XLSCell = models.TVShowTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
				e.Data = models.GetTVShowDataField(i, t)
				exp = append(exp, e)
			}
		}
		sheet = assets.CapTVShows
		break
	case api.CollectionHomeVideos:
		for i = 0; i < models.HomeVideoTableDescription.NoOfColumns; i++ {
			c.XLSCell = models.HomeVideoTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
			c.Name = models.HomeVideoTableDescription.Columns[i].Caption
			c.Column = models.HomeVideoTableDescription.Columns[i].XLSColumn
			c.Width = models.HomeVideoTableDescription.Columns[i].XLSColumnWidth
			hdr = append(hdr, c)
		}
		for _, h := range models.HomeVideoDataTable {
			j++
			for i = 0; i < models.HomeVideoTableDescription.NoOfColumns; i++ {
				e.XLSCell = models.HomeVideoTableDescription.Columns[i].XLSColumn + strconv.Itoa(j)
				e.Data = models.GetHomeVideoDataField(i, h)
				exp = append(exp, e)
			}
		}
		sheet = assets.CapHomeVideos
		break
	default:
		return
	}
	date := time2.Now().Format("2006-01-02")
	folder := settings.GetLastExportFolder()
	if folder == "" {
		folder, _ = os.UserHomeDir()
	}
	preferredFileName := assets.CapEmby + " " + sheet + " " + date + "." + assets.FileExtension
	dialog := unison.NewSaveDialog()
	dialog.SetInitialFileName(preferredFileName)
	dialog.SetInitialDirectory(folder)
	dialog.SetAllowedExtensions(assets.FileExtension)
	if dialog.RunModal() == true {
		p := dialog.Path()
		lastFolder, _ := path.Split(p)
		settings.SetLastExportFolder(lastFolder)
		err := export.XlsxExport(exp, hdr, p, sheet)
		if err != nil {
			DialogToDisplaySystemError(assets.CapError, err)
		}
	}
}
