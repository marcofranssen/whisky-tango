// +build darwin

package pm

import (
	"github.com/marcofranssen/whisky-tango/lib/pm/plugins/homebrew"
)

// NewInstaller creates a new instance of a installer
func NewInstaller() Installer {
	return &homebrew.Homebrew{}
}

// NewAppLister creates a new instance of a application lister
func NewAppLister() AppLister {
	return &homebrew.Homebrew{}
}
