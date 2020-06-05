package string;

import com.sun.org.apache.xml.internal.utils.SystemIDResolver;

import java.util.HashMap;
import java.util.Map;
import java.util.Stack;

/**
 * @author Charlie
 * @When
 * @Description Given a string containing just the characters
 * '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
 * <p>
 * An input string is valid if:
 * <p>
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 * Note that an empty string is also considered valid.
 * 既必须是对称的,类似于回文,但也不同
 * @Detail 1. 用栈来做
 * 2. 更简便的方式是,同样是栈,但是可以省略值的赋值,既将特定类型的朝左或者朝右的数据入栈,而不是用map保存
 * @Attention: // 遇到的坑:
 * 1. 是获取栈首位的元素,与数组遍历的下标index是无关的
 * 2. 在遍历stack为空时,下标忘记递增
 * @Date 创建时间：2020-02-28 16:09
 */
public class String_20_Valid_Parentheses
{
    public static boolean isValid(String s)
    {
        if (s == null || s.length() % 2 != 0) return false;

        Stack<Character> stack = new Stack<>();
        char[] chars = s.toCharArray();
        int i = 0;

        Map<Character, Character> leftMap = new HashMap<>();
        leftMap.put('(', ')');
        leftMap.put('[', ']');
        leftMap.put('{', '}');
        Map<Character, Character> rightMap = new HashMap<>();
        rightMap.put(')', '(');
        rightMap.put(']', '[');
        rightMap.put('}', '{');

        Character expect = '(';
        while (i < chars.length)
        {
            Character cur = chars[i];
            i++;
            if (stack.isEmpty())
            {
                stack.push(cur);
                continue;
            }
            Character prev = stack.get(stack.size() - 1);
            if (leftMap.containsKey(cur))
            {
                expect = leftMap.get(cur);
            } else
            {
                expect = rightMap.get(cur);
            }
            if (expect.equals(prev))
            {
                stack.pop();
            } else
            {
                stack.push(cur);
            }


        }

        return stack.isEmpty();
    }

    public boolean isValid2(String s)
    {
        Stack<Character> stack = new Stack<Character>();
        for (char c : s.toCharArray())
        {
            if (c == '(')
            {
                stack.push(')');
            } else if (c == '{')
            {
                stack.push('}');
            } else if (c == '[')
            {
                stack.push(']');
            } else if (stack.isEmpty() || stack.pop() != c)
            {
                return false;
            }
        }
        return stack.isEmpty();
    }

    public static void main(String[] args)
    {
        System.out.println(String_20_Valid_Parentheses.isValid("()"));
    }
}
