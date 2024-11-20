package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aelpxy/fresh/middlewares"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/acme/autocert"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP file server",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		folder, _ := cmd.Flags().GetString("folder")
		port, _ := cmd.Flags().GetString("port")
		domain, _ := cmd.Flags().GetString("domain")
		tls, _ := cmd.Flags().GetBool("tls")
		certDir, _ := cmd.Flags().GetString("cert-dir")

		if folder == "" {
			folder = "./"
		}

		logger := log.NewWithOptions(os.Stderr, log.Options{
			ReportCaller:    false,
			ReportTimestamp: true,
			TimeFormat:      time.Kitchen,
			Prefix:          "fresh",
		})

		fileServer := http.StripPrefix("/", http.FileServer(http.Dir(folder)))
		http.Handle("/", middlewares.LogRequest(fileServer))

		if domain != "" && tls && port == "3000" {
			port = "443"
		}

		address := domain
		if domain == "" {
			address = fmt.Sprintf(":%s", port)
		} else {
			address = fmt.Sprintf("%s:%s", domain, port)
		}

		if tls {
			if certDir == "" {
				certDir = "./certs"
			}

			if _, err := os.Stat(certDir); os.IsNotExist(err) {
				if err := os.MkdirAll(certDir, 0755); err != nil {
					logger.Fatal("failed to create certificate directory", "error", err)
				}
			}

			certManager := &autocert.Manager{
				Cache:      autocert.DirCache(certDir),
				Prompt:     autocert.AcceptTOS,
				HostPolicy: autocert.HostWhitelist(domain),
			}

			server := &http.Server{
				Addr:      address,
				Handler:   nil,
				TLSConfig: certManager.TLSConfig(),
			}

			logger.Info("launching secure server", "protocol", "https", "address", address, "cert", "Let's Encrypt")
			if err := server.ListenAndServeTLS("", ""); err != nil {
				logger.Fatal("server crashed", "error", err)
			}
		} else {
			logger.Info("launching server", "protocol", "http", "address", address)
			if err := http.ListenAndServe(address, nil); err != nil {
				logger.Fatal("server crashed", "error", err)
			}
		}
	},
}

func init() {
	serveCmd.Flags().StringP("folder", "f", "./", "Directory to serve files from")
	serveCmd.Flags().StringP("port", "p", "3000", "Port to serve the HTTP server on")
	serveCmd.Flags().StringP("domain", "d", "", "Domain name for the HTTP server")
	serveCmd.Flags().Bool("tls", false, "Enable TLS for HTTPS server")
	serveCmd.Flags().String("cert-dir", "./certs", "Directory to store Let's Encrypt certificates")

	rootCmd.AddCommand(serveCmd)
}
