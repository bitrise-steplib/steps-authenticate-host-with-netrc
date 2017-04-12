package netrcutil

import (
	"testing"

	"os"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
	"github.com/stretchr/testify/require"
)

const testCreateFileContent = `machine testhost.com
login testusername
password testpassword`

const testAppendFileContent = `machine testhost.com
login testusername
password testpassword

machine testhost2.com
login testusername2
password testpassword2`

func TestCreateFile(t *testing.T) {
	netRC := New()

	tmpDir, err := pathutil.NormalizedOSTempDirPath("__netrc_test__")
	require.NoError(t, err)

	netRC.OutputPth = tmpDir
	require.NoError(t, os.RemoveAll(netRC.OutputPth))

	writtenContent := ""
	t.Log("Test CreateFile")
	{
		netRC.AddItemModel(NetRCItemModel{Machine: "testhost.com", Login: "testusername", Password: "testpassword"})

		isExists, err := pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, false, isExists)

		err = netRC.CreateFile()
		require.NoError(t, err)

		isExists, err = pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		writtenContent, err = fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testCreateFileContent, writtenContent)
	}

	t.Log("Test Append")
	{
		netRC.ItemModels = []NetRCItemModel{NetRCItemModel{Machine: "testhost2.com", Login: "testusername2", Password: "testpassword2"}}
		err = netRC.Append()
		require.NoError(t, err)

		writtenContent, err = fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testAppendFileContent, writtenContent)
	}
}
