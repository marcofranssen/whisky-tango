// +build windows

package pm

import (
    "github.com/marcofranssen/whisky-tango/lib/installer/plugins/chocolatey"
)

// NewInstaller creates a new instance of a installer
func NewInstaller() Installer {
    return &chocolatey.Chocolatey{}
}
