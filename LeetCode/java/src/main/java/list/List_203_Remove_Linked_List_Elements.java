package list;

/**
 * @author Charlie
 * @When
 * @Description remove all elements from a linked list of integers that have value val.
 * 移除链表中所有的这个元素
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-24 16:01
 */
public class List_203_Remove_Linked_List_Elements
{
    public ListNode removeElements(ListNode head, int val)
    {
        ListNode resultNode = new ListNode(-1);
        resultNode.next = head;
        ListNode t = resultNode;
        ListNode prev = resultNode;
        while (t != null)
        {
            if (t.val == val)
            {
                prev.next = t.next;
            } else
            {
                prev = t;
            }
            t = t.next;
        }

        return resultNode.next;
    }

}
