package panicRecover

import (
	"testing"
	"time"
)

func TestDf(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"TestDf"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ticker := time.After(2 * time.Second)
			go Df()
			<-ticker
			t.Errorf("time out!\n")
		})
	}
}
