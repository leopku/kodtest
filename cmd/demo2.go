/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/go-kod/kod"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
)

type IMigration interface {
	Migrate(context.Context, int) error
}

type demo2Impl struct {
	kod.Implements[IMigration]
	kod.LazyInit
}

func (ins *demo2Impl) Migrate(context.Context, int) error {
	pp.Println("migration running...")
	return nil
}

// demo2Cmd represents the demo2 command
var demo2Cmd = &cobra.Command{
	Use:   "demo2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo2 called")
		err := kod.Run(context.Background(), func(ctx context.Context, ins *demo2Impl) error {
			pp.Println(ins)
			return nil
		})
		cobra.CheckErr(err)
	},
}

func init() {
	rootCmd.AddCommand(demo2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demo2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demo2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
