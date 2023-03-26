package firestore

import "testing"

func TestNewClient(t *testing.T) {
	_, err := NewCollection()
	if err != nil {
		t.Errorf("failed get firestore client : %v", err)
	}
}
