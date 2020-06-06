package com.basic.structure.tree;

import org.junit.Assert;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-16 09:37
 */
public class CompleteTreeTest
{

    /*
        the tree shall be like this:
                1
               / \
              2   3
             /\   /\
            4  5  6 7
           /\  /
          8 9 10

     */
//    private Integer[] arr = new Integer[]{1, 2, 4, 8, -1, -1, 9, 5, 10, -1, -1, -1, 3, 6, -1, -1, 7};
    private Integer[] arr = new Integer[]{1,  2, 3, 4, 5, 6, 7, 8, 9, 10,11};

    @Test
    public void testIteratorByArrayAndBFS()
    {
        CompleteTree completeTree = new CompleteTree();
        completeTree.buildCompleteBinaryTree(arr);
        List<Integer> resultList = new ArrayList<>();
        completeTree.inIteratorByArray(arr, 0, resultList);
        for (Integer integer : resultList)
        {
            System.out.printf("%d->", integer);
        }
        System.out.println();
//        assert Arrays.asList(1, 2, 4, 8, 9, 5, 10, 3, 6, 7).equals(resultList);
        List<Integer> bfsResultList = completeTree.BFSTree();
        for (Integer integer : bfsResultList)
        {
            System.out.printf("%d->", integer);
        }
//        assert Arrays.asList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).equals(bfsResultList);
    }


    /*
            1                           1                                  1

          2   3                        2  3                               2   3

        4  5 6  7                    4  5 6 7                           4  5 6  7
       / \    \                       \                                          \
      8  9     10                      8                                          8

     */
    private Integer[] wrongArr1 = new Integer[]{1, 2, 4, 8, -1, -1, 9, -1, -1, 5, -1, -1, 3, 6, -1, 10, -1, -1, 7};

    private Integer[] wrongArr2 = new Integer[]{1, 2, 4, -1, 8, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7};

    private Integer[] wrongArr3 = new Integer[]{1, 2, 4, -1, -1, 5, -1, -1, 3, 6, -1, -1, 7, -1, 8};


    @Test
    public void completeTreeWrongTest1()
    {
        CompleteTree completeTree = new CompleteTree();
        completeTree.buildTreeByStack(wrongArr1);
        Assert.assertFalse(completeTree.validIfCompleteTree());
    }

    @Test
    public void completeTreeWrongTest2()
    {
        CompleteTree completeTree = new CompleteTree();
        completeTree.buildTreeByStack(wrongArr2);
        Assert.assertFalse(completeTree.validIfCompleteTree());
    }

    @Test
    public void completeTreeWrongTest3()
    {
        CompleteTree completeTree = new CompleteTree();
        completeTree.buildTreeByStack(wrongArr3);
        Assert.assertFalse(completeTree.validIfCompleteTree());
    }

    @Test
    public void completeTreeTrueTest()
    {
        CompleteTree completeTree = new CompleteTree();
        completeTree.buildCompleteBinaryTree(arr);
        Assert.assertTrue(completeTree.validIfCompleteTree());
    }

}
