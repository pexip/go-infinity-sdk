package infinity_go

import (
	"crypto/tls"
	"crypto/x509"
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

type API struct {
	Log           slog.Logger
	Manager       *Manager
	ManagerClient http.Client
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

func GetManagerResponse(api API, apiPath string) (*http.Response, error) {
	// Create a new HTTP request
	// Use the scheme, host, port, and path from the Manager struct
	req, err := http.NewRequest("GET", fmt.Sprintf("%s://%s:%s%s",
		api.Manager.Scheme,
		api.Manager.Host,
		api.Manager.Port,
		apiPath),
		nil)
	if err != nil {
		api.Log.Error("Error creating request: ", "error", err)
		return nil, err
	}

	if api.Manager.Username != "" && api.Manager.Password != "" {
		// Add the Authorization header
		req.SetBasicAuth(api.Manager.Username, api.Manager.Password)
	}

	req.Header.Add("User-Agent", "Infinitygo/1.0")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := api.ManagerClient.Do(req)
	if err != nil {
		api.Log.Error("Error making request: ", "error", err)
		return nil, err
	}
	return resp, nil
}
