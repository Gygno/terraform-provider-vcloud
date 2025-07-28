---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_org"
sidebar_current: "docs-vcd-datasource-tm-org"
description: |-
  Provides a data source to read Viettel IDC Tenant Manager Organization.
---

# vcloud\_tm\_org

Provides a data source to read Viettel IDC Tenant Manager Organization.

## Example Usage

```hcl
data "vcloud_tm_org" "existing" {
  name = "my-org-name"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of organization.

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_org`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_org) resource are available.