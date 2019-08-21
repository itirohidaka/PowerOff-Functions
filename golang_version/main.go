/**
 *
 * main() will be run when you invoke this action
 *
 * @param Cloud Functions actions accept a single parameter, which must be a JSON object.
 *
 * @return The output of this action, which must be a JSON object.
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Main is the function implementing the action
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
	poweraction, ok := params["poweraction"].(string)
	if !ok {
		poweraction = "None"
	}
	tag, ok := params["tag"].(string)
	if !ok {
		tag = "None"
	}
	fmt.Println(params)

	msg := make(map[string]interface{})
	msg["username"] = username
	msg["key"] = apikey
	msg["tag"] = tag
	msg["poweraction"] = poweraction

	// can optionally log to stdout (or stderr)
	fmt.Println(username)
	fmt.Println(apikey)

	fmt.Println("Starting Functions...")

	// get the list of all VSIs from IBM Cloud account
	fmt.Println("Getting the list of VSIs...")
	url := "https://api.softlayer.com/rest/v3/SoftLayer_Account/getVirtualGuests?objectFilter={\"virtualGuests\":{\"tagReferences\":{\"tag\":{\"name\":{\"operation\":\"in\",\"options\":[{\"name\":\"data\",\"value\":[\"" + tag + "\"]}]}}}}}"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	msg["body"] = string(bodyText)

	// Acting on the VSIs
	fmt.Println("Starting the Power Action on VSIs...")
	var vsiId float64
	var vsi []map[string]interface{}
	var vsiIdStr string
	var vsiname string
	json.Unmarshal([]byte(bodyText), &vsi)
	for _, result := range vsi {
		vsiname, _ = result["hostname"].(string)
		vsiId = result["id"].(float64)
		vsiIdStr = strconv.Itoa(int(vsiId))
		fmt.Printf("Id: %s ( %s )\n", vsiIdStr, vsiname)

		//goroutine function
		go say(poweraction, vsiIdStr, username, apikey)
	}
	fmt.Println("End.")

	//say(poweraction,vsiIdStr, username, apikey)
	fmt.Println("Sleeping...")
	time.Sleep(10000 * time.Millisecond)
	return msg
}

func say(poweraction string, vsiIdStr string, username string, apikey string) {
	var url string

	if poweraction == "on" {
		url = "https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/" + vsiIdStr + "/powerOn"
	} else {
		url = "https://api.softlayer.com/rest/v3.1/SoftLayer_Virtual_Guest/" + vsiIdStr + "/powerOff"
	}
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, apikey)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	fmt.Println(vsiIdStr)
	bodyText, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bodyText))
}
