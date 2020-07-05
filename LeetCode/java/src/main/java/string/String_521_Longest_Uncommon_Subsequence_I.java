package string;

/**
 * @author Charlie
 * @When
 * @Description Given a group of two strings, you need to find the longest uncommon subsequence of this group of two strings. The longest uncommon subsequence is defined as the longest subsequence of one of these strings and this subsequence should not be any subsequence of the other strings.
 * <p>
 * A subsequence is a sequence that can be derived from one sequence by deleting some characters without changing the order of the remaining elements. Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.
 * <p>
 * The input will be two strings, and the output needs to be the length of the longest uncommon subsequence. If the longest uncommon subsequence doesn't exist, return -1.
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-10 19:21
 */
public class String_521_Longest_Uncommon_Subsequence_I
{
    public int findLUSlength(String a, String b)
    {
        if (!a.equals(b))
        {
            return Math.max(a.length(), b.length());
        } else
        {
            return -1;
        }
    }
}
