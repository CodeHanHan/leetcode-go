package offer_41_d_

import "testing"

func TestMedianFinder_FindMedian(t *testing.T) {
	mf := Constructor()
	mf.AddNum(2)
	mf.AddNum(1)
	mf.AddNum(5)
	if mf.FindMedian() != 2 {
		t.Fatal()
	}
	mf.AddNum(3)
	if mf.FindMedian() != 2.5 {
		t.Fatal()
	}
}
