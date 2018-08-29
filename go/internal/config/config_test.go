package config

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestConfigLogFile(t *testing.T) {
	testRoot := "/foo/bar"
	testCases := []struct {
		yaml   string
		path   string
		format string
	}{
		{path: "/foo/bar/gitlab-shell.log", format: "text"},
		{yaml: "log_file: my-log.log", path: "/foo/bar/my-log.log", format: "text"},
		{yaml: "log_file: /qux/my-log.log", path: "/qux/my-log.log", format: "text"},
		{yaml: "log_format: json", path: "/foo/bar/gitlab-shell.log", format: "json"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("yaml input: %q", tc.yaml), func(t *testing.T) {
			cfg := Config{RootDir: testRoot}
			if err := parseConfig([]byte(tc.yaml), &cfg); err != nil {
				t.Fatal(err)
			}

			if cfg.LogFile != tc.path {
				t.Fatalf("expected %q, got %q", tc.path, cfg.LogFile)
			}

			if cfg.LogFormat != tc.format {
				t.Fatalf("expected %q, got %q", tc.format, cfg.LogFormat)
			}
		})
	}
}

func TestExampleFile(t *testing.T) {
	content, err := ioutil.ReadFile("./testdata/config.yml.example")
	if err != nil {
		t.Fatal(err)
	}

	cfg := Config{}

	err = parseConfig(content, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	expected := Config{
		RootDir:        "",
		LogFile:        "gitlab-shell.log",
		LogFormat:      "text",
		GitlabUrl:      "http://localhost:8080",
		InternalApiUrl: "http://localhost:8080/api/v4/internal",
	}

	if cfg != expected {
		t.Fatalf("expected %v got %v", expected, cfg)
	}
}
