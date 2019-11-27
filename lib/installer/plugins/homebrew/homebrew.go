// +build darwin

package homebrew

import "fmt"

import "strings"

// Homebrew installer to install apps via Homebrew
type Homebrew struct{}

// Install installs the given applications
func (i *Homebrew) Install(apps []string) error {
	fmt.Printf("  brew install %s\n", strings.Join(apps, " "))
	return nil
}
