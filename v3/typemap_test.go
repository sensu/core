package v3

// automatically generated file, do not edit!

import (
	"testing"

	"github.com/sensu/sensu-api-tools"
)

func TestResolveClusterConfig(t *testing.T) {
	var value interface{} = new(ClusterConfig)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "ClusterConfig")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "ClusterConfig")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ClusterConfig" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveEntityConfig(t *testing.T) {
	var value interface{} = new(EntityConfig)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "EntityConfig")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "EntityConfig")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"EntityConfig" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveEntityState(t *testing.T) {
	var value interface{} = new(EntityState)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "EntityState")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "EntityState")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"EntityState" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveNamespace(t *testing.T) {
	var value interface{} = new(Namespace)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "Namespace")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "Namespace")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"Namespace" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveResourceTemplate(t *testing.T) {
	var value interface{} = new(ResourceTemplate)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "ResourceTemplate")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "ResourceTemplate")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"ResourceTemplate" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestResolveSymmetricKey(t *testing.T) {
	var value interface{} = new(SymmetricKey)
	if _, ok := value.(Resource); ok {
		raw, err := apitools.Resolve("core/v3", "SymmetricKey")
		if err != nil {
			t.Fatal(err)
		}
		resource, ok := raw.(Resource)
		if !ok {
			t.Fatal("expected Resource")
		}
		meta := resource.GetMetadata()
		if meta == nil {
			t.Fatal("nil metadata")
		}
		if meta.Labels == nil {
			t.Error("nil metadata")
		}
		if meta.Annotations == nil {
			t.Error("nil annotations")
		}
		return
	}
	_, err := apitools.Resolve("core/v3", "SymmetricKey")
	if err == nil {
		t.Fatal("expected non-nil error")
	}
	if got, want := err.Error(), `"SymmetricKey" is not a Resource`; got != want {
		t.Fatalf("unexpected error: %s", err)
	}
}
