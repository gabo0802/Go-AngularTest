package newsfeed

import (
	"fmt"
	"testing"
)

func TestAdd(test *testing.T) {
	fmt.Println("TestAdd")

	feed := New()
	feed.Add(Item{})

	if len(feed.Items) == 0 {
		test.Errorf("feed.Items is empty")

	}
}

func TestGetAll(test *testing.T) {

	feed := New()
	feed.Add(Item{})
	results := feed.GetAll()
	if len(results) != 1 {
		test.Errorf("Item was not added")
	}
}
