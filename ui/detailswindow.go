// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// UI Details window, using Unison library (c) Richard A. Wilkes
// https://github.com/richardwilkes/unison
// ---------------------------------------------------------------------------------------------------------------------

package ui

import (
	"Emby_Explorer/api"
	"Emby_Explorer/assets"
	"Emby_Explorer/models"
	"github.com/richardwilkes/unison"
	"github.com/richardwilkes/unison/enums/align"
	"github.com/richardwilkes/unison/enums/behavior"
)

const (
	textPanelWidth  float32 = 500
	textPanelHeight float32 = 200
)

var detailsWindow *unison.Window
var lastPosition unison.Point

func detailsWindowDisplay() {
	var frame unison.Rect
	var err error
	if detailsWindow != nil {
		tmp := canDisplayDetails
		detailsWindow.Dispose()
		canDisplayDetails = tmp
	}
	title := assets.AppName + " - " + assets.CapDetails
	detailsWindow, err = unison.NewWindow(title, unison.FloatingWindowOption(), unison.NotResizableWindowOption())
	if err != nil {
		panic(err)
	}
	if len(titleIcons) > 0 {
		detailsWindow.SetTitleIcons(titleIcons)
	}
	if focused := unison.ActiveWindow(); focused != nil {
		frame = focused.FrameRect()
	} else {
		frame = unison.PrimaryDisplay().Usable
	}
	content := detailsWindow.Content()
	content.RemoveAllChildren()
	content.SetLayout(&unison.FlexLayout{
		Columns:  1,
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})
	content.SetBorder(unison.NewEmptyBorder(unison.NewUniformInsets(5)))
	panel, pl, pr := newContentPanel()
	content.AddChild(panel)
	detailsWindow.Pack()
	wndFrame := detailsWindow.FrameRect()
	frame.Y += (frame.Height - wndFrame.Height) / 3
	frame.Height = wndFrame.Height
	frame.X += (frame.Width - wndFrame.Width) / 2
	frame.Width = wndFrame.Width
	if lastPosition.X > 0 {
		frame.X = lastPosition.X
	}
	if lastPosition.Y > 0 {
		frame.Y = lastPosition.Y
	}
	detailsWindow.SetFrameRect(frame.Align())
	// window closed manually by user, need "Details" button press to reactivate
	detailsWindow.WillCloseCallback = func() {
		r := detailsWindow.FrameRect()
		lastPosition = unison.NewPoint(r.X, r.Y)
		canDisplayDetails = false
	}
	if pl || pr {
		detailsWindow.ToFront()
	}
}

func newContentPanel() (*unison.Panel, bool, bool) {
	panel := unison.NewPanel()
	var ovw string
	var img *unison.Image
	var pl, pr = false, false
	if collectionType != "" {
		switch collectionType {
		case api.CollectionMovies:
			movie := models.MovieTable.SelectedRows(true)
			for _, m := range movie {
				img, _ = newImageFromBytes(m.M.MovieId)
				ovw = m.M.Overview
				break
			}
			break
		case api.CollectionTVShows:
			tvshow := models.TVShowTable.SelectedRows(true)
			for _, t := range tvshow {
				img, _ = newImageFromBytes(t.M.SeasonId)
				ovw = t.M.Overview
				break
			}
			break
		default:
		}
		panel.SetLayout(&unison.FlexLayout{
			Columns:  2,
			HSpacing: 10,
			VSpacing: unison.StdVSpacing,
		})
		panel.MarkForRedraw()
		if img != nil {
			pl = true
			imgPanel := unison.NewLabel()
			imgPanel.Drawable = img
			imgPanel.SetBorder(unison.NewEmptyBorder(unison.NewUniformInsets(1)))
			imgPanel.SetLayoutData(&unison.FlexLayout{})
			panel.AddChild(imgPanel)
		}
		if ovw != "" {
			pr = true
			textPanel := unison.NewMultiLineField()
			textPanel.SetWrap(true)
			textPanel.AutoScroll = false
			textPanel.SetFocusable(false)
			_, prefSize, _ := textPanel.Sizes(unison.Size{})
			prefSize.Width = textPanelWidth
			prefSize.Height = textPanelHeight
			textPanel.SetFrameRect(unison.Rect{Size: prefSize})
			textPanel.SetText(ovw)
			scroller := unison.NewScrollPanel()
			scroller.SetContent(textPanel, behavior.Follow, behavior.Fill)
			scroller.SetLayoutData(&unison.FlexLayoutData{
				SizeHint: prefSize,
				HAlign:   align.Fill,
				VAlign:   align.Fill,
				HGrab:    true,
				VGrab:    true,
			})
			unison.InstallDefaultFieldBorder(textPanel, scroller)
			panel.AddChild(scroller.AsPanel())
		}
	}
	return panel, pl, pr
}
