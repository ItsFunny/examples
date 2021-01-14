package string;

/**
 * @author Charlie
 * @When
 * @Description 这道题被锁了, 需要交钱才能看, 百度搜了题目
 * 既 read(char[] buf, int n)  buf代表读的数据,而n代表的是期望读的数据,
 * 返回值并不是返回读取的数据,而是返回的读取的字节数,并不是像output所示
 * @Detail 1. 简单其实就是道数学题目,当调用read4的返回值小于4的时候代表读完了,而大于4代表需要依旧读
 * @Attention:
 * @Date 创建时间：2020-03-01 16:34
 */
public class String_157_read_n_characters_given_read4
{
    public int read(char[] buf, int n)
    {

        char[] result = new char[4];

        int readCount = 0;
        while (readCount < n)
        {
            int readed = read4(buf);
            if (readed < 4)
            {
                return readed;
            } else
            {
                System.arraycopy(buf, 0, result, 4, readCount);
            }
        }
        return n;
    }

    private int read4(char[] buf)
    {
        return 4;
    }

}
