package module

import (
	"fmt"
	"github.com/alex-1900/convert-to/config"
	"github.com/alex-1900/convert-to/structs/profile/clash"
	"github.com/alex-1900/convert-to/structs/resolver"
)

type Converter struct {
	r     *Reader
	from  config.Resolver
	toStr string
}

func NewConverterFromUrl(url, from, to string) *Converter {
	r := NewReader(url, nil)
	return NewConverterFromReader(r, from, to)
}

func NewConverterFromContent(content, from, to string) *Converter {
	r := NewReader("", []byte(content))
	return NewConverterFromReader(r, from, to)
}

func NewConverterFromReader(r *Reader, from, to string) *Converter {
	var rs config.Resolver = nil
	switch from {
	case string(config.ProIDClash):
		rs = resolver.NewClash(new(clash.Clash))
		break
	default:
		return nil
	}

	return &Converter{r, rs, to}
}

func (c *Converter) Convert() (config.Profile, error) {
	var resolveFunc func() (config.Profile, error) = nil

	switch c.toStr {
	case string(config.ProIDClash):
		resolveFunc = c.from.ToClash
		break
	case string(config.ProIDSingBox):
		resolveFunc = c.from.ToSingBox
		break
	}

	if resolveFunc == nil {
		return nil, fmt.Errorf("no resolver for %s", c.toStr)
	}

	data, err := c.r.GetContent()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	err = c.from.Resolve(data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	prof, err1 := resolveFunc()
	if err1 != nil {
		fmt.Println(err)
		return nil, err
	}

	return prof, nil
}
