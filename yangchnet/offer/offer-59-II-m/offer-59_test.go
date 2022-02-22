package offer_59_II_m

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	mq := Constructor()
	mq.Push_back(1)
	mq.Push_back(2)
	fmt.Println(mq.Max_value())
	fmt.Println(mq.Pop_front())
	fmt.Println(mq.Max_value())
}
