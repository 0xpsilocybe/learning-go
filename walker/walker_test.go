package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct{
		Name string
		Input interface{}
		ExpectedCalls []string
	} {
		{
			Name: "struct with one string field",
			Input: struct {
				Name string
			}{ "Psilo" },
			ExpectedCalls: []string{ "Psilo" },
		},
		{
			Name: "struct with two string fields",
			Input: struct {
				Origin string
				Destination string
			}{ "BOM", "LHR" },
			ExpectedCalls: []string{ "BOM", "LHR" },
		},
		{
			Name: "struct with non-string field",
			Input: struct {
				Name string
				Age int
			}{ "Earth", 4700000000 },
			ExpectedCalls: []string{ "Earth" },
		},
		{
			Name: "struct with nested fields",
			Input: struct {
				Name string
				Address struct {
					City string
				}
			}{
				Name: "Noodles Company",
				Address: struct { 
					City string
				} { 
					"London",
				},
			},
			ExpectedCalls: []string{ "Noodles Company", "London" },
		},
		{
			Name: "pointer to struct",
			Input: &struct{
				Name string
			}{
				"Dolores",
			},
			ExpectedCalls: []string{ "Dolores" },
		},
		{
			Name: "slice of structs",
			Input: []struct{
				Page string
			}{
				{ "i" },
				{ "ii" },
				{ "iii" },
				{ "iv" },
				{ "v" },
			},
			ExpectedCalls: []string{ "i", "ii", "iii", "iv", "v" },
		},
		{
			Name: "array of structs",
			Input: [2]struct{
				Country string
			}{
				{ "Kongo" },
				{ "Andora" },
			},
			ExpectedCalls: []string{ "Kongo", "Andora" },
		},
		{
			Name: "map of structs",
			Input: map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			ExpectedCalls: []string{ "Bar", "Boz" },
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got[]string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			want := test.ExpectedCalls
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}

	t.Run("map of structs", func(t *testing.T) {
		input := map[string]string {
			"Foo": "Bar",
			"Baz": "Boz",
		}
		var got []string
		Walk(input, func(in string) {
			got = append(got, in)
		})
		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %s, but it didn't", haystack, needle)
	}
}

