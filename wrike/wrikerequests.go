package wrike

import (
	"net/http"
	"fmt"
	"log"
	"encoding/json"
	"bytes"
)

func DeleteWrikeRequest(path string, w *Wrike)(resp *http.Response){
	//var body []byte
	var r *http.Request
	r, err := http.NewRequest("DELETE", fmt.Sprint(w.url,path), nil)
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", w.token)))
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
	r, err = http.NewRequest("PUT", fmt.Sprint(w.url,path), bytes.NewBuffer(folderJson))
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", w.token)))
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
	r, err = http.NewRequest("POST", fmt.Sprint(w.url,path), bytes.NewBuffer(folderJson))
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", w.token)))
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
	r, err := http.NewRequest("GET", fmt.Sprint(w.url,path), nil)
	if err == nil {
		r.Header.Add("Authorization", fmt.Sprint("bearer", fmt.Sprint(" ", w.token)))
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