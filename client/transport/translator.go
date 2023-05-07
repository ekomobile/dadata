package transport

import "io"

type (
	// EncoderFactory creates new request encoder
	EncoderFactory func(w io.Writer) Encoder

	// DecoderFactory creates new response decoder
	DecoderFactory func(r io.Reader) Decoder

	// Encoder encodes request from v
	Encoder func(v interface{}) error

	// Decoder decodes response into v.
	Decoder func(v interface{}) error
)
