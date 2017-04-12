package main

import (
	"errors"
	"os"
	"strings"
	"time"

	"fmt"

	"github.com/bitrise-io/depman/pathutil"
	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-authenticate-host-with-netrc/netrcutil"
)

// ConfigsModel ...
type ConfigsModel struct {
	Host     string
	Username string
	Password string
}

//Authenticate host with netrc
func createConfigsModelFromEnvs() ConfigsModel {
	return ConfigsModel{
		Host:     os.Getenv("host"),
		Username: os.Getenv("username"),
		Password: os.Getenv("password"),
	}
}

func (configs *ConfigsModel) validate() error {
	if configs.Host == "" {
		return errors.New("No Host parameter specified")
	}
	if configs.Username == "" {
		return errors.New("No Username parameter specified")
	}
	if configs.Password == "" {
		return errors.New("No Password parameter specified")
	}
	return nil
}

func secureInput(input string) string {
	return strings.Repeat("*", len(input))
}

func (configs *ConfigsModel) print() {
	// star out sensitive fields
	log.Infof("Configs:")
	log.Printf("- Host: %s", configs.Host)
	log.Printf("- Username: %s", configs.Username)
	log.Printf("- Password: %s", secureInput(configs.Password))
}

func failf(message string, args ...interface{}) {
	log.Errorf(message, args...)
	os.Exit(1)
}

func main() {
	configs := createConfigsModelFromEnvs()

	if err := configs.validate(); err != nil {
		failf("Issue with input: %s", err)
	}

	configs.print()

	fmt.Println()

	netRC := netrcutil.New()

	log.Infof("Other configs:")
	log.Printf("- OutputPath: %s", netRC.OutputPth)

	fmt.Println()

	log.Infof("Adding host config...")
	netRC.AddItemModel(netrcutil.NetRCItemModel{Machine: configs.Host, Login: configs.Username, Password: configs.Password})
	log.Printf("%s added", configs.Host)

	fmt.Println()

	log.Infof("Writing .netrc file...")

	isExists, err := pathutil.IsPathExists(netRC.OutputPth)
	if err != nil {
		failf("Failed to check path (%s), error: %s", netRC.OutputPth, err)
	}

	if !isExists {
		log.Printf("No .netrc file found at (%s), creating new...", netRC.OutputPth)

		if err := netRC.CreateFile(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	} else {
		log.Warnf("File already exists at (%s)", netRC.OutputPth)

		backupPth := fmt.Sprintf("%s%s", strings.Replace(netRC.OutputPth, ".netrc", ".bk.netrc", -1), time.Now().Format("2006_01_02_15_04_05"))

		if originalContent, err := fileutil.ReadBytesFromFile(netRC.OutputPth); err != nil {
			failf("Failed to read file (%s), error: %s", netRC.OutputPth, err)
		} else if err := fileutil.WriteBytesToFile(backupPth, originalContent); err != nil {
			failf("Failed to write file (%s), error: %s", backupPth, err)
		} else {
			log.Printf("Backup created at: %s", backupPth)
		}

		log.Printf("Appending config to the existing .netrc file...")

		if err := netRC.Append(); err != nil {
			failf("Failed to write .netrc file, error: %s", err)
		}
	}
	log.Donef("Success")
}
