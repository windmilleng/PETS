package pets

import (
	"bytes"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/windmilleng/pets/internal/mill"
)

var DryRunCmd = &cobra.Command{
	Use: "dry-run",
	Run: func(cmd *cobra.Command, args []string) {
		file := mill.GetFilePath()

		err := mill.ExecFile(file, os.Stdout)
		if err != nil {
			fmt.Println(err)
		}
	},
}
