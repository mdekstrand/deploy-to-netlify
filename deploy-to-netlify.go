package main

import (
	goctx "context"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/netlify/open-api/v2/go/porcelain"
	"github.com/netlify/open-api/v2/go/porcelain/context"
	"github.com/sirupsen/logrus"
)

var site_path = flag.String("d", "_site", "directory from which to deploy the site")

func add_auth_token(req runtime.ClientRequest, reg strfmt.Registry) error {
	token, found := os.LookupEnv("NETLIFY_AUTH_TOKEN")
	if !found {
		return errors.New("no authentication token")
	}
	req.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", token))
	return nil
}

func main() {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	flag.Parse()

	auth_info := runtime.ClientAuthInfoWriterFunc(add_auth_token)
	ctx := context.WithAuthInfo(goctx.TODO(), auth_info)

	n := porcelain.Default
	var deploy_req porcelain.DeployOptions
	deploy_req.SiteID = os.Getenv("NETLIFY_SITE_ID")
	deploy_req.Dir = *site_path
	deploy_req.IsDraft = false

	deploy, err := n.DeploySite(ctx, deploy_req)
	if err != nil {
		println(err.Error())
		panic("deploy failed")
	}
	fmt.Printf("deployed %s\n", deploy.ID)
}
