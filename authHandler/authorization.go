package authHandler

import (
	"errors"
	"fmt"
)

func AuthorizationMiddleWare(req Request) (bool, error) {

	for _, v := range Configuration.Users {
		if v.ID == req.Id {
			fmt.Println(v.Services, req.Service)
			for _, n := range v.Services {
				fmt.Println(n.Name, n.Name == req.Service)
				if req.Service == n.Name {
					return true, nil
				}
			}
		}

	}
	return false, errors.New("not authorized for access")

}
