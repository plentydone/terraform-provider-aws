// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package amp

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	prometheusservice_sdkv1 "github.com/aws/aws-sdk-go/service/prometheusservice"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct {
	config map[string]any
}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceWorkspace,
			TypeName: "aws_prometheus_workspace",
		},
		{
			Factory:  DataSourceWorkspaces,
			TypeName: "aws_prometheus_workspaces",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAlertManagerDefinition,
			TypeName: "aws_prometheus_alert_manager_definition",
		},
		{
			Factory:  ResourceRuleGroupNamespace,
			TypeName: "aws_prometheus_rule_group_namespace",
		},
		{
			Factory:  ResourceWorkspace,
			TypeName: "aws_prometheus_workspace",
			Name:     "Workspace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.AMP
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*prometheusservice_sdkv1.PrometheusService, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return prometheusservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
