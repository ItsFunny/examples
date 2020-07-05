package stack;

import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description 用栈实现队列
 * 队列为先进先出
 * @Detail 两个栈实现队列
 * 栈为先进后出,而要做到先进先出,需要有一个stack专门push,另外一个stack专门pop
 * @Attention: 注意点就是, push的时候要把out的全部移过来, 保证顺序, pop的时候要把in的全部移过来
 * @Date 创建时间：2020-03-16 15:24
 */
public class Stack_232_Implement_Queue_using_Stacks
{
    Stack<Integer> in;
    Stack<Integer> out;

    /**
     * Initialize your data structure here.
     */
    public Stack_232_Implement_Queue_using_Stacks()
    {
        in = new Stack<>();
        out = new Stack<>();
    }

    /**
     * Push element x to the back of queue.
     */
    public void push(int x)
    {
        while (!out.isEmpty())
        {
            in.push(out.pop());
        }
        in.push(x);
    }

    /**
     * Removes the element from in front of queue and returns that element.
     */
    public int pop()
    {
        // 将in中的元素全部移动到out中
        while (!in.isEmpty())
        {
            out.push(in.pop());
        }

        return out.pop();
    }

    /**
     * Get the front element.
     */
    public int peek()
    {
        // 将in中的元素全部移动到out中
        while (!in.isEmpty())
        {
            out.push(in.pop());
        }

        return out.get(out.size() - 1);
    }

    /**
     * Returns whether the queue is empty.
     */
    public boolean empty()
    {
        return out.isEmpty() && in.isEmpty();
    }
}
