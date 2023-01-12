package v3

import (
	"strconv"
	"strings"

	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/core/v3/internal/stringutil"
)

// APIKeyFields returns a set of fields that represent that resource.
func APIKeyFields(r Resource) map[string]string {
	resource := r.(*corev2.APIKey)
	fields := map[string]string{
		"api_key.name":     resource.ObjectMeta.Name,
		"api_key.username": resource.Username,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "api_key.labels.")
	return fields
}

// AssetFields returns a set of fields that represent that resource
func AssetFields(r Resource) map[string]string {
	resource := r.(*corev2.Asset)
	fields := map[string]string{
		"asset.name":      resource.ObjectMeta.Name,
		"asset.namespace": resource.ObjectMeta.Namespace,
		"asset.filters":   strings.Join(resource.Filters, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "asset.labels.")
	return fields
}

// CheckConfigFields returns a set of fields that represent that resource
func CheckConfigFields(r Resource) map[string]string {
	resource := r.(*corev2.CheckConfig)
	fields := map[string]string{
		"check.name":           resource.ObjectMeta.Name,
		"check.namespace":      resource.ObjectMeta.Namespace,
		"check.handlers":       strings.Join(resource.Handlers, ","),
		"check.publish":        strconv.FormatBool(resource.Publish),
		"check.round_robin":    strconv.FormatBool(resource.RoundRobin),
		"check.runtime_assets": strings.Join(resource.RuntimeAssets, ","),
		"check.subscriptions":  strings.Join(resource.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "check.labels.")

	pipelineIDs := []string{}
	for _, pipeline := range resource.Pipelines {
		pipelineIDs = append(pipelineIDs, pipeline.ResourceID())
	}
	fields["check.pipelines"] = strings.Join(pipelineIDs, ",")

	return fields
}

// EntityFields returns a set of fields that represent that resource
func EntityFields(r Resource) map[string]string {
	resource := r.(*corev2.Entity)
	fields := map[string]string{
		"entity.name":          resource.ObjectMeta.Name,
		"entity.namespace":     resource.ObjectMeta.Namespace,
		"entity.deregister":    strconv.FormatBool(resource.Deregister),
		"entity.entity_class":  resource.EntityClass,
		"entity.subscriptions": strings.Join(resource.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "entity.labels.")
	return fields
}

func isSilenced(e *corev2.Event) string {
	if len(e.Check.Silenced) > 0 {
		return "true"
	}
	return "false"
}

// EventFields returns a set of fields that represent that resource
func EventFields(r Resource) map[string]string {
	resource := r.(*corev2.Event)
	fields := map[string]string{
		"event.name":                 resource.ObjectMeta.Name,
		"event.namespace":            resource.ObjectMeta.Namespace,
		"event.is_silenced":          isSilenced(resource),
		"event.check.is_silenced":    isSilenced(resource),
		"event.check.name":           resource.Check.Name,
		"event.check.handlers":       strings.Join(resource.Check.Handlers, ","),
		"event.check.publish":        strconv.FormatBool(resource.Check.Publish),
		"event.check.round_robin":    strconv.FormatBool(resource.Check.RoundRobin),
		"event.check.runtime_assets": strings.Join(resource.Check.RuntimeAssets, ","),
		"event.check.state":          resource.Check.State,
		"event.check.status":         strconv.Itoa(int(resource.Check.Status)),
		"event.check.subscriptions":  strings.Join(resource.Check.Subscriptions, ","),
		"event.entity.deregister":    strconv.FormatBool(resource.Entity.Deregister),
		"event.entity.name":          resource.Entity.ObjectMeta.Name,
		"event.entity.entity_class":  resource.Entity.EntityClass,
		"event.entity.subscriptions": strings.Join(resource.Entity.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "event.labels.")
	stringutil.MergeMapWithPrefix(fields, resource.Entity.ObjectMeta.Labels, "event.labels.")
	stringutil.MergeMapWithPrefix(fields, resource.Check.ObjectMeta.Labels, "event.labels.")
	return fields
}

// EventFilterFields returns a set of fields that represent that resource
func EventFilterFields(r Resource) map[string]string {
	resource := r.(*EventFilter)
	fields := map[string]string{
		"filter.name":           resource.ObjectMeta.Name,
		"filter.namespace":      resource.ObjectMeta.Namespace,
		"filter.action":         resource.Action,
		"filter.runtime_assets": strings.Join(resource.RuntimeAssets, ","),
	}
	utilstrings.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "filter.labels.")
	return fields
}
