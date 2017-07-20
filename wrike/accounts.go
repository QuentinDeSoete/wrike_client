package wrike

import (
	"net/http"
	"fmt"
)

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
