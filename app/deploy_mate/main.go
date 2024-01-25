package main

import (
	"DeployMateDevelopService/app/deploy_mate/config"
	"DeployMateDevelopService/app/deploy_mate/handler"
	"DeployMateDevelopService/app/deploy_mate/middleware"
	"DeployMateDevelopService/domain"
	"DeployMateDevelopService/module/deploy_mate"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/deploys-app/deploys/api/client"
	"github.com/labstack/echo/v4"
)

var (
	location      = flag.String("location", "", "Input string")
	project       = flag.String("project", "", "Input string")
	name          = flag.String("name", "", "Input string")
	to            = flag.String("to", "", "Input string")
	image         = flag.String("image", "", "Input string")
	action        = flag.String("action", "", "Input string")
	defaultString = ""
)

func main() {
	c := config.Init()
	e := echo.New()
	e.HTTPErrorHandler = middleware.EchoErrorHandler
	dmu := deploy_mate.NewUseCase()
	clientDeploy := client.Client{
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		Auth: func(r *http.Request) {
			r.SetBasicAuth(os.Getenv("DEPLOYS_AUTH_USER"), os.Getenv("DEPLOYS_AUTH_PASS"))
		},
	}

	handler.InitDeployMateHandler(e, dmu, clientDeploy)

	flag.Parse()
	if *action == defaultString {
		fmt.Println("have data is empty")
	} else {
		var err error
		var res *domain.Response
		if *location != defaultString && *project != defaultString && *name != defaultString {
			fmt.Println(*action)
			if *action == domain.ActionDeploy {
				if *to != defaultString || *image != defaultString {
					res, err = dmu.CopyAndDeploy(context.Background(), clientDeploy, domain.CopyAndDeployDto{
						Location: *location,
						Project:  *project,
						From:     *name,
						To:       *to,
						Image:    *image,
					})
				}
			} else if *action == domain.ActionDelete {
				res, err = dmu.Delete(context.Background(), clientDeploy, domain.DeploymentDeleteDto{
					Location: *location,
					Project:  *project,
					Name:     *name,
				})
			}
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(res)
			os.Exit(0)
		}
	}
	e.Logger.Fatal(e.Start(":" + c.Port))
}

// go run app/deploy_mate/main.go -location=gke.cluster-rcf2 -project=toberich-staging -name=admin-panel-fe -to=admin-panel-fe-test -image=asia-southeast1-docker.pkg.dev/scamo-group/toberich-stag/admin-panel-fe@sha256:ce4b60aa2c823ebf1df92942fcde1d2e2aa98c99d7a369827264c5732e77642 -action=deploy
// go run app/deploy_mate/main.go -location=gke.cluster-rcf2 -project=toberich-staging -name=admin-panel-fe-test -action=delete
