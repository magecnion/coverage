package helper_test

import (
	"errors"
	"fmt"
	"helper"
	"testing"
)

func TestDate(t *testing.T) {
	tcs := []struct {
		day, month, year int
		out              string
		err              error
	}{
		{1, 1, 1970, "1st of Jan 1970", nil},
		{2, 1, 1970, "2nd of Jan 1970", nil},
		{3, 1, 1970, "3rd of Jan 1970", nil},
		{1, 2, 1970, "1st of Feb 1970", nil},
		{4, 2, 1970, "4th of Feb 1970", nil},
		{4, 13, 1970, "", fmt.Errorf("invalid month")},
	}
	for _, tc := range tcs {
		dateStr, err := helper.Date(tc.day, tc.month, tc.year)
		if errors.Is(err, tc.err) == false {
			if err.Error() != tc.err.Error() {
				t.Errorf("expected error to be %v got %v", tc.err, err)
			}
		}
		if dateStr != tc.out {
			t.Errorf("expected %s but got %s", tc.out, dateStr)
		}
	}
}
