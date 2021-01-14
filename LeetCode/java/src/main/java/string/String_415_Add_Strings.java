package string;

/**
 * @author Charlie
 * @When
 * @Description Given two non-negative integers num1 and num2 represented
 * as string, return the sum of num1 and num2.
 * <p>
 * Note:
 * <p>
 * The length of both num1 and num2 is < 5100.
 * Both num1 and num2 contains only digits 0-9.
 * Both num1 and num2 does not contain any leading zero.
 * You must not use any built-in BigInteger
 * library or convert the inputs to integer directly.
 * 既 字符串型的数字相加
 * 不可以使用相关大数的类库等
 * @Detail 需要注意的就是加起来大于10的情况
 * 1. 依据数学加法的概念
 * 思路就是对字符串的所有值进行相加,字符串有长短之分,因此为了避免,短的到了末尾之后
 * 用0来代替
 * 这题忘了,对于字符串相加转数字是由套路的,/10 和%10
 * 当一个数不停的除以10: 如 12345 不停的除以10 得到的结果是: 1234,123,12,1
 * 当一个数不停的模10的时候: 如 12345 不停的模10
 * @Attention:
 * @Date 创建时间：2020-03-04 13:35
 */
public class String_415_Add_Strings
{

    public static String addStrings(String num1, String num2)
    {
        StringBuilder sb = new StringBuilder();
        int carry = 0;
        for (int i = num1.length() - 1, j = num2.length() - 1; i >= 0 || j >= 0 || carry == 1; i--, j--)
        {
            int x = i < 0 ? 0 : num1.charAt(i) - '0';
            int y = j < 0 ? 0 : num2.charAt(j) - '0';
            int t = x + y + carry;
            int mod = t % 10;
            // t/10 代表的是是否大于10,因为大于10的时候需要再+1
            int mul = t / 10;
            sb.append(mod);
            carry = mul;
        }
        return sb.reverse().toString();
    }

    public static void main(String[] args)
    {
        System.out.println(addStrings("987654321", "54321"));
    }
}
