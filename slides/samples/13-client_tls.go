tlsConfig := &tls.Config{
	Certificates: Certificates, // []*x509.Certificate
	RootCAs:      CACertPool,   // *x509.CertPool
}
tlsConfig.BuildNameToCertificate() // Builds certificate name to actual certificate map

transport := &http.Transport{TLSClientConfig: tlsConfig}

client := &http.Client{Transport: transport}

// Use client as usual, e.g.:
resp, err := client.Get("https://host.com/secure")
