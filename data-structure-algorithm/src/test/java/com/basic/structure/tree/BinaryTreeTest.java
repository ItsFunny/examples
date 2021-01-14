package com.basic.structure.tree;

import org.junit.Assert;
import org.junit.Test;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;


public class BinaryTreeTest
{
    private Integer[] arr = new Integer[]{1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7};


    /*

        The tree is looks like:
                1
            2     3
         4    5  6  7
     */

    @Test
    public void loopBuildTree()
    {
        BinaryTree tree = new BinaryTree();
        tree.setRoot(tree.loopBuildTree(tree.getRoot(), arr));
        System.out.println(tree);

    }

    @Test
    public void testStackBuildTree()
    {
        BinaryTree tree = new BinaryTree();
        tree.stackBuildTree(arr);
        System.out.println(tree);
    }


    @Test
    public void testBFSTree()
    {
        BinaryTree binaryTree = new BinaryTree();
        binaryTree.stackBuildTree(arr);
        List<Integer> result = binaryTree.BFSTree();
        for (Integer integer : result)
        {
            System.out.printf("%d->", integer);
        }

        assert Arrays.asList(1, 2, 3, 4, 5, 6, 7).equals(result);
    }

    @Test
    public void testDFSTree()
    {
        BinaryTree binaryTree = new BinaryTree();
        binaryTree.stackBuildTree(arr);
        List<Integer> result = binaryTree.DFSTree();
        for (Integer integer : result)
        {
            System.out.printf("%d->", integer);
        }
        assert Arrays.asList(1, 2, 4, 5, 3, 6, 7).equals(result);
    }


    // 先序遍历测试 --结合非递归
    @Test
    public void testPreorderTree()
    {
        BinaryTree binaryTree = new BinaryTree();
        binaryTree.stackBuildTree(arr);
        List<Integer> resultList1 = new ArrayList<>();
        binaryTree.preorderTree(binaryTree.getRoot(), resultList1);
        List<Integer> resultList2 = binaryTree.preorderTreeByStack();
        Assert.assertArrayEquals(resultList1.toArray(), resultList2.toArray());
    }

    @Test
    public void testInOrderTree()
    {
        BinaryTree binaryTree = new BinaryTree();
        binaryTree.stackBuildTree(arr);
        List<Integer> resultList = new ArrayList<>();
        binaryTree.inOrderTree(binaryTree.getRoot(), resultList);
        List<Integer> resultList2 = binaryTree.inOrderTreeByStack();
        Assert.assertArrayEquals(resultList.toArray(), resultList2.toArray());
    }

    @Test
    public void testPostOrderTree()
    {
        BinaryTree binaryTree = new BinaryTree();
        binaryTree.stackBuildTree(arr);
        List<Integer> resultList1 = new ArrayList<>();
        binaryTree.postOrderTree(binaryTree.getRoot(), resultList1);
        List<Integer> resultList2 = binaryTree.postOrderTreeByStack();
        Assert.assertArrayEquals(resultList1.toArray(), resultList2.toArray());

    }
}

