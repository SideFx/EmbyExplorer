// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// Preferences, type and access functions
// ---------------------------------------------------------------------------------------------------------------------

package settings

import "github.com/richardwilkes/unison"

type Settings struct {
	WindowRect       unison.Rect
	EmbySecure       bool
	EmbyServer       string
	EmbyPort         string
	EmbyUser         string
	EmbyPassword     []byte
	LastExportFolder string
}

var settings Settings

func SetPreferencesDetail(rect unison.Rect, secure bool, server string, port string, user string, password string) {
	settings.WindowRect = rect
	settings.EmbySecure = secure
	settings.EmbyServer = server
	settings.EmbyPort = port
	settings.EmbyUser = user
	settings.EmbyPassword = []byte(password)
}

func SetPreferences(s Settings) {
	settings = s
}

func GetPreferences() Settings {
	return settings
}

func SetLastExportFolder(path string) {
	settings.LastExportFolder = path
}

func GetLastExportFolder() string {
	return settings.LastExportFolder
}

func Valid() bool {
	return settings.EmbyServer != "" && settings.EmbyPort != "" && settings.EmbyUser != "" && len(settings.EmbyPassword) > 0 &&
		settings.WindowRect.Width > 0 && settings.WindowRect.Height > 0
}
