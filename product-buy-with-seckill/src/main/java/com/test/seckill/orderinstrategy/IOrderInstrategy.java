package com.test.seckill.orderinstrategy;

import com.test.model.Order;

/*
 * 订单插入策略,是直接db插入呢还是发布到消息队列中
 */
public interface IOrderInstrategy
{
	int createOrder(Order order);
}
