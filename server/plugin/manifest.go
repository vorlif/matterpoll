// This file is automatically generated. Do not modify it manually.

package plugin

import (
	"strings"

	"github.com/mattermost/mattermost-server/v5/model"
)

var manifest *model.Manifest

const manifestStr = `
{
  "id": "com.github.matterpoll.matterpoll",
  "name": "Matterpoll",
  "description": "Create polls and surveys directly within Mattermost.",
  "homepage_url": "https://github.com/matterpoll/matterpoll",
  "support_url": "https://github.com/matterpoll/matterpoll/issues",
  "release_notes_url": "https://github.com/matterpoll/matterpoll/releases/tag/v1.4.0",
  "icon_path": "assets/logo_dark.svg",
  "version": "1.4.0",
  "min_server_version": "5.20.0",
  "server": {
    "executables": {
      "linux-amd64": "server/dist/plugin-linux-amd64",
      "darwin-amd64": "server/dist/plugin-darwin-amd64",
      "windows-amd64": "server/dist/plugin-windows-amd64.exe"
    },
    "executable": ""
  },
  "webapp": {
    "bundle_path": "webapp/dist/main.js"
  },
  "settings_schema": {
    "header": "",
    "footer": "* To report an issue, make a suggestion, or submit a contribution, [check the repository](https://github.com/matterpoll/matterpoll).",
    "settings": [
      {
        "key": "Trigger",
        "display_name": "Trigger Word:",
        "type": "text",
        "help_text": "Trigger Word must be unique, cannot begin with a slash, and cannot contain any spaces.",
        "placeholder": "",
        "default": "poll"
      },
      {
        "key": "ExperimentalUI",
        "display_name": "Experimental UI:",
        "type": "bool",
        "help_text": "When true, Matterpoll will render poll posts with a rich UI. The rich UI is not available on the mobile app.",
        "placeholder": "",
        "default": false
      }
    ]
  }
}
`

func init() {
	manifest = model.ManifestFromJson(strings.NewReader(manifestStr))
}
