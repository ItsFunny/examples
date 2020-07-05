package array;

/**
 * @author Charlie
 * @When
 * @Description Say you have an array for which the ith element is the price of a given stock on day i.
 * <p>
 * If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.
 * <p>
 * Note that you cannot sell a stock before you buy one.
 * <p>
 * Example 1:
 * <p>
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
 * Not 7-1 = 6, as selling price needs to be larger than buying price.
 * Example 2:
 * <p>
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 * @Detail 1. 最直白的做法,时间复杂度O(n^2) ,两个for循环,每个都直接做运算,比较最小值
 * 2. 参考的网上其他思路,原理与求最短路径相同,用的是贪心算法,既假设首位为最低价格,则之后只需要更新最低价格的同时
 * 计算最大利润即可
 * @Attention:
 * @Date 创建时间：2020-02-17 10:04
 */
public class Array_121_BestTimetoBuyandSellStock
{
    public static int maxProfit(int[] prices)
    {
        int length = prices.length;
        int max = 0;
        for (int i = 0; i < length; i++)
        {
            int prv = prices[i];
            for (int j = i + 1; j < length; j++)
            {
                if (prices[j] > prv)
                {
                    max = prices[j] - prv > max ? prices[j] - prv : max;
                }
            }
        }


        return max;
    }


    public static int maxProfixWithGreedy(int[] prices)
    {
        if (prices == null||prices.length==0)
        {
            return 0;
        }
        int minPrice = prices[0];
        int maxProfit = 0;

        for (int i = 1; i < prices.length; i++)
        {
            if (prices[i] < minPrice)
            {
                minPrice = prices[i];
            } else
            {
                maxProfit = Math.max(maxProfit, prices[i] - minPrice);
            }

        }

        return maxProfit;
    }

    public static void main(String[] args)
    {
//        7,1,5,3,6,4]
        System.out.println(maxProfit(new int[]{7, 1, 5, 3, 6, 4}));
    }


}
