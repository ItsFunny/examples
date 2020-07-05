package queue;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * @author Charlie
 * @When
 * @Description Given a stream of integers and a window size, calculate the moving average of all integers in the sliding window.
 * <p>
 * For example,
 * MovingAverage m = new MovingAverage(3);
 * m.next(1) = 1
 * m.next(10) = (1 + 10) / 2
 * m.next(3) = (1 + 10 + 3) / 3
 * m.next(5) = (10 + 3 + 5) / 3
 * 初始化一个滑动窗口,大小为w,输入一系列数,
 * 求窗口内的平均数,窗口会向前滑动,当窗口填满时,将最早进入的数弹出,加入新的数.
 * 然后对余下的数作平均值计算即可
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-19 16:07
 */
public class Queue_346_Moving_Average_from_Data_Stream
{
    private List<Integer> listQueue;
    private int sum;
    private int cap;

    public Queue_346_Moving_Average_from_Data_Stream(int size)
    {
        this.listQueue = new ArrayList<>();
    }

    public double next(int v)
    {
        if (this.listQueue.size() < this.cap)
        {
            this.listQueue.add(this.listQueue.size() - 1, v);
            sum += v;
        } else
        {
            // 弹出第一个
            sum -= this.listQueue.remove(0);
        }
        return sum / this.listQueue.size();
    }

}
