package com.basic.algorithm.timewheel.jianzhioffer;

/**
 * @author joker
 * @When
 * @Description 青蛙跳台阶问题
 * 一只青蛙一次可以跳上1级台阶，也可以跳上2级。求该青蛙跳上一个n级的台阶总共有多少种跳法（先后次序不同算不同的结果）
 * @Detail 动态规划问题
 * 青蛙如果跳1次,则剩下的为F(n-1),若青蛙选择跳2级,则为F(n-2)
 * 因而,F(n)=F(n-1)+F(n-2)
 * 并且考虑到实际情况,当只有1级台阶的时候,只能跳1次,F(1)=1,而剩下2级的时候则为F(2)=2
 * 其实就是第7题的斐波那契数列问题
 * 但是与之不同的是,n=1的时候值为1,n=2的时候值为2,并且不存在从0开始,因而在for循环中是从3开始的
 * @date 创建时间：2019-05-18 20:59
 */
public class _8_frog_jump
{
    public static int JumpFloor(int target)
    {
        // 递归方式
        if (target < 0)
        {
            return -1;
        } else if (target == 1)
        {
            return 1;
        } else if (target == 2)
        {
            return 2;
        } else
        {
            return JumpFloor(target - 1) + JumpFloor(target - 2);
        }
    }

    public static int JumpFloorWithOut(int target)
    {
        if (target <= 0)
        {
            return -1;
        } else if (target == 1)
        {
            return 1;
        } else if (target == 2)
        {
            return 2;
        }
        int prev = 2, pprev = 1;
        int result = 0;

        for (int i = 3; i <= target; i++)
        {
            result = prev + pprev;
            pprev = prev;
            prev = result;
        }

        return result;
    }

    public static void main(String[] args)
    {
        System.out.println(JumpFloor(3));
        System.out.println(JumpFloorWithOut(3));
    }
}
