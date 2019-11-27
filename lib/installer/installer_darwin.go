// +build darwin

package installer

import (
    "github.com/marcofranssen/whisky-tango/lib/installer/plugins/homebrew"
)

// NewInstaller creates a new instance of a installer
func NewInstaller() Installer {
    return &homebrew.Homebrew{}
}
