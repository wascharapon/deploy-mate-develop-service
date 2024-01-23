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
	form          = flag.String("from", "", "Input string")
	to            = flag.String("to", "", "Input string")
	image         = flag.String("image", "", "Input string")
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
	e.Logger.Fatal(e.Start(":" + c.Port))

	flag.Parse()
	if *location == defaultString && *project == defaultString && *form == defaultString && *to == defaultString && *image == defaultString {
		fmt.Println("have data is empty")
	} else {
		fmt.Println("location is", *location, "project is", *project, "from is", *form, "to is", *to, "image is", *image)
		res, err := dmu.CopyAndDeploy(context.Background(), clientDeploy, domain.CopyAndDeployDto{
			Location: *location,
			Project:  *project,
			From:     *form,
			To:       *to,
			Image:    *image,
		})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			fmt.Println(res.Result)
			os.Exit(0)
		}
	}

}
