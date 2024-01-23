package domain

import (
	"context"

	"github.com/deploys-app/deploys/api/client"
)

type PullSecretsPassword struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

type PullSecrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DeployBody struct {
	Name                   *string  `json:"name"`
	FromProject            string   `json:"formProjectId"`
	ToProject              string   `json:"toProjectId"`
	GameServiceDatabaseURL string   `json:"gameServiceDBURL"`
	GameServiceURL         string   `json:"gameServiceUrl"`
	APIKey                 string   `json:"apiKey"`
	SeamlessServices       []string `json:"seamlessServices"`
	SeamlessDBHost         string   `json:"seamlessDBHost"`
	SeamlessDBPassword     string   `json:"seamlessDBPassword"`
	SeamlessDBUsername     string   `json:"seamlessDBUsername"`
	CopyEnv                bool     `json:"copyEnv"`
}

type DeploymentGetDto struct {
	Location string `json:"location" form:"location" validate:"required"`
	Project  string `json:"project" form:"project" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
}

type CopyAndDeployDto struct {
	Location string `json:"location" form:"location" validate:"required"`
	Project  string `json:"project" form:"project" validate:"required"`
	From     string `json:"from" form:"from" validate:"required"`
	To       string `json:"to" form:"to" validate:"required"`
	Image    string `json:"image" form:"image" validate:"required"`
}

type DeployMateUseCase interface {
	List(ctx context.Context, clientDeploy client.Client, project string) (*Response, error)
	Get(ctx context.Context, clientDeploy client.Client, deploymentGetDto DeploymentGetDto) (*Response, error)
	CopyAndDeploy(ctx context.Context, clientDeploy client.Client, copyAndDeployDto CopyAndDeployDto) (*Response, error)
}
