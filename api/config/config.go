package config

import (
	"log"

	"github.com/go-ini/ini"
)

type Configlist struct {
	ApiKey string
}

//読み込んだAPIキー情報を格納する構造体

var Config Configlist

//構造体を変数に代入

func init() {
	cfg, err := ini.Load("/Users/mori-ma/go/src/profile-you/api/config/config.ini") //config.ini内の情報を読み込んで変数に代入
	if err != nil {
		log.Fatalf("failed to read file\n", err)
	}
	//エラーハンドリング
	Config = Configlist{
		ApiKey: cfg.Section("OPEN_API_KEY").Key("api_key").String(),
		//構造体のApiKeyフィールドにOPEN_API_KEYセクションのapi_keyの値(APIキー)を入れる
	}
}

// 1222 削除予定
