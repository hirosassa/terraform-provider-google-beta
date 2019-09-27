// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccComputeRegionHealthCheck_regionHealthCheckTcpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckTcpExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckTcpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "tcp-region-health-check" {
  provider = "google-beta"
  name = "tcp-region-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  tcp_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckTcpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckTcpFullExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckTcpFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "tcp-region-health-check" {
  provider = "google-beta"
  name = "tcp-region-health-check%{random_suffix}"
  description = "Health check via tcp"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  tcp_health_check {
    port_name = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request = "ARE YOU HEALTHY?"
    proxy_header = "NONE"
    response = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckSslExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckSslExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckSslExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "ssl-region-health-check" {
  provider = "google-beta"
  name = "ssl-region-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  ssl_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckSslFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckSslFullExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckSslFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "ssl-region-health-check" {
  provider = "google-beta"
  name = "ssl-region-health-check%{random_suffix}"
  description = "Health check via ssl"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  ssl_health_check {
    port_name = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    request = "ARE YOU HEALTHY?"
    proxy_header = "NONE"
    response = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttpExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttpExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttpExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "http-region-health-check" {
  provider = "google-beta"
  name = "http-region-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http_health_check {
    port = "80"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttpFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttpFullExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttpFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "http-region-health-check" {
  provider = "google-beta"
  name = "http-region-health-check%{random_suffix}"
  description = "Health check via http"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http_health_check {
    port_name = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host = "1.2.3.4"
    request_path = "/mypath"
    proxy_header = "NONE"
    response = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttpsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttpsExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttpsExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "https-region-health-check" {
  provider = "google-beta"
  name = "https-region-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  https_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttpsFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttpsFullExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttpsFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "https-region-health-check" {
  provider = "google-beta"
  name = "https-region-health-check%{random_suffix}"
  description = "Health check via https"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  https_health_check {
    port_name = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host = "1.2.3.4"
    request_path = "/mypath"
    proxy_header = "NONE"
    response = "I AM HEALTHY"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttp2Example(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttp2Example(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttp2Example(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "http2-region-health-check" {
  provider = "google-beta"
  name = "http2-region-health-check%{random_suffix}"

  timeout_sec        = 1
  check_interval_sec = 1

  http2_health_check {
    port = "443"
  }
}
`, context)
}

func TestAccComputeRegionHealthCheck_regionHealthCheckHttp2FullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(10),
	}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckComputeRegionHealthCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionHealthCheck_regionHealthCheckHttp2FullExample(context),
			},
		},
	})
}

func testAccComputeRegionHealthCheck_regionHealthCheckHttp2FullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_compute_region_health_check" "http2-region-health-check" {
  provider = "google-beta"
  name = "http2-region-health-check%{random_suffix}"
  description = "Health check via http2"

  timeout_sec         = 1
  check_interval_sec  = 1
  healthy_threshold   = 4
  unhealthy_threshold = 5

  http2_health_check {
    port_name = "health-check-port"
    port_specification = "USE_NAMED_PORT"
    host = "1.2.3.4"
    request_path = "/mypath"
    proxy_header = "NONE"
    response = "I AM HEALTHY"
  }
}
`, context)
}

func testAccCheckComputeRegionHealthCheckDestroy(s *terraform.State) error {
	for name, rs := range s.RootModule().Resources {
		if rs.Type != "google_compute_region_health_check" {
			continue
		}
		if strings.HasPrefix(name, "data.") {
			continue
		}

		config := testAccProvider.Meta().(*Config)

		url, err := replaceVarsForTest(config, rs, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/healthChecks/{{name}}")
		if err != nil {
			return err
		}

		_, err = sendRequest(config, "GET", "", url, nil)
		if err == nil {
			return fmt.Errorf("ComputeRegionHealthCheck still exists at %s", url)
		}
	}

	return nil
}