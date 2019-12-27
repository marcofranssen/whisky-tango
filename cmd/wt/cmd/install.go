package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/marcofranssen/whisky-tango/lib/pm"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "installs your apps",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Printf("%s\n\nInstall apps:\n", rootCmd.Short)
			apps := viper.GetStringSlice("apps")
			cmd.Print(strings.Join(apps, " "))
			cmd.Println()

			i := pm.NewInstaller()
			err := i.Install(apps)
			if err != nil {
				cmd.PrintErrln(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(installCmd)
}
