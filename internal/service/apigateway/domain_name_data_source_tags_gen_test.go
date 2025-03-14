// Code generated by internal/generate/tagstests/main.go; DO NOT EDIT.

package apigateway_test

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/config"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfstatecheck "github.com/hashicorp/terraform-provider-aws/internal/acctest/statecheck"
	tfapigateway "github.com/hashicorp/terraform-provider-aws/internal/service/apigateway"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccAPIGatewayDomainNameDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtKey1: config.StringVariable(acctest.CtValue1),
					}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtKey1: knownvalue.StringExact(acctest.CtValue1),
					})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_NullMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:          config.StringVariable(rName),
					acctest.CtResourceTags:   nil,
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_EmptyMap(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, names.APIGatewayServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				ConfigDirectory: config.StaticDirectory("testdata/DomainName/data.tags/"),
				ConfigVariables: config.Variables{
					acctest.CtRName:          config.StringVariable(rName),
					acctest.CtResourceTags:   config.MapVariable(map[string]config.Variable{}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_DefaultTags_nonOverlapping(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.APIGatewayServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/DomainName/data.tags_defaults/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_IgnoreTags_Overlap_DefaultTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.APIGatewayServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/DomainName/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtProviderTags: config.MapVariable(map[string]config.Variable{
						acctest.CtProviderKey1: config.StringVariable(acctest.CtProviderValue1),
					}),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtProviderKey1),
					),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
					expectFullDomainNameDataSourceTags(ctx, dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtProviderKey1: knownvalue.StringExact(acctest.CtProviderValue1),
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
			},
		},
	})
}

func TestAccAPIGatewayDomainNameDataSource_tags_IgnoreTags_Overlap_ResourceTag(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_api_gateway_domain_name.test"
	rName := acctest.RandomSubdomain()
	privateKeyPEM := acctest.TLSRSAPrivateKeyPEM(t, 2048)
	certificatePEM := acctest.TLSRSAX509SelfSignedCertificatePEM(t, privateKeyPEM, rName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:   func() { acctest.PreCheck(ctx, t) },
		ErrorCheck: acctest.ErrorCheck(t, names.APIGatewayServiceID),
		Steps: []resource.TestStep{
			{
				ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
				ConfigDirectory:          config.StaticDirectory("testdata/DomainName/data.tags_ignore/"),
				ConfigVariables: config.Variables{
					acctest.CtRName: config.StringVariable(rName),
					acctest.CtResourceTags: config.MapVariable(map[string]config.Variable{
						acctest.CtResourceKey1: config.StringVariable(acctest.CtResourceValue1),
					}),
					"ignore_tag_keys": config.SetVariable(
						config.StringVariable(acctest.CtResourceKey1),
					),
					acctest.CtCertificatePEM: config.StringVariable(certificatePEM),
					acctest.CtPrivateKeyPEM:  config.StringVariable(privateKeyPEM),
				},
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(dataSourceName, tfjsonpath.New(names.AttrTags), knownvalue.MapExact(map[string]knownvalue.Check{})),
					expectFullDomainNameDataSourceTags(ctx, dataSourceName, knownvalue.MapExact(map[string]knownvalue.Check{
						acctest.CtResourceKey1: knownvalue.StringExact(acctest.CtResourceValue1),
					})),
				},
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func expectFullDomainNameDataSourceTags(ctx context.Context, resourceAddress string, knownValue knownvalue.Check) statecheck.StateCheck {
	return tfstatecheck.ExpectFullDataSourceTagsSpecTags(tfapigateway.ServicePackage(ctx), resourceAddress, &types.ServicePackageResourceTags{
		IdentifierAttribute: names.AttrARN,
	}, knownValue)
}
