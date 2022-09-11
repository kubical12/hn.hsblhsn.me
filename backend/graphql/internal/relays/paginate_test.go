//nolint:goerr113
package relays

import (
	"fmt"
	"strconv"
	"testing"
)

func intPtr(i int) *int {
	return &i
}

func strPtr(s string) *string {
	return &s
}

func TestPaginate(t *testing.T) {
	t.Parallel()
	out, err := Paginate(100, ConnectionInput{
		Before: strPtr("157"),
		After:  strPtr("132"),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(out.Start, out.End)
}

func FuzzPaginatorBeforeAfter(f *testing.F) {
	f.Add(100, 50)
	f.Add(100, 50)
	arr := make([]ConnectionInput, 100)
	total := len(arr)
	f.Fuzz(func(t *testing.T, before, after int) {
		t.Helper()
		bef := strPtr(strconv.Itoa(before))
		aft := strPtr(strconv.Itoa(after))
		resp, err := Paginate(total, ConnectionInput{
			Before: bef,
			After:  aft,
		})
		if err != nil {
			return
		}
		if resp.Start > total {
			panic(fmt.Errorf("start %d > total %d", resp.Start, total))
		}
		if resp.End > total {
			panic(fmt.Errorf("end %d > total %d", resp.End, total))
		}
		if resp.Start < 0 {
			panic(fmt.Errorf("start %d < 0", resp.Start))
		}
		if resp.End < 0 {
			panic(fmt.Errorf("end %d < 0", resp.End))
		}
		if resp.End-resp.Start > 10 {
			panic(fmt.Errorf("end %d - start %d > 10", resp.End, resp.Start))
		}
		_ = arr[resp.Start:resp.End]
	})
}

func FuzzPaginatorFirstLast(f *testing.F) {
	f.Add(0, 10)
	f.Add(10, 0)
	arr := make([]ConnectionInput, 100)
	total := len(arr)
	f.Fuzz(func(t *testing.T, first, last int) {
		t.Helper()
		fir := intPtr(first)
		las := intPtr(last)
		resp, err := Paginate(total, ConnectionInput{
			First: fir,
			Last:  las,
		})
		if err != nil {
			return
		}
		if resp.Start > total {
			panic(fmt.Errorf("start %d > total %d", resp.Start, total))
		}
		if resp.End > total {
			panic(fmt.Errorf("end %d > total %d", resp.End, total))
		}
		if resp.Start < 0 {
			panic(fmt.Errorf("start %d < 0", resp.Start))
		}
		if resp.End < 0 {
			panic(fmt.Errorf("end %d < 0", resp.End))
		}
		if resp.End-resp.Start > 10 {
			panic(fmt.Errorf("end %d - start %d > 10", resp.End, resp.Start))
		}
		_ = arr[resp.Start:resp.End]
	})
}
