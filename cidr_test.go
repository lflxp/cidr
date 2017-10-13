package cidr

import "testing"

func TestCidr_GetCidrIpRange(t *testing.T) {
	data := NewCidr("10.6.2.2/20").GetCidrIpRange()

	if data.Min == "" {
		t.Error("Min error")
	} else {
		t.Log("Min ",data.Min)
	}

	if data.Max == "" {
		t.Error("Max error")
	} else {
		t.Log("Max ",data.Max)
	}
}

func TestCidr_GetCidrHostNum(t *testing.T) {
	data := NewCidr("10.6.2.2/20").GetCidrHostNum()

	if data.Count == "" {
		t.Error("Count is nil")
	} else {
		t.Log(data.Count)
	}
}

func TestCidr_GetCidrIpMask(t *testing.T) {
	data := NewCidr("10.6.2.2/20").GetCidrIpMask()

	if data.Netmask == "" {
		t.Error("Netmask is nil")
	} else {
		t.Log(data.Netmask)
	}
}

func TestNewCidr(t *testing.T) {
	data := NewCidr("10.6.2.2/20").GetCidrIpRange()

	if data.Min == "" {
		t.Error("Min error")
	} else {
		t.Log("Min ",data.Min)
	}

	if data.Max == "" {
		t.Error("Max error")
	} else {
		t.Log("Max ",data.Max)
	}

	data.CidrIpRange = "8.8.8.8/10"
	data.GetCidrIpRange()
	if data.Min == "" {
		t.Error("Min error")
	} else {
		t.Log("Min ",data.Min)
	}

	if data.Max == "" {
		t.Error("Max error")
	} else {
		t.Log("Max ",data.Max)
	}
}