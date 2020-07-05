package array;

/**
 * @author Charlie
 * @When
 * @Description Say you have an array for which the ith element is the price of a given stock on day i.
 * <p>
 * Design an algorithm to find the maximum profit. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times).
 * <p>
 * Note: You may not engage in multiple transactions at the same time (i.e., you must sell the stock before you buy again).
 * <p>
 * Example 1:
 * <p>
 * Input: [7,1,5,3,6,4]
 * Output: 7
 * Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit = 5-1 = 4.
 * Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
 * Example 2:
 * <p>
 * Input: [1,2,3,4,5]
 * Output: 4
 * Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
 * Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are
 * engaging multiple transactions at the same time. You must sell before buying again.
 * Example 3:
 * <p>
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 * @Detail 与 121_BestTimetoBuyandSellStock 不同在于,121只能有1笔交易,而122可以有多笔交易
 * 同时有如下局限性, 卖出的日期肯定要比买的日期晚,所以 找到i下标的值Vi 的最大下标j
 * 在升序的子数组中,遇到第一个降序既卖出,同时移动下标i
 * <p>
 * 2. 看了网上的解读,是我想的太复杂了,就贼easy的 递加就行啊啊啊啊,silly silly me
 * is this question a joke?   not a damn joke
 * @Attention:
 * @Date 创建时间：2020-02-17 10:35
 */
public class Array_122_BestTimetoBuyandSellStockII
{
    public int maxProfit(int[] prices)
    {
        int maxProfit = 0;
        for (int i = 0; i < prices.length - 1; i++)
        {
            if (prices[i + 1] > prices[i]) maxProfit += prices[i + 1] - prices[i];
        }

        return maxProfit;
    }

}
