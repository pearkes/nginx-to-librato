package main

import (
	"github.com/pearkes/goconfig/config"
	"os"
	"testing"
	"time"
)

func createTestConfig() *config.Config {
	os.Mkdir("tmp", 0777)
	c := config.NewDefault()
	c.AddSection("settings")
	c.AddOption("settings", "token", "1329gj328v2n9bu2")
	c.AddOption("settings", "email", "jackpearkes@gmail.com")
	c.AddOption("settings", "source", "test")
	c.AddOption("settings", "url", "localhost:8000/nginx_status")
	c.AddOption("settings", "flush_interval", "1s")
	return c
}

func TestConfig_TestNewConf_Good(t *testing.T) {
	// Create the test configuration
	c := createTestConfig()
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	con, errs := NewConf("tmp/test.conf")
	if errs != nil {
		t.Fatalf("should not have err: %v", errs)
	}

	if con.libSource != "test" {
		t.Fatalf("source does not match:\ngot:\n%s\nexpected:\n%s\n", con.libSource, "test")
	}

	if con.libToken != "1329gj328v2n9bu2" {
		t.Fatalf("token does not match:\ngot:\n%s\nexpected:\n%s\n", con.libToken, "1329gj328v2n9bu2")
	}

	if con.libUser != "jackpearkes@gmail.com" {
		t.Fatalf("email does not match:\ngot:\n%s\nexpected:\n%s\n", con.libUser, "jackpearkes@gmail.com")
	}

	if con.url != "localhost:8000/nginx_status" {
		t.Fatalf("url does not match:\ngot:\n%s\nexpected:\n%s\n", con.url, "localhost:8000/nginx_status")
	}

	if con.rawFlushInterval != "1s" {
		t.Fatalf("flush_interval does not match:\ngot:\n%s\nexpected:\n%s\n", con.rawFlushInterval, "1s")
	}

	if con.flushInterval != time.Duration(1*time.Second) {
		t.Fatalf("flush_interval does not match:\ngot:\n%v\nexpected:\n%v\n", con.flushInterval, time.Duration(1*time.Second))
	}
}

func TestConfig_TestNewConf_Bad_Token(t *testing.T) {
	c := createTestConfig()
	c.RemoveOption("settings", "token")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 1 {
		t.Fatalf("should have err: %v", errs)
	}
}

func TestConfig_TestNewConf_Bad_Email(t *testing.T) {
	c := createTestConfig()
	c.RemoveOption("settings", "email")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 1 {
		t.Fatalf("should have err: %v", errs)
	}
	os.RemoveAll("tmp")
}

func TestConfig_TestNewConf_Bad_Source(t *testing.T) {
	c := createTestConfig()
	c.RemoveOption("settings", "source")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 1 {
		t.Fatalf("should have err: %v", errs)
	}
	os.RemoveAll("tmp")
}

func TestConfig_TestNewConf_Bad_Url(t *testing.T) {
	c := createTestConfig()
	c.RemoveOption("settings", "url")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 1 {
		t.Fatalf("should have err: %v", errs)
	}
	os.RemoveAll("tmp")
}

func TestConfig_TestNewConf_Bad_FlushInterval(t *testing.T) {
	c := createTestConfig()
	c.RemoveOption("settings", "flush_interval")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 2 {
		t.Fatalf("should have err: %v", errs)
	}
	os.RemoveAll("tmp")
}

func TestConfig_TestNewConf_Bad_FlushInterval_Format(t *testing.T) {
	c := createTestConfig()
	c.AddOption("settings", "flush_interval", "3i2gj32")
	c.WriteFile("tmp/test.conf", 0644, "Test configuration header")

	_, errs := NewConf("tmp/test.conf")
	if len(errs) != 1 {
		t.Fatalf("should have err: %v", errs)
	}
	os.RemoveAll("tmp")
}
