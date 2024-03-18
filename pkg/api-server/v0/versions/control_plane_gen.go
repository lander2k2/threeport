// generated by 'threeport-sdk codegen api-model' - do not edit

package versions

import (
	api "github.com/threeport/threeport/pkg/api"
	iapi "github.com/threeport/threeport/pkg/api-server/v0"
	v0 "github.com/threeport/threeport/pkg/api/v0"
	"reflect"
)

// AddControlPlaneDefinitionVersions adds field validation info and adds it
// to the REST API versions.
func AddControlPlaneDefinitionVersions() {
	iapi.ControlPlaneDefinitionTaggedFields[iapi.TagNameValidate] = &iapi.FieldsByTag{
		Optional:             []string{},
		OptionalAssociations: []string{},
		Required:             []string{},
		TagName:              iapi.TagNameValidate,
	}

	// parse struct and populate the FieldsByTag object
	iapi.ParseStruct(
		iapi.TagNameValidate,
		reflect.ValueOf(new(v0.ControlPlaneDefinition)),
		"",
		iapi.Translate,
		iapi.ControlPlaneDefinitionTaggedFields,
	)

	// create a version object which contains the object name and versions
	versionObj := iapi.VersionObject{
		Object:  string(v0.ObjectTypeControlPlaneDefinition),
		Version: "v0",
	}

	// add the object tagged fields to the global tagged fields map
	iapi.ObjectTaggedFields[versionObj] = iapi.ControlPlaneDefinitionTaggedFields[iapi.TagNameValidate]

	// add the object tagged fields to the rest API version
	api.AddRestApiVersion(versionObj)
}

// AddControlPlaneInstanceVersions adds field validation info and adds it
// to the REST API versions.
func AddControlPlaneInstanceVersions() {
	iapi.ControlPlaneInstanceTaggedFields[iapi.TagNameValidate] = &iapi.FieldsByTag{
		Optional:             []string{},
		OptionalAssociations: []string{},
		Required:             []string{},
		TagName:              iapi.TagNameValidate,
	}

	// parse struct and populate the FieldsByTag object
	iapi.ParseStruct(
		iapi.TagNameValidate,
		reflect.ValueOf(new(v0.ControlPlaneInstance)),
		"",
		iapi.Translate,
		iapi.ControlPlaneInstanceTaggedFields,
	)

	// create a version object which contains the object name and versions
	versionObj := iapi.VersionObject{
		Object:  string(v0.ObjectTypeControlPlaneInstance),
		Version: "v0",
	}

	// add the object tagged fields to the global tagged fields map
	iapi.ObjectTaggedFields[versionObj] = iapi.ControlPlaneInstanceTaggedFields[iapi.TagNameValidate]

	// add the object tagged fields to the rest API version
	api.AddRestApiVersion(versionObj)
}
