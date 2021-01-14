package list;

import java.util.HashMap;
import java.util.Map;

/**
 * @author Charlie
 * @When
 * @Description Given a sorted linked list,
 * delete all duplicates such that each element appear only once.
 * 在有序链表中删除重复元素
 * @Detail 1. 集合hashMap
 * 2. 既然是有序的,则后一个元素必定大于前一个元素
 * @Attention:
 * @Date 创建时间：2020-02-22 16:05
 */
public class List_83_remove_duplicates_from_sorted_list
{
    public ListNode deleteDuplicates1(ListNode head)
    {
        Map<Integer, Integer> map = new HashMap<>();
        ListNode tNode = head;
        ListNode lastNode = head;
        while (tNode != null)
        {
            if (map.containsKey(tNode.val))
            {
                lastNode.next = tNode.next;
            } else
            {
                map.put(tNode.val, 1);
                lastNode = tNode;
            }

            tNode = tNode.next;
        }

        return head;
    }

    public ListNode deleteDuplicates2(ListNode head)
    {
        ListNode tNode = head;
        while (tNode != null)
        {
            ListNode next = tNode.next;
            if (next != null && next.val == tNode.val)
            {
                tNode.next = next.next;
            } else
            {
                tNode = tNode.next;
            }
        }
        return head;
    }
}
