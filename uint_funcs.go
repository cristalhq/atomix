package atomix

import (
	"sync/atomic"
	"unsafe"
)

// AddUint is same as [atomic.AddUint32] or [atomic.AddUint64] but for uint.
func AddUint(addr *uint, delta uint) (new uint) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return uint(atomic.AddUint32((*uint32)(unsafe.Pointer(addr)), uint32(delta)))
	case 8:
		return uint(atomic.AddUint64((*uint64)(unsafe.Pointer(addr)), uint64(delta)))
	default:
		panic("atomix: uint must be 4 or 8 bytes")
	}
}

// CompareAndSwapUint is same as [atomic.CompareAndSwapUint32] or [atomic.CompareAndSwapUint64] but for uint.
func CompareAndSwapUint(addr *uint, old, new uint) (swapped bool) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return atomic.CompareAndSwapUint32((*uint32)(unsafe.Pointer(addr)), uint32(old), uint32(new))
	case 8:
		return atomic.CompareAndSwapUint64((*uint64)(unsafe.Pointer(addr)), uint64(old), uint64(new))
	default:
		panic("atomix: uint must be 4 or 8 bytes")
	}
}

// LoadUint is same as [atomic.LoadUint32] or [atomic.LoadUint64] but for uint.
func LoadUint(addr *uint) (val uint) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return uint(atomic.LoadUint32((*uint32)(unsafe.Pointer(addr))))
	case 8:
		return uint(atomic.LoadUint64((*uint64)(unsafe.Pointer(addr))))
	default:
		panic("atomix: uint must be 4 or 8 bytes")
	}
}

// StoreUint is same as [atomic.StoreUint32] or [atomic.StoreUint64] but for uint.
func StoreUint(addr *uint, val uint) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		atomic.StoreUint32((*uint32)(unsafe.Pointer(addr)), uint32(val))
	case 8:
		atomic.StoreUint64((*uint64)(unsafe.Pointer(addr)), uint64(val))
	default:
		panic("atomix: uint must be 4 or 8 bytes")
	}
}

// SwapUint is same as [atomic.SwapUint32] or [atomic.SwapUint64] but for uint.
func SwapUint(addr *uint, new uint) (old uint) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return uint(atomic.SwapUint32((*uint32)(unsafe.Pointer(addr)), uint32(new)))
	case 8:
		return uint(atomic.SwapUint64((*uint64)(unsafe.Pointer(addr)), uint64(new)))
	default:
		panic("atomix: uint must be 4 or 8 bytes")
	}
}
