package domain

import (
	"github.com/test-jagaad/internal/entity"
	"github.com/test-jagaad/internal/util"
)

type (
	MockyDom interface {
		GetMocky1() (res []entity.User, err error)
		GetMocky2() (res []entity.User, err error)
	}

	mockyDom struct{}
)

func NewMockyDom() MockyDom {
	return &mockyDom{}
}

func (d *mockyDom) GetMocky1() (res []entity.User, err error) {
	err = util.FetchJSON(entity.MockyURL1, &res)
	return res, err
}

func (d *mockyDom) GetMocky2() (res []entity.User, err error) {
	err = util.FetchJSON(entity.MockyURL2, &res)
	return res, err
}
