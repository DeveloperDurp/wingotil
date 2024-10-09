package ostype

import (
	"errors"
	"golang.org/x/sys/windows/registry"
	"runtime"
)

type WinVer struct {
	CurrentVersion string
	ProductName    string
}
type LinVer struct {
	Distro       string
	LinuxVersion string
}

func GetOsType() (interface{}, error) {

	os := runtime.GOOS
	switch os {
	case "windows":

		winVer, err := GetWinVersion()
		if err != nil {
			return nil, errors.New("Unable to detect Windows Version")
		}

		return winVer, nil

	case "darwin":
		return nil, nil
	case "linux":
		return nil, nil
	default:
		return nil, errors.New("Unable to detect OS")
	}
}

func GetWinVersion() (*WinVer, error) {

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}
	defer k.Close()

	currentVersion, _, err := k.GetStringValue("CurrentVersion")
	if err != nil {
		return nil, err
	}

	productName, _, err := k.GetStringValue("ProductName")
	if err != nil {
		return nil, err
	}

	version := &WinVer{
		CurrentVersion: currentVersion,
		ProductName:    productName,
	}
	return version, nil
}
