package yaml

import (
	"bytes"
	"io"

	"github.com/mirango/framework"
	"gopkg.in/yaml.v2"
)

type YAML struct{}

func New() *YAML {
	return &YAML{}
}

func (YAML) Encode(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (YAML) EncodeTo(w io.Writer, v interface{}) error {
	b, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}

func (YAML) EncodingType() string {
	return framework.MIME_YAML
}

func (YAML) Decode(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (YAML) DecodeFrom(r io.Reader, v interface{}) error {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(buf.Bytes(), v)
}

func (YAML) DecodingType() string {
	return framework.MIME_YAML
}

func (YAML) MIMEType() string {
	return framework.MIME_YAML
}
