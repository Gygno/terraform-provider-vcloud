---
layout: "vcd"
page_title: "Viettel IDC Cloud: vcloud_tm_edge_cluster_qos"
sidebar_current: "docs-vcd-data-source-tm-edge-cluster-qos"
description: |-
  Provides a Viettel IDC Tenant Manager Edge Cluster QoS data source.
---

# vcloud\_tm\_edge\_cluster\_qos

Provides a Viettel IDC Tenant Manager Edge Cluster QoS data source.

## Example Usage

```hcl
data "vcloud_tm_region" "demo" {
  name = "region-one"
}

data "vcloud_tm_edge_cluster" "demo" {
  name             = "my-edge-cluster"
  region_id        = data.vcloud_tm_region.demo.id
  sync_before_read = true
}

data "vcloud_tm_edge_cluster_qos" "demo" {
  edge_cluster_id = data.vcloud_tm_edge_cluster.demo.id
}
```

## Argument Reference

The following arguments are supported:

* `edge_cluster_id` - (Required) An ID of Edge Cluster. Can be looked up using
  [vcloud_tm_edge_cluster](/providers/viettelidc-provider/vcloud/latest/docs/data-sources/tm_edge_cluster) data source

## Attribute Reference

All the arguments and attributes defined in
[`vcloud_tm_edge_cluster_qos`](/providers/viettelidc-provider/vcloud/latest/docs/resources/tm_edge_cluster_qos) resource are available.