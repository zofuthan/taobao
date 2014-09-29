package response

import (
	"github.com/ohohco/taobao/api"
	. "github.com/ohohco/taobao/api/domain"
)

type ItemGetResponse struct {
	api.TaobaoResponse `json:"error_response"`
	Response           struct {
		/* 获取的商品 具体字段根据权限和设定的fields决定 */
		Item *Item `json:"item"`
	} `json:"item_get_response"`
}
