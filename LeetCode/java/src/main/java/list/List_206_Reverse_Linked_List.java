package list;

/**
 * @author Charlie
 * @When
 * @Description Reverse a singly linked list.
 * 既反转单链表
 * @Detail 1. 循环实现
 * 2. 递归实现
 * 递归实现与循环实现是一模一样的思路,就是先保存下一个节点,然后当前节点重新指向前一个节点
 * 然后先移动前一个节点到当前节点,因为对next而言,它的prev就是cur节点,然后cur节点移动到下一个节点作为新的cur节点
 * @Attention:
 * @Date 创建时间：2020-02-24 16:18
 */
public class List_206_Reverse_Linked_List
{
    /*
        1,2,3,4,5,null
        cur 指的是正向的当前节点,prev指的是正向的前一个节点
        当第一次移动的时候
        cur=1,prev=null
        需要先保存next,既 next=head.next ,这是作为新的head,因为需要继续遍历
        然后反向连接,既cur要指向prev,同时因为需要继续遍历,所以这个时候cur节点就变成了next节点的prev
        而next节点变成了cur当前节点
        所以 prev=cur ,cur=next;

     */
    public ListNode reverseList(ListNode head)
    {
        ListNode newHead = null;
        while (head != null)
        {
            // 先保存
            // 保存下一个节点作为head下一遍头节点
            ListNode next = head.next;
            // 再连
            // 连接之前的节点
            head.next = newHead;
            // 再将之前的节点往遍历的方向移动,
            newHead = head;
            // 将当前遍历的节点往前遍历
            head = next;
        }


        return newHead;
    }


    public ListNode reverseList2(ListNode head)
    {
        ListNode prev = null;
        ListNode cur = head;
        while (cur != null)
        {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }

    // 递归实现
    public ListNode reverseListWithRescure(ListNode head)
    {
        return recerseWithRescure(head, null);
    }

    private ListNode recerseWithRescure(ListNode cur, ListNode prev)
    {
        if (null == cur)
        {
            return prev;
        }
        ListNode next = cur.next;
        cur.next = prev;
        prev = cur;
        cur = next;
        return recerseWithRescure(cur, prev);
    }


}
