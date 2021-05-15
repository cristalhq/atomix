package atomix

// CacheLine of the CPU.
// See aligned_cachelineXXX.go files
const CacheLine = cacheLineBytes

type atomicType struct {
	_ incomparable
}

type incomparable [0]func()
