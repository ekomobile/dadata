package encoder

import (
	"fmt"
	"io"

	"github.com/ekomobile/dadata/v2/client/transport"
)

// RawEncoderFactory is a factory for noop-encoder that just writes to transport io.Writer.
func RawEncoderFactory() transport.EncoderFactory {
	return func(w io.Writer) transport.Encoder {
		return func(v interface{}) (err error) {
			switch value := v.(type) {
			case *string:
				_, err = w.Write([]byte(*value))
			case string:
				_, err = w.Write([]byte(value))
			case *[]byte:
				_, err = w.Write(*value)
			case []byte:
				_, err = w.Write(value)
			case io.Reader:
				_, err = io.Copy(w, value)
			case io.WriterTo:
				_, err = value.WriteTo(w)
			default:
				err = fmt.Errorf("can not encode from value: %#v", v)
			}
			return err
		}
	}
}

// RawDecoderFactory is a factory for noop-decoder that just reads from transport io.Reader.
func RawDecoderFactory() transport.DecoderFactory {
	return func(r io.Reader) transport.Decoder {
		return func(v interface{}) (err error) {
			switch value := v.(type) {
			case *[]byte:
				_, err = r.Read(*value)
			case []byte:
				_, err = r.Read(value)
			case io.Writer:
				_, err = io.Copy(value, r)
			case io.ReaderFrom:
				_, err = value.ReadFrom(r)
			default:
				err = fmt.Errorf("can not decode into value: %#v", v)
			}
			return err
		}
	}
}
