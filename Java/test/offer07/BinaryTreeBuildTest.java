package offer07;

import org.junit.Test;

import static org.junit.Assert.*;

public class BinaryTreeBuildTest {
    @Test
    public void main(){
        int[] preorder = {};
        int[] inorder = {};
        BinaryTreeBuild tree = new BinaryTreeBuild();
        BinaryTreeBuild.TreeNode t = tree.buildTree(preorder, inorder);
        System.out.println("hw");
    }

}