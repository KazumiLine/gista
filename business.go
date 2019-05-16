package gista

import (
	"github.com/aliforever/gista/constants"
	"github.com/aliforever/gista/responses"
)

type business struct {
	ig *Instagram
}

func newBusiness(i *Instagram) *business {
	return &business{ig: i}
}

func (b *business) GetBlockedMedia() (res *responses.BlockedMedia, err error) {
	res = &responses.BlockedMedia{}
	err = b.ig.client.Request(constants.BlockedMedia).GetResponse(res)
	return
}
