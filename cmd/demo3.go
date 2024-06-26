/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/leopku/kodtest/pkg/store"

	"github.com/go-kod/kod"
	"github.com/spf13/cobra"
)

type demo3Impl struct {
	kod.Implements[Demo3]
	kod.LazyInit

	// how to declare? How to specify one of implements of some interface?
	// memstore kod.Ref[store.IStore]
	// diskstore kod.Ref[store.IStore]
	store store.IStore
}

func (impl *demo3Impl) Init(context.Context) error {
	impl.store = &store.MemoryStore{}
	return nil
}

// demo3Cmd represents the demo3 command
var demo3Cmd = &cobra.Command{
	Use:   "demo3",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("demo3 called")
	},
}

func init() {
	rootCmd.AddCommand(demo3Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demo3Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demo3Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
