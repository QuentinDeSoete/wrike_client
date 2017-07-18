package main

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"bytes"
	"github.com/QuentinDeSoete/wrike_client/structs"
	"github.com/QuentinDeSoete/wrike_client/folders"
)

var VERSION = "0.1.0"
var MESSAGE = "Wrike client"

const (
	URL = "https://www.wrike.com/api/v3"
	WRIKE_TOKEN = "dC0DcjO0EnFeKCDF7dwoAWDcpwFCdk8k53uzN7E8lFjcaX4ck83neUgzUAi4ptnW-N-WFIUK"
)

func debug(data []byte, err error) {
	if err == nil {
		fmt.Printf("%s\n\n", data)
	} else {
		log.Fatalf("%s\n\n", err)
	}
}

func DeleteWrikeRequest(path string, w *Wrike)(resp *http.Response){
	//var body []byte
	var r *http.Request
	r, err := http.NewRequest("DELETE", fmt.Sprint(URL,path), nil)
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", WRIKE_TOKEN)))
		//debug(httputil.DumpRequestOut(r, true))
		resp, err = w.client.Do(r)
	}

	//if err == nil {
	//	defer resp.Body.Close()
	//	//debug(httputil.DumpResponse(resp, true))
	//	body, err = ioutil.ReadAll(resp.Body)
	//}

	//if err == nil {
	//	fmt.Printf("%s", body)
	//} else {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return resp
}

func PutWrikeRequest(path string, c interface{},  w *Wrike)(resp *http.Response){
	//var body []byte
	folderJson , err := json.Marshal(c)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	var r *http.Request
	r, err = http.NewRequest("PUT", fmt.Sprint(URL,path), bytes.NewBuffer(folderJson))
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", WRIKE_TOKEN)))
		//debug(httputil.DumpRequestOut(r, true))
		resp, err = w.client.Do(r)
	}

	//if err == nil {
	//	defer resp.Body.Close()
	//	//debug(httputil.DumpResponse(resp, true))
	//	body, err = ioutil.ReadAll(resp.Body)
	//}

	//if err == nil {
	//	fmt.Printf("%s", body)
	//} else {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return resp
}

func PostWrikeRequest(path string, c interface{},  w *Wrike)(resp *http.Response){
	//var body []byte
	folderJson , err := json.Marshal(c)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	var r *http.Request
	r, err = http.NewRequest("POST", fmt.Sprint(URL,path), bytes.NewBuffer(folderJson))
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", WRIKE_TOKEN)))
		//debug(httputil.DumpRequestOut(r, true))
		resp, err = w.client.Do(r)
	}

	//if err == nil {
	//	defer resp.Body.Close()
	//	//debug(httputil.DumpResponse(resp, true))
	//	body, err = ioutil.ReadAll(resp.Body)
	//}

	//if err == nil {
	//	fmt.Printf("%s", body)
	//} else {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return resp
}

func GetWrikeRequest(path string, w *Wrike)(resp *http.Response){
	//var body []byte
	var r *http.Request
	r, err := http.NewRequest("GET", fmt.Sprint(URL,path), nil)
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", WRIKE_TOKEN)))
		//debug(httputil.DumpRequestOut(r, true))
		resp, err = w.client.Do(r)
	}

	//if err == nil {
	//	defer resp.Body.Close()
	//	//debug(httputil.DumpResponse(resp, true))
	//	body, err = ioutil.ReadAll(resp.Body)
	//}

	//if err == nil {
	//	fmt.Printf("%s", body)
	//} else {
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	return resp
}

func (w *Wrike) QueryAccounts(id string) (m Folders){
	var resp *http.Response
	if id != "" {
		resp = GetWrikeRequest(fmt.Sprint("/accounts/", id), w)
	} else {
		resp = GetWrikeRequest("/accounts", w)
	}
	//debug(httputil.DumpResponse(resp, true))
	m = FormatResponse(resp)
	return
}

func main() {
	w := Wrike{URL,WRIKE_TOKEN, &http.Client{}}
	//var id Id
	//id.Account = "IEABJ5M4I7777777"
	//id.Folder = "IEABJ5M4I7777776"

	//m := Wrike{}.GetFoldersTree(id)
	//fmt.Printf("%+v", m.DataJson[0])
	//var id []string
	//id = []string{"IEABJ5M4I7777777","IEABJ5M4I4ENXTM3","IEABJ5M4I4ES4OJP","IEABJ5M4I4ES5XYP","IEABJ5M4I4EPYBXI","IEABJ5M4I4ERZ2SZ","IEABJ5M4I4ERZ2SZ"}
	id := ""
	m := w.QueryAccounts(id)
	fmt.Printf("%+v", m)

}