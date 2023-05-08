package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

func main() {
	dat_1 := []byte(`{
		"order_uid": "veibfwdepsjdcoef",
		"track_number": "TRAK_1",
		"entry": "UDWO",
		"delivery": {
		  "name": "Name Name_2 Name_3",
		  "phone": "+9812345678",
		  "zip": "2631234",
		  "city": "City",
		  "address": "Ulitsa Imeny Nicogo",
		  "region": "REG-1",
		  "email": "mail_mail_@gmail.com"
		},
		"payment": {
		  "transaction": "svodkododnod",
		  "request_id": "101010",
		  "currency": "RU",
		  "provider": "wbpay",
		  "amount": 2000,
		  "payment_dt": 1637901234,
		  "bank": "bank-a",
		  "delivery_cost": 150,
		  "goods_total": 12,
		  "custom_fee": 2
		},
		"items": [
		  {
			"chrt_id": 9931234,
			"track_number": "TRAK_1",
			"price": 453,
			"rid": "test_test_test",
			"name": "item name",
			"sale": 2000,
			"size": "0",
			"total_price": 400,
			"nm_id": 2381234,
			"brand": "Brand_1",
			"status": 202
		  }
		],
		"locale": "ru",
		"internal_signature": "",
		"customer_id": "testID_1",
		"delivery_service": "meest",
		"shardkey": "2",
		"sm_id": 100,
		"date_created": "2023-05-05T06:22:19Z",
		"oof_shard": "2"
	  }`)

	dat_2 := []byte(`{
		"order_uid": "ord2osdffonfocsnd",
		"track_number": "TRAK_2",
		"entry": "UDWO",
		"delivery": {
		  "name": "Name Name_3 Name_4",
		  "phone": "+9876543210",
		  "zip": "2631234",
		  "city": "City_2",
		  "address": "Ulitsa Imeny Nicogo",
		  "region": "REG-2",
		  "email": "mail_mail_@gmail.com"
		},
		"payment": {
		  "transaction": "svodkododnod",
		  "request_id": "101010",
		  "currency": "RU",
		  "provider": "wbpay",
		  "amount": 2300,
		  "payment_dt": 1637901234,
		  "bank": "bank-b",
		  "delivery_cost": 150,
		  "goods_total": 12,
		  "custom_fee": 1
		},
		"items": [
		  {
			"chrt_id": 9931234,
			"track_number": "TRAK_2",
			"price": 453,
			"rid": "test_test_test",
			"name": "item name_1",
			"sale": 2000,
			"size": "0",
			"total_price": 400,
			"nm_id": 2381234,
			"brand": "Brand_1",
			"status": 202
		  },
		{
			"chrt_id": 839275,
			"track_number": "TRAK_3",
			"price": 453,
			"rid": "test_test_test",
			"name": "item name_2",
			"sale": 2000,
			"size": "0",
			"total_price": 400,
			"nm_id": 2381234,
			"brand": "Brand_1",
			"status": 202
		  }
		],
		"locale": "ru",
		"internal_signature": "",
		"customer_id": "testID_1",
		"delivery_service": "meest",
		"shardkey": "2",
		"sm_id": 100,
		"date_created": "2023-05-05T06:22:19Z",
		"oof_shard": "2"
	  }`)
	sc, err := stan.Connect("test-cluster", "2")
	if err != nil {
		fmt.Print(err)
	}
	defer sc.Close()
	err = sc.Publish("foo", dat_1)
	err = sc.Publish("foo", dat_2)
}
