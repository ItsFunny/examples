package list;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;

/**
 * @author Charlie
 * @When
 * @Description
 * @Detail
 * @Attention:
 * @Date 创建时间：2020-02-22 15:32
 */
public class ListNode
{
    int val;
    ListNode next;

    ListNode(int x) { val = x; }


    static ListNode buildRandNodes(int limit)
    {

        Random random = new Random();
        List<ListNode> listNodes = new ArrayList<>();
        for (int i = 0; i < limit; i++)
        {
            listNodes.add(new ListNode(random.nextInt(100)));
        }
        return linkListNodes(listNodes);
    }

    public static ListNode buildInValues(int[] values)
    {
        List<ListNode> listNodes = new ArrayList<>();
        for (int i = 0; i < values.length; i++)
        {
            listNodes.add(new ListNode(values[i]));
        }
        return linkListNodes(listNodes);
    }

    private static ListNode linkListNodes(List<ListNode> listNodes)
    {
        for (int i = 0; i < listNodes.size() - 1; i++)
        {
            listNodes.get(i).next = listNodes.get(i + 1);
        }

        return listNodes.get(0);
    }

    static ListNode buildIteratorNodes(int maxValue)
    {
        List<ListNode> listNodes = new ArrayList<>();
        for (int i = 1; i <= maxValue; i++)
        {
            listNodes.add(new ListNode(i));
        }
        return linkListNodes(listNodes);
    }
}
