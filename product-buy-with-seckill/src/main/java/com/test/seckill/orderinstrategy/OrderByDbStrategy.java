/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午2:33:00
* 
*/
package com.test.seckill.orderinstrategy;

import org.springframework.beans.factory.annotation.Autowired;

import com.test.dao.OrderDao;
import com.test.model.Order;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月23日 下午2:33:00
*/
public class OrderByDbStrategy extends AbstractOrderInstrategy
{
	@Autowired
	private OrderDao orderDao;
	
	@Override
	protected int doCreate(Order order)
	{
		return  orderDao.createOrder(order);
	}

}
