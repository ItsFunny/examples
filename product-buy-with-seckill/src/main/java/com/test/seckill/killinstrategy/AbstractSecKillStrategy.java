/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午2:53:52
* 
*/
package com.test.seckill.killinstrategy;

import java.math.BigDecimal;
import java.util.Map;

import org.springframework.beans.factory.annotation.Autowired;

import com.test.factory.OrderFactory;
import com.test.model.Order;
import com.test.model.ProductDTO;
import com.test.seckill.orderinstrategy.IOrderInstrategy;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月23日 下午2:53:52
 */
@Slf4j
public abstract class AbstractSecKillStrategy implements SecKillInstrantegy
{
	protected IOrderInstrategy orderInstrategy;

	public abstract void doKill(ProductDTO productDTO, long userId, Map<String, Object> params);
	
	//失败时候的回滚操作
	public abstract void rollBack(ProductDTO productDTO,long userId,Map<String, Object>params);

	@Override
	public void kill(ProductDTO productDTO, long userId, Map<String, Object> params)
	{
		doKill(productDTO, userId, params);
		if (params.containsKey("error"))
			return;
		// 生成订单,如果采用串行的方式,数据库操作不可以在这里进行,应该抛到消息队列中
		Order order = OrderFactory.createOrder(productDTO, userId, 0);
		log.info("[秒杀商品]生成订单{}",order);
		int validCount = orderInstrategy.createOrder(order);
		if (validCount == 0)
		{
			params.put("error", "insert order error,plz try later");
			rollBack(productDTO, userId,params);
		}else {
			params.put("orderId", order.getOrderId());
		}
	}

	public IOrderInstrategy getOrderInstrategy()
	{
		return orderInstrategy;
	}

	@Autowired
	public void setOrderInstrategy(IOrderInstrategy orderInstrategy)
	{
		this.orderInstrategy = orderInstrategy;
	}

}
