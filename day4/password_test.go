package main

import "testing"

func TestPassword1(t *testing.T) {
	if !is_password(112233) {
		t.Errorf("is_password(112233) = false; want true")
	}
}

func TestPassword2(t *testing.T) {
	if is_password(123444) {
		t.Errorf("is_password(123444) = true; want false")
	}
}

func TestPassword3(t *testing.T) {
	if !is_password(111122) {
		t.Errorf("is_password(111122) = false; want true")
	}
}

func TestPassword4(t *testing.T) {
	if is_password(111111) {
		t.Errorf("is_password(111111) = true; want false")
	}
}

func TestPassword5(t *testing.T) {
	if is_password(223450) {
		t.Errorf("is_password(223450) = true; want false")
	}
}

func TestPassword6(t *testing.T) {
	if is_password(441111) {
		t.Errorf("is_password(111144) = true; want false")
	}
}

func TestPassword7(t *testing.T) {
	if !is_password(111448) {
		t.Errorf("is_password(111144) = false; want true")
	}
}

func TestPassword8(t *testing.T) {
	if !is_password(119999) {
		t.Errorf("is_password(119999) = false; want true")
	}
}
