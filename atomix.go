package atomix

// atomicType is a helper struct to forbid some operations on types.
type atomicType struct {
	noCopy
	noEqual
}

// noCopy may be embedded into structs which must not be copied after the first use.
//
// See https://github.com/golang/go/issues/8005#issuecomment-190753527
type noCopy struct{}

func (*noCopy) Lock() {}

// noEqual is a not comparable struct to embed, if you want to prevent == and != usage.
type noEqual struct {
	_ [0]func()
}
