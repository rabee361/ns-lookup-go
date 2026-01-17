package cmd

import (
	"context"
	"fmt"
	"crypto/tls"
	"time"

	"github.com/urfave/cli/v3"
) 

func CheckSSL(ctx context.Context, cmd *cli.Command) error {
	config := &tls.Config{}
	if cmd.NArg() > 0 {
		domain := cmd.Args().Get(0)
		conn, err := tls.Dial("tcp", domain+":443", config)
		if err != nil {
			return err
		}
		certs := conn.ConnectionState()
		for _, cert := range certs.PeerCertificates {
			fmt.Println("Issuer: " ,cert.Issuer.CommonName)
			fmt.Println("Not After: " ,cert.NotAfter)
			fmt.Println("Not Before: " ,cert.NotBefore)
			fmt.Println("Days until Expired: ", time.Until(cert.NotAfter).Hours() / 24)
			fmt.Println("Signature Algorithm: " ,cert.SignatureAlgorithm)
			if len(cert.Subject.Names) > 1 {
				fmt.Println("Subject: " ,cert.Subject.Names[1].Value)
			}
			fmt.Println("Version: " ,cert.Version)
		}
		defer conn.Close()
		return nil
	}
	return nil
}
