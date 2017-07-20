package main

import (
	"net/http"
	"fmt"
	"log"
	. "github.com/QuentinDeSoete/wrike_client/wrike"
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

func main() {
	w := NewWrike(URL,WRIKE_TOKEN,&http.Client{})
	//var id Id
	//id.Account = "IEABJ5M4I7777777"
	//id.Folder = "IEABJ5M4I7777776"

	//m := w.GetFoldersTree(id)
	//fmt.Printf("%+v", m.DataJson[0])
	//var id []string
	//id = []string{"IEABJ5M4I7777777","IEABJ5M4I4ENXTM3","IEABJ5M4I4ES4OJP","IEABJ5M4I4ES5XYP","IEABJ5M4I4EPYBXI","IEABJ5M4I4ERZ2SZ","IEABJ5M4I4ERZ2SZ"}
	id := ""
	m := w.QueryAccounts(id)
	fmt.Printf("%+v", m)

}