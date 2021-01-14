package array;

import org.junit.Test;

public class TwoSumNum2Test
{

    @Test
    public void testForce()
    {
        int[] nums = {3,2,4};
        int targe = 6;
        int[] ints = TwoSumNum2.forceMatch(nums, targe);

        for (int anInt : ints)
        {

            System.out.println(anInt);
        }
    }

}
