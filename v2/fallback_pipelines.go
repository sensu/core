package v2

import (
	"errors"
	"net/url"
	"path"

	stringsutil "github.com/sensu/core/v2/internal/stringutil"
)

const (
	// FallbackPipelinesResource is the name of this resource type
	FallbackPipelinesResource = "fallback_pipelines"
)

// GetObjectMeta returns the object metadata for the resource.
func (f *FallbackPipelines) GetObjectMeta() ObjectMeta {
	return f.ObjectMeta
}

// SetObjectMeta sets the object metadata for the resource.
func (f *FallbackPipelines) SetObjectMeta(meta ObjectMeta) {
	f.ObjectMeta = meta
}

// SetNamespace sets the namespace of the resource.
func (f *FallbackPipelines) SetNamespace(namespace string) {
	f.Namespace = namespace
}

// StorePrefix returns the path prefix to this resource in the store.
func (f *FallbackPipelines) StorePrefix() string {
	return FallbackPipelinesResource
}

// RBACName describes the name of the resource for RBAC purposes.
func (f *FallbackPipelines) RBACName() string {
	return FallbackPipelinesResource
}

// URIPath gives the path component of a fallback URI.
func (f *FallbackPipelines) URIPath() string {
	if f.Namespace == "" {
		return path.Join(URLPrefix, FallbackPipelinesResource, url.PathEscape(f.Name))
	}
	return path.Join(URLPrefix, "namespaces", url.PathEscape(f.Namespace), FallbackPipelinesResource, url.PathEscape(f.Name))
}

// Validate checks if the fields in the resource are valid.
func (f *FallbackPipelines) Validate() error {
	if err := ValidateName(f.ObjectMeta.Name); err != nil {
		return errors.New("name " + err.Error())
	}
	if f.ObjectMeta.Namespace == "" {
		return errors.New("namespace must be set")
	}
	return nil
}

// FallbackPipelinesFields returns a set of fields that represent that resource.
func FallbackPipelinesFields(r Resource) map[string]string {
	resource := r.(*FallbackPipelines)
	fields := map[string]string{
		"fallbackPipelines.name":      resource.ObjectMeta.Name,
		"fallbackPipelines.namespace": resource.ObjectMeta.Namespace,
	}
	stringsutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "fallbackPipelines.labels.")
	return fields
}

// Fields returns a set of fields that represent that resource.
func (f *FallbackPipelines) Fields() map[string]string {
	return FallbackPipelinesFields(f)
}

// FixtureFallbackPipelines returns a testing fixture for a FallbackPipelines object.
func FixtureFallbackPipelines(name, namespace string) *FallbackPipelines {
	return &FallbackPipelines{
		ObjectMeta: NewObjectMeta(name, namespace),
		Pipelines:  []*ResourceReference{},
	}
}

// FixtureFallbackPipelineReference returns a testing fixture for a ResourceReference
// object referencing a corev2.FallbackPipeline.
func FixtureFallbackPipelineReference(name string) *ResourceReference {
	return &ResourceReference{
		APIVersion: "core/v2",
		Type:       "FallbackPipelines",
		Name:       name,
	}
}
