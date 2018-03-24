package main

import (
        "github.com/eaydogan/palto/cli"
)

func main() {
        options := &cli.ScanOption{IP: "", StartPort: 2379, StopPort: 2379}
        cli.Scan(options)
}