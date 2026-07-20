package blockchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
)

func SignTx(priv *ecdsa.PrivateKey, tx *Transaction)error{
	h:=TxSiginingHash(tx)
	r, s,err :=ecdsa.Sign(rand.Reader, priv, h)
	if err != nil {return err}
	tx.SigR = r.Bytes()
	tx.SigS = s.Bytes()
	tx.PubKey = elliptic.Marshal(priv.Curve, priv.PublicKey.X,priv.PublicKey.Y)
	return nil
}

func VerifyTx(tx *Transaction)bool{
	h:=TxSiginingHash(tx)
	x,y :=elliptic.Unmarshal(elliptic.P256(),tx.PubKey)
	pub := &ecdsa.PublicKey{Curve:elliptic.P256(),X:x,Y:y}
	r := new(big.Int).SetBytes(tx.SigR)
	s := new(big.Int).SetBytes(tx.SigS)
	return ecdsa.Verify(pub,h,r,s)
}

func TxSiginingHash(tx *Transaction)[]byte {
	s := fmt.Sprintf("%s|%s|%d|%d",
tx.Sender, tx.Recipient, tx.Amount,tx.Nonce)
	h:=sha256.Sum256([]byte(s))
	return h[:]
}

func TxLeafHash(tx Transaction) []byte{
	s := fmt.Sprintf("%s|%s|%d|%d|%x|%x|%x",tx.Sender,tx.Recipient,tx.Amount,tx.Nonce,tx.PubKey,tx.SigR,tx.SigS)
	h:=sha256.Sum256([]byte(s))
	return h[:]
}
	


		func LoadPrivKey(path string)(*ecdsa.PrivateKey, error){
			data, err := os.ReadFile(path)
		if err != nil{
			return nil, err
		}
		var obj struct {PrivateKeyHex string}
		if err := json.Unmarshal(data,&obj); err !=nil{
			return nil, err
		}
		d, ok := new(big.Int).SetString(obj.PrivateKeyHex, 16)
		if !ok {
			return nil, fmt.Errorf("invalid hex")
		}

		priv := new(ecdsa.PrivateKey)
		priv.PublicKey.Curve = elliptic.P256()
		priv.D = d
		priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(d.Bytes())
		return priv, nil
	}