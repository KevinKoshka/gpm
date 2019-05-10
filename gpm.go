package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

type InstallPack struct {
	Github bool
	Name   string
}

//IsGithub Standardizes github packages.
func (pack *InstallPack) IsGithub() {
	if pack.Github == true {
		pack.Name = "github.com/" + pack.Name
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "gpm"
	app.Usage = "Go Package Manager"
	app.Author = "Kevin Coscarelli"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:   "add",
			Usage:  "Adds package to 'gpmConfig.json'",
			Action: addPackage,
		}, {
			Name:   "install",
			Usage:  "Installs the packages listed in 'gpmConfig.json'",
			Action: installPackages,
		}, {
			Name:   "remove",
			Usage:  "Removes package from 'gpmConfig.json'",
			Action: removePackage,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func readJSON() (data []InstallPack) {
	file, _ := ioutil.ReadFile("gpmConfig.json")
	data = []InstallPack{}
	_ = json.Unmarshal(file, &data)
	if len(data) > 0 {
		return data
	} else {
		return nil
	}
}

func writeJSON(data []InstallPack) error {
	fileData, writeJSONerr := json.MarshalIndent(data, "", "    ")
	ioutil.WriteFile("gpmConfig.json", fileData, 0644)
	if writeJSONerr != nil {
		return writeJSONerr
	} else {
		return nil
	}
}

func addPackage(c *cli.Context) error {
	//First, we read and format the config file.
	data := readJSON()
	newPackage := InstallPack{true, c.Args().First()}
	newPackage.IsGithub()

	//Second, we check for duplicated entries.
	for _, entry := range data {
		if entry.Github == newPackage.Github && entry.Name == newPackage.Name {
			return errors.New("\u001b[31m" + newPackage.Name + " is already listed.")
		}
	}

	//Third, we append the new data and write the config file.
	data = append(data, newPackage)
	return writeJSON(data)
}

func installPackages(c *cli.Context) error {
	//First, we read and format the config file.
	var cmd *exec.Cmd
	data := readJSON()
	if data == nil {
		return errors.New("\u001b[31m" + "There are no packages to install.")
	}
	stdOutput, err := cmd.CombinedOutput()
	for _, entry := range data {
		cmd = exec.Command("go", "get", entry.Name)
	}
	fmt.Println(string(stdOutput))
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return err
	} else {
		return nil
	}
}

func removePackage(c *cli.Context) error {
	data := readJSON()
	byePackage := InstallPack{true, c.Args().First()}
	byePackage.IsGithub()
	if data == nil {
		return errors.New("\u001b[31m" + "There are no packages to remove.")
	}
	for i, entry := range data {
		if entry.Github == byePackage.Github && entry.Name == byePackage.Name {
			data = append(data[:i], data[i+1:]...)
		}
		break
	}
	return writeJSON(data)
}
