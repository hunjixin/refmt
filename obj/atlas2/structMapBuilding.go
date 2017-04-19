package atlas

import (
	"reflect"
	"strings"
)

func (x *BuilderCore) StructMap() *BuilderStructMap {
	return &BuilderStructMap{x.entry}
}

type BuilderStructMap struct {
	entry *AtlasEntry
}

func (x *BuilderStructMap) Complete() AtlasEntry {
	return *x.entry
}

/*
	Add a field to the mapping based on its name.

	Given a struct:

		struct{
			X int
			Y struct{ Z int }
		}

	`AddField("X", {"x", ...}) will cause that field to be serialized as key "x";
	`AddField("Y.Z", {"z", ...})` will cause that *nested* field to be serialized
	as key "z" in the same object (e.g. "x" and "z" will be siblings).

	Returns the mutated builder for convenient call chaining.

	If the fieldName string doesn't map onto the structure type info,
	a panic will be raised.
*/
func (x *BuilderStructMap) AddField(fieldName string, mapping StructMapEntry) *BuilderStructMap {
	fieldNameSplit := strings.Split(fieldName, ".")
	rr, err := fieldNameToReflectRoute(nil, fieldNameSplit)
	if err != nil {
		panic(err) // REVIEW: now that we have the builder obj, we could just curry these into it until 'Complete' is called (or, thus, 'MustComplete'!).
	}
	mapping.reflectRoute = rr
	x.entry.StructMap.Fields = append(x.entry.StructMap.Fields, mapping)
	return x
}

func fieldNameToReflectRoute(rt reflect.Type, fieldNameSplit []string) (rr reflectRoute, err error) {
	for _, fn := range fieldNameSplit {
		rf, ok := rt.FieldByName(fn)
		if !ok {
			return nil, ErrStructureMismatch{rt.Name(), "does not have field named " + fn}
		}
		rr = append(rr, rf.Index...)
	}
	return rr, nil
}
