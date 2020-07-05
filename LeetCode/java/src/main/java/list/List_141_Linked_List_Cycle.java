package list;

/**
 * @author Charlie
 * @When
 * @Description Given a linked list, determine if it has a cycle in it.
 * <p>
 * To represent a cycle in the given linked list,
 * we use an integer pos which represents the position (0-indexed) in the linked
 * list where tail connects to. If pos is -1, then there is no cycle in the linked list
 * 既判断是否有环
 * @Detail
 * 1. 快慢指针
 * @Attention:
 * @Date 创建时间：2020-02-23 10:12
 */
public class List_141_Linked_List_Cycle
{

    public boolean hasCycle(ListNode head)
    {
        if (head==null)return false;
        ListNode slow = head;
        ListNode fast = head.next;
        while (slow != null && fast != null)
        {
            if (fast == slow)
            {
                return true;
            }
            fast = fast.next;
            slow = slow.next;
            if (null != fast)
            {
                fast = fast.next;
            } else
            {
                return false;
            }
        }

        return false;
    }
}
