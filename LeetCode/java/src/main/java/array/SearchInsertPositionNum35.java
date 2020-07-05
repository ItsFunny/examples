package array;

/**
 * @author joker
 * @When
 * @Description 返回target所处的下标, 如果不存在, 则插入返回插入的下标(插入的时候要按顺序插入)
 * @Detail 注意边界问题
 * 1. 最简单的方式是遍历的同时,记录小于它的最大值的下标
 * 边界问题,要注意 参数为空的情况,以及值处于边界(首部和尾部的情况:既无相同的值,同时首位和末位都是小于的最大值)
 * @Learn 对于数组中查询的, 尤其如果是有序的, 可以充分参考二分查找, 解题步骤就是low和high, low
 * 一直都会是增加,而high则一直都会是减少,当low==high的时候意味着要么找到了,要么就是处于插入的位置
 * @date 创建时间：2019-12-10 18:18
 */
public class SearchInsertPositionNum35
{
    public int searchInsert(int[] nums, int target)
    {
        int low = 0, high = nums.length, mid = 0;
        while (low < high)
        {
            mid = low + (high - low) >> 1;
            // 因为插入的时候是要顺序插入的,要找到的值为小于target的最大值,因此当大于的时候都要移动
            if (target <= nums[mid])
            {
                high = mid;
            } else
            {
                low = mid + 1;
            }
        }
        return low;
    }
}
