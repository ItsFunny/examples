package string;

import java.net.Inet4Address;

/**
 * @author Charlie
 * @When
 * @Description Given two binary strings, return their sum (also a binary string).
 * <p>
 * The input strings are both non-empty and contains only characters 1 or 0.
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-29 20:04
 */
public class String_67_Add_Binary_重看
{

    public String addBinary(String a, String b)
    {
        StringBuilder sb = new StringBuilder();
        int i = a.length() - 1, j = b.length() -1, carry = 0;
        while (i >= 0 || j >= 0) {
            int sum = carry;
            if (j >= 0) sum += b.charAt(j--) - '0';
            if (i >= 0) sum += a.charAt(i--) - '0';
            sb.append(sum % 2);
            carry = sum / 2;
        }
        if (carry != 0) sb.append(carry);
        return sb.reverse().toString();
    }
}
