package waybackurls

import (
	"os"
	"strings"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

type testValues struct {
	Domain  string
	Domains []string
	Output  string
}

func NewTestingValues() testValues {
	domains := strings.Split(os.Getenv("DOMAINS"), ",")
	return testValues{
		Domain:  domains[0],
		Domains: domains,
		Output:  os.Getenv("OUTPUT"),
	}

}

var tv = NewTestingValues()

func TestWaybackURLS(t *testing.T) {
	t.Log(tv.Domains)
	wurls := WaybackURLS(tv.Domains, "")
	t.Log(wurls)
}
