package stack;

import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description Given a string S of lowercase letters, a duplicate removal consists of choosing two adjacent and equal letters, and removing them.
 * <p>
 * We repeatedly make duplicate removals on S until we no longer can.
 * <p>
 * Return the final string after all such duplicate removals have been made.  It is guaranteed the answer is unique.
 * 移除相连的元素,
 * @Detail 题目没理解, 刚开始理解为删除直至成为回文字符串, 并不是, 而是一直不停地删除, 知道剩下的元素不是相同的
 * 1. 解题思路就是很明确,用栈,如果与之前的相同,则把之前的弹出(此时栈中还存着以前的数据,并且题目要求相邻)
 * @Attention:
 * @Date 创建时间：2020-03-18 16:46
 */
public class Stack_1047_Remove_All_Adjacent_Duplicates_In_String
{
    public String removeDuplicates(String S)
    {
        Stack<Character> stack = new Stack<>();
        for (int i = 0; i < S.length(); i++)
        {
            if (!stack.isEmpty() && stack.peek().equals(S.charAt(i)))
            {
                // 两个相同,则出栈
                stack.pop();
            } else
            {
                stack.push(S.charAt(i));
            }
        }

        StringBuilder sb = new StringBuilder();
        for (Character character : stack)
        {
            sb.append(character);
        }
        return sb.toString();

    }



}
