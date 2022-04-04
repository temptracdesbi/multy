package network_interface

import (
	"fmt"
	"github.com/multycloud/multy/api/proto/commonpb"
	"github.com/multycloud/multy/resources/common"
	"github.com/multycloud/multy/validate"
)

const AzureResourceName = "azurerm_network_interface"

type AzureNetworkInterface struct {
	*common.AzResource `hcl:",squash" default:"name=azurerm_network_interface"`
	IpConfigurations   []AzureIpConfiguration `hcl:"ip_configuration,blocks"`
}

func (nic AzureNetworkInterface) GetId(cloud commonpb.CloudProvider) string {
	if cloud == common.AZURE {
		return fmt.Sprintf("${%s.%s.id}", AzureResourceName, nic.ResourceId)
	}
	validate.LogInternalError("cloud %s is not supported for this resource type ", cloud)
	return ""
}

type AzureIpConfiguration struct {
	Name                       string `hcl:"name"`
	PrivateIpAddressAllocation string `hcl:"private_ip_address_allocation"`
	SubnetId                   string `hcl:"subnet_id" hcle:"omitempty"`
	PublicIpAddressId          string `hcl:"public_ip_address_id,expr" hcle:"omitempty"`
	Primary                    bool   `hcl:"primary" hcle:"omitempty"`
}
