package mirrors_test

import (
	"testing"

	"github.com/farbodahm/lets-go/fastestMirrorFinder/pkg/mirrors"
)

func TestGettingMirrors(t *testing.T) {
	expected := []string{
		"http://ftp.am.debian.org/debian", "http://ftp.au.debian.org/debian",
		"http://ftp.at.debian.org/debian", "http://ftp.by.debian.org/debian",
		"http://ftp.be.debian.org/debian", "http://ftp.br.debian.org/debian",
		"http://ftp.ca.debian.org/debian", "http://ftp.cn.debian.org/debian",
		"http://ftp.hr.debian.org/debian", "http://ftp.cz.debian.org/debian",
		"http://ftp.dk.debian.org/debian", "http://ftp.sv.debian.org/debian",
		"http://ftp.ee.debian.org/debian", "http://ftp.fr.debian.org/debian",
		"http://ftp2.de.debian.org/debian", "http://ftp.de.debian.org/debian",
		"http://ftp.no.debian.org/debian", "http://ftp.hk.debian.org/debian",
		"http://ftp.hu.debian.org/debian", "http://ftp.is.debian.org/debian",
		"http://ftp.it.debian.org/debian", "http://ftp.jp.debian.org/debian",
		"http://ftp.kr.debian.org/debian", "http://ftp.lt.debian.org/debian",
		"http://ftp.md.debian.org/debian", "http://ftp.nc.debian.org/debian",
		"http://ftp.nl.debian.org/debian", "http://ftp.nz.debian.org/debian",
		"http://ftp.pl.debian.org/debian", "http://ftp.pt.debian.org/debian",
		"http://ftp.ru.debian.org/debian", "http://ftp.sk.debian.org/debian",
		"http://ftp.si.debian.org/debian", "http://ftp.es.debian.org/debian",
		"http://ftp.ch.debian.org/debian", "http://ftp.tw.debian.org/debian",
		"http://ftp.tr.debian.org/debian", "http://ftp.fi.debian.org/debian",
		"http://ftp.se.debian.org/debian",
	}

	got := mirrors.GetMirrorsList()

	if len(got) != len(expected) {
		t.Error(
			"Len(expected) is", len(expected),
			"but len(got) is", len(got),
		)
	}

	expectedMap := make(map[string]bool)
	gotMap := make(map[string]bool)
	for _, url := range expected {
		expectedMap[url] = true
	}
	for _, url := range got {
		gotMap[url] = true
	}

	for url := range gotMap {
		if _, err := expectedMap[url]; !err {
			t.Error("got unexpected mirror", url)
		}
	}
}
