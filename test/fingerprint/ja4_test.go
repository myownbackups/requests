package main

import (
	"log"
	"net/textproto"
	"slices"
	"testing"

	"github.com/gospider007/requests"
)

func TestOrderHeaders(t *testing.T) {
	orderKeys := []string{
		"Accept-Encoding",
		"Accept",
		"Sec-Ch-Ua-Mobile",
		"Sec-Ch-Ua-Platform",
	}

	resp, err := requests.Get(nil, "https://tools.scrapfly.io/api/fp/anything", requests.RequestOption{
		ClientOption: requests.ClientOption{
			OrderHeaders: orderKeys,
			// Headers: headers,
		},
		// ForceHttp1: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := resp.Json()
	if err != nil {
		t.Fatal(err)
	}
	header_order := jsonData.Find("ordered_headers_key")
	if !header_order.Exists() {
		t.Fatal("not found akamai")
	}
	i := -1
	for _, key := range header_order.Array() {
		// log.Print(key)
		kk := textproto.CanonicalMIMEHeaderKey(key.String())
		if slices.Contains(orderKeys, kk) {
			i2 := slices.Index(orderKeys, textproto.CanonicalMIMEHeaderKey(kk))
			if i2 < i {
				// log.Print(header_order)
				t.Fatal("not equal")
			}
			i = i2
		}
	}
}
func TestOrderHeaders2(t *testing.T) {

	headers := map[string]any{
		"Accept-Encoding":    "gzip, deflate, br",
		"Accept":             "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Sec-Ch-Ua-Mobile":   "?0",
		"Sec-Ch-Ua-Platform": `"Windows"`,
	}
	orderHeaders := []string{
		"Accept-Encoding",
		"Accept",
		"User-Agent",
		"Accept-Language",
		"Sec-Ch-Ua",
		"Sec-Ch-Ua-Mobile",
		"Sec-Ch-Ua-Platform",
	}
	resp, err := requests.Get(nil, "https://tools.scrapfly.io/api/fp/anything", requests.RequestOption{
		ClientOption: requests.ClientOption{

			Headers:      headers,
			OrderHeaders: orderHeaders,
		},
		// ForceHttp1: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	jsonData, err := resp.Json()
	header_order := jsonData.Find("ordered_headers_key")
	if !header_order.Exists() {
		t.Fatal("not found akamai")
	}
	i := -1
	log.Print(header_order)
	// log.Print(headers.Keys())
	kks := []string{}
	for _, kk := range orderHeaders {
		kks = append(kks, textproto.CanonicalMIMEHeaderKey(kk))
	}
	for _, key := range header_order.Array() {
		kk := textproto.CanonicalMIMEHeaderKey(key.String())
		if slices.Contains(kks, kk) {
			i2 := slices.Index(kks, textproto.CanonicalMIMEHeaderKey(kk))
			if i2 < i {
				log.Print(header_order)
				t.Fatal("not equal")
			}
			i = i2
		}
	}
}
