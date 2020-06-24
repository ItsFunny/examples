package com.basic.reference.soft;

import lombok.Data;

import java.lang.ref.ReferenceQueue;
import java.lang.ref.SoftReference;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-01 07:01
 */
@Data
public class SoftReferenceInfo<K,T> extends SoftReference<T>
{
    private K key;

    public SoftReferenceInfo(T referent)
    {
        super(referent);
    }

//    public SoftReferenceInfo(T referent, ReferenceQueue<? super T> q) { super(referent, q); }


    public SoftReferenceInfo( K key,T referent, ReferenceQueue<? super T> q)
    {
        super(referent, q);
        this.key = key;
    }
}
