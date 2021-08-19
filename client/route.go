package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RoutesResponse struct {
	AdvertisedRoutes []string `json:"advertisedRoutes"`
	EnabledRoutes    []string `json:"enabledRoutes"`
}

type RoutesParam struct {
	Routes []string `json:"routes"`
}

func (c *Client) GetRoutes(deviceId string) (available, enabled []string, err error) {
	req, err := http.NewRequest("GET", c.target.String()+basePath+"device/"+deviceId+"/routes", nil)
	if err != nil {
		return nil, nil, err
	}
	req.SetBasicAuth(c.apiToken, "")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != 200 {
		err = fmt.Errorf("unexpected http status code of %d", res.StatusCode)
		return nil, nil, err
	}
	bytes, err := ioutil.ReadAll(res.Body)
	var resObj = RoutesResponse{}
	err = json.Unmarshal(bytes, &resObj)
	if err != nil {
		return nil, nil, err
	}
	return resObj.AdvertisedRoutes, resObj.EnabledRoutes, nil
}

func (c *Client) EnableAllRoutes(deviceId string) error {
	avail, _, err := c.GetRoutes(deviceId)
	if err != nil {
		return err
	}
	params := RoutesParam{Routes: avail}
	body, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", c.target.String()+basePath+"device/"+deviceId+"/routes", bytes.NewBuffer(body))
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
