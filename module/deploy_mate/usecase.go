package deploy_mate

import (
	"DeployMateDevelopService/domain"
	"context"
	"encoding/json"
	"net/http"

	"github.com/deploys-app/deploys/api"
	"github.com/deploys-app/deploys/api/client"
)

type useCase struct {
}

var (
	defaultPort             = 8080
	defaultWorkloadIdentity = ""
	minReplicas             = 1
	maxReplicas             = 4
	defaultProtocol         = api.DeploymentProtocol("http")
)

func NewUseCase() domain.DeployMateUseCase {
	return &useCase{}
}

func (u *useCase) List(ctx context.Context, clientDeploy client.Client, project string) (*domain.Response, error) {
	deployment, err := clientDeploy.Deployment().List(ctx, &api.DeploymentList{Project: project})

	resp := &domain.Response{
		Title: domain.CheckServerStatus, Status: http.StatusOK, Description: domain.Success, Result: deployment,
	}

	if err != nil {
		resp.Description = domain.NotFound
		resp.Status = http.StatusFound
	}

	return resp, nil
}

func (u *useCase) Get(ctx context.Context, clientDeploy client.Client, deploymentGet domain.DeploymentGetDto) (*domain.Response, error) {
	deployment, err := clientDeploy.Deployment().Get(ctx, &api.DeploymentGet{
		Location: deploymentGet.Location,
		Project:  deploymentGet.Project,
		Name:     deploymentGet.Name,
	})

	resp := &domain.Response{
		Title: domain.GetAccountDetail, Status: http.StatusOK, Description: domain.Success, Result: deployment,
	}

	if err != nil {
		resp.Description = domain.NotFound
		resp.Status = http.StatusFound
	}

	return resp, nil
}

func (u *useCase) CopyAndDeploy(ctx context.Context, clientDeploy client.Client, copyAndDeployDto domain.CopyAndDeployDto) (*domain.Response, error) {
	project, err := u.Get(ctx, clientDeploy, domain.DeploymentGetDto{
		Location: copyAndDeployDto.Location,
		Project:  copyAndDeployDto.Project,
		Name:     copyAndDeployDto.From,
	})

	resp := &domain.Response{
		Title: domain.CopyAndDeploy, Status: http.StatusOK, Description: domain.Success, Result: project,
	}

	if err != nil {
		resp.Description = domain.NotFound
		resp.Status = http.StatusFound
		return resp, nil
	}

	deploymentItem := &api.DeploymentItem{}

	jsonData, err := json.Marshal(project.Result)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, &deploymentItem)
	if err != nil {
		return nil, err
	}

	deploymentDeploy := api.DeploymentDeploy{
		Location:         copyAndDeployDto.Location,
		Project:          copyAndDeployDto.Project,
		Name:             copyAndDeployDto.To,
		Type:             api.DeploymentTypeWebService,
		Protocol:         &defaultProtocol,
		Port:             &defaultPort,
		MinReplicas:      &minReplicas,
		MaxReplicas:      &maxReplicas,
		AddEnv:           deploymentItem.Env,
		Image:            deploymentItem.Image,
		WorkloadIdentity: &defaultWorkloadIdentity,
	}

	deployment, err := clientDeploy.Deployment().Deploy(ctx, &deploymentDeploy)

	resp.Result = deployment

	if err != nil {
		resp.Description = domain.Failed
		resp.Status = http.StatusBadRequest
	}

	return resp, nil
}
