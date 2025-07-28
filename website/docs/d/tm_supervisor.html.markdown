---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_supervisor"
sidebar_current: "docs-vcd-datasource-tm-supervisor"
description: |-
  Provides a data source to read Supervisors in Viettel IDC Tenant Manager.
---

# vcloud\_tm\_supervisor

Provides a data source to read Supervisors in Viettel IDC Tenant Manager.

## Example Usage

```hcl
data "vcloud_vcenter" "one" {
  name = "vcenter-one"
}

data "vcloud_tm_supervisor" "one" {
  name       = "my-supervisor-name"
  vcenter_id = data.vcloud_vcenter.one.id

  depends_on = [vcloud_vcenter.one]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of Supervisor
* `vcenter_id` - (Required) vCenter server ID that contains this Supervisor

## Attribute Reference

* `region_id` - Region ID that consumes this Supervisor
