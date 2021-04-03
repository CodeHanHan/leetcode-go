package offer07;

import sun.reflect.generics.tree.Tree;
import java.util.List;
import java.util.ArrayList;

public class BinaryTreeBuild {
    public static class TreeNode{
         int val;
         TreeNode left;
         TreeNode right;
         public TreeNode(int v){
             val = v;
         }
    }
    //@大威版
//    public TreeNode buildTree(int[] preorder, int[] inorder) {
//        if(preorder == null || inorder == null) return null;
//        if(preorder.length == 0 || inorder.length == 0) return null;
//        if(preorder.length != inorder.length) return null;
//        TreeNode root = new TreeNode(preorder[0]);
//        for(int i = 0; i < inorder.length; i++){
//            if(preorder[0] == inorder[i]){
//                root.left = buildTree(Arrays.copyOfRange(preorder,1,i + 1),
//                        Arrays.copyOfRange(inorder,0,i));
//                root.right = buildTree(Arrays.copyOfRange(preorder,i + 1,preorder.length),
//                        Arrays.copyOfRange(inorder,i + 1,inorder.length));
//            }
//        }
//        return root;
//    }


    public TreeNode buildTree(int[] preorder, int[] inorder) {
        /*
         * @Des: 根据给定的前序和中序序列，构造出对应的二叉树
         * @Para:
         *      preorder：前序序列
         *      inorder: 中序序列
         * @Return:
         *      二叉树
         **/
        ArrayList<Integer> pre = new ArrayList<>();
        ArrayList<Integer> in = new ArrayList<>();
        for(int value: preorder){   //复制到ArrayList里，为了切片与indexof
            pre.add(value);
        }
        for(int value: inorder){
            in.add(value);
        }
        return build(pre.subList(0, pre.size()), in.subList(0, in.size()));//这里切片是为了转化类型

        // 之所以使用Arraylist,是想使用其切片功能与根据值返回下标功能，但是查看源码后发现其返回下标
        // 功能就是遍历对比，这样子的话，还不如自己遍历，然后切片使用CopyofRange()
    }

    private TreeNode build(List<Integer> pre, List<Integer> in){
        if(pre.size() < 1){
            return null;
        }
        TreeNode head = new TreeNode(pre.get(0));
        int root = in.indexOf(pre.get(0));  //当前子树的根节点位置
        //根节点
        head.left = build( pre.subList(1, root+1), in.subList(0, root));
        head.right = build(pre.subList(root+1, pre.size()), in.subList(root+1,in.size()));
        return head;
    }
}
