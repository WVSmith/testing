package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/mediaFORGE/partner_common/common"
	"github.com/mediaFORGE/partner_common/providers/mediamath"
)

func main() {
	c := mediamath.NewConfig()
	c.APIKey = "4423bc5079b405e62ca4212f8ebf6d48"
	c.Username = "api@mediaforge.com"
	c.Password = "mediaforge10"
	a := common.NewAgent(false)
	mm := mediamath.NewMediaMath(c, a)

	concept := &mediamath.Concept{
		AdvertiserID: 156383,
		Name:         "test03",
	}

	concept, err := mm.CreateConcept(concept)
	if nil != err {
		fmt.Println(err)
	}
	spew.Dump(concept)
}
