package folders

import (
	"net/http"
	"fmt"
	"github.com/QuentinDeSoete/wrike_client/structs"
)

func (w *Wrike) GetFoldersTree(id Id) (m Folders){
	var resp *http.Response
	if id.Account != "" {
		resp = GetWrikeRequest(fmt.Sprint("/accounts/", fmt.Sprint(id.Account,"/folders")), w)
	} else if id.Folder != "" {
		resp = GetWrikeRequest(fmt.Sprint("/folders/", fmt.Sprint(id.Folder,"/folders")), w)
	} else {
		resp = GetWrikeRequest("/folders", w)
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
		resp = GetWrikeRequest(fmt.Sprint("/folders/",concat), w)
		m = FormatResponse(resp)
	}
	return
}

func (w *Wrike) ModifyFolder (id string, c ModifyFolder) (m Folders) {
	resp := PutWrikeRequest(fmt.Sprint("/folders/", id), c , w)
	m = FormatResponse(resp)
	return
}

func (w *Wrike) CopyFolder (id string, c CopyFolder) (m Folders) {
	resp := PostWrikeRequest(fmt.Sprint("/copy_folder/",id), c, w)
	m = FormatResponse(resp)
	return
}

func (w *Wrike) CreateFolder (id string, c CreateFolder) (m Folders) {
	resp := PostWrikeRequest(fmt.Sprint("/folders/",fmt.Sprint(id, "/folders")), c, w)
	m = FormatResponse(resp)
	return
}

func (w *Wrike) DeleteFolder (id string) (m Folders) {
	resp := DeleteWrikeRequest(fmt.Sprint("/folders/", id), w)
	m = FormatResponse(resp)
	return
}