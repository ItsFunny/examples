package hashtable;

/**
 * @author Charlie
 * @When
 * @Description Count the number of prime numbers less than a non-negative number, n.
 * 既在指定范围内,获取质数的个数
 * @Detail 质数是指除了1和他本身外无其他数字
 * 1. 如何计算质数:
 * @Attention:
 * @Date 创建时间：2020-03-20 16:48
 */
public class HashTable_204_Count_Primes
{
    public int countPrimes(int n)
    {
        boolean[] notPrime = new boolean[n];
        int count = 0;
        for (int i = 2; i < n; i++)
        {
            if (notPrime[i] == false)
            {
                count++;
                // 如果和其他数相乘在此范围内则说明该值非质数
                for (int j = 2; i * j < n; j++)
                {
                    notPrime[i * j] = true;
                }
            }
        }

        return count;
    }
}
