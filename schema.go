package nuget

import (
	"encoding/json"
	"time"
)

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

type CatalogLeaf struct {
	Id                       string          `json:"@id"`
	Type                     []string        `json:"@type"`
	Authors                  string          `json:"authors"`
	CatalogCommitId          string          `json:"catalog:commitId"`
	CatalogCommitTimeStamp   string          `json:"catalog:commitTimeStamp"`
	Copyright                string          `json:"copyright"`
	Created                  string          `json:"created"`
	Description              string          `json:"description"`
	IconUrl                  string          `json:"iconUrl"`
	PackageId                string          `json:"id"`
	IsPrerelease             bool            `json:"isPrerelease"`
	LastEdited               string          `json:"lastEdited"`
	Listed                   bool            `json:"listed"`
	PackageHash              string          `json:"packageHash"`
	PackageHashAlgorithm     string          `json:"packageHashAlgorithm"`
	PackageSize              int             `json:"packageSize"`
	ProjectUrl               string          `json:"projectUrl"`
	Published                string          `json:"published"`
	RequireLicenseAcceptance bool            `json:"requireLicenseAcceptance"`
	VerbatimVersion          string          `json:"verbatimVersion"`
	Version                  string          `json:"version"`
	DependencyGroups         json.RawMessage `json:"dependencyGroups"`
	PackageEntries           json.RawMessage `json:"packageEntries"`
}
