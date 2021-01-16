// Copyright (C) 2020  Germán Fuentes Capella

package datetime

import (
	"testing"
	"time"
)

func TestNewFromTime(t *testing.T) {
	expectdate, err := time.Parse("2006-01-02", "2020-01-30")
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	dt := NewFromTime(expectdate)
	gotdate := dt.Time()
	if !expectdate.Equal(gotdate) {
		t.Fatalf("expected %v; got %v", expectdate, gotdate)
	}
}
