package main

import (
    "fmt"
    "os"

    svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
    "github.com/tonynovatsky/metalchain_project/vmetalchaind/app"
    "github.com/tonynovatsky/metalchain_project/vmetalchaind/cmd/vmetalchaind/cmd"
)

func main() {
    rootCmd, _ := cmd.NewRootCmd()

    if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
