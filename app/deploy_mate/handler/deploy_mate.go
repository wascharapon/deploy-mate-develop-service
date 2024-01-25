package handler

import (
	"DeployMateDevelopService/domain"
	"DeployMateDevelopService/pkg/api"

	"github.com/deploys-app/deploys/api/client"
	"github.com/labstack/echo/v4"
)

type deployMateHandler struct {
	DeployMateUseCase domain.DeployMateUseCase
	clientDeploy      client.Client
}

func InitDeployMateHandler(e *echo.Echo, DeployMateUseCase domain.DeployMateUseCase, clientDeploy client.Client) {
	h := &deployMateHandler{
		DeployMateUseCase,
		clientDeploy,
	}
	dm := e.Group("/deploy-mate")
	dm.GET("/list/:project", h.list)
	dm.POST("/get", h.get)
	dm.POST("/copyAndDeploy", h.copyAndDeploy)
	dm.POST("/delete", h.delete)
}

func (dmh *deployMateHandler) list(c echo.Context) error {
	project := c.Param("project")
	res, err := dmh.DeployMateUseCase.List(c.Request().Context(), dmh.clientDeploy, project)
	if err != nil {
		return err
	}
	return c.JSON(res.Status, res)
}

func (dmh *deployMateHandler) get(c echo.Context) error {
	var dto domain.DeploymentGetDto
	if err := api.BindQueryParams(c, &dto); err != nil {
		return err
	}
	if err := c.Bind(&dto); err != nil {
		return domain.ErrorBindStructFailed.SetMessage(domain.GetAccountDetail)
	}
	res, err := dmh.DeployMateUseCase.Get(c.Request().Context(), dmh.clientDeploy, dto)
	if err != nil {
		return err
	}
	return c.JSON(res.Status, res)
}

func (dmh *deployMateHandler) copyAndDeploy(c echo.Context) error {
	var dto domain.CopyAndDeployDto
	if err := api.BindQueryParams(c, &dto); err != nil {
		return err
	}
	if err := c.Bind(&dto); err != nil {
		return domain.ErrorBindStructFailed.SetMessage(domain.GetAccountDetail)
	}
	res, err := dmh.DeployMateUseCase.CopyAndDeploy(c.Request().Context(), dmh.clientDeploy, dto)
	if err != nil {
		return err
	}
	return c.JSON(res.Status, res)
}

func (dmh *deployMateHandler) delete(c echo.Context) error {
	var dto domain.DeploymentDeleteDto
	if err := api.BindQueryParams(c, &dto); err != nil {
		return err
	}
	if err := c.Bind(&dto); err != nil {
		return domain.ErrorBindStructFailed.SetMessage(domain.GetAccountDetail)
	}
	res, err := dmh.DeployMateUseCase.Delete(c.Request().Context(), dmh.clientDeploy, dto)
	if err != nil {
		return err
	}
	return c.JSON(res.Status, res)
}
