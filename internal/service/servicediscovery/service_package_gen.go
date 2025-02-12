// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

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
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_service_discovery_instance",
		},
		{
			Factory:  ResourcePrivateDNSNamespace,
			TypeName: "aws_service_discovery_private_dns_namespace",
		},
		{
			Factory:  ResourcePublicDNSNamespace,
			TypeName: "aws_service_discovery_public_dns_namespace",
		},
		{
			Factory:  ResourceService,
			TypeName: "aws_service_discovery_service",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceDiscovery
}

var ServicePackage = &servicePackage{}
