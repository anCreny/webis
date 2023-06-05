/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"awesomeProject/Printers"
	"github.com/spf13/cobra"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"os/exec"
	"runtime"
)

var operationSystem string
var path string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "webis",
	Short: "CLI for deployment of MVC imitator",
}

var newCmd = &cobra.Command{
	Use:   "new [project name]",
	Short: "A command for creating new web application.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if path == "" {
			installPath, err := os.Getwd()
			if err != nil {
				Printers.ShowError("Caused some error:", err)
			}

			path = installPath

		}

		var projectName = args[0]
		if _, err := os.Stat(path); os.IsNotExist(err) {
			Printers.ShowError("Flagged path doesn't exist in your system")
			return
		}

		var fullProjectPath string

		switch operationSystem {
		case "windows":
			fullProjectPath = path + "\\" + projectName
		default:
			fullProjectPath = path + "/" + projectName
		}

		if _, err := os.Stat(fullProjectPath); !os.IsNotExist(err) {
			Printers.ShowError("Folder with the same name already exists in the directory " + path)
			return
		}

		os.Mkdir(fullProjectPath, os.ModePerm)
		Printers.ShowLog(projectName + " folder made")

		os.Chdir(fullProjectPath)

		Printers.ShowLog("Starting project files downloading")
		_, err := git.PlainClone(fullProjectPath, false, &git.CloneOptions{
			URL:      "https://github.com/anCreny/ASP.NET_Imitator.git",
			Progress: os.Stdout,
		})

		if err != nil {
			Printers.ShowError(err)
			return
		}
		Printers.ShowLog("All files of the pattern downloaded")

		switch operationSystem {
		case "windows":
			if _, err := exec.Command("cmd", "/C", "go", "mod", "init", projectName).Output(); err != nil {
				Printers.ShowError(err)
				return
			} else {
				Printers.ShowLog("Go module initiated")
			}
			if _, err := exec.Command("cmd", "/c", "go", "get", "github.com/anCreny/WebIsland@latest").Output(); err != nil {
				Printers.ShowError(err)
				return
			} else {
				Printers.ShowLog("Framework downloaded")
			}
			if err := os.RemoveAll(fullProjectPath + "\\.git"); err != nil {
				Printers.ShowError(err)
				return
			}
		default:
			if _, err := exec.Command("go mod init " + projectName).Output(); err != nil {
				Printers.ShowError(err)
				return
			} else {
				Printers.ShowLog("Go module initiated")
			}
			if _, err := exec.Command("go get github.com/anCreny/WebIsland@latest").Output(); err != nil {
				Printers.ShowError(err)
				return
			} else {
				Printers.ShowLog("Framework downloaded")
			}
			if err := os.RemoveAll(fullProjectPath + "/.git"); err != nil {
				Printers.ShowError(err)
				return
			}
		}

		Printers.ShowOk("Everything done successfully. Now you can execute \"go run main.go\" to start your first server!")

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	operationSystem = runtime.GOOS
	newCmd.Flags().StringVarP(&path, "path", "p", "", "Path for deployment")
	rootCmd.AddCommand(newCmd)
}
