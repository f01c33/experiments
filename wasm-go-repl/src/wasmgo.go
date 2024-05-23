// +build js,wasm

package sum

//export sum
func sum(a, b int32) int32 {
	return a + b
}
