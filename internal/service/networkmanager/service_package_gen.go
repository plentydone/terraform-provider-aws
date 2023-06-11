// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package networkmanager

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	networkmanager_sdkv1 "github.com/aws/aws-sdk-go/service/networkmanager"
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
			Factory:  DataSourceConnection,
			TypeName: "aws_networkmanager_connection",
		},
		{
			Factory:  DataSourceConnections,
			TypeName: "aws_networkmanager_connections",
		},
		{
			Factory:  DataSourceCoreNetworkPolicyDocument,
			TypeName: "aws_networkmanager_core_network_policy_document",
		},
		{
			Factory:  DataSourceDevice,
			TypeName: "aws_networkmanager_device",
		},
		{
			Factory:  DataSourceDevices,
			TypeName: "aws_networkmanager_devices",
		},
		{
			Factory:  DataSourceGlobalNetwork,
			TypeName: "aws_networkmanager_global_network",
		},
		{
			Factory:  DataSourceGlobalNetworks,
			TypeName: "aws_networkmanager_global_networks",
		},
		{
			Factory:  DataSourceLink,
			TypeName: "aws_networkmanager_link",
		},
		{
			Factory:  DataSourceLinks,
			TypeName: "aws_networkmanager_links",
		},
		{
			Factory:  DataSourceSite,
			TypeName: "aws_networkmanager_site",
		},
		{
			Factory:  DataSourceSites,
			TypeName: "aws_networkmanager_sites",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAttachmentAccepter,
			TypeName: "aws_networkmanager_attachment_accepter",
		},
		{
			Factory:  ResourceConnectAttachment,
			TypeName: "aws_networkmanager_connect_attachment",
			Name:     "Connect Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConnectPeer,
			TypeName: "aws_networkmanager_connect_peer",
			Name:     "Connect Peer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceConnection,
			TypeName: "aws_networkmanager_connection",
			Name:     "Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceCoreNetwork,
			TypeName: "aws_networkmanager_core_network",
			Name:     "Core Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceCoreNetworkPolicyAttachment,
			TypeName: "aws_networkmanager_core_network_policy_attachment",
		},
		{
			Factory:  ResourceCustomerGatewayAssociation,
			TypeName: "aws_networkmanager_customer_gateway_association",
		},
		{
			Factory:  ResourceDevice,
			TypeName: "aws_networkmanager_device",
			Name:     "Device",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceGlobalNetwork,
			TypeName: "aws_networkmanager_global_network",
			Name:     "Global Network",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceLink,
			TypeName: "aws_networkmanager_link",
			Name:     "Link",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceLinkAssociation,
			TypeName: "aws_networkmanager_link_association",
		},
		{
			Factory:  ResourceSite,
			TypeName: "aws_networkmanager_site",
			Name:     "Site",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceSiteToSiteVPNAttachment,
			TypeName: "aws_networkmanager_site_to_site_vpn_attachment",
			Name:     "Site To Site VPN Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTransitGatewayConnectPeerAssociation,
			TypeName: "aws_networkmanager_transit_gateway_connect_peer_association",
		},
		{
			Factory:  ResourceTransitGatewayPeering,
			TypeName: "aws_networkmanager_transit_gateway_peering",
			Name:     "Transit Gateway Peering",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceTransitGatewayRegistration,
			TypeName: "aws_networkmanager_transit_gateway_registration",
		},
		{
			Factory:  ResourceTransitGatewayRouteTableAttachment,
			TypeName: "aws_networkmanager_transit_gateway_route_table_attachment",
			Name:     "Transit Gateway Route Table Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
		{
			Factory:  ResourceVPCAttachment,
			TypeName: "aws_networkmanager_vpc_attachment",
			Name:     "VPC Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "arn",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.NetworkManager
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*networkmanager_sdkv1.NetworkManager, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return networkmanager_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
