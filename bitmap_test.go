package bitmap_test

import (
	"bytes"
	"fmt"
	"net"
	"testing"

	"github.com/anexia-it/bitmap"
	"github.com/stretchr/testify/assert"
)

func TestIsSet(t *testing.T) {
	t.Run("OneByte", func(t *testing.T) {
		firstBitSet := []byte{0x1}
		assert.True(t, bitmap.IsSet(firstBitSet, 0))
		assert.False(t, bitmap.IsSet(firstBitSet, 1))

		secondBitSet := []byte{0x2}
		assert.False(t, bitmap.IsSet(secondBitSet, 0))
		assert.True(t, bitmap.IsSet(secondBitSet, 1))

		allBitsSet := []byte{0xff}
		for i := 0; i < len(allBitsSet)*8; i++ {
			assert.True(t, bitmap.IsSet(allBitsSet, uint(i)), "bit should be set", i)
		}
	})

	t.Run("TwoBytes", func(t *testing.T) {
		firstByteFirstBitSet := []byte{0x0, 0x1}
		assert.True(t, bitmap.IsSet(firstByteFirstBitSet, 0))
		assert.False(t, bitmap.IsSet(firstByteFirstBitSet, 1))

		firstByteSecondBitSet := []byte{0x0, 0x2}
		assert.False(t, bitmap.IsSet(firstByteSecondBitSet, 0))
		assert.True(t, bitmap.IsSet(firstByteSecondBitSet, 1))

		secondByteFirstBitSet := []byte{0x1, 0x0}
		assert.True(t, bitmap.IsSet(secondByteFirstBitSet, 8))

		secondByteSecondBitSet := []byte{0x2, 0x0}
		assert.True(t, bitmap.IsSet(secondByteSecondBitSet, 9))
	})
}

func TestSet(t *testing.T) {
	t.Run("OneByte", func(t *testing.T) {
		b := []byte{0x00}

		bitmap.Set(b, 0)
		bitmap.Set(b, 7)
		assert.EqualValues(t, []byte{0x81}, b)
	})

	t.Run("TwoBytes", func(t *testing.T) {
		b := []byte{0x00, 0x00}
		bitmap.Set(b, 0)
		bitmap.Set(b, 7)
		bitmap.Set(b, 9)
		bitmap.Set(b, 14)
		assert.EqualValues(t, []byte{0x42, 0x81}, b)
	})
}

func TestClear(t *testing.T) {
	t.Run("OneByte", func(t *testing.T) {
		b := []byte{0xff}
		bitmap.Clear(b, 0)
		bitmap.Clear(b, 7)
		assert.EqualValues(t, []byte{0x7e}, b)
	})

	t.Run("TwoBytes", func(t *testing.T) {
		b := []byte{0xf0, 0x0f}
		bitmap.Clear(b, 0)
		bitmap.Clear(b, 7)
		bitmap.Clear(b, 8)
		bitmap.Clear(b, 15)
		assert.EqualValues(t, []byte{0x70, 0x0e}, b)
	})
}

func TestMask(t *testing.T) {
	t.Run("LengthMismatch", func(t *testing.T) {
		b := make([]byte, 8)
		mask := make([]byte, 16)

		masked, err := bitmap.Mask(b, mask)
		assert.Nil(t, masked)
		assert.EqualError(t, err, "mismatching bit lengths: 64 vs 128")
	})

	t.Run("OK", func(t *testing.T) {
		t.Run("Prefix", func(t *testing.T) {
			for i := 0; i < 32; i++ {
				t.Run(fmt.Sprintf("%02dBits", i), func(t *testing.T) {
					b := bytes.Repeat([]byte{0xff}, 4)
					mask := net.CIDRMask(i+1, 32)

					masked, err := bitmap.Mask(b, mask)
					assert.NoError(t, err)
					assert.EqualValues(t, mask, masked)
				})
			}
		})

		t.Run("SingleBit", func(t *testing.T) {
			for i := 0; i < 32; i++ {
				t.Run(fmt.Sprintf("Bit%02d", i), func(t *testing.T) {
					b := bytes.Repeat([]byte{0xff}, 4)
					mask := make([]byte, 4)
					bitmap.Set(mask, uint(i))

					masked, err := bitmap.Mask(b, mask)
					assert.NoError(t, err)
					assert.EqualValues(t, mask, masked)
				})
			}
		})
	})
}
