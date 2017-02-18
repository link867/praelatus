package cli

import (
	"fmt"

	"github.com/praelatus/praelatus/config"
	"github.com/urfave/cli"
)

func showConfig(c *cli.Context) error {
	fmt.Println(config.Cfg)
	return nil
}
