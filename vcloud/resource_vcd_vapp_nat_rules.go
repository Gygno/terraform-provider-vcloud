package vcloud

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vmware/go-vcloud-director/v3/govcd"
	"github.com/vmware/go-vcloud-director/v3/types/v56"
)

const (
	allowTrafficInPolicy  = "allowTrafficIn"
	allowTrafficPolicy    = "allowTraffic"
	ipTranslationNatType  = "ipTranslation"
	portForwardingNatType = "portForwarding"
)

func resourceVcdVappNetworkNatRules() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVappNetworkNatRulesCreate,
		DeleteContext: resourceVAppNetworkNatRulesDelete,
		ReadContext:   resourceVappNetworkNatRulesRead,
		UpdateContext: resourceVappNetworkNatRulesUpdate,
		Importer: &schema.ResourceImporter{
			StateContext: vappNetworkNatRulesImport,
		},

		Schema: map[string]*schema.Schema{
			"org": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The name of organization to use, optional if defined at provider " +
					"level. Useful when connected as sysadmin working across different organizations",
			},
			"vdc": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The name of VDC to use, optional if defined at provider level",
			},
			"vapp_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "vApp identifier",
			},
			"network_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "vApp network identifier",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enable or disable NAT service. Default is `true`.",
			},
			"nat_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{portForwardingNatType, ipTranslationNatType}, false),
				Description:  "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding).",
			},
			"enable_ip_masquerade": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic.",
			},
			"rule": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the rule. Can be used to track syslog messages.",
						},
						"mapping_mode": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"automatic", "manual"}, false),
							Description:  "Mapping mode. One of: `automatic`, `manual`",
						},
						"vm_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "VM to which this rule applies.",
						},
						"vm_nic_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "VM NIC ID to which this rule applies.",
						},
						"external_ip": {
							Type:         schema.TypeString,
							Optional:     true,
							Computed:     true,
							ValidateFunc: validation.IsIPAddress,
							Description:  "External IP address to forward to or External IP address to map to VM",
						},
						"external_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "External port to forward.",
						},
						"forward_to_port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Internal port to forward.",
						},
						"protocol": {
							Type:         schema.TypeString,
							Optional:     true,
							ValidateFunc: validation.StringInSlice([]string{"TCP", "UDP", "TCP_UDP"}, false),
							Description:  "Protocol to forward. One of: `TCP` (forward TCP packets), `UDP` (forward UDP packets), `TCP_UDP` (forward TCP and UDP packets).",
						},
					},
				},
			},
		},
	}
}

func resourceVappNetworkNatRulesCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceVappNetworkNatRulesUpdate(ctx, d, meta)
}

func resourceVappNetworkNatRulesUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vcdClient := meta.(*VCDClient)

	_, vdc, err := vcdClient.GetOrgAndVdcFromResource(d)
	if err != nil {
		return diag.Errorf(errorRetrievingOrgAndVdc, err)
	}

	vappId := d.Get("vapp_id").(string)
	vapp, err := vdc.GetVAppById(vappId, false)
	if err != nil {
		return diag.Errorf("error finding vApp. %s", err)
	}
	vcdClient.lockParentVappWithName(d, vapp.VApp.Name)
	defer vcdClient.unLockParentVappWithName(d, vapp.VApp.Name)

	networkId := d.Get("network_id").(string)
	natType := d.Get("nat_type").(string)
	netRules, err := expandVappNetworkNatRules(d, vapp, natType)
	if err != nil {
		return diag.Errorf("error expanding NAT rules: %s", err)
	}
	policy := allowTrafficInPolicy
	if d.Get("enable_ip_masquerade").(bool) && natType == portForwardingNatType {
		policy = allowTrafficPolicy
	}
	vappNetwork, err := vapp.UpdateNetworkNatRules(networkId, netRules, d.Get("enabled").(bool),
		natType, policy)
	if err != nil {
		log.Printf("[INFO] Error setting NAT rules: %s", err)
		return diag.Errorf("error setting NAT rules: %s", err)
	}

	if vappNetwork.Configuration.Features.FirewallService != nil &&
		!vappNetwork.Configuration.Features.FirewallService.IsEnabled &&
		d.Get("enabled").(bool) {
		logForScreen("vcd_vapp_nat_rules", "WARNING: for NAT rules to work, firewall has to be enabled. It can be enabled using vcd_vapp_firewall_rules")
	}

	d.SetId(vappNetwork.ID)

	return resourceVappNetworkNatRulesRead(ctx, d, meta)
}

