package string;

/**
 * @author Charlie
 * @When
 * @Description Given a non-empty string check if it
 * can be constructed by taking a substring
 * of it and appending multiple copies of the substring together.
 * You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.
 * 既判断 这个字符串能否由其所属的n个子串所构建
 * @Detail 1. 编码的第一印象,这种题大概率是先直接返回true,在代码中判断不符合的条件
 * 既然是n个,则必定只需要截取一半即可
 * @Attention:
 * @Date 创建时间：2020-03-05 16:56
 */
public class String_459_Repeated_Substring_Pattern
{

    public boolean repeatedSubstringPattern(String str)
    {
        int l = str.length();
        for (int i = l / 2; i >= 1; i--)
        {
            // 既然是成对出现,则必然刚好是n个  ,不可能是小数点个
            if (l % i == 0)
            {
                //  如 12/4=3 ,3个组成一个
                int m = l / i;
                String subS = str.substring(0, i);
                StringBuilder sb = new StringBuilder();
                for (int j = 0; j < m; j++)
                {
                    sb.append(subS);
                }
                if (sb.toString().equals(str)) return true;
            }
        }
        return false;
    }

}
