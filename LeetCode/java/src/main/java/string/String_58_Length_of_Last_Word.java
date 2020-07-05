package string;

/**
 * @author Charlie
 * @When
 * @Description Given a string s consists of upper/lower-case alphabets and
 * empty space characters ' ', return the length of last word (last word means the last appearing word if we loop from left to right) in the string.
 * If the last word does not exist, return 0.
 * <p>
 * Note: A word is defined as a maximal substring
 * consisting of non-space characters only.
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-29 19:53
 */
public class String_58_Length_of_Last_Word
{
    public int lengthOfLastWord(String s)
    {
        if (s == null || s.length() == 0) return 0;
        String[] s1 = s.split(" ");
        if (s1.length == 0)
        {
            return 0;
        }
        return s1[s1.length - 1].length();
//        int length = 0;
//        for (int i = 0; i < s1.length; i++)
//        {
//            if (s1[i].length() > length)
//            {
//                length = s1[i].length();
//            }
//        }
//        return length;
    }

    public static void main(String[] args)
    {
        String[] s = " ".split(" ");
        for (String s1 : s)
        {
            System.out.println(s1);
        }
        System.out.println(s);
    }
}
