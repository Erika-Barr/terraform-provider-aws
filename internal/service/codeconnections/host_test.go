// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codeconnections_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/codeconnections/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfcodeconnections "github.com/hashicorp/terraform-provider-aws/internal/service/codeconnections"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccCodeConnectionsHost_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Host
	resourceName := "aws_codeconnections_host.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.CodeConnectionsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckHostDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "codeconnections", regexache.MustCompile("host/.+$")),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrID, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "provider_endpoint", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "provider_type", string(types.ProviderTypeGithubEnterpriseServer)),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCodeConnectionsHost_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Host
	resourceName := "aws_codeconnections_host.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.CodeConnectionsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckHostDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig_basic(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					acctest.CheckFrameworkResourceDisappears(ctx, acctest.Provider, tfcodeconnections.ResourceHost, resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccCodeConnectionsHost_vpc(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Host
	resourceName := "aws_codeconnections_host.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.CodeConnectionsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckHostDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig_vpcNoCertificate(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "codeconnections", regexache.MustCompile("host/.+$")),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrID, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "provider_endpoint", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "provider_type", string(types.ProviderTypeGithubEnterpriseServer)),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.0.security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.0.subnet_ids.#", "2"),
					resource.TestCheckNoResourceAttr(resourceName, "vpc_configuration.0.tls_certificate"),
					resource.TestCheckResourceAttrSet(resourceName, "vpc_configuration.0.vpc_id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccHostConfig_vpc(rName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrRegionalARN(ctx, resourceName, names.AttrARN, "codeconnections", regexache.MustCompile("host/.+$")),
					resource.TestCheckResourceAttrPair(resourceName, names.AttrID, resourceName, names.AttrARN),
					resource.TestCheckResourceAttr(resourceName, names.AttrName, rName),
					resource.TestCheckResourceAttr(resourceName, "provider_endpoint", "https://example.com"),
					resource.TestCheckResourceAttr(resourceName, "provider_type", string(types.ProviderTypeGithubEnterpriseServer)),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.0.security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.0.subnet_ids.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "vpc_configuration.0.tls_certificate", "-----BEGIN CERTIFICATE-----\nMIID2jCCAsKgAwIBAgIJAJ58TJVjU7G1MA0GCSqGSIb3DQEBBQUAMFExCzAJBgNV\nBAYTAlVTMREwDwYDVQQIEwhDb2xvcmFkbzEPMA0GA1UEBxMGRGVudmVyMRAwDgYD\nVQQKEwdDaGFydGVyMQwwCgYDVQQLEwNDU0UwHhcNMTcwMTMwMTkyMDA4WhcNMjYx\nMjA5MTkyMDA4WjBRMQswCQYDVQQGEwJVUzERMA8GA1UECBMIQ29sb3JhZG8xDzAN\nBgNVBAcTBkRlbnZlcjEQMA4GA1UEChMHQ2hhcnRlcjEMMAoGA1UECxMDQ1NFMIIB\nIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv6dq6VLIImlAaTrckb5w3X6J\nWP7EGz2ChGAXlkEYto6dPCba0v5+f+8UlMOpeB25XGoai7gdItqNWVFpYsgmndx3\nvTad3ukO1zeElKtw5oHPH2plOaiv/gVJaDa9NTeINj0EtGZs74fCOclAzGFX5vBc\nb08ESWBceRgGjGv3nlij4JzHfqTkCKQz6P6pBivQBfk62rcOkkH5rKoaGltRHROS\nMbkwOhu2hN0KmSYTXRvts0LXnZU4N0l2ms39gmr7UNNNlKYINL2JoTs9dNBc7APD\ndZvlEHd+/FjcLCI8hC3t4g4AbfW0okIBCNG0+oVjqGb2DeONSJKsThahXt89MQID\nAQABo4G0MIGxMB0GA1UdDgQWBBQKq8JxjY1GmeZXJjfOMfW0kBIzPDCBgQYDVR0j\nBHoweIAUCqvCcY2NRpnmVyY3zjH1tJASMzyhVaRTMFExCzAJBgNVBAYTAlVTMREw\nDwYDVQQIEwhDb2xvcmFkbzEPMA0GA1UEBxMGRGVudmVyMRAwDgYDVQQKEwdDaGFy\ndGVyMQwwCgYDVQQLEwNDU0WCCQCefEyVY1OxtTAMBgNVHRMEBTADAQH/MA0GCSqG\nSIb3DQEBBQUAA4IBAQAWifoMk5kbv+yuWXvFwHiB4dWUUmMlUlPU/E300yVTRl58\np6DfOgJs7MMftd1KeWqTO+uW134QlTt7+jwI8Jq0uyKCu/O2kJhVtH/Ryog14tGl\n+wLcuIPLbwJI9CwZX4WMBrq4DnYss+6F47i8NCc+Z3MAiG4vtq9ytBmaod0dj2bI\ng4/Lac0e00dql9RnqENh1+dF0V+QgTJCoPkMqDNAlSB8vOodBW81UAb2z12t+IFi\n3X9J3WtCK2+T5brXL6itzewWJ2ALvX3QpmZx7fMHJ3tE+SjjyivE1BbOlzYHx83t\nTeYnm7pS9un7A/UzTDHbs7hPUezLek+H3xTPAnnq\n-----END CERTIFICATE-----\n"),
					resource.TestCheckResourceAttrSet(resourceName, "vpc_configuration.0.vpc_id"),
				),
			},
		},
	})
}

