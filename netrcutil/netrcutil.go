package netrcutil

import (
	"fmt"
	"path/filepath"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
)

const netrcDefaultFileName = ".netrc"

// NetRCItemModel ...
type NetRCItemModel struct {
	Machine  string
	Login    string
	Password string
}

// NetRCModel ...
type NetRCModel struct {
	OutputPth  string
	ItemModels []NetRCItemModel
}

// New ...
func New() *NetRCModel {
	netRCPth := filepath.Join(pathutil.UserHomeDir(), netrcDefaultFileName)
	return &NetRCModel{OutputPth: netRCPth}
}

// AddItemModel ...
func (netRCModel *NetRCModel) AddItemModel(itemModels ...NetRCItemModel) {
	netRCModel.ItemModels = append(netRCModel.ItemModels, itemModels...)
}

// CreateFile ...
func (netRCModel *NetRCModel) CreateFile() error {
	EOF := func(i, len int) string {
		if i != len-1 {
			return "\n\n"
		}
		return ""
	}

	netRCFileContent := ""
	for i, itemModel := range netRCModel.ItemModels {
		netRCFileContent += fmt.Sprintf("machine %s\nlogin %s\npassword %s%s", itemModel.Machine, itemModel.Login, itemModel.Password, EOF(i, len(netRCModel.ItemModels)))
	}

	return fileutil.WriteStringToFile(netRCModel.OutputPth, netRCFileContent)
}

// Append ...
func (netRCModel *NetRCModel) Append() error {
	EOF := func(i, len int) string {
		if i != len-1 {
			return "\n\n"
		}
		return ""
	}

	netRCFileContent := ""
	for i, itemModel := range netRCModel.ItemModels {
		netRCFileContent += fmt.Sprintf("machine %s\nlogin %s\npassword %s%s", itemModel.Machine, itemModel.Login, itemModel.Password, EOF(i, len(netRCModel.ItemModels)))
	}

	return fileutil.AppendStringToFile(netRCModel.OutputPth, fmt.Sprintf("\n\n%s", netRCFileContent))
}
