package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mrmarble/deco"
)

const ENV_ROUTER_ADDR = "ROUTER_ADDR"
const ENV_ROUTER_PASSWD = "ROUTER_PASSWD"

func main() {
	var err error

	// sample code to reverse payload from webui
	// localStorage.getItem('encryptorAES') ==> "k=1723427984598410&i=1723427984598385"
	// aesKeys := utils.AESKey{}
	// aesKeys.Key = []byte("1723427984598410")
	// aesKeys.Iv = []byte("1723427984598385")

	// s, _ := url.QueryUnescape("Q6iL%2FAhNpPCOlppWDiTIbNIIJsDlm0VKuWjF3uvxq8k%3D")
	// res, err := utils.AES256Decrypt(s, aesKeys)
	// fmt.Println("res", res)
	// fmt.Println("err", err)
	// os.Exit(0)

	c := deco.New(os.Getenv(ENV_ROUTER_ADDR))
	err = c.Authenticate(os.Getenv(ENV_ROUTER_PASSWD))
	if err != nil {
		log.Fatal(err.Error())
	}

	devices, err := c.DeviceList()
	if err != nil {
		log.Fatal(err.Error())
	}

	deviceMacs := []string{}
	for _, d := range devices.Result.DeviceList {
		deviceMacs = append(deviceMacs, d.MAC)
	}
	fmt.Println(deviceMacs)
	_, err = c.Reboot(deviceMacs...)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func printPerformance(c *deco.Client) {
	fmt.Println("[+] Permormance")
	result, err := c.Performance()
	if err != nil {
		log.Fatal(err.Error())
	}
	// Print response as json
	jsonData, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(jsonData))
}

func printDevices(c *deco.Client) {
	fmt.Println("[+] Clients")
	result, err := c.ClientList()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, device := range result.Result.ClientList {

		fmt.Printf("%s\tOnline: %t\n", device.Name, device.Online)
	}
}

func printDecos(c *deco.Client) {
	fmt.Println("[+] Devices")
	result, err := c.DeviceList()
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, device := range result.Result.DeviceList {
		fmt.Printf("%s\tStatus: %s\n", device.DeviceIP, device.InetStatus)
	}
}
