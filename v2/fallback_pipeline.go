package v2

import (
	"errors"
	"net/url"
	"path"

	stringsutil "github.com/sensu/core/v2/internal/stringutil"
)

const (
	// PipelinesResource is the name of this resource type
	FallbackPipelineResource = "fallbackPipeline"
)

// GetObjectMeta returns the object metadata for the resource.
func (f *FallbackPipeline) GetObjectMeta() ObjectMeta {
	return f.ObjectMeta
}

// SetObjectMeta sets the object metadata for the resource.
func (f *FallbackPipeline) SetObjectMeta(meta ObjectMeta) {
	f.ObjectMeta = meta
}

// SetNamespace sets the namespace of the resource.
func (f *FallbackPipeline) SetNamespace(namespace string) {
	f.Namespace = namespace
}

// StorePrefix returns the path prefix to this resource in the store.
func (f *FallbackPipeline) StorePrefix() string {
	return FallbackPipelineResource
}

// RBACName describes the name of the resource for RBAC purposes.
func (f *FallbackPipeline) RBACName() string {
	return FallbackPipelineResource
}

// URIPath gives the path component of a pipeline URI.
func (f *FallbackPipeline) URIPath() string {
	if f.Namespace == "" {
		return path.Join(URLPrefix, FallbackPipelineResource, url.PathEscape(f.Name))
	}
	return path.Join(URLPrefix, "namespaces", url.PathEscape(f.Namespace), FallbackPipelineResource, url.PathEscape(f.Name))
}

// // Validate checks if a pipeline resource passes validation rules.
func (f *FallbackPipeline) Validate() error {
	if err := ValidateName(f.ObjectMeta.Name); err != nil {
		return errors.New("name " + err.Error())
	}

	if f.ObjectMeta.Namespace == "" {
		return errors.New("namespace must be set")
	}
	//Manisha
	//for _, workflow := range f.Fallbackflows {
	//	if err := workflow.Validate(); err != nil {
	//		return fmt.Errorf("workflow %w", err)
	//	}
	//}

	return nil
}

// FallbackPipelineFields returns a set of fields that represent that resource.
func FallbackPipelineFields(r Resource) map[string]string {
	resource := r.(*FallbackPipeline)
	fields := map[string]string{
		"fallbackPipeline.name":      resource.ObjectMeta.Name,
		"fallbackPipeline.namespace": resource.ObjectMeta.Namespace,
	}
	stringsutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "fallbackPipeline.labels.")
	return fields
}

// Fields returns a set of fields that represent that resource.
func (f *FallbackPipeline) Fields() map[string]string {
	return FallbackPipelineFields(f)
}

// // FixturePipeline returns a testing fixture for a Pipeline object.
func FixtureFallbackPipeline(name, namespace string) *FallbackPipeline {
	return &FallbackPipeline{
		ObjectMeta: NewObjectMeta(name, namespace),
		Pipelist:   []*ResourceReference{},
	}
}

// // FixturePipelineReference returns a testing fixture for a ResourceReference
// // object referencing a corev2.Pipeline.
func FixtureFallbackPipelineReference(name string) *ResourceReference {
	return &ResourceReference{
		APIVersion: "core/v2",
		Type:       "FallbackPipeline",
		Name:       name,
	}
}
