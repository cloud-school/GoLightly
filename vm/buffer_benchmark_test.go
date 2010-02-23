package vm
import "testing"
import "container/vector"

func BenchmarkBufferAt(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.At(0) }
}

func BenchmarkBufferSet(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Set(0, 1) }
}

func BenchmarkBufferAdd(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Add(0, 1) }
}

func BenchmarkBufferSubtract(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Subtract(0, 1) }
}

func BenchmarkBufferMultiply(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 2)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Multiply(0, 1) }
}

func BenchmarkBufferDivide(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 987654321)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Divide(0, 1) }
}

func BenchmarkBufferIncrement(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferDecrement(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Decrement(0) }
}

func BenchmarkBufferNegate(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Negate(0) }
}

func BenchmarkBufferShiftLeft(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftLeft(0, 1) }
}

func BenchmarkBufferShiftRight(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.ShiftRight(0, 1) }
}

func BenchmarkBufferInvert(b *testing.B) {
    b.StopTimer()
	buffer := defaultBuffer()
	buffer.Set(0, 100)
    b.StartTimer()
    for i := 0; i < b.N; i++ { buffer.Invert(0) }
}

func BenchmarkValueAt(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.At(0) }
}

func BenchmarkValueSet(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
    b.StartTimer()
    for i := 0; i < b.N; i++ { v.Set(0, 1) }
}

func BenchmarkValueAdd(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) + 1) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) + 1 }
}

func BenchmarkValueSubtract(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(0)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) - 1) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) - 1 }
}

func BenchmarkValueMultiply(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(2)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) * 27) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) * 27 }
}

func BenchmarkValueDivide(b *testing.B) {
    b.StopTimer()
	v := new(vector.Vector)
	v.Push(987654321)
	a := *v
    b.StartTimer()
//    for i := 0; i < b.N; i++ { v.Set(0, v.At(0).(int) / 4) }
    for i := 0; i < b.N; i++ { a[0] = a[0].(int) / 1 }
}