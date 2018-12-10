package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"

	"github.com/rs/xid"
	"google.golang.org/grpc"
)

func GetXuid() string {
	id := xid.New()
	return id.String()
}
func GetGRPCConn() *grpc.ClientConn {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	return conn
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
