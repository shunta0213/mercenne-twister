package tausworthe_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/shunta0213/mercenne-twister/tausworthe"
)

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func TestSeed(t *testing.T) {
	t.Run("first call", func(t *testing.T) {
		tausworthe.Seed(1)
		for i := 0; i < 10; i++ {
			t.Log(tausworthe.Bool())
		}
	})

	t.Run("randomness check", func(t *testing.T) {
		for i := 0; i < 25; i++ {
			file, _ := os.Create(fmt.Sprintf("data/seed_random/seed_random_%d.txt", i))
			defer file.Close()
			tausworthe.Seed(1 << i)
			for j := 0; j < 1<<10; j++ {
				file.WriteString(fmt.Sprint(tausworthe.Uint64(), "\n"))
				t.Log(tausworthe.Uint64())
			}
		}
	})
}

func TestBool(t *testing.T) {
	t.Run("first call", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Log(tausworthe.Bool())
		}
	})

	// for verifying randomness
	t.Run("gen 2^15 numbers", func(t *testing.T) {
		t.SkipNow()

		file, err := os.OpenFile("data/bool/tausworthe_fuzz.txt", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			t.Fatal(err)
		}

		for i := 0; i < 1<<15; i++ {
			file.WriteString(fmt.Sprint(Btoi(tausworthe.Bool()), "\n"))
			t.Log(tausworthe.Bool())
		}
	})
}

func TestUint64(t *testing.T) {
	t.Run("first call", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			t.Log(tausworthe.Uint64())
		}
	})

	// for verifying randomness
	t.Run("gen 2^15 numbers", func(t *testing.T) {
		t.SkipNow()

		file, err := os.Create("data/uint64/fuzz.txt")
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 1<<15; i++ {
			file.WriteString(fmt.Sprint(tausworthe.Uint64(), "\n"))
			t.Log(tausworthe.Uint64())
		}
	})
}

func BenchmarkBool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tausworthe.Bool()
	}
}

func BenchmarkUint64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tausworthe.Uint64()
	}
}
