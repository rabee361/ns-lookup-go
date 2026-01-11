package main

import (
	"bytes"
	"context"
	"io"
	"ns/cmd"
	"os"
	"strings"
	"testing"

	"github.com/urfave/cli/v3"
)

func TestQueryDNS(t *testing.T) {
	// Test A record
	ip, err := cmd.QueryDNS("google.com", "8.8.8.8", "A")
	if err != nil {
		t.Errorf("QueryDNS failed for A record: %v", err)
	}
	if ip == "" {
		t.Error("Expected IP address for A record, got empty string")
	}

	// Test MX record
	mx, err := cmd.QueryDNS("google.com", "8.8.8.8", "MX")
	if err != nil {
		t.Errorf("QueryDNS failed for MX record: %v", err)
	}
	if mx == "" {
		t.Error("Expected MX record, got empty string")
	}
}

func TestLookupCommand(t *testing.T) {
	app := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:   "lookup",
				Action: cmd.Lookup,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "type",
						Value: "A",
					},
				},
			},
		},
	}

	// Capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run command
	// Note: We use "app" as the first argument as checking arg[0] is often skipped or treated as program name
	err := app.Run(context.Background(), []string{"pulse", "lookup", "--type", "A", "example.com"})

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "the ip for example.com is:") {
		t.Errorf("Unexpected output: %s", output)
	}
}

func TestScanPortCommand(t *testing.T) {
	app := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:   "scan",
				Action: cmd.ScanPort,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "port",
						Value: "all",
					},
				},
			},
		},
	}

	// Test specific port (80) on google.com which should be open
	// Capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := app.Run(context.Background(), []string{"pulse", "scan", "--port", "80", "google.com"})

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("Scan Command failed: %v", err)
	}

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// Either "Port open: 80" or check if it actually connected
	// Since google.com:80 should be open, we expect success.
	if !strings.Contains(output, "Port open: 80") {
		// Fallback check in case of network issues, but ideally it should work.
		// If it says "Port closed", the test fails, which is correct for google.com:80
		t.Errorf("Expected port 80 to be open on google.com, output: %s", output)
	}
}
