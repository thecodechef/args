package cmd

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.com/yosssi/gcss"
)

var (
	rootCmd = &cobra.Command{
		Use:   "args",
		Short: "the arguments command",
		Long:  `the arguments command`,
		Run: func() {
			cssPath, err := gcss.CompileFile("./test.gcss")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.ServeFile(w, r, cssPath)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func init() {
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
}
