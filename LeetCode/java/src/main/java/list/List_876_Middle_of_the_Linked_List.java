package list;

import com.sun.xml.internal.bind.v2.model.core.ID;

/**
 * @author Charlie
 * @When
 * @Description Given a non-empty,
 * singly linked list with head node head,
 * return a middle node of linked list.
 * <p>
 * If there are two middle nodes, return the second middle node.
 * 截断链表返回后半程
 * @Detail
 * 利用快慢指针,当快指针走到末尾的时候,慢指针刚好走到中间位置 ,既 index>>1 的位置
 * @Attention:
 * @Date 创建时间：2020-02-26 16:04
 */
public class List_876_Middle_of_the_Linked_List
{
    public static ListNode middleNode(ListNode head)
    {
        if (head == null) return null;
        ListNode fast = head.next;
        ListNode slow = head;
        int linkCount = 2;
        while (true)
        {
            if (fast == null) break;
            if (fast.next == null)
            {
                linkCount += 1;
                break;
            }
            fast = fast.next.next;
            slow = slow.next;
            if (fast != null)
            {
                linkCount += 2;
            }
        }

        if (linkCount % 2 == 0)
        {
            return slow;
        }

        return slow.next;
    }

    public static void main(String[] args)
    {
        ListNode listNode = ListNode.buildIteratorNodes(6);
        ListNode listNode1 = List_876_Middle_of_the_Linked_List.middleNode(listNode);
        System.out.println(listNode1.val);
    }
}

