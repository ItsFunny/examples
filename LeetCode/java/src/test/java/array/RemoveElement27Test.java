package array;

import org.junit.Test;

import static org.junit.Assert.*;

public class RemoveElement27Test
{

    @Test
    public void removeElement()
    {
        RemoveElement27 removeElement27=new RemoveElement27();
        int[] arrs=new int[]{2};
        int n = removeElement27.removeElement(arrs, 3);
        for(int i=0;i<n;i++)
        {
            if (arrs[i]==5)
            {
                throw  new RuntimeException("bad resp");
            }
        }
    }
}