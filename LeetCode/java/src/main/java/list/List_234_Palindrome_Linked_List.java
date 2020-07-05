package list;

/**
 * @author Charlie
 * @When
 * @Description Given a singly linked list,
 * determine if it is a palindrome.
 * 判断是否是回文链表
 * 1. 什么是回文字符串:
 * 简单的一句话概括就是关于中心左右对称的字符串。
 * 例如：ABCBA或者AACCAA是回文字符串；ABCCA或者AABBCC不是回文字符串。
 * @Detail 似乎只能通过反转来实现, 起初卡在了限制条件 时间复杂度O(n) ,并且空间复杂度O(1)
 * 通过快慢指针+反转来实现
 * 快慢指针的核心就是每一次快指针都比慢指针快一步,最终慢指针会处于中间
 * 1,2,3,4,5,6,7
 * @Attention:
 * @Date 创建时间：2020-02-25 16:07
 */
public class List_234_Palindrome_Linked_List
{
    public boolean isPalindrome(ListNode head)
    {
        if (null == head) return true;
        ListNode fast = head.next;
        ListNode slow = head;
        while (fast != null && fast.next != null)
        {
            fast = fast.next.next;
            slow = slow.next;
        }
        if (null != slow)
        {
            slow = slow.next;
        }
        slow = reverse(slow);
        fast = head;
        // 进行判断
        while (slow != null && fast != null)
        {
            if (slow.val != fast.val)
            {
                return false;
            }
            slow = slow.next;
            fast = fast.next;
        }
        return true;
    }

    private ListNode reverse(ListNode head)
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

    public static void main(String[] args)
    {

    }
}