func resourceVAppNetworkNatRulesDelete(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vcdClient := meta.(*VCDClient)

	_, vdc, err := vcdClient.GetOrgAndVdcFromResource(d)
	if err != nil {
		return diag.Errorf(errorRetrievingOrgAndVdc, err)
	}

	vappId := d.Get("vapp_id").(string)
	vapp, err := vdc.GetVAppById(vappId, false)
	if err != nil {
		return diag.Errorf("error finding vApp. %s", err)
	}

	vcdClient.lockParentVappWithName(d, vapp.VApp.Name)
	defer vcdClient.unLockParentVappWithName(d, vapp.VApp.Name)

	err = vapp.RemoveAllNetworkNatRules(d.Get("network_id").(string))
	if err != nil {
		log.Printf("[INFO] Error deleting NAT rules: %s", err)
		return diag.Errorf("error deleting NAT rules: %s", err)
	}

	return nil
}

func resourceVappNetworkNatRulesRead(_ context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vcdClient := meta.(*VCDClient)

	_, vdc, err := vcdClient.GetOrgAndVdcFromResource(d)
	if err != nil {
		return diag.Errorf(errorRetrievingOrgAndVdc, err)
	}

	vappId := d.Get("vapp_id").(string)
	vapp, err := vdc.GetVAppById(vappId, false)
	if err != nil {
		return diag.Errorf("error finding vApp. %s", err)
	}

	vappNetwork, err := vapp.GetVappNetworkById(d.Get("network_id").(string), false)
	if err != nil {
		if govcd.ContainsNotFound(err) {
			log.Printf("vApp network not found. Removing from state file: %s", err)
			d.SetId("")
			return nil
		}
		return diag.Errorf("error finding vApp network. %s", err)
	}

	var rules []map[string]interface{}
	if vappNetwork.Configuration.Features == nil || vappNetwork.Configuration.Features.NatService == nil {
		log.Print("no Nat rules found.")
		dSet(d, "rule", rules)
	}

	for _, rule := range vappNetwork.Configuration.Features.NatService.NatRule {
		singleRule := make(map[string]interface{})
		singleRule["id"] = rule.ID
		switch vappNetwork.Configuration.Features.NatService.NatType {
		case portForwardingNatType:
			singleRule["external_port"] = rule.VMRule.ExternalPort
			singleRule["vm_nic_id"] = rule.VMRule.VMNicID
			singleRule["forward_to_port"] = rule.VMRule.InternalPort
			singleRule["protocol"] = rule.VMRule.Protocol
			singleRule["vm_id"] = getVmIdFromVmVappLocalId(vapp, rule.VMRule.VAppScopedVMID)
		case ipTranslationNatType:
			singleRule["vm_nic_id"] = rule.OneToOneVMRule.VMNicID
			singleRule["external_ip"] = rule.OneToOneVMRule.ExternalIPAddress
			singleRule["mapping_mode"] = rule.OneToOneVMRule.MappingMode
			singleRule["vm_id"] = getVmIdFromVmVappLocalId(vapp, rule.OneToOneVMRule.VAppScopedVMID)
		}
		rules = append(rules, singleRule)
	}
	dSet(d, "enabled", vappNetwork.Configuration.Features.NatService.IsEnabled)
	if vappNetwork.Configuration.Features.NatService.NatType == portForwardingNatType &&
		vappNetwork.Configuration.Features.NatService.Policy == allowTrafficPolicy {
		dSet(d, "enable_ip_masquerade", true)
	} else if vappNetwork.Configuration.Features.NatService.NatType == portForwardingNatType &&
		vappNetwork.Configuration.Features.NatService.Policy == allowTrafficInPolicy {
		dSet(d, "enable_ip_masquerade", false)
	}
	dSet(d, "nat_type", vappNetwork.Configuration.Features.NatService.NatType)
	err = d.Set("rule", rules)
	if err != nil {
		return diag.FromErr(err)
	}

	if vappNetwork.Configuration.Features.FirewallService != nil &&
		!vappNetwork.Configuration.Features.FirewallService.IsEnabled &&
		d.Get("enabled").(bool) {
		logForScreen("vcd_vapp_firewall_rules", "WARNING: for NAT rules to work, firewall has to be enabled. It can be enabled using vcd_vapp_firewall_rules\n")
	}

	return nil
}

