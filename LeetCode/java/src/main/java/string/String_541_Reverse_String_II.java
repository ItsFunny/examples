package string;

/**
 * @author Charlie
 * @When
 * @Description Given a string and an integer k,
 * you need to reverse the first k characters for every 2k characters counting from the start of the string.
 * If there are less than k characters left, reverse all of them.
 * If there are less than 2k but greater than or equal to k characters,
 * then reverse the first k characters and left the other as original.
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-11 16:05
 */
public class String_541_Reverse_String_II
{
    public static String reverseStr(String s, int k)
    {
        int length = s.length();
        int over = k << 1;
        int times = length / over;
        int left = length % over;
        StringBuilder sb = new StringBuilder();
        int index = 0;
        for (int i = 0; i < times; i += over)
        {
            // 对这里面的所有字符串反转
            sb.append(s, index, index + k);
            sb.reverse();
            sb.append(s, index + k, index + k + k);
            index = i + k + k;
        }
        if (left < k)
        {
//            全部反转
            StringBuilder sb2 = new StringBuilder();
            sb.append(sb2.append(s, index, length - 1).reverse());
        } else
        {
            // 只对k个反转
            StringBuilder sb2 = new StringBuilder();
            sb.append(sb2.append(s, index, index + k).reverse());
            sb.append(s, index + k, length);

        }
        return sb.toString();
    }


    public String reverseStr2(String s, int k) {
        char[] arr = s.toCharArray();
        int n = arr.length;
        int i = 0;
        while(i < n) {
            int j = Math.min(i + k - 1, n - 1);
            swap(arr, i, j);
            i += 2 * k;
        }
        return String.valueOf(arr);
    }
    private void swap(char[] arr, int l, int r) {
        while (l < r) {
            char temp = arr[l];
            arr[l++] = arr[r];
            arr[r--] = temp;
        }
    }
}