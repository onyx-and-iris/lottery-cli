package lottery

import "testing"

func TestDrawUnique(t *testing.T) {
	for _, tc := range []struct {
		name   string
		count  int
		maxNum int
	}{
		{"small", 3, 10},
		{"exact", 5, 5},
		{"larger", 10, 50},
	} {
		t.Run(tc.name, func(t *testing.T) {
			for range 100 {
				nums := drawUnique(tc.count, tc.maxNum)
				if len(nums) != tc.count {
					t.Fatalf("got len=%d, want %d", len(nums), tc.count)
				}

				seen := make(map[int]struct{}, len(nums))
				prev := 0
				for _, n := range nums {
					if n < 1 || n > tc.maxNum {
						t.Fatalf("value %d out of range [1,%d]", n, tc.maxNum)
					}
					if n < prev {
						t.Fatalf("values not sorted: %v", nums)
					}
					if _, ok := seen[n]; ok {
						t.Fatalf("duplicate value %d in %v", n, nums)
					}
					prev = n
					seen[n] = struct{}{}
				}
			}
		})
	}
}
