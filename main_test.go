package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		wantLen    int
		expectZero bool
	}{
		{"zero size", 0, 0, true},
		{"negative size", -10, 0, true},
		{"one element", 1, 1, false},
		{"ten elements", 10, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			require.Equal(t, tt.wantLen, len(data), "длина слайса должна совпадать")

			if tt.expectZero {
				assert.Empty(t, data, "для size<=0 ожидаем пустой слайс")
			} else {
				for i, v := range data {
					assert.GreaterOrEqualf(t, v, 0, "элемент[%d] должен быть >=0", i)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{42}, 42},
		{"all equal", []int{5, 5, 5, 5}, 5},
		{"ascending", []int{1, 2, 3, 4, 5}, 5},
		{"descending", []int{5, 4, 3, 2, 1}, 5},
		{"mixed", []int{7, 2, 9, 3, 8}, 9},
		{"negatives", []int{-5, -1, -3}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			assert.Equal(t, tt.want, got, fmt.Sprintf("maximum(%v)", tt.data))
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"one element", []int{100}, 100},
		{"small < chunks", []int{3, 1, 4}, 4},
		{"exact chunks", []int{5, 1, 7, 3, 9, 2, 8, 4}, 9},
		{"larger slice", []int{2, 9, 1, 8, 3, 7, 4, 6, 5, 0}, 9},
		{"mixed values", []int{10, -2, 5, 22, 17, 3, 19, 8, 14}, 22},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			assert.Equal(t, tt.want, got, fmt.Sprintf("maxChunks(%v)", tt.data))
		})
	}
}
