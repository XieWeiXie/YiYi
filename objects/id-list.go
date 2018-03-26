package objects

import "fmt"

type ListData struct {
	Response string `json: "res"`
	Data     string `json: "data"`
}

func (L ListData) String() string {
	return fmt.Sprintf(`
You Get One List From One App:

     Response: %s,
     Data:     %s
`, L.Response, L.Data)
}
