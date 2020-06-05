package stack;

/**
 * @author Charlie
 * @When
 * @Description A valid parentheses string is either empty (""), "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.  For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.
 * <p>
 * A valid parentheses string S is primitive if it is nonempty, and there does not exist a way to split it into S = A+B, with A and B nonempty valid parentheses strings.
 * <p>
 * Given a valid parentheses string S, consider its primitive decomposition: S = P_1 + P_2 + ... + P_k, where P_i are primitive valid parentheses strings.
 * <p>
 * Return S after removing the outermost parentheses of every primitive string in the primitive decomposition of S.
 * 去除最外层的括号对
 * @Detail 没思路
 * 看了别人的其实好像挺简单的,核心就是根据要求,既然 () 是成对出现的,则意味着最终 ( 和 ) 的数量会一致
 * @Attention:
 * @Date 创建时间：2020-03-18 16:06
 */
public class Stack_1021_Remove_Outermost_Parentheses
{
    public String removeOuterParentheses(String S)
    {
        int left = 0, right = 0, index = 0;
        StringBuilder res = new StringBuilder();
        for (int i = 0; i < S.length(); i++)
        {
            if (S.charAt(i) == '(')
            {
                left++;
            } else
            {
                right++;
            }
            if (left == right)
            {
                res.append(S.substring(index + 1, i));
                index = i + 1;
                left = right = 0;
            }
        }


        return res.toString();
    }

}
