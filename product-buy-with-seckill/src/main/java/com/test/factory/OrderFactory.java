/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午2:57:22
* 
*/
package com.test.factory;

import java.math.BigDecimal;

import com.test.model.Order;
import com.test.model.ProductDTO;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月23日 下午2:57:22
*/
public class OrderFactory
{
	public static Order createOrder(ProductDTO productDTO,long userId,int status)
	{
		Order order = new Order();
		order.setOrderPayment(productDTO.getProductPrice().multiply(new BigDecimal(productDTO.getBuyNum())));
		order.setOrderStatus(status);
		order.setUserId(userId);
		return order;
	}
}
