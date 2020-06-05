package stack;

import java.util.LinkedList;
import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description You're now a baseball game point recorder.
 * <p>
 * Given a list of strings, each string can be one of the 4 following types:
 * <p>
 * Integer (one round's score): Directly represents the number of points you get in this round.
 * "+" (one round's score): Represents that the points you get in this round are the sum of the last two valid round's points.
 * "D" (one round's score): Represents that the points you get in this round are the doubled data of the last valid round's points.
 * "C" (an operation, which isn't a round's score): Represents the last valid round's points you get were invalid and should be removed.
 * Each round's operation is permanent and could have an impact on the round before and the round after.
 * <p>
 * You need to return the sum of the points you could get in all the rounds.
 * @Detail 1. 就是看题意解题这种即可,但是被局限于tag 了,既然是stack的标签,因此一直在纠结stack如何实现
 * 其实用队列是最快的,或者说不是队列,而是可以通过下标获取元素的数据结构即可
 * @Attention: 1. linkedlist用peek 而不是用get,原因在于,get会抛出错误,而peek并不会,只会返回null
 * @Date 创建时间：2020-03-15 16:22
 */
public class Stack_682_Baseball_Game
{
    public int calPoints(String[] ops)
    {
        int currentAmount = 0;
        LinkedList<Integer> l = new LinkedList<>();
        for (String op : ops)
        {
            if (op.equals("C"))
            {
                currentAmount -= l.removeLast();
            } else if (op.equals("D"))
            {
                l.addLast(l.getLast() << 1);
                currentAmount += l.peekLast();
            } else if (op.equals("+"))
            {
                l.addLast(l.getLast() + l.get(l.size() - 2));
                currentAmount += l.peekLast();
            } else
            {
                int i = Integer.parseInt(op);
                l.addLast(i);
                currentAmount += i;
            }

        }

        return currentAmount;

    }

    public static void main(String[] args)
    {
        Stack_682_Baseball_Game g = new Stack_682_Baseball_Game();
        System.out.println(g.calPoints(new String[]{"5", "-2", "4", "C", "D", "9", "+", "+"}));
    }

}
