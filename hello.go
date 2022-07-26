package say

import "fmt"

type Greeter struct {
	Format string
}

func (g Greeter) Hello(name string) string {
	if g.Format == "" {
		g.Format = "Hello %s"
	}
	return fmt.Sprintf(g.Format, name)
}
