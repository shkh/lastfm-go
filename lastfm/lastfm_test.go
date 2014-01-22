package lastfm

import (
	"testing"
)

func compareMap(from, to map[string]string) (result bool) {
	result = true
	for key, val1 := range from {
		if val2, ok := to[key]; !ok || val1 != val2 {
			result = false
			break
		}
	}
	return
}

func TestFormatArgs(t *testing.T) {
	tests := []struct {
		in1 P
		in2 P
		out map[string]string
	}{
		{
			P{"tags": []string{"t0", "t1", "t2", "t3", "t4", "t5", "t6", "t7", "t8", "t9", "t10", "t11", "t12"}},
			P{"normal": []string{"tags"}},
			map[string]string{"tags": "t0,t1,t2,t3,t4,t5,t6,t7,t8,t9"},
		},
		{
			P{"tags": "t"},
			P{"normal": []string{"tags"}},
			map[string]string{"tags": "t"},
		},
		{
			P{"artist": []string{"a0", "a1", "a2"}},
			P{"indexing": []string{"artist"}},
			map[string]string{"artist[0]": "a0", "artist[1]": "a1", "artist[2]": "a2"},
		},
		{
			P{"artist": "a"},
			P{"indexing": []string{"artist"}},
			map[string]string{"artist[0]": "a"},
		},
		{
			P{"artist": []string{"a0", "a1", "a2"}, "recipient": []string{"r0", "r1", "r2"}},
			P{"indexing": []string{"artist"}, "normal": []string{"recipient"}},
			map[string]string{"artist[0]": "a0", "artist[1]": "a1", "artist[2]": "a2", "recipient": "r0,r1,r2"},
		},
		{
			P{"tags": []string{"t0", "t1", "t2"}, "artist": []string{"a0", "a1"}, "id": 10, "name": "kumakichi"},
			P{"normal": []string{"id", "name", "tags"}, "indexing": []string{"artist"}},
			map[string]string{"tags": "t0,t1,t2", "artist[0]": "a0", "artist[1]": "a1", "id": "10", "name": "kumakichi"},
		},
	}

	for i, tt := range tests {
		if result, _ := formatArgs(tt.in1, tt.in2); !compareMap(tt.out, result) {
			t.Errorf("case %d", i+1)
			t.Fail()
		}
	}
}
