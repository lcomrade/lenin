# Examples
## Get paste
Replace the value of the `pasteName` variable with the name of your paste.

```
package main

import(
	"github.com/lcomrade/lenin"
	"fmt"
)

func main() {
	// Base API URL
	baseURL := "https://paste.lcomrade.su/api"
	
	// Paste name
	pasteName := "PASTE_NAME"

	// Do request
	paste, err := lenin.Get(pasteName, baseURL)
	if err != nil {
		panic(err)
	}

	// Print result
	fmt.Println("Name:", paste.Name)
	fmt.Println("Title:", paste.Info.Title)
	fmt.Println("CreateTime:", paste.Info.CreateTime)
	fmt.Println("DeleteTime:", paste.Info.DeleteTime)
	fmt.Println("## TEXT ##")
	fmt.Println(paste.Text)
}
```
