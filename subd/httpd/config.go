package httpd

import (
	"bytes"
	"fmt"
	"strconv"
	"text/tabwriter"

	"github.com/megamsys/libgo/cmd"
)

type Config struct {
	Enabled     bool   `toml:"enabled"`
	BindAddress string `toml:"bind_address"`
}

func (c Config) String() string {
	w := new(tabwriter.Writer)
	var b bytes.Buffer
	w.Init(&b, 0, 8, 0, '\t', 0)
	b.Write([]byte(cmd.Colorfy("Config:", "white", "", "bold") + "\t" +
		cmd.Colorfy("httpd", "green", "", "") + "\n"))
	b.Write([]byte("Enabled" + "\t" + strconv.FormatBool(c.Enabled) + "\n"))
	b.Write([]byte("BindAddress" + "\t" + c.BindAddress + "\n"))
	fmt.Fprintln(w)
	w.Flush()
	return b.String()
}

func NewConfig() *Config {
	return &Config{
		Enabled:     true,
		BindAddress: "localhost:7777",
	}
}
