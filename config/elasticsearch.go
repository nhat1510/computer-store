package config

import (
    "log"
    "github.com/elastic/go-elasticsearch/v8"
)

var ES *elasticsearch.Client

func ConnectElasticSearch() {
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        log.Fatalf("❌ Lỗi kết nối ElasticSearch: %s", err)
    }
    ES = es
    log.Println("✅ Kết nối ElasticSearch thành công")
}