func TestAccCodeConnectionsHost_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Host
	resourceName := "aws_codeconnections_host.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.CodeConnectionsServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckConnectionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccHostConfig_tags1(rName, acctest.CtKey1, acctest.CtValue1),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccHostConfig_tags2(rName, acctest.CtKey1, acctest.CtValue1Updated, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "2"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey1, acctest.CtValue1Updated),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
			{
				Config: testAccHostConfig_tags1(rName, acctest.CtKey2, acctest.CtValue2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckHostExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsPercent, "1"),
					resource.TestCheckResourceAttr(resourceName, acctest.CtTagsKey2, acctest.CtValue2),
				),
			},
		},
	})
}

func testAccCheckHostExists(ctx context.Context, n string, v *types.Host) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).CodeConnectionsClient(ctx)

		output, err := tfcodeconnections.FindHostByARN(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckHostDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).CodeConnectionsClient(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_codeconnections_host" {
				continue
			}

			_, err := tfcodeconnections.FindHostByARN(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("CodeStar Connections Host %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccHostVPCBaseConfig(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  count             = 2
  availability_zone = data.aws_availability_zones.available.names[count.index]
  cidr_block        = cidrsubnet(aws_vpc.test.cidr_block, 8, count.index + 2)
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_security_group" "test" {
  name   = %[1]q
  vpc_id = aws_vpc.test.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
  ingress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = %[1]q
  }
}
`, rName))
}

func testAccHostConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_codeconnections_host" "test" {
  name              = %[1]q
  provider_endpoint = "https://example.com"
  provider_type     = "GitHubEnterpriseServer"
}
`, rName)
}

func testAccHostConfig_vpcNoCertificate(rName string) string {
	return acctest.ConfigCompose(testAccHostVPCBaseConfig(rName), fmt.Sprintf(`
resource "aws_codeconnections_host" "test" {
  name              = %[1]q
  provider_endpoint = "https://example.com"
  provider_type     = "GitHubEnterpriseServer"
  vpc_configuration {
    security_group_ids = [aws_security_group.test.id]
    subnet_ids         = aws_subnet.test[*].id
    vpc_id             = aws_vpc.test.id
  }
}
`, rName))
}

