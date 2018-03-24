package restful

import "jenlix.com/apps/identity-provider/common"

type RestResource interface {
	Add(link common.Link)
}
