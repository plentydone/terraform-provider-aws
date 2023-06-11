// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package servicediscovery

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	servicediscovery_sdkv1 "github.com/aws/aws-sdk-go/service/servicediscovery"
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
			Factory:  DataSourceDNSNamespace,
			TypeName: "aws_service_discovery_dns_namespace",
		},
		{
			Factory:  DataSourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
		},
		{
			Factory:  DataSourceService,
			TypeName: "aws_service_discovery_service",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
			Name:     "HTTP Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_service_discovery_instance",
		},
		{
			Factory:  ResourcePrivateDNSNamespace,
			TypeName: "aws_service_discovery_private_dns_namespace",
			Name:     "Private DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourcePublicDNSNamespace,
			TypeName: "aws_service_discovery_public_dns_namespace",
			Name:     "Public DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceService,
			TypeName: "aws_service_discovery_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceDiscovery
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*servicediscovery_sdkv1.ServiceDiscovery, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return servicediscovery_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
