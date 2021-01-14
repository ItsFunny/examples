package com.basic.algorithm.timewheel.jianzhioffer;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.List;
import java.util.Stack;

/**
 * @author joker
 * @When
 * @Description 输入一个链表，按链表值从尾到头的顺序返回一个ArrayList。,即可以认为是反转链表
 * @Detail 就是先后顺序进行更换, 毫无疑问stack是非常好的选择,
 * 不过既然能用stack,那必然能用递归来实现
 * @date 创建时间：2019-05-16 22:54
 */

class ListNode
{
    int val;
    ListNode next = null;

    ListNode(int val)
    {
        this.val = val;
    }
}

public class _3_print_from_tail_to_head
{
    public static ArrayList<Integer> printListFromTailToHead(ListNode listNode)
    {
        if (null == listNode) return null;
        ArrayList<Integer> result = new ArrayList<>();
        Stack<Integer> stack = new Stack<>();

        ListNode tempNode = listNode;
        while (null != tempNode)
        {
            stack.push(tempNode.val);
            tempNode = tempNode.next;
        }
        while (!stack.isEmpty())
        {
            result.add(stack.pop());
        }
        return result;
    }

    static List<Integer> l = new ArrayList<>();

    public static List<Integer> printListReverse(ListNode listNode)
    {
        if (null != listNode)
        {
            printListReverse(listNode.next);
            l.add(listNode.val);
        }
        return l;
    }

    public static void main(String[] args)
    {
        ListNode listNode = new ListNode(67);
        ListNode listNode1 = new ListNode(0);
        ListNode listNode2 = new ListNode(24);
        ListNode listNode3 = new ListNode(58);
        listNode.next = listNode1;
        listNode1.next = listNode2;
        listNode2.next = listNode3;
        List<Integer> arrayList = _3_print_from_tail_to_head.printListReverse(listNode);
        for (Integer integer : arrayList)
        {
            System.out.println(integer);
        }
    }
}
