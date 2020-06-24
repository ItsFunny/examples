package com.basic.algorithm.timewheel.jianzhioffer;

/**
 * @author joker
 * @When
 * @Description 斐波那契数列, 既F(n)=F(n-1)+F(n-2)+....
 * 大家都知道斐波那契数列，现在要求输入一个整数n，请你输出斐波那契数列的第n项（从0开始，第0项为0）
 * @Detail 这道题目如果用递归的话是非常easy的, 但是肯定会栈溢出, 但是既然能用递归的地方, 肯定能用循环或者栈来代替,
 * 但是因为基于斐波那契数列的特性,会有多余的计算,用栈恐怕也无法避免,所以直接采取循环来处理
 * 在这里,采用定义2个变量的方式代替F(n-1)和F(n-2),这里唯一需要注意的地方就是,起始是从2开始的,也就意味着prevRes=1
 * 同时我们需要尤其注意,是从0项开始的,所以在for循环中i=2,其实是3的位置
 * @date 创建时间：2019-05-18 20:25
 */
public class _7_FiberNaqi
{
    public static int Fibonacci(int n)
    {
        if (n <= 0)
        {
            return 0;
        } else if (n == 1)
        {
            return 1;
        }

        return Fibonacci(n - 1) + Fibonacci(n - 2);
    }

    public static int Fibonacci2(int n)
    {
        if (n == 0)
        {
            return 0;
        } else if (n == 1)
        {
            return 1;
        }
        int prevRes = 1, prevPrevRes = 0;
        int result = 0;
        for (int i = 2; i <= n; i++)
        {
            result = prevRes + prevPrevRes;
            prevPrevRes = prevRes;
            prevRes = result;
        }

        return result;
    }

    public static void main(String[] args)
    {
        System.out.println(_7_FiberNaqi.Fibonacci2(1));
        System.out.println(_7_FiberNaqi.Fibonacci2(2));
        System.out.println(_7_FiberNaqi.Fibonacci2(3));


    }
}
