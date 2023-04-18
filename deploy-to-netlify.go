package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/netlify/open-api/v2/go/porcelain"
)

var site_path = flag.String("d", "_site", "directory from which to deploy the site")

func main() {
	flag.Parse()

	n := porcelain.Default
	var deploy_req porcelain.DeployOptions
	deploy_req.SiteID = os.Getenv("NETLIFY_SITE_ID")
	deploy_req.Dir = *site_path
	deploy_req.IsDraft = false

	deploy, err := n.DeploySite(nil, deploy_req)
	if err != nil {
		panic(fmt.Sprintf("deploy failed: %s", err))
	}
	fmt.Printf("deployed %s\n", deploy.ID)
}
