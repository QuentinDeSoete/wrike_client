package wrike

import (
	"net/http"
	"fmt"
)

func (w *Wrike) GetFoldersTree(id Id) (m Folders){
	var resp *http.Response
	if id.Account != "" {
		resp = GetWrikeRequest(fmt.Sprint("/accounts/", fmt.Sprint(id.Account,"/wrike")), w)
	} else if id.Folder != "" {
		resp = GetWrikeRequest(fmt.Sprint("/wrike/", fmt.Sprint(id.Folder,"/wrike")), w)
	} else {
		resp = GetWrikeRequest("/wrike", w)
	}
	//debug(httputil.DumpResponse(resp, true))
	m = FormatResponse(resp)
	return
}

func (w *Wrike) GetFolders(id []string) (m Folders){
	var resp *http.Response
	concat := id[0]
	if len(id) <= 100 {
		for i := 1; i < len(id); i++ {
			concat += ","
			concat += id[i]
		}
		resp = GetWrikeRequest(fmt.Sprint("/wrike/",concat), w)
		m = FormatResponse(resp)
	}
	return
}

func (w *Wrike) ModifyFolder (id string, c ModifyFolder) (m Folders) {
	resp := PutWrikeRequest(fmt.Sprint("/wrike/", id), c , w)
	m = FormatResponse(resp)
	return
}

func (w *Wrike) CopyFolder (id string, c CopyFolder) (m Folders) {
	resp := PostWrikeRequest(fmt.Sprint("/copy_folder/",id), c, w)
	m = FormatResponse(resp)
	return
}

func (w *Wrike) CreateFolder (id string, c CreateFolder) (m Folders) {
	resp := PostWrikeRequest(fmt.Sprint("/wrike/",fmt.Sprint(id, "/wrike")), c, w)
	m = FormatResponse(resp)
	return
}
