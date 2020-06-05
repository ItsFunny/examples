package array;

import org.junit.Test;

import static org.junit.Assert.*;

public class RemoveDuplicatesNum26Test
{

    @Test
    public void removeDuplicates()
    {
        RemoveDuplicatesNum26 removeDuplicatesNum26 = new RemoveDuplicatesNum26();
        int i = removeDuplicatesNum26.removeDuplicates(new Integer[]{1, 2, 3, 3, 3, 4, 5, 5});
        System.out.println(i);
    }

}