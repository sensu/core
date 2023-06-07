package v3

// automatically generated file, do not edit!

import (
	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/sensu-api-tools"
)

func init() {
	for alias, v := range typeMap {
		if _, ok := v.(Resource); ok {
			apitools.RegisterType(
				"core/v3",
				v,
				apitools.WithAlias(alias),
				apitools.WithResolveHook(resolveResource),
			)
		}
	}
}

// typeMap is used to dynamically look up data types from strings.
var typeMap = map[string]interface{}{
	"cluster_config":    &ClusterConfig{},
	"entity_config":     &EntityConfig{},
	"entity_state":      &EntityState{},
	"namespace":         &Namespace{},
	"resource_template": &ResourceTemplate{},
	"symmetric_key":     &SymmetricKey{},
}

func resolveResource(v interface{}) {
	resource, ok := v.(Resource)
	if !ok {
		return
	}
	resource.SetMetadata(&corev2.ObjectMeta{
		Labels:      make(map[string]string),
		Annotations: make(map[string]string),
	})
}
