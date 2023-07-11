package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-authenticate-host-with-netrc/netrcutil"
)

// ConfigsModel ...
type ConfigsModel struct {
	Host     string
	Username string
	Password string
}

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

	configs.print()
	fmt.Println()

	if err := configs.validate(); err != nil {
		failf("Issue with input: %s", err)
	}

	netRC := netrcutil.New()

	log.Infof("Other configs:")
	log.Printf("- OutputPath: %s", netRC.OutputPth)

	fmt.Println()

	log.Infof("Adding host config...")

	if err := netRC.CreateOrUpdateFile(netrcutil.NetRCItemModel{Machine: configs.Host, Login: configs.Username, Password: configs.Password}); err != nil {
		failf("Failed to update .netrc file: %s", err)
	}

	log.Donef("Success")
}
