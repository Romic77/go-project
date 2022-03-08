package main

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type testCase struct {
		str  string
		sep  string
		want []string
	}

	testGroup := map[string]testCase{"case1": {
		str:  "abcd",
		sep:  "b",
		want: []string{"a", "cd"},
	},
		"case2": {
			str:  "abcd",
			sep:  "bc",
			want: []string{"a", "d"},
		},
		"case3": {
			str:  "a:b:c:d",
			sep:  ":",
			want: []string{"a", "b", "c", "d"},
		}}

	/*for _, value := range testGroup {
		gots := split(value.str, value.sep)
		if !reflect.DeepEqual(value.want, gots) {
			t.Fatalf("expect:%#v got:%#v\n", value.want, gots)
		}
	}*/

	for key, value := range testGroup {
		t.Run(key, func(t *testing.T) {
			gots := split(value.str, value.sep)
			if !reflect.DeepEqual(value.want, gots) {
				t.Fatalf("expect:%#v got:%#v\n", value.want, gots)
			}
		})
	}
}

//单元测试覆盖率是  go test -cover
//单元测试覆盖率以文件生成 go test -cover -coverprofile=cover.out
//单元测试覆盖率以html展示 go tool cover -html=cover.out

//基准测试go test -bench=Split
//go test -bench=Split -benchmem 查看申请内存
func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		split("abcdef", "b")
	}
}
