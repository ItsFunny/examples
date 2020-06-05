package list;

import java.util.List;
import java.util.logging.Level;

/**
 * @author Charlie
 * @When
 * @Description 合并2个有序链表
 * Merge two sorted linked lists and return it as a new list.
 * The new list should be made by splicing together the nodes of the first two lists.
 * @Detail 1. O(n^2)的时间复杂度,2个while循环
 * 2. 采用递归的方式,这种方式我理解有点问题
 * 3. 不采用递归,采用临时指针的形式,遍历判断,类似于申请新的数组,然后排序
 * @Attention:
 * 链表移动题,可以定义2个临时指针,其中一个临时指针去移动赋值(curNode),
 * 而另外一个指针则可以用来当返回值(resultNode,直接返回resultNode.next即可),
 * @Date 创建时间：2020-02-22 15:11
 */
public class List_21_Merge_Two_Sorted_Lists
{
    public static ListNode mergeTwoLists(ListNode l1, ListNode l2)
    {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        if (l1.val < l2.val)
        {
            // 题目要求的是,升序
            // 如果l2的值比l1的值大,但是不确定l1后面的值是否比l2的小
            // 所以需要找到l1中小于l2的最大值
            l1.next = mergeTwoLists(l1.next, l2);
            return l1;
        } else
        {
            // 同理: l1>l2 但是无法确定 l2中往后的值是否比l1的这个值小
            // 所以需要在l2中找到小于l1的最大值
            l2.next = mergeTwoLists(l1, l2.next);
            return l2;
        }
    }

    public static ListNode mergeTwoLists2(ListNode l1, ListNode l2)
    {
        if (l1 == null) return l2;
        if (l2 == null) return l1;
        ListNode resultNode = new ListNode(0);
        ListNode curNode = resultNode;
        while (l1 != null && l2 != null)
        {
            if (l1.val < l2.val)
            {
                curNode.next = l1;
                l1 = l1.next;
            } else
            {
                curNode.next = l2;
                l2 = l2.next;
            }
            // 移动首部临时指针
            curNode = curNode.next;
        }
        curNode.next = l1 == null ? l2 : l1;
        return resultNode.next;
    }

    public static void main(String[] args)
    {
        ListNode l1 = ListNode.buildRandNodes(4);
        ListNode l2 = ListNode.buildRandNodes(4);
        mergeTwoLists(l1, l2);

    }
}
