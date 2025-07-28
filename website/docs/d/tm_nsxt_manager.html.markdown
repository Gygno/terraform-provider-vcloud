---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_nsxt_manager"
sidebar_current: "docs-vcd-data-source-tm-nsxt-manager"
description: |-
  Provides a data source for available Tenant Manager NSX-T manager.
---

# vcloud\_tm\_nsxt\_manager

Provides a data source for available Tenant Manager NSX-T manager.

## Example Usage 

```hcl
data "vcloud_tm_nsxt_manager" "main" {
  name = "nsxt-manager-one"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) NSX-T manager name

## Attribute reference

* `id` - ID of the manager
* `href` - Full URL of the manager

All attributes defined in
[`vcloud_tm_nsxt_manager`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_nsxt_manager#attribute-reference)
are supported.
