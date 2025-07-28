---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_org_vdc"
sidebar_current: "docs-vcd-datasource-tm-org-vdc"
description: |-
  Provides a data source to manage Viettel IDC Tenant Manager Organization VDC.
---

# vcloud\_tm\_org\_vdc

Provides a data source to manage Viettel IDC Tenant Manager Organization VDC.

## Example Usage

```hcl
data "vcloud_tm_org_vdc" "test" {
  name = "my-tm-org-vdc"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A name for the existing Org VDC
* `org_id` - (Required) An ID for the parent Org

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_org_vdc`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_org_vdc) resource are available.
