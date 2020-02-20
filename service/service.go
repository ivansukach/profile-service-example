package service

import (
	"fmt"
	"github.com/ivansukach/profile-service-example/repositories"
	log "github.com/sirupsen/logrus"
)

type ProfileService struct {
	r *repositories.Repository
}

func (ps ProfileService) CreateUser() error {
	return nil
}

func Init() {
	prf := ProfileService{}
	prf.CreateUser()
	prs := new(ProfileService)
	prs2 := new(ProfileService)
	var prs3 ProfileService
	ms := new(ProfileService)
	ms.r = prs.r
	prs2 = prs
	prs3 = *prs
	fmt.Println(prs2)
	fmt.Println(prs3)
}
