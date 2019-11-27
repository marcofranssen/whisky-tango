package installer

// Installer installs given applications
type Installer interface {
	Install(apps []string) error
}
