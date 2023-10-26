package bootstrap

import (
	"github.com/neerubhandari/ecom-backend/infrastructure"
	"github.com/neerubhandari/ecom-backend/routes"
	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	infrastructure.Module,
	routes.Module,
)
