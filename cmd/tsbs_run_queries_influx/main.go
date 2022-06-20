// tsbs_run_queries_influx speed tests InfluxDB using requests from stdin.
//
// It reads encoded Query objects from stdin, and makes concurrent requests
// to the provided HTTP endpoint. This program has no knowledge of the
// internals of the endpoint.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/blagojts/viper"
	"github.com/spf13/pflag"
	"github.com/timescale/tsbs/internal/utils"
	"github.com/timescale/tsbs/pkg/query"
)

// Program option vars:
var (
	daemonUrls []string
	chunkSize  uint64
	// InfluxDB v2
	organization string
	token        string
	bucketId     string
	orgId        string
)

// Global vars:
var (
	runner *query.BenchmarkRunner
)

// Parse args:
func init() {
	var config query.BenchmarkRunnerConfig
	config.AddToFlagSet(pflag.CommandLine)
	var csvDaemonUrls string

	pflag.String("urls", "http://localhost:8086", "Daemon URLs, comma-separated. Will be used in a round-robin fashion.")
	pflag.Uint64("chunk-response-size", 0, "Number of series to chunk results into. 0 means no chunking.")
	// InfluxDB v2
	pflag.String("organization", "", "Organization name (InfluxDB v2).")
	pflag.String("token", "", "Authentication token (InfluxDB v2).")

	pflag.Parse()

	err := utils.SetupConfigFile()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode config: %s", err))
	}

	csvDaemonUrls = viper.GetString("urls")
	chunkSize = viper.GetUint64("chunk-response-size")
	organization = viper.GetString("organization")
	token = viper.GetString("token")

	// Validate args
	daemonUrls = strings.Split(csvDaemonUrls, ",")
	if len(daemonUrls) == 0 {
		log.Fatal("missing 'urls' flag")
	}
	log.Printf("daemon URLs: %v\n", daemonUrls)

	if organization == "" {
		log.Fatal("organization must be specified for InfluxDB 2.x")
	}
	if token == "" {
		log.Fatal("token must be specified for InfluxDB 2.x")
	}
	organizations, err := listOrgs2(daemonUrls[0], organization)
	if err != nil {
		log.Fatalf("error listing organizations: %v", err)
	}
	orgId, _ = organizations[organization]
	if orgId == "" {
		log.Fatalf("organization '%s' not found", organization)
	}

	runner = query.NewBenchmarkRunner(config)
}

func main() {
	runner.Run(&query.HTTPPool, newProcessor)
}

type processor struct {
	w    *HTTPClient
	opts *HTTPClientDoOptions
}

func newProcessor() query.Processor { return &processor{} }

func (p *processor) Init(workerNumber int) {
	p.opts = &HTTPClientDoOptions{
		ContentType:          "application/vnd.flux", //application/json  /  application/vnd.flux
		Accept:               "application/csv",
		Debug:                runner.DebugLevel(),
		PrettyPrintResponses: runner.DoPrintResponses(),
		chunkSize:            chunkSize,
		database:             runner.DatabaseName(),
		Organization:         organization,
		AuthToken:            token,
		Path:                 []byte(fmt.Sprintf("/api/v2/query?orgID=%s", orgId)), // query path is empty for 2.x in generated queries
		// TODO: Add new stuff (token, etc) here
	}
	url := daemonUrls[workerNumber%len(daemonUrls)]
	p.w = NewHTTPClient(url)
}

func (p *processor) ProcessQuery(q query.Query, _ bool) ([]*query.Stat, error) {

	//opts.ContentType = "application/vnd.flux"
	//opts.Accept = "application/csv"
	//opts.Path = []byte(fmt.Sprintf("/api/v2/query?orgID=%s", orgId)) // query path is empty for 2.x in generated queries

	hq := q.(*query.HTTP)
	//hq.Path = p.opts.Path
	//fmt.Println(" ----- ====== HQ:")
	//fmt.Println(hq)
	lag, err := p.w.Do(hq, p.opts)
	if err != nil {
		return nil, err
	}
	stat := query.GetStat()
	stat.Init(q.HumanLabelName(), lag)
	return []*query.Stat{stat}, nil
}

func listOrgs2(daemonUrl string, orgName string) (map[string]string, error) {
	u := fmt.Sprintf("%s/api/v2/orgs", daemonUrl)
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("listOrgs2 newRequest error: %s", err.Error())
	}
	req.Header.Add("Authorization", fmt.Sprintf("Token %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("listOrgs2 GET error: %s", err.Error())
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("listOrgs2 GET status code: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("listOrgs2 readAll error: %s", err.Error())
	}

	type listingType struct {
		Orgs []struct {
			Id   string
			Name string
		}
	}
	var listing listingType
	err = json.Unmarshal(body, &listing)
	if err != nil {
		return nil, fmt.Errorf("listOrgs unmarshal error: %s", err.Error())
	}

	ret := make(map[string]string)
	for _, org := range listing.Orgs {
		ret[org.Name] = org.Id
	}
	return ret, nil
}
