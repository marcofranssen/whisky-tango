package cmd

import (
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configCmd = &cobra.Command{
		Use:   "config",
		Short: "Get current configuration as json output",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Print(sprintConfig())
		},
	}
)

func init() {
	rootCmd.AddCommand(configCmd)
}

func sprintConfig() string {
	sb := new(strings.Builder)
	sb.WriteString("\nconfig:\n")
	writeSettings(sb, viper.AllSettings())
	return sb.String()
}

func writeSettings(w io.Writer, settings map[string]interface{}) {
	tw := tabwriter.NewWriter(w, 0, 0, 2, ' ', 0)
	writeRecurseSetting(tw, "", settings)
	tw.Flush()
}

func writeRecurseSetting(w io.Writer, prefix string, settings map[string]interface{}) {
	// maps can not guarantee same order, below is a trick to always print settings in same order.
	keys := make([]string, 0, len(settings))
	for k := range settings {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		v := settings[k]
		t := reflect.TypeOf(v)
		if t.Kind() == reflect.Map {
			if prefix != "" {
				writeRecurseSetting(w, prefix+"."+k, v.(map[string]interface{}))
			} else {
				writeRecurseSetting(w, k, v.(map[string]interface{}))
			}
		} else {
			if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
				slice := strings.Replace(fmt.Sprintf("%q", v), "\" \"", "\",\"", -1)
				fmt.Fprintf(w, "\t%s:\t%s\n", prefix+"."+k, slice)
			} else {
				fmt.Fprintf(w, "\t%s:\t%v\n", prefix+"."+k, v)
			}
		}
	}
}
