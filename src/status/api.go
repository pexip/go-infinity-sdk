package infinity_go

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Manager struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Scheme   string `yaml:"scheme"`
}

func CreateTLSClient(log slog.Logger) *http.Client {
	pool := x509.NewCertPool()
	certFile := "./ca-certificates.crt"
	fi, err := os.ReadFile(certFile)
	if err != nil {
		log.Warn(fmt.Sprintf("Could not open %s for reading CAs", certFile))
		return nil
	}
	ok := pool.AppendCertsFromPEM(fi)
	if !ok {
		log.Warn("Certificates were not parsed correctly")
		return nil
	}
	client := &http.Client{
		Timeout: 20 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{RootCAs: pool},
		},
	}
	return client
}

func basicAuth(username string, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func NewPexipClient(client *http.Client, log slog.Logger, m *Manager, apiPath string) (*http.Response, error) {
	// Create a new HTTP request
	// Use the scheme, host, port, and path from the Manager struct
	req, err := http.NewRequest("GET", fmt.Sprintf("%s://%s:%s%s", m.Scheme, m.Host, m.Port, apiPath), nil)
	if err != nil {
		log.Error("Error creating request: ", "error", err)
		return nil, err
	}

	if m.Username != "" && m.Password != "" {
		// Add the Authorization header
		req.Header.Add("Authorization", "Basic "+basicAuth(m.Username, m.Password))
	}

	req.Header.Add("User-Agent", "Pexip Metric Exporter")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Error("Error making request: ", "error", err)
		return nil, err
	}
	return resp, nil
}
