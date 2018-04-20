package metadata

import (
	"crypto"
	"encoding/hex"
	"time"

	"github.com/icstglobal/go-icst/content"
)

//ContentMetadata represents metadata of a content
type ContentMetadata struct {
	Hash  string `json:"hash,required"`
	Date  int64  `json:"date,required"`
	Owner UserID `json:"owner,required"`
}

// UserID identify a user by its public key
type UserID struct {
	PublicKey crypto.PublicKey `json:"public_key,required"`
	//crypto algorithm to apply the publick key, like AES, ecdsa
	Crypto string `json:"crypto,required"`
}

//FromContent creates ContentMetadata instance from the content
func FromContent(c *content.Content) ContentMetadata {
	return ContentMetadata{
		Hash:  hex.EncodeToString(c.Hash),
		Date:  time.Now().Unix(),
		Owner: UserID{c.Owner.PrivateKey.Public(), "ecdsa"},
	}
}