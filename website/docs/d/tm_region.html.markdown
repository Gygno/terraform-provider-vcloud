---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_region"
sidebar_current: "docs-vcd-datasource-tm-region"
description: |-
  Provides a data source to read Regions in Viettel IDC Tenant Manager.
---

# vcloud\_tm\_region

Provides a data source to read Regions in Viettel IDC Tenant Manager.

## Example Usage

```hcl
data "vcloud_tm_region" "one" {
  name = "region-one"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A name of existing Region

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_region`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_region) resource are available.
