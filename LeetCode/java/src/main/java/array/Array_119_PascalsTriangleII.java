package array;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail 参考https://blog.csdn.net/weixin_33744854/article/details/94703410
 * 之前卡在 L23, 原因在于临时变量位置定义错误,i 下标的值=i-1 + i 的值 而 i+1的值=i+i+2的值,因此
 * i下标的值会被取代,所以要保存i下标的值
 * @Attention:
 * @Date 创建时间：2020-02-15 15:46
 */
public class Array_119_PascalsTriangleII
{
    public static List<Integer> getRow(int rowIndex)
    {
        ArrayList<Integer> result = new ArrayList<Integer>();
        result.add(1);
        for (int i = 1; i <= rowIndex; i++)
        {
            int tmp = 1;
            for (int j = 1; j < i; j++)
            {
                int midtmp = result.get(j);
                result.set(j, tmp + midtmp);
                tmp = midtmp;
            }
            result.add(1);
        }
        return result;
    }
}
