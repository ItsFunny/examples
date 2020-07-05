package string;

/**
 * @author Charlie
 * @When
 * @Description Write a function to find the
 * longest common prefix string amongst an array of strings.
 * <p>
 * If there is no common prefix, return an empty string ""
 * @Detail 找到最常出现的子串
 * 这题没思路,一直卡在,以为需要自己造轮子,感觉这个轮子好复杂,然后发现可以直接用String.indexOf来逐个
 * 判断即可
 * 也有陷阱,陷阱在题目上,是前缀,不是包含,所以要indexOf==0来判断
 * @Attention:
 * @Date 创建时间：2020-02-27 16:44
 */
public class String_14_Longest_Common_Prefix
{
    public String longestCommonPrefix(String[] strs)
    {
        if (strs == null || strs.length == 0) return "";
        String common = strs[0];
        for (int i = 1; i < strs.length; i++)
        {
            while (strs[i].indexOf(common) != 0)
            {
                common = common.substring(0, common.length() - 1);
            }

        }

        return common;
    }
}
