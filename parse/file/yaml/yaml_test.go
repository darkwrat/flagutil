package yaml

import (
	"context"
	"testing"

	yaml "gopkg.in/yaml.v2"

	"github.com/gobwas/flagutil/parse"
	"github.com/gobwas/flagutil/parse/file"
	"github.com/gobwas/flagutil/parse/testutil"
)

func TestYAML(t *testing.T) {
	testutil.TestParser(t, func(values testutil.Values, fs parse.FlagSet) error {
		p := file.Parser{
			Lookup: file.BytesLookup(marshal(values)),
			Syntax: new(Syntax),
		}
		return p.Parse(context.Background(), fs)
	})
}

func marshal(values testutil.Values) []byte {
	bts, err := yaml.Marshal(values)
	if err != nil {
		panic(err)
	}
	return bts
}
