package main

import (
	"encoding/json"
	"fmt"
	"github.com/deepch/go-onvif"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Result []string `json:"result"`
}

func fetchDevices() ([]string, error) {
	devices, err := onvif.StartDiscovery(100 * time.Millisecond)
	if err != nil {
		return nil, err
	}

	var ipAddresses []string
	for _, device := range devices {
		addr := strings.TrimPrefix(device.XAddr, "http://")
		addr = strings.TrimSuffix(addr, "/device_service")
		addr = strings.TrimSuffix(addr, "/onvif")
		ipAddresses = append(ipAddresses, addr)
	}

	return ipAddresses, nil
}

func handleGetAllOnvifCameras(w http.ResponseWriter, r *http.Request) {
	ipAddresses, err := fetchDevices()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	response := Response{Result: ipAddresses}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/get_all_onvif_cameras/", handleGetAllOnvifCameras)
	http.ListenAndServe(":7654", nil)
}
