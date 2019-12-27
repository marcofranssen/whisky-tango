// +build darwin

package pm

import (
	"github.com/marcofranssen/whisky-tango/lib/pm/plugins/homebrew"
)

// NewInstaller creates a new instance of a installer
func NewInstaller() Installer {
	return &homebrew.Homebrew{}
}
