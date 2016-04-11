package yaml

import (
	"io"

	"github.com/wlMalk/mirango/framework"
	"gopkg.in/yaml.v2"
)

type Renderer struct{}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) MimeType() string {
	return framework.MIME_YAML
}

func (r *Renderer) RenderToWriter(wr io.Writer, c framework.Context, data interface{}) error {
	b, err := r.Render(c, data)
	if err != nil {
		return err
	}
	_, err = wr.Write(b)
	return err
}

func (r *Renderer) Render(c framework.Context, data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}
