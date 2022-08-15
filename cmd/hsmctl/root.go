package hsmctl

import (
 "fmt"
 "os"
 "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "hsmctl",
    Short: "Usefull cli tool to interact with HSM hosts",
    Long: "",
    Run: func(cmd *cobra.Command, args []string) {

    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
        os.Exit(1)
    }
}
