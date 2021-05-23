package file

import (
	"errors"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type File struct {
	Path string

	credentials map[string]string
}

func init() {
	authenticators.RegisterAuthenticator("file", New)
}

func New(config map[string]string) (authenticators.Authenticator, error) {
	fileAuth := File{}

	// check if path is given
	if path, ok := config["path"]; !ok {
		return nil, errors.New("FileAuthenticator: no file path given")
	} else {
		fileAuth.Path = path

		if fileBytes, err := ioutil.ReadFile(path); err != nil {
			return nil, err
		} else {
			if err := yaml.Unmarshal(fileBytes, &fileAuth.credentials); err != nil {
				return nil, err
			}
		}
	}

	return &fileAuth, nil
}

func (f *File) Authenticate(username, password string) bool {
	if foundPassword, ok := f.credentials[username]; ok {
		return password == foundPassword
	}
	return false
}

func (f *File) GetTenantID(username string) string {
	return username
}
