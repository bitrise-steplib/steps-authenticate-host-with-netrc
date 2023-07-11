package netrcutil

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bitrise-io/go-utils/fileutil"
	"github.com/bitrise-io/go-utils/pathutil"
)

const testCreateFileContent = `machine testhost.com
	login testusername
	password testpassword
`

const testAppendFileContent = `machine testhost.com
	login testusername
	password testpassword


machine testhost2.com
	login testusername2
	password testpassword2
`

func TestCreateFile(t *testing.T) {
	tmpDir, err := pathutil.NormalizedOSTempDirPath("__netrc_test__")
	require.NoError(t, err)
	defer func() {
		require.NoError(t, os.RemoveAll(tmpDir))
	}()

	t.Log("Test CreateFile")
	{
		netRC := New()
		netRC.OutputPth = filepath.Join(tmpDir, ".netrc")

		netRC.AddItemModel(NetRCItemModel{Machine: "testhost.com", Login: "testusername", Password: "testpassword"})

		isExists, err := pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, false, isExists)

		err = netRC.CreateFile()
		require.NoError(t, err)

		isExists, err = pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		writtenContent, err := fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testCreateFileContent, writtenContent)
	}

	t.Log("Test Append")
	{
		netRC := New()
		netRC.OutputPth = filepath.Join(tmpDir, ".netrc")

		isExists, err := pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		netRC.AddItemModel(NetRCItemModel{Machine: "testhost2.com", Login: "testusername2", Password: "testpassword2"})
		err = netRC.Append()
		require.NoError(t, err)

		writtenContent, err := fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testAppendFileContent, writtenContent)
	}
}

func TestCreateOrUpdateFile(t *testing.T) {
	t.Run("when no .netrc file is present", func(t *testing.T) {
		netRC, _, cleanup := setup(t)
		defer cleanup()

		isExists, err := pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, false, isExists)

		err = netRC.CreateOrUpdateFile(NetRCItemModel{Machine: "testhost.com", Login: "testusername", Password: "testpassword"})
		require.NoError(t, err)

		isExists, err = pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		writtenContent, err := fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testCreateFileContent, writtenContent)
	})

	t.Run("when a .netrc file is already present", func(t *testing.T) {
		netRC, tmpDir, cleanup := setup(t)
		defer cleanup()

		// set up existing .netrc file
		// a different client is used on the same path
		// to avoid issues with its internal state like existing .netrc entries
		netRCForExistingFile := New()
		netRCForExistingFile.OutputPth = netRC.OutputPth
		err := netRCForExistingFile.CreateOrUpdateFile(NetRCItemModel{Machine: "testhost.com", Login: "testusername", Password: "testpassword"})
		require.NoError(t, err)

		isExists, err := pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		err = netRC.CreateOrUpdateFile(NetRCItemModel{Machine: "testhost2.com", Login: "testusername2", Password: "testpassword2"})
		require.NoError(t, err)

		isExists, err = pathutil.IsPathExists(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, true, isExists)

		writtenContent, err := fileutil.ReadStringFromFile(netRC.OutputPth)
		require.NoError(t, err)
		require.Equal(t, testAppendFileContent, writtenContent)

		// verify backup file was created
		backupFile := ""
		found := 0
		err = filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && path != netRC.OutputPth {
				backupFile = path
				found += 1
			}
			return nil
		})
		require.NoError(t, err)

		require.Equal(t, 1, found)
		backupContent, err := fileutil.ReadStringFromFile(backupFile)
		require.NoError(t, err)
		require.Equal(t, testCreateFileContent, backupContent)
	})
}

func setup(t *testing.T) (*NetRCModel, string, func()) {
	tmpDir, err := pathutil.NormalizedOSTempDirPath("__netrc_test__")
	require.NoError(t, err)

	netRC := New()
	netRC.OutputPth = filepath.Join(tmpDir, ".netrc")

	return netRC, tmpDir, func() {
		require.NoError(t, os.RemoveAll(tmpDir))
	}
}
