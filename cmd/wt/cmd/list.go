package cmd

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/marcofranssen/whisky-tango/lib/pm"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "list your installed apps",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Printf("%s\n\nInstalled apps:\n", rootCmd.Short)
			i := pm.NewAppLister()
			apps, err := i.List()
			if err != nil {
				cmd.PrintErrln(err)
				return
			}
			cmd.Print(strings.Join(apps, " "))
			cmd.Println()
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)
}
