import (
    "encoding/pem"
    "encoding/x509"
    "errors"
    "fmt"
    "time"
)

// DecodeCreator decodes the byte array representation of the creator into an x509 certificate.
func DecodeCreator(creatorBytes []byte) (*x509.Certificate, error) {
    // Decode the PEM-encoded certificate
    block, _ := pem.Decode(creatorBytes)
    if block == nil {
        return nil, errors.New("failed to decode PEM block")
    }

    // Parse the certificate
    cert, err := x509.ParseCertificate(block.Bytes)
    if err != nil {
        return nil, err
    }

    return cert, nil
}

// ExtractPublicKey extracts the public key from the certificate.
func ExtractPublicKey(cert *x509.Certificate) interface{} {
    return cert.PublicKey
}

// ExtractCertificateAttributes extracts the attributes from the certificate.
func ExtractCertificateAttributes(cert *x509.Certificate) map[string]string {
    attributes := make(map[string]string)
    for _, attr := range cert.Subject.Organization {
        attributes["Organization"] = attr
    }
    // Add more attribute fields as needed
    return attributes
}

// ExtractIssuerInformation extracts information about the issuer from the certificate.
func ExtractIssuerInformation(cert *x509.Certificate) string {
    return cert.Issuer.String()
}

// ExtractValidityPeriod extracts the validity period (start and end dates) from the certificate.
func ExtractValidityPeriod(cert *x509.Certificate) (time.Time, time.Time) {
    return cert.NotBefore, cert.NotAfter
}

// ExtractSignature extracts the digital signature from the certificate.
func ExtractSignature(cert *x509.Certificate) []byte {
    return cert.Signature
}

func main() {
    // Sample creator byte array
    creatorBytes := []byte("-----BEGIN CERTIFICATE-----\nMIIBIjCB...-----END CERTIFICATE-----\n")

    // Decode the creator
    cert, err := DecodeCreator(creatorBytes)
    if err != nil {
        fmt.Println("Error decoding creator:", err)
        return
    }


    // Certificate got from the DecodeCreator() can have these properties
    // Print the output (decoded certificate)
    fmt.Println("Output (decoded certificate):")
    fmt.Printf("Subject: %s\n", cert.Subject)
    fmt.Printf("Issuer: %s\n", cert.Issuer)
    fmt.Printf("Serial Number: %s\n", cert.SerialNumber)
    fmt.Printf("Not Before: %s\n", cert.NotBefore)
    fmt.Printf("Not After: %s\n", cert.NotAfter)
    // Add more fields as needed


    
    // Extract and print each type of information
    publicKey := ExtractPublicKey(cert)
    fmt.Println("Public Key:", publicKey)

    attributes := ExtractCertificateAttributes(cert)
    fmt.Println("Certificate Attributes:", attributes)

    issuerInfo := ExtractIssuerInformation(cert)
    fmt.Println("Issuer Information:", issuerInfo)

    validityStart, validityEnd := ExtractValidityPeriod(cert)
    fmt.Println("Validity Start:", validityStart)
    fmt.Println("Validity End:", validityEnd)

    signature := ExtractSignature(cert)
    fmt.Println("Signature:", signature)
}
