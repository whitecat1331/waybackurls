package waybackurls

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

func TestWaybackURLS(t *testing.T) {
	domains := []string{os.Getenv("DOMAIN")}
	fmt.Println(domains)
	wurls := WaybackURLS([]string{""})
	fmt.Println(wurls)
}
