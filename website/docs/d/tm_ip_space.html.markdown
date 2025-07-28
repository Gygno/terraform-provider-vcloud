---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_ip_space"
sidebar_current: "docs-vcd-data-source-tm-ip-space"
description: |-
  Provides a Viettel IDC Tenant Manager IP Space data source.
---

# vcloud\_tm\_ip\_space

Provides a Viettel IDC Tenant Manager IP Space data source.

## Example Usage

```hcl
data "vcloud_tm_region" "demo" {
  name = "demo-region"
}

data "vcloud_tm_ip_space" "demo" {
  name      = "demo-ip-space"
  region_id = data.vcloud_tm_region.region.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of IP Space
* `region_id` - (Required) The Region ID that has this IP Space definition. Can be looked up using
  [`vcloud_tm_region`](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_region)

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_ip_space`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_ip_space) resource are available.
