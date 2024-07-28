package nuget

import "time"

const TimestampFormat = time.RFC3339Nano

type Services struct {
	Version   string            `json:"version"`
	Resources []ServiceResource `json:"resources"`
}
type ServiceResource struct {
	Id      string `json:"@id"`
	Type    string `json:"@type"`
	Comment string `json:"comment"`
}

type Catalog struct {
	Id               string        `json:"@id"`
	CommitId         string        `json:"commitId"`
	CommitTimeStamp  string        `json:"commitTimeStamp"`
	Count            int           `json:"count"`
	NugetLastCreated string        `json:"nuget:lastCreated"`
	NugetLastDeleted string        `json:"nuget:lastDeleted"`
	NugetLastEdited  string        `json:"nuget:lastEdited"`
	Items            []CatalogItem `json:"items"`
}

type CatalogItem struct {
	Id              string `json:"@id"`
	Type            string `json:"@type"`
	CommitId        string `json:"commitId"`
	CommitTimestamp string `json:"commitTimeStamp"`
	Count           int    `json:"count"`
}
