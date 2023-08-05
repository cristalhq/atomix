package atomix

import (
	"sync/atomic"
	"unsafe"
)

// AddInt is same as [atomic.AddInt32] or [atomic.AddInt64] but for int type.
func AddInt(addr *int, delta int) (new int) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return int(atomic.AddInt32((*int32)(unsafe.Pointer(addr)), int32(delta)))
	case 8:
		return int(atomic.AddInt64((*int64)(unsafe.Pointer(addr)), int64(delta)))
	default:
		panic("atomix: int must be 4 or 8 bytes")
	}
}

// CompareAndSwapInt is same as [atomic.CompareAndSwapInt32] or [atomic.CompareAndSwapInt64] but for int type.
func CompareAndSwapInt(addr *int, old, new int) (swapped bool) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(addr)), int32(old), int32(new))
	case 8:
		return atomic.CompareAndSwapInt64((*int64)(unsafe.Pointer(addr)), int64(old), int64(new))
	default:
		panic("atomix: int must be 4 or 8 bytes")
	}
}

// LoadInt is same as [atomic.LoadInt32] or [atomic.LoadInt64] but for int type.
func LoadInt(addr *int) (val int) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return int(atomic.LoadInt32((*int32)(unsafe.Pointer(addr))))
	case 8:
		return int(atomic.LoadInt64((*int64)(unsafe.Pointer(addr))))
	default:
		panic("atomix: int must be 4 or 8 bytes")
	}
}

// StoreInt is same as [atomic.StoreInt32] or [atomic.StoreInt64] but for int type.
func StoreInt(addr *int, val int) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		atomic.StoreInt32((*int32)(unsafe.Pointer(addr)), int32(val))
	case 8:
		atomic.StoreInt64((*int64)(unsafe.Pointer(addr)), int64(val))
	default:
		panic("atomix: int must be 4 or 8 bytes")
	}
}

// SwapInt is same as [atomic.SwapInt32] or [atomic.SwapInt64] but for int type.
func SwapInt(addr *int, new int) (old int) {
	switch unsafe.Sizeof(*addr) {
	case 4:
		return int(atomic.SwapInt32((*int32)(unsafe.Pointer(addr)), int32(new)))
	case 8:
		return int(atomic.SwapInt64((*int64)(unsafe.Pointer(addr)), int64(new)))
	default:
		panic("atomix: int must be 4 or 8 bytes")
	}
}
