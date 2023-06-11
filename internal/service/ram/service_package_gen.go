// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package ram

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	ram_sdkv1 "github.com/aws/aws-sdk-go/service/ram"
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
			Factory:  DataSourceResourceShare,
			TypeName: "aws_ram_resource_share",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourcePrincipalAssociation,
			TypeName: "aws_ram_principal_association",
		},
		{
			Factory:  ResourceResourceAssociation,
			TypeName: "aws_ram_resource_association",
		},
		{
			Factory:  ResourceResourceShare,
			TypeName: "aws_ram_resource_share",
			Name:     "Resource Share",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceResourceShareAccepter,
			TypeName: "aws_ram_resource_share_accepter",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.RAM
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*ram_sdkv1.RAM, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return ram_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
