package   包 utils

import (   导入(
    "crypto/ecdsa"
    "crypto/elliptic"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "math/big"   “数学/大”
)

func GenerateKeyPair() (*ecdsa.PrivateKey, string) {
    privKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
    pubKey := append(privKey.PublicKey.X.Bytes(), privKey.PublicKey.Y.Bytes()...)
    hash := sha256.Sum256(pubKey)
    address := hex.EncodeToString(hash[len(hash)-20:])
    return privKey, "0x" + address
}

func SignMessage(priv *ecdsa.PrivateKey, msg []byte) (r, s *big.Int, err error) {
    hash := sha256.Sum256(msg)
    return ecdsa.Sign(rand.Reader, priv, hash[:])
}

func VerifySignature(pub *ecdsa.PublicKey, msg []byte, r, s *big.Int) bool {
    hash := sha256.Sum256(msg)
    return ecdsa.Verify(pub, hash[:], r, s)
}


