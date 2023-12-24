package authorization

import "testing"

func TestSign(t *testing.T) {
	t.Log(sign("test", "1700000000", "123456"))
}
