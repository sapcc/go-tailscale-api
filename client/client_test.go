package client

import (
	"net/url"
	"os"
	"testing"
)

func TestClient_all(t *testing.T) {
	url, err := url.Parse("https://tailscale.global.cloud.sap")
	if err != nil {
		t.Fatal(err)
	}
	client, err := New(*url, os.Getenv("API_KEY"), "sap.com")
	if err != nil {
		t.Fatal(err)
	}
	devices, err := client.ListDevices()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("device len: %d", len(devices))
	id := devices[0].Id
	device, err := client.GetDevice(id)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("device name: %s", device.Name)
	err = client.EnableAllRoutes(device.Id)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("successfully set routes")
}
