package string;

/**
 * @author Charlie
 * @When
 * @Description Write a function that reverses a string.
 * The input string is given as an array of characters char[].
 * <p>
 * Do not allocate extra space for another array,
 * you must do this by modifying the input array in-place with O(1) extra memory.
 * <p>
 * You may assume all the characters consist of printable ascii characters.
 * 字符串反转
 * 1. 不可以申请额外内存,既O(1)的空间复杂度,只能在原先的操作上继续
 * @Detail
 * 基于数组的特性,直接基于下标交换即可
 * @Attention:
 * @Date 创建时间：2020-03-02 15:25
 */
public class String_344_Reverse_String
{
    public static void reverseString(char[] s)
    {
        int midIndex = s.length >> 1;
        int last = s.length - 1;
        int start = 0;
        int change;
        for (; start < midIndex; start++)
        {
            change = last - start;
            if (s[start] == s[change])
            {
                continue;
            }
            char temp = s[start];
            s[start] = s[change];
            s[change] = temp;
        }
//        int end = s.length - 1;
//        while (start < end)
//        {
//            if (s[start] == s[end])
//            {
//                continue;
//            }
//            char temp = s[start];
//            s[start] = s[end];
//            s[end] = temp;
//            start++;
//            end--;
//        }
    }

    public static void main(String[] args)
    {
        char[] a = new char[]{'h', 'e', 'l', 'l', 'o'};
//        a=new char[]{'a','b','c','d'};
        reverseString(a);
        StringBuilder sb = new StringBuilder();
        for (char c : a)
        {
            sb.append(c);
        }
        System.out.println(sb.toString());
    }
}
