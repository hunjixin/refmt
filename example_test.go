package refmt_test

import (
	"bytes"
	"fmt"

	"github.com/polydawn/refmt"
	"github.com/polydawn/refmt/obj/atlas"
)

func ExampleJsonEncodeAtlasDefaults() {
	type MyType struct {
		X string
		Y int
	}

	MyType_AtlasEntry := atlas.BuildEntry(MyType{}).
		StructMap().Autogenerate().
		Complete()

	atl := atlas.MustBuild(
		MyType_AtlasEntry,
		// this is a vararg... stack more entries here!
	)

	var buf bytes.Buffer
	encoder := refmt.NewAtlasedJsonEncoder(&buf, atl)
	err := encoder.Marshal(MyType{"a", 1})
	fmt.Println(buf.String())
	fmt.Printf("%v\n", err)

	// Output:
	// {"x":"a","y":1}
	// <nil>
}

func ExampleJsonEncodeAtlasCustom() {
	type MyType struct {
		X string
		Y int
	}

	MyType_AtlasEntry := atlas.BuildEntry(MyType{}).
		StructMap().
		AddField("X", atlas.StructMapEntry{SerialName: "overrideName"}).
		// and no "Y" mapping at all!
		Complete()

	atl := atlas.MustBuild(
		MyType_AtlasEntry,
		// this is a vararg... stack more entries here!
	)

	var buf bytes.Buffer
	encoder := refmt.NewAtlasedJsonEncoder(&buf, atl)
	err := encoder.Marshal(MyType{"a", 1})
	fmt.Println(buf.String())
	fmt.Printf("%v\n", err)

	// Output:
	// {"overrideName":"a"}
	// <nil>
}
