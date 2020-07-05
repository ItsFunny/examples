package string;

/**
 * @author Charlie
 * @When
 * @Description 很简单的一道题,
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-03-12 16:35
 */
public class String_551_Student_Attendance_Record_I
{
    public boolean checkRecord(String s)
    {
        return !s.matches(".*LLL.*|.*A.*A.*");
    }
}
