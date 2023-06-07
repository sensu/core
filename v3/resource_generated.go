package v3

// automatically generated file, do not edit!

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"path"
	"reflect"

	corev2 "github.com/sensu/core/v2"
)

// implement validator to add validation to Validate methods
type validator interface {
	validate() error
}

// implement storeNamer to override StoreName methods
type storeNamer interface {
	storeName() string
}

// implement rbacNamer to override RBACName methods
type rbacNamer interface {
	rbacName() string
}

// implement uriPather to override URIPath methods
type uriPather interface {
	uriPath() string
}

type getMetadataer interface {
	GetMetadata() *corev2.ObjectMeta
}

func uriPath(typename, namespace, name string) string {
	if namespace == "" {
		return path.Join("api", "core", "v3", typename, url.PathEscape(name))
	}
	return path.Join("api", "core", "v3", "namespaces", url.PathEscape(namespace), typename, url.PathEscape(name))
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (c *ClusterConfig) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(c))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for ClusterConfig. It will be
// overridden if there is a method for ClusterConfig called "storeName".
func (c *ClusterConfig) StoreName() string {
	var iface interface{} = c
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "cluster_configs"
}

// RBACName returns the RBAC name for ClusterConfig. It will be overridden if
// there is a method for ClusterConfig called "rbacName".
func (c *ClusterConfig) RBACName() string {
	var iface interface{} = c
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "cluster_configs"
}

// URIPath returns the URI path for ClusterConfig. It will be overridden if
// there is a method for ClusterConfig called uriPath.
func (c *ClusterConfig) URIPath() string {
	var iface interface{} = c
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("cluster-configs", "", "")
	}
	return uriPath("cluster-configs", meta.Namespace, meta.Name)
}

// Validate validates the ClusterConfig. If the ClusterConfig has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// ClusterConfig called validate, then it will be used to cooperatively
// validate the ClusterConfig.
func (c *ClusterConfig) Validate() error {
	if c == nil {
		return errors.New("nil ClusterConfig")
	}
	var iface interface{} = c
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid ClusterConfig: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid ClusterConfig: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (c *ClusterConfig) UnmarshalJSON(msg []byte) error {
	type Clone ClusterConfig
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*c = *(*ClusterConfig)(&clone)
	var iface interface{} = c
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a ClusterConfig.
func (c *ClusterConfig) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "ClusterConfig",
	}
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (e *EntityConfig) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(e))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for EntityConfig. It will be
// overridden if there is a method for EntityConfig called "storeName".
func (e *EntityConfig) StoreName() string {
	var iface interface{} = e
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "entity_configs"
}

// RBACName returns the RBAC name for EntityConfig. It will be overridden if
// there is a method for EntityConfig called "rbacName".
func (e *EntityConfig) RBACName() string {
	var iface interface{} = e
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "entity_configs"
}

// URIPath returns the URI path for EntityConfig. It will be overridden if
// there is a method for EntityConfig called uriPath.
func (e *EntityConfig) URIPath() string {
	var iface interface{} = e
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("entity-configs", "", "")
	}
	return uriPath("entity-configs", meta.Namespace, meta.Name)
}

// Validate validates the EntityConfig. If the EntityConfig has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// EntityConfig called validate, then it will be used to cooperatively
// validate the EntityConfig.
func (e *EntityConfig) Validate() error {
	if e == nil {
		return errors.New("nil EntityConfig")
	}
	var iface interface{} = e
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid EntityConfig: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid EntityConfig: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (e *EntityConfig) UnmarshalJSON(msg []byte) error {
	type Clone EntityConfig
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*e = *(*EntityConfig)(&clone)
	var iface interface{} = e
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a EntityConfig.
func (e *EntityConfig) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "EntityConfig",
	}
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (e *EntityState) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(e))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for EntityState. It will be
// overridden if there is a method for EntityState called "storeName".
func (e *EntityState) StoreName() string {
	var iface interface{} = e
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "entity_states"
}

// RBACName returns the RBAC name for EntityState. It will be overridden if
// there is a method for EntityState called "rbacName".
func (e *EntityState) RBACName() string {
	var iface interface{} = e
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "entity_states"
}

// URIPath returns the URI path for EntityState. It will be overridden if
// there is a method for EntityState called uriPath.
func (e *EntityState) URIPath() string {
	var iface interface{} = e
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("entity-states", "", "")
	}
	return uriPath("entity-states", meta.Namespace, meta.Name)
}

// Validate validates the EntityState. If the EntityState has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// EntityState called validate, then it will be used to cooperatively
// validate the EntityState.
func (e *EntityState) Validate() error {
	if e == nil {
		return errors.New("nil EntityState")
	}
	var iface interface{} = e
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid EntityState: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid EntityState: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (e *EntityState) UnmarshalJSON(msg []byte) error {
	type Clone EntityState
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*e = *(*EntityState)(&clone)
	var iface interface{} = e
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a EntityState.
func (e *EntityState) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "EntityState",
	}
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (n *Namespace) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(n))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for Namespace. It will be
// overridden if there is a method for Namespace called "storeName".
func (n *Namespace) StoreName() string {
	var iface interface{} = n
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "namespaces"
}

// RBACName returns the RBAC name for Namespace. It will be overridden if
// there is a method for Namespace called "rbacName".
func (n *Namespace) RBACName() string {
	var iface interface{} = n
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "namespaces"
}

// URIPath returns the URI path for Namespace. It will be overridden if
// there is a method for Namespace called uriPath.
func (n *Namespace) URIPath() string {
	var iface interface{} = n
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("namespaces", "", "")
	}
	return uriPath("namespaces", meta.Namespace, meta.Name)
}

