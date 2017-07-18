package structures

import (
	"net/http"
	"io/ioutil"
	"log"
	"encoding/json"
)

type Wrike struct{
	url string
	token string
	client *http.Client
}

type ModifyFolder struct{
	Title string
	Description string
	AddParents []string
	RemoveParents []string
	AddShareds []string
	RemoveShareds []string
	Metadata []string
	Restore bool
	CustomFields []string
	CustomColumns []string
	Project Project
}

type CopyFolder struct{
	Parent string
	Title string
	CopyDescriptions bool
	CopyResponsibles bool
	AddResponsibles []string
	RemoveResponsibles []string
	CopyCustomFields bool
	CopyCustomStatuses bool
	CopyStatuses bool
	CopyParents bool
	RescheduleDate string
	RescheduleMode string
	EntryLimit int
}

type CreateFolder struct{
	Title string
	Description string
	Shareds []string
	Metadata []string
	CustomFields []string
	CustomColumns []string
	Project Project
}

type Project struct {
	AuthorId string        `json:"authorId"`
	OwnerIds []string      `json:"ownerIds"`
	Status string          `json:"status"`
	CreatedDate string     `json:"createdDate"`
	StartDate string       `json:"startDate"`
	EndDate string         `json:"endDate"`
}

type Data struct {
	Id string              `json:"id"`
	Title string           `json:"title"`
	ChildIds []string      `json:"childIds"`
	Scope string           `json:"scope"`
	Project Project        `json:"project"`
	CreatedDate string     `json:"createdDate"`
	UpdatedDate string     `json:"updatedDate"`
	BriefDescription string`json:"briefDescription"`
	Description string     `json:"description"`
	Color string           `json:"color"`
	SharedIds []string     `json:"sharedIds"`
	ParentIds []string     `json:"parentIds"`
	SuperParentIds []string`json:"superParentIds"`
	HasAttachments bool    `json:"hasAttachments"`
	Permalink string       `json:"permalink"`
	WorkflowId string      `json:"workflowId"`
	Metadata []string      `json:"metadata"`
	CustomFields []string  `json:"customFields"`
}

type Folders struct {
	Kind string            `json:"kind"`
	DataJson []Data        `json:"data"`
}

type Id struct {
	Account string
	Folder string
}

func FormatResponse(resp *http.Response) (m Folders){
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	//fmt.Printf("%s", body)
	err = json.Unmarshal(body, &m)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	//fmt.Printf("%+v", m.DataJson[0])
	return
}