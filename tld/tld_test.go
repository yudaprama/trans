package tld

import (
	"testing"
)

var testCases = []struct {
	domain      string
	expectedTLD string
}{
	// Empty string.
	{"", ""},

	// id
	{"id", "id"},
	{"aaa.id", "id"},
	{"www.aaa.id", "id"},
	{"mod.id", "id"},
	{"www.mod.id", "id"},
	{"sch.id", "sch.id"},
	{"mod.sch.id", "sch.id"},
	{"www.sch.id", "sch.id"},
	{"co.id", "co.id"},
	{"www.co.id", "co.id"},
	{"blogspot.co.id", "blogspot.co.id"},
	{"blogspot.nic.id", "id"},
	{"blogspot.sch.id", "sch.id"},

	// ao
	{"ao", "ao"},
	{"www.ao", "ao"},
	{"pb.ao", "pb.ao"},
	{"www.pb.ao", "pb.ao"},
	{"www.xxx.yyy.zzz.pb.ao", "pb.ao"},

	// ar
	{"ar", "ar"},
	{"www.ar", "ar"},
	{"nic.ar", "ar"},
	{"www.nic.ar", "ar"},
	{"com.ar", "com.ar"},
	{"www.com.ar", "com.ar"},
	{"blogspot.com.ar", "blogspot.com.ar"},
	{"www.blogspot.com.ar", "blogspot.com.ar"},
	{"www.xxx.yyy.zzz.blogspot.com.ar", "blogspot.com.ar"},
	{"logspot.com.ar", "com.ar"},
	{"zlogspot.com.ar", "com.ar"},
	{"zblogspot.com.ar", "com.ar"},

	// arpa
	{"arpa", "arpa"},
	{"www.arpa", "arpa"},
	{"urn.arpa", "urn.arpa"},
	{"www.urn.arpa", "urn.arpa"},
	{"www.xxx.yyy.zzz.urn.arpa", "urn.arpa"},

	// jp
	{"jp", "jp"},
	{"kobe.jp", "jp"},
	{"c.kobe.jp", "c.kobe.jp"},
	{"b.c.kobe.jp", "c.kobe.jp"},
	{"a.b.c.kobe.jp", "c.kobe.jp"},
	{"city.kobe.jp", "kobe.jp"},
	{"www.city.kobe.jp", "kobe.jp"},
	{"kyoto.jp", "kyoto.jp"},
	{"test.kyoto.jp", "kyoto.jp"},
	{"ide.kyoto.jp", "ide.kyoto.jp"},
	{"b.ide.kyoto.jp", "ide.kyoto.jp"},
	{"a.b.ide.kyoto.jp", "ide.kyoto.jp"},

	// tw
	{"tw", "tw"},
	{"aaa.tw", "tw"},
	{"www.aaa.tw", "tw"},
	{"xn--czrw28b.aaa.tw", "tw"},
	{"edu.tw", "edu.tw"},
	{"www.edu.tw", "edu.tw"},
	{"xn--czrw28b.edu.tw", "edu.tw"},
	{"xn--czrw28b.tw", "xn--czrw28b.tw"},
	{"www.xn--czrw28b.tw", "xn--czrw28b.tw"},
	{"xn--uc0atv.xn--czrw28b.tw", "xn--czrw28b.tw"},
	{"xn--kpry57d.tw", "tw"},

	// uk
	{"uk", "uk"},
	{"aaa.uk", "uk"},
	{"www.aaa.uk", "uk"},
	{"mod.uk", "uk"},
	{"www.mod.uk", "uk"},
	{"sch.uk", "uk"},
	{"mod.sch.uk", "mod.sch.uk"},
	{"www.sch.uk", "www.sch.uk"},
	{"co.uk", "co.uk"},
	{"www.co.uk", "co.uk"},
	{"blogspot.co.uk", "blogspot.co.uk"},
	{"blogspot.nic.uk", "uk"},
	{"blogspot.sch.uk", "blogspot.sch.uk"},

	// рф (xn--p1ai)
	{"xn--p1ai", "xn--p1ai"},
	{"aaa.xn--p1ai", "xn--p1ai"},
	{"www.xxx.yyy.xn--p1ai", "xn--p1ai"},

	// *.bd
	{"bd", "bd"},
	{"www.bd", "www.bd"},
	{"xxx.www.bd", "www.bd"},
	{"zzz.bd", "zzz.bd"},
	{"www.zzz.bd", "zzz.bd"},
	{"www.xxx.yyy.zzz.bd", "zzz.bd"},

	// *.ck
	{"ck", "ck"},
	{"www.ck", "ck"},
	{"xxx.www.ck", "ck"},
	{"zzz.ck", "zzz.ck"},
	{"www.zzz.ck", "zzz.ck"},
	{"www.xxx.yyy.zzz.ck", "zzz.ck"},

	// myjino.ru
	{"myjino.ru", "myjino.ru"},
	{"aaa.myjino.ru", "myjino.ru"},
	{"bbb.ccc.myjino.ru", "myjino.ru"},
	{"hosting.ddd.myjino.ru", "myjino.ru"},
	{"landing.myjino.ru", "myjino.ru"},
	{"www.landing.myjino.ru", "www.landing.myjino.ru"},
	{"spectrum.vps.myjino.ru", "spectrum.vps.myjino.ru"},

	// *.uberspace.de
	{"uberspace.de", "de"},
	{"aaa.uberspace.de", "aaa.uberspace.de"},
	{"bbb.ccc.uberspace.de", "ccc.uberspace.de"},

	// .nosuchtld.
	{"nosuchtld", "nosuchtld"},
	{"foo.nosuchtld", "nosuchtld"},
	{"bar.foo.nosuchtld", "nosuchtld"},
}

func TestGetTLD(t *testing.T) {
	for _, tc := range testCases {
		actualTLD := GetTLD(tc.domain)
		if actualTLD != tc.expectedTLD {
			t.Errorf("%q: actual (%q), expected (%q)", tc.domain, actualTLD, tc.expectedTLD)
		}
	}
}
