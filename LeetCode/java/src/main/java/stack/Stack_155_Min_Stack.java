package stack;

/**
 * @author Charlie
 * @When
 * @Description Design a stack that supports push, pop, top,
 * and retrieving the minimum element in constant time.
 * <p>
 * push(x) -- Push element x onto stack.
 * pop() -- Removes the element on top of the stack.
 * top() -- Get the top element.
 * getMin() -- Retrieve the minimum element in the stack.
 * @Detail 手动实现一个栈, 栈为先进后出类型
 * @Attention:
 * 1. 唯一需要注意只有条件中的,以O(1)的时间查找到最小值,意味着需要保存下标或者是最小值
 * @Date 创建时间：2020-03-17 16:02
 */
public class Stack_155_Min_Stack
{
    private Integer[] values;
    private int index;
    private int minIndex;

    public Stack_155_Min_Stack()
    {
        this.values = new Integer[8];
    }

    public void push(int x)
    {
        ensureCap();
        if (null == this.values[minIndex] || x < this.values[minIndex])
        {
            minIndex = index;
        }
        this.values[index++] = x;
    }

    private void ensureCap()
    {
        if (this.index >= values.length >> 1)
        {
            // 扩容
            this.cap();
        }
    }

    private void cap()
    {
        Integer[] newValues = new Integer[this.values.length << 1];
        System.arraycopy(this.values, 0, newValues, 0, this.values.length);
        this.values = newValues;
    }

    public void pop()
    {
        this.values[--index] = null;

        if (this.values[minIndex] == null)
        {
            // 说明删除的是最后一个元素
            // 则此时minIndex直接设置为0,让其重新查找最小值
            this.minIndex = 0;
        }
        // 判断大小
        for (int i = this.minIndex; i < index; i++)
        {
            if (this.values[i] < this.values[minIndex])
            {
                this.minIndex = i;
            }
        }
    }

    public int top()
    {
        return this.values[index - 1];
    }

    public int getMin()
    {
        return this.values[this.minIndex];
    }

    public static void main(String[] args)
    {
        Stack_155_Min_Stack min_stack = new Stack_155_Min_Stack();
        min_stack.push(-2);
        min_stack.push(0);
        min_stack.push(-3);
        System.out.println(min_stack.getMin());
        min_stack.pop();
        System.out.println(min_stack.top());

        System.out.println(min_stack.getMin());

    }
}
