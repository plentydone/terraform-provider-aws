// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package elb

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	elb_sdkv1 "github.com/aws/aws-sdk-go/service/elb"
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
			Factory:  DataSourceLoadBalancer,
			TypeName: "aws_elb",
		},
		{
			Factory:  DataSourceHostedZoneID,
			TypeName: "aws_elb_hosted_zone_id",
		},
		{
			Factory:  DataSourceServiceAccount,
			TypeName: "aws_elb_service_account",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAppCookieStickinessPolicy,
			TypeName: "aws_app_cookie_stickiness_policy",
		},
		{
			Factory:  ResourceLoadBalancer,
			TypeName: "aws_elb",
			Name:     "Classic Load Balancer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceAttachment,
			TypeName: "aws_elb_attachment",
		},
		{
			Factory:  ResourceCookieStickinessPolicy,
			TypeName: "aws_lb_cookie_stickiness_policy",
		},
		{
			Factory:  ResourceSSLNegotiationPolicy,
			TypeName: "aws_lb_ssl_negotiation_policy",
		},
		{
			Factory:  ResourceBackendServerPolicy,
			TypeName: "aws_load_balancer_backend_server_policy",
		},
		{
			Factory:  ResourceListenerPolicy,
			TypeName: "aws_load_balancer_listener_policy",
		},
		{
			Factory:  ResourcePolicy,
			TypeName: "aws_load_balancer_policy",
		},
		{
			Factory:  ResourceProxyProtocolPolicy,
			TypeName: "aws_proxy_protocol_policy",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ELB
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*elb_sdkv1.ELB, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return elb_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
