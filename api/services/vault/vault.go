package vault

import (
	"github.com/multycloud/multy/api/proto/common"
	"github.com/multycloud/multy/api/proto/resources"
	"github.com/multycloud/multy/api/services"
	"github.com/multycloud/multy/api/util"
	"github.com/multycloud/multy/db"
)

type VaultService struct {
	Service services.Service[*resources.CloudSpecificVaultArgs, *resources.VaultResource]
}

func (s VaultService) Convert(resourceId string, args []*resources.CloudSpecificVaultArgs) *resources.VaultResource {
	var result []*resources.CloudSpecificVaultResource
	for _, r := range args {
		result = append(result, &resources.CloudSpecificVaultResource{
			CommonParameters: util.ConvertCommonParams(r.CommonParameters),
			Name:             r.Name,
		})
	}

	return &resources.VaultResource{
		CommonParameters: &common.CommonResourceParameters{ResourceId: resourceId},
		Resources:        result,
	}
}

func (s VaultService) NewArg() *resources.CloudSpecificVaultArgs {
	return &resources.CloudSpecificVaultArgs{}
}

func (s VaultService) Nil() *resources.VaultResource {
	return nil
}

func NewVaultService(database *db.Database) VaultService {
	ni := VaultService{
		Service: services.Service[*resources.CloudSpecificVaultArgs, *resources.VaultResource]{
			Db:         database,
			Converters: nil,
		},
	}
	ni.Service.Converters = &ni
	return ni
}