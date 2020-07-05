package array;

import org.junit.Test;

import static org.junit.Assert.*;

public class SearchInsertPositionNum35Test
{

    @Test
    public void searchInsert()
    {
        SearchInsertPositionNum35 searchInsertPositionNum35 = new SearchInsertPositionNum35();
        int[] ints = new int[]{1, 3, 5, 6};
        int i = searchInsertPositionNum35.searchInsert(ints, 2);
        System.out.println(i
        );
    }
}