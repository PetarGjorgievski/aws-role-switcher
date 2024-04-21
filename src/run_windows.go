// +build windows

package main

import (
    "bytes"
    "os/exec"
)

func runAwsRunasAndWait(profile string) (string, error) {
    cmd := exec.Command("aws-runas", "-e", profile)
    var out bytes.Buffer
    cmd.Stdout = &out
    if err := cmd.Start(); err != nil {
        return "", err
    }

    if err := cmd.Wait(); err != nil {
        return "", err
    }

    return out.String(), nil
}
