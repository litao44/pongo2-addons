package addons

import (
	"testing"

	"github.com/flosch/pongo2"
	"github.com/stretchr/testify/assert"
)

func Test_CustomFilter(t *testing.T) {
	as := assert.New(t)

	testFloat := 123456789.12345678
	testTime := "2018-05-02T15:04:05Z"

	data := map[string]interface{}{
		"f": testFloat,
		"t": testTime,
	}

	template := `{{ f|comma }}`
	o, err := execute(template, data)
	as.Nil(err)
	as.Equal("123,456,789", o)

	template = `{{ f|comma:"2" }}`
	o, err = execute(template, data)
	as.Nil(err)
	as.Equal("123,456,789.12", o)

	template = `{{ f|comma:"-1" }}`
	_, err = execute(template, data)
	as.NotNil(err)

	template = `{{ t|date_s:"2006-01-02" }}`
	o, err = execute(template, data)
	as.Nil(err)
	as.Equal("2018-05-02", o)
}

func execute(template string, object map[string]interface{}) (string, error) {
	tpl, err := pongo2.FromString(template)
	if err != nil {
		return "", err
	}

	out, err := tpl.Execute(pongo2.Context(object))
	if err != nil {
		return "", err
	}

	return out, nil
}