// Validate validates the Namespace. If the Namespace has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// Namespace called validate, then it will be used to cooperatively
// validate the Namespace.
func (n *Namespace) Validate() error {
	if n == nil {
		return errors.New("nil Namespace")
	}
	var iface interface{} = n
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid Namespace: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid Namespace: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (n *Namespace) UnmarshalJSON(msg []byte) error {
	type Clone Namespace
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*n = *(*Namespace)(&clone)
	var iface interface{} = n
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a Namespace.
func (n *Namespace) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "Namespace",
	}
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (r *ResourceTemplate) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(r))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for ResourceTemplate. It will be
// overridden if there is a method for ResourceTemplate called "storeName".
func (r *ResourceTemplate) StoreName() string {
	var iface interface{} = r
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "resource_templates"
}

// RBACName returns the RBAC name for ResourceTemplate. It will be overridden if
// there is a method for ResourceTemplate called "rbacName".
func (r *ResourceTemplate) RBACName() string {
	var iface interface{} = r
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "resource_templates"
}

// URIPath returns the URI path for ResourceTemplate. It will be overridden if
// there is a method for ResourceTemplate called uriPath.
func (r *ResourceTemplate) URIPath() string {
	var iface interface{} = r
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("resource-templates", "", "")
	}
	return uriPath("resource-templates", meta.Namespace, meta.Name)
}

// Validate validates the ResourceTemplate. If the ResourceTemplate has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// ResourceTemplate called validate, then it will be used to cooperatively
// validate the ResourceTemplate.
func (r *ResourceTemplate) Validate() error {
	if r == nil {
		return errors.New("nil ResourceTemplate")
	}
	var iface interface{} = r
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid ResourceTemplate: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid ResourceTemplate: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (r *ResourceTemplate) UnmarshalJSON(msg []byte) error {
	type Clone ResourceTemplate
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*r = *(*ResourceTemplate)(&clone)
	var iface interface{} = r
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a ResourceTemplate.
func (r *ResourceTemplate) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "ResourceTemplate",
	}
}

// SetMetadata sets the provided metadata on the type. If the type does not
// have any metadata, nothing will happen.
func (s *SymmetricKey) SetMetadata(meta *corev2.ObjectMeta) {
	// The function has to use reflection, since not all of the generated types
	// will have metadata.
	value := reflect.Indirect(reflect.ValueOf(s))
	field := value.FieldByName("Metadata")
	if !field.CanSet() {
		return
	}
	field.Set(reflect.ValueOf(meta))
}

// StoreName returns the store name for SymmetricKey. It will be
// overridden if there is a method for SymmetricKey called "storeName".
func (s *SymmetricKey) StoreName() string {
	var iface interface{} = s
	if prefixer, ok := iface.(storeNamer); ok {
		return prefixer.storeName()
	}
	return "symmetric_keys"
}

// RBACName returns the RBAC name for SymmetricKey. It will be overridden if
// there is a method for SymmetricKey called "rbacName".
func (s *SymmetricKey) RBACName() string {
	var iface interface{} = s
	if namer, ok := iface.(rbacNamer); ok {
		return namer.rbacName()
	}
	return "symmetric_keys"
}

// URIPath returns the URI path for SymmetricKey. It will be overridden if
// there is a method for SymmetricKey called uriPath.
func (s *SymmetricKey) URIPath() string {
	var iface interface{} = s
	if pather, ok := iface.(uriPather); ok {
		return pather.uriPath()
	}
	metaer, ok := iface.(getMetadataer)
	if !ok {
		return ""
	}
	meta := metaer.GetMetadata()
	if meta == nil {
		return uriPath("symmetric-keys", "", "")
	}
	return uriPath("symmetric-keys", meta.Namespace, meta.Name)
}

// Validate validates the SymmetricKey. If the SymmetricKey has metadata,
// it will be validated via ValidateMetadata. If there is a method for
// SymmetricKey called validate, then it will be used to cooperatively
// validate the SymmetricKey.
func (s *SymmetricKey) Validate() error {
	if s == nil {
		return errors.New("nil SymmetricKey")
	}
	var iface interface{} = s
	if resource, ok := iface.(Resource); ok {
		meta := resource.GetMetadata()
		if err := ValidateMetadata(meta); err != nil {
			return fmt.Errorf("invalid SymmetricKey: %s", err)
		}
		if gr, ok := iface.(GlobalResource); ok && gr.IsGlobalResource() {
			if err := ValidateGlobalMetadata(meta); err != nil {
				return fmt.Errorf("invalid SymmetricKey: %s", err)
			}
		}
	}
	if validator, ok := iface.(validator); ok {
		if err := validator.validate(); err != nil {
			return err
		}
	}
	return nil
}

// UnmarshalJSON is provided in order to ensure that metadata labels and
// annotations are never nil.
func (s *SymmetricKey) UnmarshalJSON(msg []byte) error {
	type Clone SymmetricKey
	var clone Clone
	if err := json.Unmarshal(msg, &clone); err != nil {
		return err
	}
	*s = *(*SymmetricKey)(&clone)
	var iface interface{} = s
	var meta *corev2.ObjectMeta
	if metaer, ok := iface.(getMetadataer); ok {
		meta = metaer.GetMetadata()
	}
	if meta != nil {
		if meta.Labels == nil {
			meta.Labels = make(map[string]string)
		}
		if meta.Annotations == nil {
			meta.Annotations = make(map[string]string)
		}
	}
	return nil
}

// GetTypeMeta gets the type metadata for a SymmetricKey.
func (s *SymmetricKey) GetTypeMeta() corev2.TypeMeta {
	return corev2.TypeMeta{
		APIVersion: "core/v3",
		Type:       "SymmetricKey",
	}
}
