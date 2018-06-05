package cmd

import (
	"mockapi/api"
	"fmt"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:"api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("api")
	},
}

var apiRunCmd = &cobra.Command{
	Use:"run",
	Run: func(cmd *cobra.Command, args []string) {
		api.Run()
	},
}


func init(){
	apiCmd.AddCommand(apiRunCmd)
}