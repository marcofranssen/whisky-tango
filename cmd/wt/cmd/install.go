package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/marcofranssen/whisky-tango/lib/installer"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "installs your apps",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Printf("%s\n\nInstall apps:\n", rootCmd.Short)
			apps := viper.GetStringSlice("apps")
			i := installer.NewInstaller()
			err := i.Install(apps)
			if err != nil {
				cmd.PrintErr(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}
