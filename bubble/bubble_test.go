package bubble

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{4, 1, 7, 9, 2, 5, 3}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 ||
		values[4] != 5 || values[5] != 7 || values[6] != 9 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 4 5 7 9")
	}
}