func testAccHostConfig_vpc(rName string) string {
	return acctest.ConfigCompose(testAccHostVPCBaseConfig(rName), fmt.Sprintf(`
resource "aws_codeconnections_host" "test" {
  name              = %[1]q
  provider_endpoint = "https://example.com"
  provider_type     = "GitHubEnterpriseServer"
  vpc_configuration {
    security_group_ids = [aws_security_group.test.id]
    subnet_ids         = aws_subnet.test[*].id
    tls_certificate    = "-----BEGIN CERTIFICATE-----\nMIID2jCCAsKgAwIBAgIJAJ58TJVjU7G1MA0GCSqGSIb3DQEBBQUAMFExCzAJBgNV\nBAYTAlVTMREwDwYDVQQIEwhDb2xvcmFkbzEPMA0GA1UEBxMGRGVudmVyMRAwDgYD\nVQQKEwdDaGFydGVyMQwwCgYDVQQLEwNDU0UwHhcNMTcwMTMwMTkyMDA4WhcNMjYx\nMjA5MTkyMDA4WjBRMQswCQYDVQQGEwJVUzERMA8GA1UECBMIQ29sb3JhZG8xDzAN\nBgNVBAcTBkRlbnZlcjEQMA4GA1UEChMHQ2hhcnRlcjEMMAoGA1UECxMDQ1NFMIIB\nIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAv6dq6VLIImlAaTrckb5w3X6J\nWP7EGz2ChGAXlkEYto6dPCba0v5+f+8UlMOpeB25XGoai7gdItqNWVFpYsgmndx3\nvTad3ukO1zeElKtw5oHPH2plOaiv/gVJaDa9NTeINj0EtGZs74fCOclAzGFX5vBc\nb08ESWBceRgGjGv3nlij4JzHfqTkCKQz6P6pBivQBfk62rcOkkH5rKoaGltRHROS\nMbkwOhu2hN0KmSYTXRvts0LXnZU4N0l2ms39gmr7UNNNlKYINL2JoTs9dNBc7APD\ndZvlEHd+/FjcLCI8hC3t4g4AbfW0okIBCNG0+oVjqGb2DeONSJKsThahXt89MQID\nAQABo4G0MIGxMB0GA1UdDgQWBBQKq8JxjY1GmeZXJjfOMfW0kBIzPDCBgQYDVR0j\nBHoweIAUCqvCcY2NRpnmVyY3zjH1tJASMzyhVaRTMFExCzAJBgNVBAYTAlVTMREw\nDwYDVQQIEwhDb2xvcmFkbzEPMA0GA1UEBxMGRGVudmVyMRAwDgYDVQQKEwdDaGFy\ndGVyMQwwCgYDVQQLEwNDU0WCCQCefEyVY1OxtTAMBgNVHRMEBTADAQH/MA0GCSqG\nSIb3DQEBBQUAA4IBAQAWifoMk5kbv+yuWXvFwHiB4dWUUmMlUlPU/E300yVTRl58\np6DfOgJs7MMftd1KeWqTO+uW134QlTt7+jwI8Jq0uyKCu/O2kJhVtH/Ryog14tGl\n+wLcuIPLbwJI9CwZX4WMBrq4DnYss+6F47i8NCc+Z3MAiG4vtq9ytBmaod0dj2bI\ng4/Lac0e00dql9RnqENh1+dF0V+QgTJCoPkMqDNAlSB8vOodBW81UAb2z12t+IFi\n3X9J3WtCK2+T5brXL6itzewWJ2ALvX3QpmZx7fMHJ3tE+SjjyivE1BbOlzYHx83t\nTeYnm7pS9un7A/UzTDHbs7hPUezLek+H3xTPAnnq\n-----END CERTIFICATE-----\n"
    vpc_id             = aws_vpc.test.id
  }
}
`, rName))
}

func testAccHostConfig_tags1(rName string, tagKey1 string, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_codeconnections_host" "test" {
  name              = %[1]q
  provider_endpoint = "https://example.com"
  provider_type     = "GitHubEnterpriseServer"

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccHostConfig_tags2(rName string, tagKey1 string, tagValue1 string, tagKey2 string, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_codeconnections_host" "test" {
  name              = %[1]q
  provider_endpoint = "https://example.com"
  provider_type     = "GitHubEnterpriseServer"

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}
