package model

type AmazonTokenCred struct {
	AccessKeyId  string `json:"access_key_id"`
	SecretKeyId  string `json:"secret_key_id"`
	AssociateTag string `json:"associate_tag"`
}
