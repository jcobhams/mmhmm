package config

import "testing"

func TestConfig(t *testing.T) {
	extras := []StaticValue{
		{Key: "KEYA", Value: "TRUE"},
		{Key: "KEYB", Value: "VALB"},
		{Key: "KEYC", Value: "VALC"},
		{Key: "SERVER_PORT", Value: "1010"},
	}
	c := New(NewStaticProvider(extras...))

	if 1010 != c.GetInt("SERVER_PORT") {
		t.Error("Expected SERVER_PORT to be 1010")
	}
	if "VALB" != c.GetString("KEYB") {
		t.Error("Expected KEYB to be VALB")
	}

}
