package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
   {
      "addresses": [
        "100.91.252.79"
      ],
      "authorized": true,
      "blocksIncomingConnections": false,
      "clientVersion": "1.12.3-td91ea7286-ge1bbbd90c",
      "created": "2021-08-18T13:20:39Z",
      "expires": "2022-02-14T13:20:39Z",
      "hostname": "tail01.cc.eu-de-1.cloud.sap",
      "id": "56479",
      "isExternal": false,
      "keyExpiryDisabled": true,
      "lastSeen": "2021-08-19T10:14:26Z",
      "machineKey": "mkey:3d089f7572af5af814bd24cc681176a6d8e6a8e3f3527e7c12b343fdd92e3131",
      "name": "tail01-cc-eu-de-1-cloud-sap.sap.com",
      "nodeKey": "nodekey:331616bbf18945315dde77e0f6aebbdd5a698d754b2e0d00fd2c50196b9bfc4e",
      "os": "linux",
      "updateAvailable": false,
      "user": "dmitri.fedotov@sap.com"
    },
*/

type Device struct {
	Authorized        bool   `json:"authorized"`
	ClientVersion     string `json:"clientVersion"`
	Hostname          string `json:"hostname"`
	Id                string `json:"id"`
	KeyExpiryDisabled bool   `json:"keyExpiryDisabled"`
	Name              string `json:"name"`
	Os                string `json:"os"`
	User              string `json:"user"`
}

type DeviceResponse struct {
	Devices []Device `json:"devices"`
}

func (c *Client) ListDevices() ([]Device, error) {
	req, err := http.NewRequest("GET", c.target.String()+basePath+"tailnet/"+c.tailNet+"/devices", nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiToken, "")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("unexpected http status code of %d", res.StatusCode)
		return nil, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resObj = DeviceResponse{}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, err
	}
	return resObj.Devices, nil
}

func (c *Client) GetDevice(id string) (*Device, error) {
	req, err := http.NewRequest("GET", c.target.String()+basePath+"device/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.apiToken, "")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("unexpected http status code of %d", res.StatusCode)
		return nil, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var device = Device{}
	err = json.Unmarshal(bytes, &device)
	if err != nil {
		return nil, err
	}
	return &device, nil
}

func (c *Client) UpdateDevice(device Device) error {
	return nil
}

func (c *Client) DeleteDevice(id string) error {
	req, err := http.NewRequest("DELETE", c.target.String()+basePath+"device/"+id, nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.apiToken, "")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("unexpected http status code of %d", res.StatusCode)
		return err
	}
	return nil
}
