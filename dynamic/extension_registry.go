package dynamic

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/golang/protobuf/proto"

	"github.com/jhump/protoreflect/desc"
)

// ExtensionRegistry is a registry of known extension fields. This is used to parse
// extension fields encountered when de-serializing a dynamic message.
type ExtensionRegistry struct {
	includeDefault bool
	mu             sync.RWMutex
	exts           map[string]map[int32]*desc.FieldDescriptor
}

// NewExtensionRegistryWithDefaults is a registry that includes all "default" extensions,
// which are those that are statically linked into the current program (e.g. registered by
// protoc-generated code via proto.RegisterExtension). Extensions explicitly added to the
// registry will override any default extensions that are for the same extendee and have the
// same tag number and/or name.
func NewExtensionRegistryWithDefaults() *ExtensionRegistry {
	return &ExtensionRegistry{includeDefault: true}
}

// AddExtensionDesc adds the given extensions to the registry.
func (r *ExtensionRegistry) AddExtensionDesc(exts ...*proto.ExtensionDesc) error {
	flds := make([]*desc.FieldDescriptor, len(exts))
	for i, ext := range exts {
		fd, err := desc.LoadFieldDescriptorForExtension(ext)
		if err != nil {
			return err
		}
		flds[i] = fd
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.exts == nil {
		r.exts = map[string]map[int32]*desc.FieldDescriptor{}
	}
	for _, fd := range flds {
		r.putExtensionLocked(fd)
	}
	return nil
}

// AddExtension adds the given extensions to the registry.
func (r *ExtensionRegistry) AddExtension(exts ...*desc.FieldDescriptor) error {
	for _, ext := range exts {
		if !ext.IsExtension() {
			return fmt.Errorf("Given field is not an extension: %s", ext.GetFullyQualifiedName())
		}
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.exts == nil {
		r.exts = map[string]map[int32]*desc.FieldDescriptor{}
	}
	for _, ext := range exts {
		r.putExtensionLocked(ext)
	}
	return nil
}

// AddExtensionsFromFile adds to the registry all extension fields defined in the given file descriptor.
func (r *ExtensionRegistry) AddExtensionsFromFile(fd *desc.FileDescriptor) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.exts == nil {
		r.exts = map[string]map[int32]*desc.FieldDescriptor{}
	}
	for _, ext := range fd.GetExtensions() {
		r.putExtensionLocked(ext)
	}
	for _, msg := range fd.GetMessageTypes() {
		r.addExtensionsFromMessageLocked(msg)
	}
}

func (r *ExtensionRegistry) addExtensionsFromMessageLocked(md *desc.MessageDescriptor) {
	for _, ext := range md.GetNestedExtensions() {
		r.putExtensionLocked(ext)
	}
	for _, msg := range md.GetNestedMessageTypes() {
		r.addExtensionsFromMessageLocked(msg)
	}
}

func (r *ExtensionRegistry) putExtensionLocked(fd *desc.FieldDescriptor) {
	msgName := fd.GetOwner().GetFullyQualifiedName()
	m := r.exts[msgName]
	if m == nil {
		m = map[int32]*desc.FieldDescriptor{}
		r.exts[msgName] = m
	}
	m[fd.GetNumber()] = fd
}

// FindExtension queries for the extension field with the given extendee name (must be a fully-qualified
// message name) and tag number. If no extension is known, nil is returned.
func (r *ExtensionRegistry) FindExtension(messageName string, tagNumber int32) *desc.FieldDescriptor {
	if r == nil {
		return nil
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	fd := r.exts[messageName][tagNumber]
	if fd == nil && r.includeDefault {
		ext := getDefaultExtensions(messageName)[tagNumber]
		if ext != nil {
			fd, _ = desc.LoadFieldDescriptorForExtension(ext)
		}
	}
	return fd
}

// FindExtensionByName queries for the extension field with the given extendee name (must be a fully-qualified
// message name) and field name (must also be a fully-qualified extension name). If no extension is known, nil
// is returned.
func (r *ExtensionRegistry) FindExtensionByName(messageName string, fieldName string) *desc.FieldDescriptor {
	if r == nil {
		return nil
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, fd := range r.exts[messageName] {
		if fd.GetFullyQualifiedName() == fieldName {
			return fd
		}
	}
	if r.includeDefault {
		for _, ext := range getDefaultExtensions(messageName) {
			fd, _ := desc.LoadFieldDescriptorForExtension(ext)
			if fd.GetFullyQualifiedName() == fieldName {
				return fd
			}
		}
	}
	return nil
}

func getDefaultExtensions(messageName string) map[int32]*proto.ExtensionDesc {
	t := proto.MessageType(messageName)
	if t != nil {
		msg := reflect.Zero(t).Interface().(proto.Message)
		return proto.RegisteredExtensions(msg)
	}
	return nil
}

// AllExtensionsForType returns all known extension fields for the given extendee name (must be a
// fully-qualified message name).
func (r *ExtensionRegistry) AllExtensionsForType(messageName string) []*desc.FieldDescriptor {
	if r == nil {
		return []*desc.FieldDescriptor(nil)
	}
	r.mu.RLock()
	defer r.mu.RUnlock()
	flds := r.exts[messageName]
	var ret []*desc.FieldDescriptor
	if r.includeDefault {
		exts := getDefaultExtensions(messageName)
		if len(exts) > 0 || len(flds) > 0 {
			ret = make([]*desc.FieldDescriptor, 0, len(exts)+len(flds))
		}
		for tag, ext := range exts {
			if _, ok := flds[tag]; ok {
				// skip default extension and use the one explicitly registered instead
				continue
			}
			fd, _ := desc.LoadFieldDescriptorForExtension(ext)
			if fd != nil {
				ret = append(ret, fd)
			}
		}
	} else if len(flds) > 0 {
		ret = make([]*desc.FieldDescriptor, 0, len(flds))
	}

	for _, ext := range flds {
		ret = append(ret, ext)
	}
	return ret
}
