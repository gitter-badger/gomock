package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hlts2/gomock/pkg/gomock"
	"gopkg.in/yaml.v2"

	cli "github.com/spf13/cobra"
)

var runCmd = &cli.Command{
	Use:   "run",
	Short: "Start API mock server",
	Run: func(cmd *cli.Command, args []string) {
		if err := run(cmd, args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	},
}

var file string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&file, "set", "s", "config.yml", "set config file")
}

func run(cmd *cli.Command, args []string) error {
	d, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var conf gomock.Config

	err = yaml.Unmarshal(d, &conf)
	if err != nil {
		return err
	}

	return gomock.RunServer(conf)
}
