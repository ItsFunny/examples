package com.basic.algorithm.timewheel.jianzhioffer;

import java.util.Stack;

/**
 * @author joker
 * @When
 * @Description 用两个栈实现一个队列
 * @Detail 队列是先进先出的, 而栈是先进后出的
 * 核心就是入的时候,用于都是往某个特定的栈入,而出的时候先将这个入的栈中的元素全部压入到另外一个栈中
 * 记忆方法: 栈实现队列,他出他入
 * @date 创建时间：2019-05-17 20:06
 */
public class _5_two_stack_impl_queue
{
    class SpecialQueue
    {
        private Stack<Integer> pushStack;
        private Stack<Integer> popStack;

        public SpecialQueue()
        {
            this.pushStack = new Stack<>();
            this.popStack = new Stack<>();
        }

        public void push(int value)
        {
            this.pushStack.push(value);
        }

        public int pop()
        {
            if (!this.popStack.isEmpty())
            {
                return this.popStack.pop();
            }
            if (this.pushStack.isEmpty())
            {
                throw new RuntimeException("null");
            }
            while (!this.pushStack.isEmpty())
            {
                this.popStack.push(this.pushStack.pop());
            }
            return this.popStack.pop();
        }
    }

    public void test()
    {
        SpecialQueue specialQueue = new SpecialQueue();
        specialQueue.push(1);
        specialQueue.push(2);
        specialQueue.push(3);
        System.out.println(specialQueue.pop());
        System.out.println(specialQueue.pop());
        System.out.println(specialQueue.pop());
    }


    public static void main(String[] args)
    {
        new _5_two_stack_impl_queue().test();

    }

}
