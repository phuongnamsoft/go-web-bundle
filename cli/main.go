package cli

import (
	"flag"

	"github.com/phuongnamsoft/go-web-bundle/migrations"
)

func main() {
	migrate := flag.Bool("migrate", false, "Update db structure")
	if *migrate {
		migrations.Migrate()
	}
}
