# paragraphxyz-tools

Tools that help work with paragraph.xyz data

`convert` is a package that can convert the json structure stored on arweave to markdown

```go
package main

import (
	"fmt"

	"github.com/gabrielayuso/paragraphxyz-tools/convert"
)

func main() {
	md, err := convert.PostJSONToMarkdown([]byte(`{"type":"doc","content":[{"type":"paragraph","content":[{"type":"text","text":"Hello, world!"}]}]}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(md)
}
```
