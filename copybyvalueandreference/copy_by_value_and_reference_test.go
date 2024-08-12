package copybyvalueandreference

import (
	"testing"
)

const (
	IterValue = 10
)

func BenchmarkCopyByValSmallStruct(b *testing.B) {
	obj := JustOneField{43}
	for i := 0; i < b.N; i++ {
		for j := 0; j < IterValue; j++ {
			copyByValueForSmallStruct(obj)
		}
	}
}

func BenchmarkCopyByRefSmallStruct(b *testing.B) {
	obj := &JustOneField{43}
	for i := 0; i < b.N; i++ {
		for j := 0; j < IterValue; j++ {
			copyByRefForSmallStruct(obj)
		}
	}
}

func BenchmarkCopyByValLargeStruct(b *testing.B) {
	obj := ManyFields{43, 54, 65, 67, 54, 65, 65, 22}
	for i := 0; i < b.N; i++ {
		for j := 0; j < IterValue; j++ {
			copyByValueForLargeStruct(obj)
		}
	}
}

func BenchmarkCopyByRefLargeStruct(b *testing.B) {
	obj := &ManyFields{43, 54, 65, 67, 54, 65, 65, 22}
	for i := 0; i < b.N; i++ {
		for j := 0; j < IterValue; j++ {
			copyByRefForLargeStruct(obj)
		}
	}
}
