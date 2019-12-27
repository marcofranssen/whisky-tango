// +build windows

package pm

import (
	"github.com/marcofranssen/whisky-tango/lib/pm/plugins/chocolatey"
)

// NewInstaller creates a new instance of a installer
func NewInstaller() Installer {
	return &chocolatey.Chocolatey{}
}

// NewAppLister creates a new instance of a application lister
func NewAppLister() AppLister {
	return &chocolatey.Chocolatey{}
}
