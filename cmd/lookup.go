package cmd

import (
    "fmt"
    "context"
	"net"
	"time"

    "github.com/urfave/cli/v3"
)


func Lookup(ctx context.Context, cmd *cli.Command) error {
		domain := "alnoor-hajj.com"
		record_type := "A"
		if cmd.NArg() > 0 {
			domain = cmd.Args().Get(0)
		}
		if cmd.String("type") == "TXT" {
			record_type = "TXT"
		}
		if cmd.String("type") == "MX" {
			record_type = "MX"
		}
		if cmd.String("type") == "NS" {
			record_type = "NS"
		}
		if cmd.String("type") == "CNAME" {
			record_type = "CNAME"
		}
		result ,err := QueryDNS(domain, "8.8.8.8", record_type)

		if err != nil {
			fmt.Print("Error")
			return err
		}
		fmt.Printf("the ip for %s is: %s", domain, result)
		return nil
	}

func QueryDNS(domain, dnsServerIP string, record_type string) (string, error) {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
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

	result := ""
	var err error

	if record_type == "A" {
		result, err = queryA(ctx, domain, resolver)
	} else if record_type == "CNAME" {
		result, err = queryCNAME(ctx, domain, resolver)
	} else if record_type == "NS" {
		result, err = queryNS(ctx, domain, resolver)
	} else if record_type == "TXT" {
		result, err = queryTXT(ctx, domain, resolver)
	} else if record_type == "MX" {
		result, err = queryMX(ctx, domain, resolver)
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
	return result[0].String(), err
}

func queryTXT(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupTXT(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0], err
}

func queryMX(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	result, err := resolver.LookupMX(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0].Host, err
}

func queryNS(ctx context.Context, domain string, resolver *net.Resolver) (string, error) {
	
	result, err := resolver.LookupNS(ctx, domain)
	if err != nil {
		return "", err
	}
	return result[0].Host, err
}