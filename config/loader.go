package config

import (
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

//Conf global configure instance
var Conf config

type config struct {
	Process process
	Path    path
	SFTP    sftp
}

type process struct {
	Backup bool
}

type path struct {
	LogPath      string
	ResultReport string
	CsvResources string
	SFTPDest     string
}

type sftp struct {
	Host string
	Port string
	User string
	Pass string
}

/*LoadConfig Loads the configurations parameters stored on
* the configuration file (config.toml)
 */
func LoadConfig(configFile string) error {

	if _, err := toml.DecodeFile(configFile, &Conf); err != nil {
		log.Fatal("Load config failed", err)
		return err
	}
	return nil
}
