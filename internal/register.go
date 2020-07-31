package internal

import (
	"github.com/Floor-Gang/authserver/pkg"
)

func (a *AuthServer) Register(args *pkg.Feature, res *pkg.RegisterResponse) error {
	name := args.Name
	a.Features[name] = args

	*res = pkg.RegisterResponse{
		Serving: a.config.Guild,
	}

	return nil
}
