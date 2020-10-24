package plugin

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/plugin/modelgen"
)

func init() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}
	// Attaching the mutation function onto modelgen plugin
	p := modelgen.Plugin{
		MutateHook: mutateHook,
	}
	err = api.Generate(cfg,
		api.NoPlugins(),
		api.AddPlugin(&p),
	)
}

func mutateHook(b *modelgen.ModelBuild) *modelgen.ModelBuild {

	gorm(b)

	return b
}
