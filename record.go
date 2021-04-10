package gdcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

type getRecordRequestParam struct {
	Zone                 string
	Name                 string
	ZoneID               string
	ZoneCurrentVersionID string
}

type addRecordRequestParam struct {
	Zone                 string
	Address              string
	Name                 string
	TTL                  int
	Type                 string
	Alias                bool
	Cname                string
	Data                 string
	Nsdname              string
	Prio                 int
	Exchange             string
	ZoneID               string
	ZoneCurrentVersionID string
	RecordID             string
}

type removeRecordRequestParam struct {
	Zone                 string
	Name                 string
	Type                 string
	ZoneID               string
	ZoneCurrentVersionID string
	RecordID             string
}

type updateRecordRequestParam struct {
	Zone                 string
	Address              string
	Name                 string
	TTL                  int
	Type                 string
	Alias                bool
	Cname                string
	Data                 string
	Nsdname              string
	Prio                 int
	Exchange             string
	ZoneID               string
	ZoneCurrentVersionID string
	RecordID             string
}

type recordResponseBody struct {
	Name    string `json:"name"`
	Records []struct {
		Address  string `json:"address"`
		Nsdname  string `json:"nsdname"`
		Data     string `json:"data"`
		Cname    string `json:"cname"`
		Prio     int    `json:"prio"`
		Exchange string `json:"exchange"`
	} `json:"records"`
	EnableAlias bool   `json:"enable_alias"`
	Type        string `json:"type"`
	ID          string `json:"id"`
	TTL         int    `json:"ttl"`
}

type addRecordRequestBody struct {
	Name        string      `json:"name"`
	TTL         int         `json:"ttl"`
	EnableAlias bool        `json:"enable_alias"`
	Type        string      `json:"type"`
	Records     interface{} `json:"records"`
}

type addParamsecordRequestBody struct {
	Address string `json:"address"`
}

type txtRecordRequestBody struct {
	Data string `json:"data"`
}

type cnameRecordRequestBody struct {
	Cname string `json:"cname"`
}

type mxRecordRequestBody struct {
	Prio     int    `json:"prio"`
	Exchange string `json:"exchange"`
}

func RecordCommand() cli.Command {
	return cli.Command{
		Name:  "record",
		Usage: "Manage records",
		Subcommands: []cli.Command{
			{
				Name:  "ls",
				Usage: "List records",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "name, n",
						Usage: "Specify an domain",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name.", "USAGE: gdcli record ls <ZONE NAME>")
						return nil
					}

					var getParams getRecordRequestParam
					getParams.Zone = c.Args().Get(0)
					getParams.Name = c.String("name")

					getParams.Execute()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a new record",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "ip",
						Usage: "Specify an IP Address",
					},
					cli.StringFlag{
						Name:  "name",
						Usage: "Specify an domain",
					},
					cli.StringFlag{
						Name:  "ttl",
						Value: "300",
						Usage: "Specify an TTL",
					},
					cli.StringFlag{
						Name:  "type",
						Value: "A",
						Usage: "Specify an record type",
					},
					cli.StringFlag{
						Name:  "data",
						Usage: "Specify an data",
					},
					cli.StringFlag{
						Name:  "cname",
						Usage: "Specify an reference domain",
					},
					cli.StringFlag{
						Name:  "prio",
						Value: "10",
						Usage: "Specify an MX record priority",
					},
					cli.StringFlag{
						Name:  "exchange",
						Usage: "Specify an MX record exhange",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name.", "USAGE: gdcli record add <ZONE NAME> --type A --ip x.x.x.x --name www --ttl 300")
						os.Exit(2)
					}
					var addParams addRecordRequestParam
					addParams.Zone = c.Args().Get(0)
					addParams.Address = c.String("ip")
					addParams.Name = c.String("name")
					addParams.Type = c.String("type")
					addParams.Data = c.String("data")
					addParams.Cname = c.String("cname")
					addParams.Exchange = c.String("exchange")

					if prio, err := strconv.Atoi(c.String("prio")); err == nil {
						addParams.Prio = prio
					}

					if ttl, err := strconv.Atoi(c.String("ttl")); err == nil {
						addParams.TTL = ttl
					}

					addParams.Execute()
					return nil
				},
			},
			{
				Name:  "rm",
				Usage: "Remove one record",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "type, t",
						Value: "A",
						Usage: "Specify an record type",
					},
					cli.StringFlag{
						Name:  "name, n",
						Usage: "Specify an domain",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name.", "USAGE: gdcli record rm <ZONE NAME> --type A --name www")
						os.Exit(2)
					}
					var rRRP removeRecordRequestParam
					rRRP.Zone = c.Args().Get(0)
					rRRP.Name = c.String("name")
					rRRP.Type = c.String("type")

					rRRP.Execute()
					return nil
				},
			},
			{
				Name:  "update",
				Usage: "Update a record",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "ip",
						Usage: "Specify an IP Address",
					},
					cli.StringFlag{
						Name:  "name",
						Usage: "Specify an domain",
					},
					cli.StringFlag{
						Name:  "ttl",
						Value: "300",
						Usage: "Specify an TTL",
					},
					cli.StringFlag{
						Name:  "type",
						Value: "A",
						Usage: "Specify an record type",
					},
					cli.StringFlag{
						Name:  "data",
						Usage: "Specify an data",
					},
					cli.StringFlag{
						Name:  "cname",
						Usage: "Specify an reference domain",
					},
					cli.StringFlag{
						Name:  "prio",
						Value: "10",
						Usage: "Specify an MX record priority",
					},
					cli.StringFlag{
						Name:  "exchange",
						Usage: "Specify an MX record exhange",
					},
				},
				Action: func(c *cli.Context) error {
					if len(c.Args()) == 0 {
						fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name", "USAGE: gdcli record update <ZONE NAME> --type A --name www --ttl 300")
						os.Exit(2)
					}
					var upParams updateRecordRequestParam
					upParams.Zone = c.Args().Get(0)
					upParams.Address = c.String("ip")
					upParams.Name = c.String("name")
					upParams.Type = c.String("type")
					upParams.Data = c.String("data")
					upParams.Cname = c.String("cname")
					upParams.Exchange = c.String("exchange")

					if prio, err := strconv.Atoi(c.String("prio")); err == nil {
						upParams.Prio = prio
					}

					if ttl, err := strconv.Atoi(c.String("ttl")); err == nil {
						upParams.TTL = ttl
					}

					upParams.Execute()
					return nil
				},
			},
		},
	}
}

