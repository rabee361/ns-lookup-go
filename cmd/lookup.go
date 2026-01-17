package cmd

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/urfave/cli/v3"
)

func Lookup(ctx context.Context, cmd *cli.Command) error {
	domain := "alnoor-hajj.com"
	recordType := "A"
	if cmd.NArg() > 0 {
		domain = cmd.Args().Get(0)
	}

	// Override recordType if flag is set
	if typeFlag := cmd.String("type"); typeFlag != "" {
		recordType = typeFlag
	}

	result, err := QueryDNS(domain, "8.8.8.8", recordType)
	if err != nil {
		return err
	}
	fmt.Print(result)
	return nil
}

func QueryDNS(domain, dnsServerIP string, recordType string) (string, error) {
	dialer := &net.Dialer{
		Timeout:   5 * time.Second,
		LocalAddr: nil,
	}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			return dialer.DialContext(ctx, "udp", dnsServerIP+":53")
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var result string
	var err error

	switch recordType {
	case "A":
		result, err = queryA(ctx, domain, resolver)
	case "CNAME":
		result, err = queryCNAME(ctx, domain, resolver)
	case "NS":
		result, err = queryNS(ctx, domain, resolver)
	case "TXT":
		result, err = queryTXT(ctx, domain, resolver)
	case "MX":
		result, err = queryMX(ctx, domain, resolver)
	default:
		// Default to A record
		result, err = queryA(ctx, domain, resolver)
	}
	return result, err
}

func queryCNAME(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupCNAME(ctx, domain)
	if err != nil {
		return "", err
	}
	return result, err
}

func queryA(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupIPAddr(ctx, domain)
	if err != nil {
		return "", err
	}
	// Return the first IP found
	return result[0].String(), err
}

func queryTXT(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupTXT(ctx, domain)
	if err != nil {
		return "", err
	}
	// Return the first TXT record found
	return result[0], err
}

func queryMX(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupMX(ctx, domain)
	if err != nil {
		return "", err
	}
	// Return the host of the first MX record
	return result[0].Host, err
}

func queryNS(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupNS(ctx, domain)
	if err != nil {
		return "", err
	}
	// Return the host of the first NS record
	return result[0].Host, err
}
