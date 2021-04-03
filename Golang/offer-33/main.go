package main

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	root := postorder[len(postorder)-1] //根节点
	var lastLeft int
	for i := len(postorder) - 2; i >= -1; i-- { // 定位左右子树
		if i == -1 { //没有左子树
			lastLeft = -1
			break
		}
		if postorder[i] < root {
			lastLeft = i
			break
		}
	}

	lChild := postorder[0 : lastLeft+1]                //左子树
	rChild := postorder[lastLeft+1 : len(postorder)-1] //右子树

	for i := 0; i < len(lChild); i++ {
		if root < lChild[i] {
			return false
		}
	}
	for i := 0; i < len(rChild); i++ {
		if root > rChild[i] {
			return false
		}
	}

	return verifyPostorder(lChild) && verifyPostorder(rChild)
}

//func judge(postorder []int, lastLeft int)bool{
//    root := postorder[len(postorder)]  //根节点
//    lChild := postorder[0: lastLeft+1]      //左子树
//    rChild := postorder[lastLeft+1: len(postorder)-1]       //右子树
//    for i := 0; i < len(lChild); i++ {
//        if root < lChild[i] {
//            return false
//        }
//    }
//    for i := 0; i < len(rChild); i++ {
//        if root > rChild[i] {
//            return false
//        }
//    }
//    return judge(postorder)
//
//}