func (getParams *getRecordRequestParam) Execute() {
	zones, err := getZones()
	if err != nil {
		fmt.Printf("getZones() %s\n", err)
		os.Exit(2)
	}

	if len(zones) == 0 {
		fmt.Printf("%s\n", "Zone is not registered.")
		os.Exit(0)
	}

	for _, z := range zones {
		if z.Name == getParams.Zone {
			getParams.ZoneID = z.ID
			getParams.ZoneCurrentVersionID = z.CurrentVersionID
			records, err := getParams.get()
			if err != nil {
				fmt.Printf("get() %s\n", err)
				os.Exit(2)
			}

			if getParams.Name != "" {
				for _, r := range records {
					if r.Name == getParams.Name+"."+z.Name+"." {
						r.display()
						os.Exit(0)
					}
				}
				fmt.Printf("%s %s %s\n", "Record", getParams.Name+"."+z.Name, "not found.")
			} else {
				for _, r := range records {
					r.display()
				}
			}

			os.Exit(0)
		}
	}

	fmt.Printf("%s %s %s\n", "Zone", getParams.Zone, "not found.")
	os.Exit(2)
}

func (getParams *getRecordRequestParam) get() (rR []recordResponseBody, err error) {
	conf := *getDNSConfig()
	req, err := generateGetRecordRequest(getParams, conf.URL)
	if err != nil {
		return rR, err
	}

	generateRequestWithToken(req, conf.User, conf.Password)

	res, err := doRequest(req)
	if err != nil {
		return rR, err
	}

	if err = decodeRequestBody(res, &rR); err != nil {
		return rR, err
	}

	return rR, nil
}

func generateGetRecordRequest(getParams *getRecordRequestParam, baseURL string) (*http.Request, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s/versions/%s/records", baseURL, getParams.ZoneID, getParams.ZoneCurrentVersionID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (addParams *addRecordRequestParam) Execute() {
	if len(addParams.Zone) == 0 {
		fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name.", "USAGE: gdcli record add <ZONE NAME> --type A --ip x.x.x.x --name www")
		os.Exit(2)
	}

	if addParams.Type == "A" || addParams.Type == "AAAA" {
		if len(addParams.Address) == 0 {
			fmt.Printf("\n%s\n\n%s\n", "Please specify the --ip option.", "USAGE: gdcli record add <ZONE NAME> --type A --ip x.x.x.x --name www")
			os.Exit(2)
		}
	}

	zones, err := getZones()
	if err != nil {
		fmt.Printf("getZones() %s\n", err)
		os.Exit(2)
	}

	for _, z := range zones {
		if z.Name == addParams.Zone {
			addParams.ZoneID = z.ID
			addParams.ZoneCurrentVersionID = z.CurrentVersionID
			record, err := addParams.add()
			if err != nil {
				fmt.Printf("%s\n", err)
				os.Exit(2)
			}
			record.display()
			os.Exit(0)
		}
	}

	fmt.Printf("%s %s %s\n", "Zone", addParams.Zone, "not found.")
	os.Exit(2)
}

