package osvwrapper

import (
	"net/http"
	"bytes"
	"io"
	"github.com/synk-labs/parlay/lib/ecosystems"
	"github.com/package-url/packageurl-go"
)

func OSVQuery(purl string) (string, error) {
	purl_t, err := packageurl.FromString(purl)
	if err != nil {
		return "", err
	}

	eco_resp, err := ecosystems.GetPackageData(purl_t)
	if err != nil || eco_resp.JSON200 == nil || eco_resp.JSON200.RepositoryUrl == nil {
		return "", err
	}

	queryurl := "https://api.osv.dev/v1/query"
	jsonquery := []byte("{\"package\": {\"purl\": \"" + purl + "\"}}")
	response, err := http.Post(queryurl, "application/json", bytes.NewBuffer(jsonquery))
	if err != nil || response.StatusCode != http.StatusOK{
		return "", err
	}
	responsebody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	vuln_report := "{\"OSV\": " + string(responsebody) + "}"

	defer response.Body.Close()
	return vuln_report, nil
}


