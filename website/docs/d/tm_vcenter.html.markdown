---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_vcenter"
sidebar_current: "docs-vcd-data-source-tm-vcenter"
description: |-
  Provides a data source for vCenter server.
---

# vcloud\_tm\_vcenter

Provides a data source for vCenter server.

Supported in provider *v3.0+*


## Example Usage

```hcl
data "vcloud_tm_vcenter" "vc" {
  name = "vcenter-one"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) vCenter name

## Attribute reference

All attributes defined in
[`vcloud_tm_vcenter`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_vcenter#attribute-reference) are
supported.
