package constants

import "sync"

const (
	ERR_UPLOAD_FAILED  = -401
	ERR_DECODE_CONTENT = -402
)

var once sync.Once

func init() {
	once.Do(Init)
}

func Init() {
	codeErr[ERR_UPLOAD_FAILED] = "upload err"
	codeErr[ERR_DECODE_CONTENT] = "decode data error"
}
