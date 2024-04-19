package main
import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"github.com/atotto/clipboard"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program> <aws-profile>")
		os.Exit(1)
	}
	awsProfile := os.Args[1] 

	switch runtime.GOOS {
	case "linux", "darwin": // darwin is macOS
		checkAndInstall("aws-runas", []string{"which", "aws-runas"})

		if runtime.GOOS == "linux" {
			checkAndInstall("xclip", []string{"sudo", "apt-get", "install", "-y", "xclip"})
		} 

	case "windows":
		checkAndInstall("aws-runas.exe", []string{"where", "aws-runas.exe"})
	}
	exportCommands := runAwsRunasAndWait(awsProfile)
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

func checkAndInstall(programName string, installCmd []string) {
	if _, err := exec.LookPath(programName); err != nil {
		fmt.Printf("%s is not installed. Please install %s and retry.\n", programName, programName)
		if len(installCmd) > 1 {
			fmt.Printf("Attempting to install %s...\n", programName)
			runCommand(exec.Command(installCmd[0], installCmd[1:]...))
		}
		os.Exit(1)
	}
}

func runAwsRunasAndWait(profile string) string {
	cmd := exec.Command("aws-runas", "-e", profile)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} 
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting aws-runas:", err)
		return ""
	}

	err := cmd.Wait()
	if err != nil {
		fmt.Println("aws-runas completed with error:", err)
		return ""
	}

	return strings.TrimSpace(out.String())
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

func runCommand(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to execute command: %s\n", err)
		os.Exit(1)
	}
}
