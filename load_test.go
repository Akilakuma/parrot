package parrot

import (
	"fmt"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

// Config 設定檔結構
type Config struct {
	ListenGRPCPort string `config:"API_LISTEN_GRPC"`
	ListenHTTPPort string `config:"API_LISTEN_HTTP_PORT"`
	TargetGRPC     string `config:"TARGET_ADDR"`
	MasterDB       string `config:"DB_READ"`
	SlaveDB        string `config:"DB_WRITE"`
}

func TestReadEnv(t *testing.T) {

	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("import file error")
	}
	var config Config

	c, err := ReadEnv(config)
	if err != nil {
		log.Println("read error:", err)
	}

	fmt.Printf("%#v", c)

	// result
	// `parrot.Config{
	// 	ListenGRPCPort:":8801",
	// 	ListenHTTPPort:":8080",
	// 	TargetGRPC:"127.0.0.1:8802",
	// 	MasterDB:"root:abd123@tcp(127.0.0.1:3306)/finance?charset=utf8&parseTime=True&loc=Local",
	// 	SlaveDB:"root:abd123@tcp(127.0.0.1:3306)/finance?charset=utf8&parseTime=True&loc=Local"
	// }`
}
