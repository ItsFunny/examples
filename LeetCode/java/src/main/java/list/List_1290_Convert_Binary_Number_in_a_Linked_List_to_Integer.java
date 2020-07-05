package list;

/**
 * @author Charlie
 * @When
 * @Description Given head which is a reference node to a singly-linked list.
 * The value of each node in the linked list is either 0 or 1.
 * The linked list holds the binary representation of a number.
 * <p>
 * Return the decimal value of the number in the linked list.
 * 既将 链表转换为整数
 * @Detail 1. 链表反转
 * 2. 大神的思路:  直接用或运算即可  ,其实不需要考虑  (1<<i) , `关键是 这个val的值要么为0,要么为1` ,所以我们可以直接通过| 运算构建出这个 0,1数组
 * @Attention:
 * @Date 创建时间：2020-02-26 16:34
 */
public class List_1290_Convert_Binary_Number_in_a_Linked_List_to_Integer
{
    public static int getDecimalValue(ListNode head)
    {
        head = reverse(head);
        int c = 0;
        int v = 0;
        while (head != null)
        {
            v += head.val * (1 << c);
            head = head.next;
            c++;
        }


        return v;
    }

    private static ListNode reverse(ListNode head)
    {
        ListNode cur = head;
        ListNode prev = null;
        while (cur != null)
        {
            ListNode next = cur.next;
            cur.next = prev;
            prev = cur;
            cur = next;
        }
        return prev;
    }

    public int getDecimalValue2(ListNode head)
    {
        int ans = 0;
        while (head != null)
        {
            ans = (ans << 1) | head.val;
            head = head.next;
        }
        return ans;
    }

    public static void main(String[] args)
    {
        ListNode listNode = ListNode.buildInValues(new int[]{1, 0, 1});
        System.out.println(getDecimalValue(listNode));
    }
}