func expandVappNetworkNatRules(d *schema.ResourceData, vapp *govcd.VApp, natType string) ([]*types.NatRule, error) {

	var natRules []*types.NatRule
	for _, singleRule := range d.Get("rule").([]interface{}) {
		configuredRule := singleRule.(map[string]interface{})
		if natType == portForwardingNatType {
			rule := &types.NatRule{
				VMRule: &types.NatVMRule{
					ExternalPort: configuredRule["external_port"].(int),
					VMNicID:      configuredRule["vm_nic_id"].(int),
					InternalPort: configuredRule["forward_to_port"].(int),
					Protocol:     configuredRule["protocol"].(string),
				},
			}
			vm, err := vapp.GetVMById(configuredRule["vm_id"].(string), false)
			if err != nil {
				return nil, fmt.Errorf("error fetchining VM: %s", err)
			}
			rule.VMRule.VAppScopedVMID = vm.VM.VAppScopedLocalID
			natRules = append(natRules, rule)
		} else if natType == ipTranslationNatType {
			rule := &types.NatRule{
				OneToOneVMRule: &types.NatOneToOneVMRule{
					MappingMode: configuredRule["mapping_mode"].(string),
					VMNicID:     configuredRule["vm_nic_id"].(int),
				},
			}
			externalIp := configuredRule["external_ip"].(string)
			if externalIp != "" {
				rule.OneToOneVMRule.ExternalIPAddress = &externalIp
			}

			vm, err := vapp.GetVMById(configuredRule["vm_id"].(string), false)
			if err != nil {
				return nil, fmt.Errorf("error fetchining VM: %s", err)
			}
			rule.OneToOneVMRule.VAppScopedVMID = vm.VM.VAppScopedLocalID
			natRules = append(natRules, rule)
		}
	}

	return natRules, nil
}

// vappNetworkNatRulesImport is responsible for importing the resource.
// The following steps happen as part of import
// 1. The user supplies `terraform import _resource_name_ _the_id_string_` command
// 2. `_the_id_string_` contains a dot formatted path to resource as in the example below
// 3. The functions splits the dot-formatted path and tries to lookup the object
// 4. If the lookup succeeds it set's the ID field for `_resource_name_` resource in state file
// (the resource must be already defined in .tf config otherwise `terraform import` will complain)
// 5. `terraform refresh` is being implicitly launched. The Read method looks up all other fields
// based on the known ID of object.
//
// Example resource name (_resource_name_): vcd_vapp_nat_rules.my_existing_nat_rules
// Example import path (_the_id_string_): org.my_existing_vdc.vapp_name.network_name or org.my_existing_vdc.vapp_id.network_id
// Note: the separator can be changed using Provider.import_separator or variable VCD_IMPORT_SEPARATOR
func vappNetworkNatRulesImport(_ context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return vappNetworkRuleImport(d, meta, "vcd_vapp_nat_rules")
}
