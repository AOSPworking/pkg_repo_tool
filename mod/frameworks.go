package mod

// FrameworkAllowModuleType is module types collection.
// frameworks allow.
var FrameworkAllowModuleType map[string]bool = map[string]bool{
	"cc_library_headers":     true,
	"cc_library_shared":      true,
	"cc_library_static":      true,
	"cc_library":             true,
	"cc_binary":              true,
	"cc_library_host_shared": true,

	"ndk_library":        true,
	"ndk_headers":        true,
	"prebuilt_etc":       true,
	"python_binary_host": true,

	"java_library":     true,
	"java_sdk_library": true,
}
