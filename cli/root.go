
package cli

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
)

var dataDir = "data"

var rootCli = &cobra.Command{
    Use:   "toychain",
    Short: "Toy blockchain CLI",
}

func Execute() {
    if err := rootCli.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

