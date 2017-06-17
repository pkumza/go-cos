package cos

import (
	"context"
	"encoding/xml"
	"net/http"
)

// BucketCORSRule ...
type BucketCORSRule struct {
	ID             string   `xml:"ID,omitempty"`
	AllowedMethods []string `xml:"AllowedMethod"`
	AllowedOrigins []string `xml:"AllowedOrigin"`
	AllowedHeaders []string `xml:"AllowedHeader,omitempty"`
	MaxAgeSeconds  int      `xml:"MaxAgeSeconds,omitempty"`
	ExposeHeaders  []string `xml:"ExposeHeader,omitempty"`
}

// BucketGetCORSResult ...
type BucketGetCORSResult struct {
	XMLName xml.Name          `xml:"CORSConfiguration"`
	Rules   []*BucketCORSRule `xml:"CORSRule,omitempty"`
}

// GetCORS ...
//
// Get Bucket CORS实现跨域访问配置读取。
//
// https://www.qcloud.com/document/product/436/8274
func (s *BucketService) GetCORS(ctx context.Context,
	authTime *AuthTime) (*BucketGetCORSResult, *Response, error) {
	var res BucketGetCORSResult
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.BucketURL,
		uri:      "/?cors",
		method:   http.MethodGet,
		authTime: authTime,
		result:   &res,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return &res, resp, err
}

// BucketPutCORSOptions ...
type BucketPutCORSOptions struct {
	XMLName xml.Name          `xml:"CORSConfiguration"`
	Rules   []*BucketCORSRule `xml:"CORSRule,omitempty"`
}

// PutCORS ...
//
// Put Bucket CORS实现跨域访问设置，您可以通过传入XML格式的配置文件实现配置，文件大小限制为64 KB。
//
// https://www.qcloud.com/document/product/436/8279
func (s *BucketService) PutCORS(ctx context.Context,
	authTime *AuthTime, opt *BucketPutCORSOptions) (*Response, error) {
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.BucketURL,
		uri:      "/?cors",
		method:   http.MethodPut,
		authTime: authTime,
		body:     opt,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return resp, err
}

// DeleteCORS ...
//
// Delete Bucket CORS实现跨域访问配置删除。
//
// https://www.qcloud.com/document/product/436/8283
func (s *BucketService) DeleteCORS(ctx context.Context,
	authTime *AuthTime) (*Response, error) {
	sendOpt := sendOptions{
		baseURL:  s.client.BaseURL.BucketURL,
		uri:      "/?cors",
		method:   http.MethodDelete,
		authTime: authTime,
	}
	resp, err := s.client.send(ctx, &sendOpt)
	return resp, err
}