func (addParams *addRecordRequestParam) add() (*recordResponseBody, error) {
	conf := *getDNSConfig()
	var rR recordResponseBody
	reqParams := generateAddRecordRequestBody(addParams)
	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		return nil, err
	}

	req, err := generateAddRecordRequest(addParams, conf.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	generateRequestWithToken(req, conf.User, conf.Password)

	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	if err = decodeRequestBody(res, &rR); err != nil {
		return nil, err
	}

	return &rR, nil
}

func generateAddRecordRequest(addParams *addRecordRequestParam, baseURL string, reqBody *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/%s/versions/%s/records", baseURL, addParams.ZoneID, addParams.ZoneCurrentVersionID),
		reqBody,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func generateAddRecordRequestBody(addParams *addRecordRequestParam) *addRecordRequestBody {
	var addParamsRB addRecordRequestBody
	if len(addParams.Name) == 0 {
		addParamsRB.Name = addParams.Zone + "."
	} else {
		addParamsRB.Name = addParams.Name + "." + addParams.Zone + "."
	}
	addParamsRB.TTL = addParams.TTL
	addParamsRB.Type = addParams.Type
	addParamsRB.EnableAlias = false

	if addParams.Type == "A" || addParams.Type == "AAAA" {
		addParamsRB.Records = []addParamsecordRequestBody{
			{
				Address: addParams.Address,
			},
		}
	} else if addParams.Type == "TXT" {
		addParamsRB.Records = []txtRecordRequestBody{
			{
				Data: addParams.Data,
			},
		}
	} else if addParams.Type == "CNAME" {
		addParamsRB.Records = []cnameRecordRequestBody{
			{
				Cname: addParams.Cname,
			},
		}
	} else if addParams.Type == "MX" {
		addParamsRB.Records = []mxRecordRequestBody{
			{
				Prio:     addParams.Prio,
				Exchange: addParams.Exchange,
			},
		}
	}

	return &addParamsRB
}

func (upParams *updateRecordRequestParam) Execute() {
	if len(upParams.Zone) == 0 {
		fmt.Printf("\n%s\n\n%s\n", "Please specify the zone name.", "USAGE: gdcli record add <ZONE NAME> --type A --ip x.x.x.x --name www")
		os.Exit(2)
	}

	if upParams.Type == "A" || upParams.Type == "AAAA" {
		if len(upParams.Address) == 0 {
			fmt.Printf("\n%s\n\n%s\n", "Please specify the --ip option.", "USAGE: gdcli record add <ZONE NAME> --type A --ip x.x.x.x --name www")
			os.Exit(2)
		}
	}

	zones, err := getZones()
	if err != nil {
		fmt.Printf("getZones() %s\n", err)
		os.Exit(2)
	}

	for _, z := range zones {
		if z.Name == upParams.Zone {
			var getParams getRecordRequestParam
			getParams.ZoneID = z.ID
			getParams.ZoneCurrentVersionID = z.CurrentVersionID
			records, err := getParams.get()
			if err != nil {
				fmt.Printf("Execute() %s\n", err)
				os.Exit(2)
			}

			if len(records) == 0 {
				fmt.Printf("%s\n", "Record not found.")
				os.Exit(0)
			}

			for _, r := range records {
				fqdn := upParams.Name + "." + upParams.Zone + "."
				if len(upParams.Name) == 0 {
					fqdn = upParams.Zone + "."
				}

				if r.Name == fqdn && upParams.Type == r.Type {
					upParams.ZoneID = z.ID
					upParams.ZoneCurrentVersionID = z.CurrentVersionID
					upParams.RecordID = r.ID
					record, err := upParams.update()
					if err != nil {
						fmt.Printf("update() %s\n", err)
						os.Exit(2)
					}

					record.display()
					os.Exit(0)
				}
			}
		}
	}
}

func (upParams *updateRecordRequestParam) update() (*recordResponseBody, error) {
	conf := *getDNSConfig()
	var rR recordResponseBody
	reqParams := generateUpdateRequestBody(upParams)
	reqBody, err := json.Marshal(reqParams)
	if err != nil {
		return nil, err
	}

	req, err := generateUpdateRequest(upParams, conf.URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	generateRequestWithToken(req, conf.User, conf.Password)

	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	if err = decodeRequestBody(res, &rR); err != nil {
		return nil, err
	}

	return &rR, nil
}

func generateUpdateRequest(upParams *updateRecordRequestParam, baseURL string, reqBody *bytes.Buffer) (*http.Request, error) {
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/%s/versions/%s/records/%s", baseURL, upParams.ZoneID, upParams.ZoneCurrentVersionID, upParams.RecordID),
		reqBody,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func generateUpdateRequestBody(upParams *updateRecordRequestParam) *addRecordRequestBody {
	var addParamsRB addRecordRequestBody
	addParamsRB.Name = upParams.Name + "." + upParams.Zone + "."
	if len(upParams.Name) == 0 {
		addParamsRB.Name = upParams.Zone + "."
	} else {
		addParamsRB.Name = upParams.Name + "." + upParams.Zone + "."
	}
	addParamsRB.TTL = upParams.TTL
	addParamsRB.Type = upParams.Type
	addParamsRB.EnableAlias = false

	if upParams.Type == "A" || upParams.Type == "AAAA" {
		addParamsRB.Records = []addParamsecordRequestBody{
			{
				Address: upParams.Address,
			},
		}
	} else if upParams.Type == "TXT" {
		addParamsRB.Records = []txtRecordRequestBody{
			{
				Data: upParams.Data,
			},
		}
	} else if upParams.Type == "CNAME" {
		addParamsRB.Records = []cnameRecordRequestBody{
			{
				Cname: upParams.Cname,
			},
		}
	} else if upParams.Type == "MX" {
		addParamsRB.Records = []mxRecordRequestBody{
			{
				Prio:     upParams.Prio,
				Exchange: upParams.Exchange,
			},
		}
	}

	return &addParamsRB
}

func (rmParams *removeRecordRequestParam) Execute() {
	zones, err := getZones()
	if err != nil {
		fmt.Printf("getZones() %s\n", err)
		os.Exit(2)
	}

	for _, z := range zones {
		if z.Name == rmParams.Zone {
			var getParams getRecordRequestParam
			getParams.ZoneID = z.ID
			getParams.ZoneCurrentVersionID = z.CurrentVersionID
			records, err := getParams.get()
			if err != nil {
				fmt.Printf("Execute() %s\n", err)
				os.Exit(2)
			}

			if len(records) == 0 {
				fmt.Printf("%s\n", "Record not found.")
				os.Exit(0)
			}

			for _, r := range records {
				fqdn := rmParams.Name + "." + rmParams.Zone + "."
				if len(rmParams.Name) == 0 {
					fqdn = rmParams.Zone + "."
				}
				if r.Name == fqdn && rmParams.Type == r.Type {
					rmParams.ZoneID = z.ID
					rmParams.ZoneCurrentVersionID = z.CurrentVersionID
					rmParams.RecordID = r.ID
					record, err := rmParams.remove()
					if err != nil {
						fmt.Printf("rmParams.remove() %s\n", err)
						os.Exit(2)
					}

					record.display()
					os.Exit(0)
				}
			}

			fmt.Printf("%s %s %s\n", "Record", rmParams.Name+"."+rmParams.Zone, "not found.")
			os.Exit(0)
		}
	}
}

func (rmParams *removeRecordRequestParam) remove() (*recordResponseBody, error) {
	conf := *getDNSConfig()
	req, err := generateRemoveRecordRequest(rmParams, conf.URL)
	if err != nil {
		return nil, err
	}

	generateRequestWithToken(req, conf.User, conf.Password)

	res, err := doRequest(req)
	if err != nil {
		return nil, err
	}

	var rR recordResponseBody
	if err = decodeRequestBody(res, &rR); err != nil {
		return nil, err
	}

	return &rR, nil
}

func generateRemoveRecordRequest(rR *removeRecordRequestParam, baseURL string) (*http.Request, error) {
	req, err := http.NewRequest(
		"DELETE",
		fmt.Sprintf("%s/%s/versions/%s/records/%s", baseURL, rR.ZoneID, rR.ZoneCurrentVersionID, rR.RecordID),
		nil,
	)
	if err != nil {
		return nil, err
	}

	return req, nil
}
