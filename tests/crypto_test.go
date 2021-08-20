package tests

import (
	"fmt"
	"testing"

	"github.com/kbrownehs18/gotool/cryptoutils"
)

func TestRSA(t *testing.T) {
	publicKey := `
-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBAMylYl8Llx0T4o6ygtjAEPUtUTzkSLTWNKWtYCleF2yD7pSuAhBfamlI
3HMgF54501bAwqgPugqmIRHvdot8hUFH0+l5jOD45mTTYC98QwxUo6XUDWKHX8Lf
jKaq1rD0avFlV9V4eMoNUNEbiRvtqdBbtjiyhGdkM/Gux6OYW895AgMBAAE=
-----END RSA PUBLIC KEY-----
`
	privateKey := `
-----BEGIN RSA PRIVATE KEY-----
MIICXgIBAAKBgQDMpWJfC5cdE+KOsoLYwBD1LVE85Ei01jSlrWApXhdsg+6UrgIQ
X2ppSNxzIBeeOdNWwMKoD7oKpiER73aLfIVBR9PpeYzg+OZk02AvfEMMVKOl1A1i
h1/C34ymqtaw9GrxZVfVeHjKDVDRG4kb7anQW7Y4soRnZDPxrsejmFvPeQIDAQAB
AoGAa5h0vRYh8GbZEj+m7gElqVKXSNWZqIKLEaSzT2yqEiLXmJJYgXU5RHvLdDgm
UsmCZTVZ4vTJ0vl/n6dwg2wHvtbSh4g7Zb4tCvimDRSY6HAf82YcrwtNGlS5wHzy
bkBskYb8u7n9Ohys7dcxQtmXfcO7ESX+3StP8syoEky2uoECQQDxv7g/MClIadAv
yK3J453BBxE5c+GIxgbY0R0Axvbil/rxbhgbCZXBppMjgpg3jcWtoW8hogIzSlSN
0AEpkHYRAkEA2LW84CVLFHmhR2uwbH95rvyUy+5kjAGM4TDAjlmNhASA2Kfi5MlJ
bQnKJ/ezHOyu4Pt35l5g3aJEWT1cdJi66QJBAOlEu/6M9GjhYXeaRseWkPRfY2ly
vd+CZbz1Gu1TD4taZ1RrjWsZdp3jo/sR2ttQO7ztFxT3BPSE9s3YNibrNGECQQC1
+f6uFoLyoaR97f9LTMxo1e85RGmoa9DadO7tWmQMnR95T5mnVyPSbWsVntoIivPb
Ny+bAlvDIXTVn0JZIrupAkEAznsvuxSZeTzmoxFEJ/9mcpw/YjyX858ts4RsocFO
ot7EvW+2jo93Suooj3EKA97ASXJtUG26LcZbi1JYAWVXYg==
-----END RSA PRIVATE KEY-----
`

	b := []byte("scnjl")
	bb, err := cryptoutils.RSAEncode(b, []byte(publicKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bb))
	bc, err := cryptoutils.RSADecode(bb, []byte(privateKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(bc))

	s, err := cryptoutils.RSABase64Encode(b, []byte(publicKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(s)
	ss, err := cryptoutils.RSABase64Decode(s, []byte(privateKey))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ss)
}

func TestHash(t *testing.T) {
	s := "scnjl"
	fmt.Println(cryptoutils.SHA1(s))
}
