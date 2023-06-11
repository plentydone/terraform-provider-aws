// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package kms

import (
	"context"
	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	kms_sdkv1 "github.com/aws/aws-sdk-go/service/kms"
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
			Factory:  DataSourceAlias,
			TypeName: "aws_kms_alias",
		},
		{
			Factory:  DataSourceCiphertext,
			TypeName: "aws_kms_ciphertext",
		},
		{
			Factory:  DataSourceCustomKeyStore,
			TypeName: "aws_kms_custom_key_store",
		},
		{
			Factory:  DataSourceKey,
			TypeName: "aws_kms_key",
		},
		{
			Factory:  DataSourcePublicKey,
			TypeName: "aws_kms_public_key",
		},
		{
			Factory:  DataSourceSecret,
			TypeName: "aws_kms_secret",
		},
		{
			Factory:  DataSourceSecrets,
			TypeName: "aws_kms_secrets",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAlias,
			TypeName: "aws_kms_alias",
		},
		{
			Factory:  ResourceCiphertext,
			TypeName: "aws_kms_ciphertext",
		},
		{
			Factory:  ResourceCustomKeyStore,
			TypeName: "aws_kms_custom_key_store",
		},
		{
			Factory:  ResourceExternalKey,
			TypeName: "aws_kms_external_key",
			Name:     "External Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceGrant,
			TypeName: "aws_kms_grant",
		},
		{
			Factory:  ResourceKey,
			TypeName: "aws_kms_key",
			Name:     "Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceKeyPolicy,
			TypeName: "aws_kms_key_policy",
		},
		{
			Factory:  ResourceReplicaExternalKey,
			TypeName: "aws_kms_replica_external_key",
			Name:     "Replica External Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
		{
			Factory:  ResourceReplicaKey,
			TypeName: "aws_kms_replica_key",
			Name:     "Replica Key",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "id",
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.KMS
}

func (p *servicePackage) Configure(config map[string]any) {
	p.config = config
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context) (*kms_sdkv1.KMS, error) {
	sess := p.config["session"].(*session_sdkv1.Session)

	return kms_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(p.config["endpoint"].(string))})), nil
}

var ServicePackage = &servicePackage{}
