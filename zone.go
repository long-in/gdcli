package gdcli

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/urfave/cli"
)

type zoneResponseBody struct {
	Name             string `json:"name"`
	CurrentVersionID string `json:"current_version_id"`
	ID               string `json:"id"`
	CurrentVersion   struct {
		LastModifiedAt time.Time `json:"last_modified_at"`
		Name           string    `json:"name"`
		CreatedAt      time.Time `json:"created_at"`
		ID             string    `json:"id"`
		Editable       bool      `json:"editable"`
	} `json:"current_version"`
}

type getZoneRequestParams struct {
	Zone string
}

func ZoneCommand() cli.Command {
	return cli.Command{
		Name:  "zone",
		Usage: "Manage Zones",
		Subcommands: []cli.Command{
			{
				Name:  "ls",
				Usage: "List zones",
				Action: func(c *cli.Context) error {
					var sZ getZoneRequestParams
					sZ.Zone = c.Args().Get(0)
					sZ.Execute()
					return nil
				},
			},
		},
	}
}

func (sZ *getZoneRequestParams) Execute() {
	zones, err := getZones()
	if err != nil {
		fmt.Printf("getZones() %s\n", err)
		os.Exit(2)
	}

	if len(zones) == 0 {
		fmt.Printf("%s\n", "Zone is not registered.")
		os.Exit(0)
	}

	if sZ.Zone != "" {
		for _, z := range zones {
			if z.Name == sZ.Zone {
				z.display()
				os.Exit(0)
			}
		}
	} else {
		for _, z := range zones {
			z.display()
		}
		os.Exit(0)
	}
}

func getZones() (zRB []zoneResponseBody, err error) {
	conf := *getDNSConfig()
	req, err := generateGetZoneRequest(conf.URL)
	if err != nil {
		return nil, err
	}

	generateRequestWithToken(req, conf.User, conf.Password)

	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	if err = decodeRequestBody(res, &zRB); err != nil {
		return zRB, err
	}

	return zRB, nil
}

func generateGetZoneRequest(baseURL string) (*http.Request, error) {
	req, err := http.NewRequest("GET", baseURL, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
