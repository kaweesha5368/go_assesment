package blockchain_test

import(
	"testing"
	"go_assesment/internal/blockchain"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/elliptic"
)

func TestSignAndVerifyTx(t *testing.T){
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil{
		t.Fatalf("failed to generate key %v", err)
	}

	tx := &blockchain.Transaction{
		Sender: "alice",
		Recipient: "bob",
		Amount: 50,
		Nonce: 1,
	}

	if err := blockchain.SignTx(priv, tx); err != nil{
		t.Fatalf("failed to sign tx : %v", err)
	}

	if !blockchain.VerifyTx(tx){
		t.Errorf("expected signature to be valid, but got invalis")
	}
}