package Top100;

/**
 * @author Charlie
 * @When
 * @Description You are climbing a stair case. It takes n steps to reach to the top.
 * <p>
 * Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
 * <p>
 * Note: Given n will be a positive integer.
 * 青蛙跳台阶问题
 * @Detail 动态规划问题:
 * @Attention:
 * @Date 创建时间：2020-04-08 16:23
 */
public class Top100_Easy_70_Climbing_Stairs
{
    public int climbStairs(int n)
    {
        // base cases
        if (n <= 0) return 0;
        if (n == 1) return 1;
        if (n == 2) return 2;

        int one_step_before = 2;
        int two_steps_before = 1;
        int all_ways = 0;

        for (int i = 2; i < n; i++)
        {
            all_ways = one_step_before + two_steps_before;
            two_steps_before = one_step_before;
            one_step_before = all_ways;
        }
        return all_ways;
    }

}
