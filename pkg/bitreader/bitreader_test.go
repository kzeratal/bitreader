package bitreader

import "testing"

func TestNewReader(t *testing.T) {
	str := "This is a string."
	bitreader := NewReader([]byte(str))
	if string(bitreader.bytes) != str {
		t.Fail()
	}
}

func TestRead9Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	_, err := bitreader.Read(9)
	if err == nil {
		t.Fail()
	}
}

func TestRead8Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	if byte, err := bitreader.Read(8); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 255 {
		t.Error("should be 255")
		t.Fail()
	}
	if byte, err := bitreader.Read(8); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 0 {
		t.Error("should be 0")
		t.Fail()
	}
}

func TestRead7Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	if byte, err := bitreader.Read(7); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 127 {
		t.Error("should be 127")
		t.Fail()
	}
	if byte, err := bitreader.Read(7); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 1 {
		t.Error("should be 1")
		t.Fail()
	}
}

func TestRead6Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	if byte, err := bitreader.Read(6); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 63 {
		t.Error("should be 63")
		t.Fail()
	}
	if byte, err := bitreader.Read(6); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 3 {
		t.Error("should be 3")
		t.Fail()
	}
}

func TestRead5Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	if byte, err := bitreader.Read(5); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 31 {
		t.Error("should be 31")
		t.Fail()
	}
	if byte, err := bitreader.Read(5); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 7 {
		t.Error("should be 7")
		t.Fail()
	}
}

func TestRead4Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	for i := 0; i < 2; i++ {
		if byte, err := bitreader.Read(4); err != nil {
			t.Error("should not raise error")
			t.Fail()
		} else if int(byte) != 15 {
			t.Error("should be 15")
			t.Fail()
		}
	}
	if byte, err := bitreader.Read(4); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 0 {
		t.Error("should be 0")
		t.Fail()
	}
}

func TestRead3Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	for i := 0; i < 2; i++ {
		if byte, err := bitreader.Read(3); err != nil {
			t.Error("should not raise error")
			t.Fail()
		} else if int(byte) != 7 {
			t.Error("should be 7")
			t.Fail()
		}
	}
	if byte, err := bitreader.Read(3); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 3 {
		t.Error("should be 3")
		t.Fail()
	}
}

func TestRead2Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	for i := 0; i < 4; i++ {
		if byte, err := bitreader.Read(2); err != nil {
			t.Error("should not raise error")
			t.Fail()
		} else if int(byte) != 3 {
			t.Error("should be 3")
			t.Fail()
		}
	}
	if byte, err := bitreader.Read(2); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 0 {
		t.Error("should be 0")
		t.Fail()
	}
}

func TestRead1Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	for i := 0; i < 8; i++ {
		if byte, err := bitreader.Read(1); err != nil {
			t.Error("should not raise error")
			t.Fail()
		} else if int(byte) != 1 {
			t.Error("should be 1")
			t.Fail()
		}
	}
	if byte, err := bitreader.Read(1); err != nil {
		t.Error("should not raise error")
		t.Fail()
	} else if int(byte) != 0 {
		t.Error("should be 0")
		t.Fail()
	}
}

func TestRead0Bits(t *testing.T) {
	bitreader := NewReader([]byte{255})
	_, err := bitreader.Read(0)
	if err == nil {
		t.Fail()
	}
}