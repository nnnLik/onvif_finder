package main

import (
        "encoding/json"
        "fmt"
        "github.com/deepch/go-onvif"
        "net/http"
        "os"
        "strings"
        "time"
)

func fetchDevices() ([]string, error) {
        devices, err := onvif.StartDiscovery(100*time.Millisecond)
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

func getLastCameraIP(ipAddresses []string) []string {
        serverIP := os.Getenv("IP")
        extraIP_1 := "10.20.100.40"
        extraIP_2 := "10.20.100.41"
        if serverIP != "" {
                ipAddresses = append(ipAddresses, serverIP)
        }
        ipAddresses = append(ipAddresses, extraIP_1)
        ipAddresses = append(ipAddresses, extraIP_2)
        return ipAddresses
}

func handleGetAllOnvifCameras(w http.ResponseWriter, r *http.Request) {
        ipAddresses, err := fetchDevices()
        if err != nil {
                fmt.Println("Error: ", err)
                return
        }
        ipAddresses = getLastCameraIP(ipAddresses)

        json.NewEncoder(w).Encode(ipAddresses)
}

func main() {
        http.HandleFunc("/get_all_onvif_cameras/", handleGetAllOnvifCameras)
        http.ListenAndServe(":7654", nil)
}
