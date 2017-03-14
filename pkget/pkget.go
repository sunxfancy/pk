package pkget

import (
	"io"
	"net/http"
	"os"

	"github.com/jinzhu/configor"
)

// Pkget : a package manager for linux
type Pkget struct {
	ServerLists     []string
	CachePath       string `default:"/tmp/pkget/"`
	InstallPath     string `default:"/Pkget/"`
	UserInstallPath string `default:"~/Pkget/"`
}

// NewPkget : constructor
func NewPkget() *Pkget {
	var p = new(Pkget)
	return p
}

func (p *Pkget) loadConfig() error {
	return configor.Load(p, "~/.config/pkget.yml", "/etc/pkget.yml")
}

func (p Pkget) getPkg(url string) {

}

func (p Pkget) checkVersion(version string) {

}

func (p Pkget) downloadFile(filepath string, url string) (err error) {
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
