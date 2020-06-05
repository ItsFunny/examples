package array;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.SortedMap;

/**
 * @author Charlie
 * @When
 * @Description Given a list of words and two words word1 and word2,
 * return the shortest distance between these two words in the list.
 * <p>
 * For example,
 * Assume that words = ["practice", "makes", "perfect", "coding", "makes"].
 * <p>
 * Given word1 = “coding”, word2 = “practice”, return 3.
 * Given word1 = "makes", word2 = "coding", return 1.
 * <p>
 * Note:
 * You may assume that word1 does not equal to word2, and word1 and word2 are both in the list.
 * @Detail 1. 时间复杂度O(n^2)的方式,将符合key1和key2的值存储到2个list中,再2个list进行for循环遍历判断值
 * 2. 采取2个临时变量,分别保存key1和key2的下标值,每次2个值都有值则比较下当前的最小值
 * @Attention:
 * 经验: 求2个值的最短路径,可以采用3个变量,1个变量保存值,另外2个变量保存的是出现的位置,每次比较都触发一次min的比较
 * @Date 创建时间：2020-02-20 15:57
 */
public class Array_243_Shorted_Word_Distance
{
    public int shortedDistance(List<String> values, String key1, String key2)
    {
        int min = -1;
        int index1 = -1;
        int index2 = -1;
        for (int i = 0; i < values.size(); i++)
        {
            if (values.get(i).equals(key1))
            {
                index1 = i;
            } else if (values.get(i).equals(key2))
            {
                index2 = i;
            }
            if (index1 != -1 && index2 != -1)
            {
                min = Math.min(min, Math.abs(index1 - index2));
            }
        }
        return min;
    }

}
