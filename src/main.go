package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/atotto/clipboard"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: <program> <aws-profile>")
        os.Exit(1)
    }
    awsProfile := os.Args[1]

    exportCommands, err := runAwsRunasAndWait(awsProfile)
    if err != nil {
        fmt.Println("Error running aws-runas:", err)
        os.Exit(1)
    }

    if exportCommands == "" {
        fmt.Println("No export commands were generated.")
        os.Exit(1)
    }

    exportCommands = filterAndSetExports(exportCommands)
    fmt.Printf("The following export commands are set for the profile %s:\n%s\n", awsProfile, exportCommands)
    
    if err := clipboard.WriteAll(exportCommands); err != nil {
        fmt.Println("Failed to copy to clipboard:", err)
        os.Exit(1)
    }

    fmt.Println("AWS credentials exported and copied to clipboard.")
}

func filterAndSetExports(output string) string {
    var filteredExports []string
    for _, line := range strings.Split(output, "\n") {
        if strings.HasPrefix(line, "export ") {
            filteredExports = append(filteredExports, line)
            parts := strings.SplitN(line, "=", 2)
            if len(parts) == 2 {
                key := strings.TrimPrefix(parts[0], "export ")
                value := strings.Trim(parts[1], "\"")
                os.Setenv(key, value) 
            }
        }
    }
    return strings.Join(filteredExports, "\n")
}
