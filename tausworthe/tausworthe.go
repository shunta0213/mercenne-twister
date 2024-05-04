package tausworthe

const (
	P = 607
	Q = 273
)

type Tausworthe struct {
	// seed of the random number generator
	seed int

	// array of random numbers
	num [P]bool

	// current index of the array
	index uint
}

var tausworthe *Tausworthe

func init() {
	tausworthe = &Tausworthe{
		seed:  1,
		num:   [P]bool{},
		index: 0,
	}

	for i := 0; i < P-1; i++ {
		tausworthe.num[i] = true
	}

	for i := 0; i < i<<22; i++ {
		tausworthe.next()
	}
}

func Seed(seed int) {
	tausworthe.Seed(seed)
}

func Bool() bool {
	return tausworthe.Bool()
}

func Uint64() uint64 {
	return tausworthe.Uint64()
}

func Uint32() uint32 {
	return tausworthe.Uint32()
}

func Uint16() uint16 {
	return tausworthe.Uint16()
}

func Uint8() uint8 {
	return tausworthe.Uint8()
}

func (t *Tausworthe) Seed(seed int) {
	t.seed = seed
	for i := 0; i < seed; i++ {
		tausworthe.next()
	}
}

func (t *Tausworthe) Bool() bool {
	t.next()
	return t.num[t.index]
}

func (t *Tausworthe) Uint64() uint64 {
	t.next()
	var result uint64
	for i := 0; i < 64; i++ {
		if t.num[(t.index+uint(i))%P] {
			result |= 1 << i
		}
	}
	return result
}

func (t *Tausworthe) Uint32() uint32 {
	t.next()
	var result uint32
	for i := 0; i < 32; i++ {
		if t.num[(t.index+uint(i))%P] {
			result |= 1 << i
		}
	}
	return result
}

func (t *Tausworthe) Uint16() uint16 {
	t.next()
	var result uint16
	for i := 0; i < 16; i++ {
		if t.num[(t.index+uint(i))%P] {
			result |= 1 << i
		}
	}
	return result
}

func (t *Tausworthe) Uint8() uint8 {
	t.next()
	var result uint8
	for i := 0; i < 8; i++ {
		if t.num[(t.index+uint(i))%P] {
			result |= 1 << i
		}
	}
	return result
}

func (t *Tausworthe) next() {
	t.index = (t.index + 1) % P
	t.num[t.index] = t.num[(t.index+Q)%P] != t.num[t.index]
}
