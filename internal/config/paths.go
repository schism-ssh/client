package config

import (
	"os"
	"path/filepath"
)

const DefaultLocalFile = "schism_cli.toml"

/*
 * MustPath returns a valid path to a config file in a specific order
 *  1. The path to the file represented by `localFile` arg
 *  2. /etc/schism/cli.toml
 *  3. $HOME/.config/schism/cli.toml
 *  4. ./`DefaultLocalFile` if nothing else is found
 *
 * If none of these files exist, returns an empty string and an error
 */
func MustPath(localFile string) (confPath string, err error) {
	if len(localFile) > 1 {
		confPath, err = MustLocalPath(localFile)
		if err == nil {
			return
		}
	}
	confPath, err = mustConfFile(SystemPath())
	if err == nil {
		return
	}
	confPath, err = MustUserPath()
	if err == nil {
		return
	}
	return MustLocalPath(DefaultLocalFile)
}

/*
 * MustUserPath returns a path to a config in the users home directory
 *   it returns an error if the file doesn't exist
 *   or the homedir is not configured in your environment
 */
func MustUserPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return mustConfFile(UserPath(homeDir))
}

/*
 * MustLocalPath returns the full path to the given file, or an error if:
 *  1. the file doesn't exist
 *  2. there's an issue getting the |`Abs` path|
 */
func MustLocalPath(localFile string) (string, error) {
	localConf, err := filepath.Abs(localFile)
	if err != nil {
		return "", err
	}
	return mustConfFile(localConf)
}

/*
 * SystemPath just returns the path, no errors here
 */
func SystemPath() string {
	return filepath.Join(
		"/etc",
		"schism",
		"cli.toml",
	)
}

/*
 * UserPath takes a user's homeDir and returns a path to a potential config there.
 */
func UserPath(homeDir string) string {
	return filepath.Join(
		homeDir,
		".config",
		"schism",
		"cli.toml",
	)
}

/* mustConfFile validates the existence of a given `configPath`
 *   If the file doesn't exist, the error from `os.Stat` is returned,
 *   otherwise return the full path to the file
 */
func mustConfFile(configPath string) (string, error) {
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		return "", err
	}
	return configPath, nil
}
