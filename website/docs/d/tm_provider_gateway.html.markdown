---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_provider_gateway"
sidebar_current: "docs-vcd-data-source-tm-provider-gateway"
description: |-
  Provides a Viettel IDC Tenant Manager Provider Gateway data source.
---

# vcloud\_tm\_provider\_gateway

Provides a Viettel IDC Tenant Manager Provider Gateway data source.

## Example Usage

```hcl
data "vcloud_tm_region" "demo" {
  name = "region-one"
}

data "vcloud_tm_provider_gateway" "demo" {
  name      = "Demo Provider Gateway"
  region_id = data.vcloud_tm_region.demo.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of Provider Gateway
* `region_id` - (Required) An ID of Region. Can be looked up using
  [vcloud_tm_region](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_region) data source


## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_provider_gateway`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_provider_gateway)
resource are available.