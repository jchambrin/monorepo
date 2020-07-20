package lookup

import (
	"testing"
)

var builds = []string{"/xxx/aaa/project1/BUILD", "/xxx/bbb/project2/BUILD", "/xxx/aaa/project1/aaa/project1.2/BUILD"}
var charts = []string{"/xxx/aaa/project1/charts/Chart.yaml", "/xxx/bbb/project2/Chart.yaml", "/xxx/aaa/project1/aaa/project1.2/charts/Helm/Chart.yaml", "/aaa/project1/charts/Chart.yaml"}
var lookupResults = &Results{
	Builds: builds,
	Charts: charts,
}

func TestSubCharts(t *testing.T) {
	results := lookupResults.GetProjectCharts("/xxx/aaa/project1/")
	if len(results) != 1 || results[0] != "/xxx/aaa/project1/charts/Chart.yaml" {
		t.Fail()
	}

}
