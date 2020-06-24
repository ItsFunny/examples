package com.basic.structure.map;

import java.util.Iterator;
import java.util.Map;
import java.util.Set;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-20 11:33
 */
public class MapForeach
{
    // 通过key遍历value
    public void foreachMapByKey(Map<String, String> map)
    {
        Set<String> keys = map.keySet();
        for (String key : keys)
        {
            System.out.println(map.get(key));
        }
    }

    // 直接遍历所有的value
    public void foreachMapByValue(Map<String, String> map)
    {
        for (String value : map.values())
        {
            System.out.println(value);
        }
    }

    // 通过map.entrySet 的iterator遍历所有的key和value
    public void foreachMapByIterator(Map<String, String> map)
    {
        Iterator<Map.Entry<String, String>> entryIterator = map.entrySet().iterator();
        while (entryIterator.hasNext())
        {
            Map.Entry<String, String> objectEntry = entryIterator.next();
            System.out.printf("key=%s,value=%s", objectEntry.getKey(), objectEntry.getValue());
        }
    }

    // 通过map.entrySet遍历所有的key和value
    public void foreachMapByEntrySet(Map<String, String> map)
    {
        Set<Map.Entry<String, String>> entrySet = map.entrySet();
        for (Map.Entry<String, String> entry : entrySet)
        {
            System.out.printf("key=%s,value=%s", entry.getKey(), entry.getValue());
            System.out.println();
        }
    }

}
