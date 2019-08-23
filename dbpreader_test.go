package dbpreader

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	results, err := Query("Go_(programming_language)")
	if err != nil {
		panic(err)
	}
	fmt.Println(results)

	fmt.Println("============================")
	fmt.Println("Exact Resource:")
	resource := results.FindResource("Go_(programming_language)")
	fmt.Println(resource)

	fmt.Println("============================")
	fmt.Println("WikiData Item:")
	fmt.Println(resource.GetWikiDataItem())
}