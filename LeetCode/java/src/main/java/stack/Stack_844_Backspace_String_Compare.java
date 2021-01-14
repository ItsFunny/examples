package stack;

import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description Share
 * Given two strings S and T,
 * return if they are equal when both are typed into empty text editors.
 * # means a backspace character.
 * 既 # 能取消一个 字符
 * @Detail
 * 1. 利用两个栈来取消字符即可
 * @Attention:
 * @Date 创建时间：2020-03-15 16:02
 */
public class Stack_844_Backspace_String_Compare
{
    public boolean backspaceCompare(String S, String T)
    {
        Stack<Character> ss = new Stack<>();
        Stack<Character> st = new Stack<>();

        for (int i = 0; i < S.length(); i++)
        {
            if (S.charAt(i) == '#')
            {
                if (!ss.isEmpty())
                {
                    ss.pop();
                }
            } else
            {
                ss.push(S.charAt(i));
            }
        }

        for (int i = 0; i < T.length(); i++)
        {
            if (T.charAt(i) == '#')
            {
                if (!st.isEmpty())
                {
                    st.pop();
                }
            } else
            {

                st.push(T.charAt(i));
            }
        }
        if (ss.size() != st.size())
        {
            return false;
        }
        for (int i = 0; i < ss.size(); i++)
        {
            if (!ss.pop().equals(st.pop()))
            {
                return false;
            }
        }


        return true;
    }

    public static void main(String[] args)
    {
        Stack_844_Backspace_String_Compare s = new Stack_844_Backspace_String_Compare();
        System.out.println(s.backspaceCompare("ab#c", "ad#c"));

    }
}
