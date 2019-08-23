# dbpreader
Read dbpedia pages into golang type and structs. Simple and easy to use.

### Install
```
go get -u github.com/Navid2zp/dbpreader
```


### Example

Page: `http://dbpedia.org/page/Go_(programming_language)`
```
import (
	"fmt"
	"dbpreader"
)

func main() {
	results, err := dbpreader.Query("Go_(programming_language)")
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("Exact Resource:")
	resource := results.FindResource("Go_(programming_language)")
	fmt.Println(resource)

	fmt.Println("WikiData Item:")
	fmt.Println(resource.GetWikiDataItem())
}
```

#### Methods:

There is some small useful methods that you can use.
I might add more along the way.


**_FindResource:_**

The json version of dbpedia returns a list of results
This method will find and return the exact resource

```
resource := results.FindResource("Go_(programming_language)")
fmt.Println(resource)
```

**_GetWikiDataItem:_**

Finds the wikidata item url

```
resource.GetWikiDataItem()
```

License
----

MIT
