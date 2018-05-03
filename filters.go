package addons

import (
	"fmt"
	"time"

	humanize "github.com/dustin/go-humanize"
	"github.com/flosch/pongo2"
)

type Value = pongo2.Value
type Error = pongo2.Error

var AsValue = pongo2.AsValue

func init() {
	pongo2.RegisterFilter("comma", filterComma)
	pongo2.RegisterFilter("date_s", filterStringDateFormat)
}

// {{ vale|comma:"2" }} value 123456.789 -> 123,456.78
func filterComma(in *Value, param *Value) (*Value, *Error) {
	decimals := 0
	if !param.IsNil() {
		decimals = param.Integer()
	}

	if decimals < 0 {
		return nil, &Error{
			Sender:   "filter:comma",
			ErrorMsg: fmt.Sprintf("Filter input argument invalid, must be positive"),
		}
	}

	val := humanize.CommafWithDigits(in.Float(), decimals)
	return AsValue(val), nil
}

func filterStringDateFormat(in *Value, param *Value) (*Value, *Error) {
	val := in.String()

	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, &Error{
			Sender:   "filter:date_s",
			ErrorMsg: fmt.Sprintf("Filter input argument invalid: %v", err.Error()),
		}
	}

	return AsValue(t.Format(param.String())), nil
}
