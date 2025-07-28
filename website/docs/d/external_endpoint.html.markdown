---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_external_endpoint"
sidebar_current: "docs-vcd-data-source-external-endpoint"
description: |-
  Provides a data source to read External Endpoints. An External Endpoint holds information for the
  HTTPS endpoint which requests will be proxied to when using an API Filter.
---

# vcloud\_external\_endpoint

Supported in provider *v3.14+*.

Provides a data source to read External Endpoints. An External Endpoint holds information for the
HTTPS endpoint which requests will be proxied to when using a [`vcloud_api_filter`](/providers/viettelidc-provider/vcloud/latest/docs/resources/api_filter).

~> Only `System Administrator` can use this data source.

## Example Usage

```hcl
data "vcloud_external_endpoint" "external_endpoint1" {
  vendor  = "vmware"
  name    = "my-endpoint"
  version = "1.0.0"
}
```

## Argument Reference

* `vendor` - (Required) The vendor name of the External Endpoint
* `name` - (Required) The name of the External Endpoint
* `version` - (Required) The version of the External Endpoint

## Attribute Reference

All the remaining arguments from [the `vcloud_external_endpoint` resource](/providers/viettelidc-provider/vcloud/latest/docs/resources/external_endpoint)
are available as read-only.
