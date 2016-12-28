package transport

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/mitch000001/go-hbci/segment"
)

type Transport interface {
	Do(*Request) (*Response, error)
}

type TransportFunc func(*Request) (*Response, error)

func (fn TransportFunc) Do(req *Request) (*Response, error) {
	return fn(req)
}

type Middleware func(Transport) Transport

type Request struct {
	URL              string
	MarshaledMessage []byte
}

func ReadResponse(marshaledMessage []byte, request *Request) (*Response, error) {
	extractor := segment.NewSegmentExtractor(marshaledMessage)
	_, err := extractor.Extract()
	if err != nil {
		return nil, err
	}
	response := &Response{
		Request:           request,
		MarshaledResponse: marshaledMessage,
		SegmentExtractor:  extractor,
	}
	return response, nil
}

type Response struct {
	*segment.SegmentExtractor
	Request           *Request
	MarshaledResponse []byte
}

func (h *Response) IsEncrypted() bool {
	return h.FindSegment("HNVSK") != nil
}
