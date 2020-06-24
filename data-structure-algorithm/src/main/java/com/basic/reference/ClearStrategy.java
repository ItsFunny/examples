package com.basic.reference;

import java.lang.ref.ReferenceQueue;

public interface ClearStrategy<T>
{
    void clear(ReferenceQueue<T> queue);

}
