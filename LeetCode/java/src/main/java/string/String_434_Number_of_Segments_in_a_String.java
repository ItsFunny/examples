package string;

/**
 * @author Charlie
 * @When
 * @Description Count the number of segments in a string,
 * where a segment is defined to be a contiguous sequence of non-space characters.
 * <p>
 * Please note that
 * the string does not contain any non-printable characters.
 * 获取一段字符串中非空的有效串的个数
 * <p>
 * 理解错提示了
 * 什么时候统计次数会+1:
 * 当该字符不为空,同时上一个字符为空,代表着是一个新的单词,则++
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-04 16:07
 */
public class String_434_Number_of_Segments_in_a_String
{
    public int countSegments(String s)
    {
        if (s.equals(" ")) return 0;
        int count = 0;
        for (int i = 0; i < s.length(); i++)
        {
            if (s.charAt(i) != ' ' && (i == 0 || s.charAt(i - 1) == ' '))
            {
                count++;
            }
        }
        return count;
    }
}
