package backtrack

import (
	"bytes"
	"fmt"
)

/**
* 0-1背包问题
* 有一个背包，背包总的承载重量是 W kg。
* 现在有 n 个物品，每个物品的重量不等，并且不可分割。
* 现在期望选择几件物品，装载到背包中。在不超过背包所能装载重量的前提下，如何让背包中物品的总重量最大？
**/

type Bag struct {
	Goods  []uint32
	Cap    uint32
	Weight uint32
}

func (b *Bag) Load(items []uint32) (goods []uint32, maxWeight uint32) {

	var fn func(i int, items []uint32)
	fn = func(i int, items []uint32) {
		if b.Weight >= b.Cap || i == len(items) { // 装满了或装到最后一个了
			if b.Weight > maxWeight { // 超过了原来装载的最大值
				fmt.Println(b)
				goods = append([]uint32{}, b.Goods...)
				maxWeight = b.Weight
			}
			return
		}

		fn(i+1, items) // 不装当前物品

		// 装载当前物品
		if b.Weight+items[i] <= b.Cap { // 剪枝
			// 装载
			b.Goods = append(b.Goods, items[i])
			b.Weight += items[i]

			// 继续下一个判断
			fn(i+1, items)

			// 恢复状态
			b.Goods = b.Goods[:len(b.Goods)-1]
			b.Weight -= items[i]
		}
	}

	fn(0, items)

	return
}

func (b *Bag) String() string {
	buf := bytes.Buffer{}

	buf.WriteString(fmt.Sprintf("%v; ", b.Goods))
	buf.WriteString(fmt.Sprintf("%d/%d", b.Weight, b.Cap))

	return buf.String()
}
