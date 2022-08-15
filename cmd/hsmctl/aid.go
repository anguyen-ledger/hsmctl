package hsmctl

import (
    "fmt"
    //"os"
    "github.com/spf13/cobra"
    "github.com/anguyen-ledger/hsmctl/internal/aid"
)

var (
    arr []string
    aidCmd = &cobra.Command{
    Use:   "aid",
    //Aliases: []string{"att"},
    Short:  "commands to interact with attest_xx files",
    //Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("Error: aid command cannot be used without an action, please use the helper")
    },
    }
    aid_generateCmd = &cobra.Command{
    Use:   "generate",
    //Aliases: []string{"att"},
    Short:  "generate attest_xx files from yaml",
    //Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        //fmt.Println(cmd.Flags().GetString("term"))
	aid.Generate(arr)
    },
    }
)

func init() {
    rootCmd.AddCommand(aidCmd)
    aidCmd.AddCommand(aid_generateCmd)
    //reverseCmd.PersistentFlags().StringArray("term", []string{""} , "A search term for a dad joke.")
    aid_generateCmd.PersistentFlags().StringArrayVarP(&arr, "file", "f", []string{}, "path to yaml file(s) containing CID and AID info")
    aid_generateCmd.MarkPersistentFlagRequired("file")
}
