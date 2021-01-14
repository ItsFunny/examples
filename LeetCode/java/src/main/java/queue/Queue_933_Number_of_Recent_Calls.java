package queue;

import java.util.LinkedList;
import java.util.Queue;

/**
 * @author Charlie
 * @When
 * @Description Write a class RecentCounter to count recent requests.
 * <p>
 * It has only one method: ping(int t), where t represents some time in milliseconds.
 * <p>
 * Return the number of pings that have been made from 3000 milliseconds ago until now.
 * <p>
 * Any ping with time in [t - 3000, t] will count, including the current ping.
 * <p>
 * It is guaranteed that every call to ping uses a strictly larger value of t than before.
 * 题目没看懂
 * <p>
 * 既统计值在[t-3000,t]范围内的次数
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-19 17:21
 */
public class Queue_933_Number_of_Recent_Calls
{
    private Queue<Integer> queue;

    public Queue_933_Number_of_Recent_Calls()
    {
        this.queue = new LinkedList<>();
    }

    public int ping(int t)
    {
        while (!queue.isEmpty())
        {
            if (queue.peek() + 3000 >= t) break;
            queue.poll();
        }
        queue.add(t);
        return queue.size();
    }
}
