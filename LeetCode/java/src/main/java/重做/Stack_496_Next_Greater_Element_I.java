package 重做;

import java.util.HashMap;
import java.util.Map;
import java.util.ResourceBundle;
import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description You are given two arrays (without duplicates) nums1 and nums2 where nums1’s elements are subset of nums2.
 * Find all the next greater numbers for nums1's elements in the corresponding places of nums2.
 * <p>
 * The Next Greater Number of a number x in nums1 is the first greater number to its right in nums2. If it does not exist, output -1 for this number.
 * 1. 既 sum1 中的元素为sum2中的子元素
 * 2. 除元素需要大于sum1中的外并且sum2中的下标也需要大于sum1中的下标
 * 3. 并且这个元素需要为 大于这个元素的最小值
 * <p>
 * 并不是存储下标,而是直接存储这个值,也并非是大于这个元素的最小值,大于它的第一个值即可
 * @Detail 1. 最简单的方式是直接2个for循环一个一个匹配
 * 题目没搞懂
 * 存储的不是下标,是元素,但是这道题难度不应该是easy 而应该是medium
 * 未解决
 * @Attention:
 * @Date 创建时间：2020-03-17 16:27
 */
public class Stack_496_Next_Greater_Element_I
{
    /*
        大神的思路:
        Key observation:
Suppose we have a decreasing sequence followed by a greater number
For example [5, 4, 3, 2, 1, 6] then the greater number 6 is the next greater element for all previous numbers in the sequence

We use a stack to keep a decreasing sub-sequence, whenever we see a number x greater than stack.peek() we pop all elements less than x and for all the popped ones, their next greater element is x
For example [9, 8, 7, 3, 2, 1, 6]
The stack will first contain [9, 8, 7, 3, 2, 1] and then we see 6 which is greater than 1 so we pop 1 2 3 whose next greater element should be 6
     */
    public static int[] nextGreaterElement(int[] findNums, int[] nums)
    {
        Map<Integer, Integer> map = new HashMap<>(); // map from x to next greater element of x
        Stack<Integer> stack = new Stack<>();
        for (int num : nums)
        {
            // 如果栈中的元素小于父串的元素
            while (!stack.isEmpty() && stack.peek() < num)
            // 因为是降序,并且题目要求为找到这个值的next greater值,则这个pop的值与num的值必然是满足条件的
            {
                map.put(stack.pop(), num);
            }
            // 栈中的元素大于 这个数,则插入,这样既可以构造一个 降序的stack
            stack.push(num);
        }


        for (int i = 0; i < findNums.length; i++)
            findNums[i] = map.getOrDefault(findNums[i], -1);
        return findNums;
    }

}
