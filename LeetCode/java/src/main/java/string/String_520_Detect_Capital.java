package string;

/**
 * @author Charlie
 * @When
 * @Description Given a word, you need to judge whether the usage of capitals in it is right or not.
 * <p>
 * We define the usage of capitals in a word to be right when one of the following cases holds:
 * <p>
 * All letters in this word are capitals, like "USA".
 * All letters in this word are not capitals, like "leetcode".
 * Only the first letter in this word is capital, like "Google".
 * Otherwise, we define that this word doesn't use capitals in a right way.
 * 既全是大写的,或者全是小写的,或者只有首个是大写的
 * @Detail 1. 遍历判断
 * 2. 既然题目没要求的话,直接内置api解决是最快的
 * @Attention: a~z : 97~122
 * A~Z : 65~90
 * 1. 如果第一个是小写的话,则后面的必须为全是小写
 * @Date 创建时间：2020-03-10 18:54
 */
public class String_520_Detect_Capital
{

    public boolean detectCapitalUse(String word)
    {
        if (word.length() < 2) return true;
        if (word.toUpperCase().equals(word)) return true;
        if (word.substring(1).toLowerCase().equals(word.substring(1))) return true;
        return false;
    }

    public static void main(String[] args)
    {
        System.out.println((int) 'Z');
    }
}
