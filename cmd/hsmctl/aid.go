package hsmctl

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/anguyen-ledger/hsmctl/internal/aid"
    "path/filepath"
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
    Short:  "generate attest_xx files from yaml",
    Run: func(cmd *cobra.Command, args []string) {
	hash,_ := cmd.Flags().GetString("hash")
	path,_ := cmd.Flags().GetString("path")
	path = filepath.Join(path,"attest_%d")
        aid.Generate(arr,hash,path)
    },
    }
)

func init() {
    rootCmd.AddCommand(aidCmd)
    aidCmd.AddCommand(aid_generateCmd)
    aid_generateCmd.PersistentFlags().StringArrayVarP(&arr, "file", "f", []string{}, "path to yaml file(s) containing CID and AID info")
    aid_generateCmd.PersistentFlags().StringP("hash", "", "0300564c54004f5247%08X", "hexa code used to generate attest files")
    aid_generateCmd.PersistentFlags().StringP("path", "p", "/srv/BlueHSMServer/data/params/global", "path to store attest files, it must exists")
    aid_generateCmd.MarkPersistentFlagRequired("file")
}
