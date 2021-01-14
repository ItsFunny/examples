package list;

/**
 * @author Charlie
 * @When
 * @Description Write a program to find the node at which the intersection of
 * two singly linked lists begins.
 * For example, the following two linked lists:
 * 既返回2个链表共通的元素的起始,类似于寻找子串
 * @Detail 1. 没思路
 * 2. 这题也算考察JVM了,虽然 l1=new ListNode(1) 与 l2=new ListNode(1) l1==l2 返回false
 * 但是根据题意以及图解,其实headA和headB 相同的部分,其实占用的是同一部分内存,所以之前卡在的点是
 * 截取相同子串,不仅第一个值要相同,后面的值也要相同,卡在了如何判断后面的值也相同的方法上,所以只需要判断
 * 相等即可
 * @Attention: If the two linked lists have no intersection at all, return null.
 * The linked lists must retain their original structure after the function returns.
 * You may assume there are no cycles anywhere in the entire linked structure.
 * Your code should preferably run in O(n) time and use only O(1) memory.
 * 1. 无则返回null
 * 2. 返回的结构体不能更改属性
 * 3. 无循环
 * 4. 空间复杂度O(1),时间复杂度O(n),既不可以申请额外内存,并且一个for循环搞定
 * @Date 创建时间：2020-02-23 10:25
 * <p>
 * 总结: 考察2点:
 * 1. 内存分配概念,链表相等与字符串相等不同,链表相等可以通过==直接判断,而字符串则需要所有字符相等
 * 2. 不同长度的链表A,B如何在O(n)的时间复杂度上遍历完,2个临时指针若为空,则移动到互相的指针头即可:此时是快慢指针(
 * 特殊情况是当链表为环的时候,会必然为有相似,因为此时是快慢指针,但是题目无环)
 */
public class List_160_Intersection_of_Two_Linked_Lists
{
    public ListNode getIntersectionNode(ListNode headA, ListNode headB)
    {
        if (headA == null || headB == null) return null;
        ListNode t1 = headA;
        ListNode t2 = headB;
        while (t1 != t2)
        {
            // 进行移动
            if (t1 != null)
            {
                t1 = t1.next;
            } else
            {
                t1=headB;
            }
            if (t2!=null){
                t2=t2.next;
            }else{
                t2=headA;
            }
        }


        return t1;
    }

}
