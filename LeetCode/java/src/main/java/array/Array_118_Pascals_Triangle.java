package array;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description 就是创建一个三角形即可
 * @Detail
 * 很简单,只需要掌握规律即可, 这个下标i的值=上一行的i-1下标的值+上一行i下标的值
 * @Attention:
 * @Date 创建时间：2020-02-15 15:33
 */
public class Array_118_Pascals_Triangle
{

    public List<List<Integer>> generate(int numRows)
    {
        List<List<Integer>> result = new ArrayList<>();

        for (int i = 0; i < numRows; i++)
        {
            Integer[] array = new Integer[i + 1];
            array[0] = 1;
            array[array.length - 1] = 1;
            for (int j = 1; j < array.length - 1; j++)
            {
                array[j] = result.get(i - 1).get(j - 1) + result.get(i - 1).get(j);
            }
            result.add(Arrays.asList(array));

        }

        return result;
    }

}
