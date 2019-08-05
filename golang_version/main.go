/*
*
* Golang version of PowerOff/On VSI
* Use the followin example as a Trigger JSON Output
{
  "username": "1234567_itiro@br.ibm.com",
  "key": "dfasdflasdfhjkahsdkflhaksdhjfklhaksjdhfklahskjfhakshdfkj",
  "vsiname":"itiroteste02",
  "poweraction":"off"
}
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func Main(params map[string]interface{}) map[string]interface{} {
	// parse the input JSON
	username, ok := params["username"].(string)
	if !ok {
		username = "None"
	}
	apikey, ok := params["key"].(string)
	if !ok {
		apikey = "None"
	}
	vsiname, ok := params["vsiname"].(string)
	if !ok {
		vsiname = "None"
	}
	poweraction, ok := params["poweraction"].(string)
	if !ok {
		poweraction = "None"
	}

	// building the return msg
	msg := make(map[string]interface{})
	msg["username"] = username
	msg["key"] = apikey
	msg["vsiname"] = vsiname
	msg["poweraction"] = poweraction

	// get the list of all VSIs from IBM Cloud account
	fmt.Println("getting the list of VSIs...")
	url := "https://api.softlayer.com/rest/v3.1/SoftLayer_Account/getVirtualGuests"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	// finding the VSI ID
	fmt.Println("finding the VSI ID using hostname...")
	var vsiId float64
	var vsi []map[string]interface{}
	var hostname string
	json.Unmarshal([]byte(bodyText), &vsi)
	for _, result := range vsi {
		hostname, _ = result["hostname"].(string)
		if hostname == vsiname {
			vsiId = result["id"].(float64)
		}
	}
	var vsiIdStr string = strconv.Itoa(int(vsiId))
	fmt.Printf("Id: %s ( %s )\n", vsiIdStr, vsiname)
	msg["vsiid"] = vsiIdStr // including the vsiID into the return msg

	// make the request to IBM Cloud api / SoftLayer api
	fmt.Println("sending the http request...")
	if poweraction == "on" {
		url = "https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/" + vsiIdStr + "/powerOn"
	} else {
		url = "https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/" + vsiIdStr + "/powerOff"
	}
	client = &http.Client{}
	req, err = http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, apikey)
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	bodyText, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyText))
	fmt.Println("End.")

	// return the output JSON
	return msg
}
