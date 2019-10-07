package gdcli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/urfave/cli"
)

type config struct {
	Services []configDetail `json:"services"`
}

type configDetail struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func ConfigCommand() cli.Command {
	return cli.Command{
		Name:  "config",
		Usage: "Manage config",
		Subcommands: []cli.Command{
			{
				Name:  "init",
				Usage: "Init Config File",
				Action: func(c *cli.Context) error {
					initConfig()
					return nil
				},
			},
		},
	}
}

func initConfig() {
	conf := config{
		Services: []configDetail{
			{
				Name:     "dns",
				URL:      "https://api.gis.gehirn.jp/dns/v1/zones",
				User:     "",
				Password: "",
			},
		},
	}

	data, err := json.Marshal(conf)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	file, err := os.Create(os.Getenv("HOME") + "/.gehirun.json.sample")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}
	defer file.Close()

	file.Write(([]byte)(string(data)))

	fmt.Printf("%s: %s\n", "Create Config File", os.Getenv("HOME")+"/.gehirun.json.sample")
}

func getDNSConfig() *configDetail {
	conf := loadConfig()
	var cD configDetail
	for _, s := range conf.Services {
		if s.Name == "dns" {
			cD = s
		}
	}
	return &cD
}

func loadConfig() (conf config) {
	f, err := ioutil.ReadFile(os.Getenv("HOME") + "/.gehirun.json")
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	if err := json.Unmarshal(f, &conf); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(2)
	}

	return conf
}
