package pm

// Installer installs given applications
type Installer interface {
	Install(apps []string) error
}

// AppLister lists installed applications
type AppLister interface {
	List() ([]string, error)
}
