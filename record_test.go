package gdcli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGenerateGetRecordRequest(t *testing.T) {
	ep := "https://api.gis.gehirn.jp/dns/v1/zones"
	var getParams getRecordRequestParam
	getParams.ZoneID = "xxxxxxxx-yyyy-xxxx-yyyy-xxxxxxxxxxxx"
	getParams.ZoneCurrentVersionID = "00000000-1111-2222-333333333333"

	req, err := generateGetRecordRequest(&getParams, ep)

	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "GET" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("%s/%s/versions/%s/records", ep, getParams.ZoneID, getParams.ZoneCurrentVersionID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	if req.Body != nil {
		t.Errorf("Unexpected request body.")
	}
}

func TestGenerateAddRecordRequest(t *testing.T) {
	ep := "https://api.gis.gehirn.jp/dns/v1/zones"
	var aR addRecordRequestParam
	aR.ZoneID = "xxxxxxxx-yyyy-xxxx-yyyy-xxxxxxxxxxxx"
	aR.ZoneCurrentVersionID = "00000000-1111-2222-333333333333"
	aR.Name = "test"
	aR.Zone = "example.com"
	aR.Address = "192.168.10.1"
	aR.TTL = 300
	aR.Type = "A"

	params := generateAddRecordRequestBody(&aR)
	body, err := json.Marshal(params)
	if err != nil {
		t.Errorf("Unexpected error JSON Marshal. %s", err)
	}

	req, err := generateAddRecordRequest(&aR, ep, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "POST" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("%s/%s/versions/%s/records", ep, aR.ZoneID, aR.ZoneCurrentVersionID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	defer req.Body.Close()
	if string(b) != fmt.Sprintf(`{"name":"%s","ttl":%d,"enable_alias":false,"type":"%s","records":[{"address":"%s"}]}`, params.Name, aR.TTL, aR.Type, aR.Address) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGenerateUpdateRecordRequest(t *testing.T) {
	ep := "https://api.gis.gehirn.jp/dns/v1/zones"
	var upParams updateRecordRequestParam
	upParams.ZoneID = "xxxxxxxx-yyyy-xxxx-yyyy-xxxxxxxxxxxx"
	upParams.ZoneCurrentVersionID = "00000000-1111-2222-333333333333"
	upParams.RecordID = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"
	upParams.Name = "test"
	upParams.Zone = "example.com"
	upParams.Address = "192.168.10.1"
	upParams.TTL = 300
	upParams.Type = "A"

	params := generateUpdateRequestBody(&upParams)
	body, err := json.Marshal(params)
	if err != nil {
		t.Errorf("Unexpected error JSON Marshal. %s", err)
	}

	req, err := generateUpdateRequest(&upParams, ep, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "PUT" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("%s/%s/versions/%s/records/%s", ep, upParams.ZoneID, upParams.ZoneCurrentVersionID, upParams.RecordID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Errorf("Failed to read request body. %s", err)
	}
	defer req.Body.Close()
	if string(b) != fmt.Sprintf(`{"name":"%s","ttl":%d,"enable_alias":false,"type":"%s","records":[{"address":"%s"}]}`, params.Name, upParams.TTL, upParams.Type, upParams.Address) {
		t.Errorf("Unexpected request body. %s", string(b))
	}
}

func TestGeneratupParamsemoveRecordRequest(t *testing.T) {
	ep := "https://api.gis.gehirn.jp/dns/v1/zones"
	var rmParams removeRecordRequestParam
	rmParams.ZoneID = "xxxxxxxx-yyyy-xxxx-yyyy-xxxxxxxxxxxx"
	rmParams.ZoneCurrentVersionID = "00000000-1111-2222-333333333333"
	rmParams.RecordID = "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee"

	req, err := generateRemoveRecordRequest(&rmParams, ep)

	if err != nil {
		t.Errorf("Unexpected error occurs. %s", err)
	}
	if req.Method != "DELETE" {
		t.Errorf("Unexpected request method. %s", req.Method)
	}
	if req.URL.String() != fmt.Sprintf("%s/%s/versions/%s/records/%s", ep, rmParams.ZoneID, rmParams.ZoneCurrentVersionID, rmParams.RecordID) {
		t.Errorf("Unexpected request path. %s", req.URL.String())
	}
	if req.Body != nil {
		t.Errorf("Unexpected request body.")
	}
}
