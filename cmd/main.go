package main

import (
	"github.com/nemo9nemo/mds-api-server/internal/common"
	"github.com/nemo9nemo/mds-api-server/internal/server"
)

func main() {
	common.LoadConfig()
	server.Start()
}
