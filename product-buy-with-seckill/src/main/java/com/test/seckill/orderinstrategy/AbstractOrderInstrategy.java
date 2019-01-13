/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午2:36:30
* 
*/
package com.test.seckill.orderinstrategy;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.data.redis.core.StringRedisTemplate;

import com.test.model.Order;
import com.test.mq.MQEventPublisher;
import com.test.mq.MQEventPublisher.AppEvent;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月23日 下午2:36:30
 */
@Slf4j
public abstract class AbstractOrderInstrategy implements IOrderInstrategy
{
	protected StringRedisTemplate stringRedisTemplate;

	protected MQEventPublisher publisher;

	protected abstract int doCreate(Order order);

	private void beforeTermited(AppEvent event)
	{
		log.info("[AbstractOrderInstrategy#beforeTermited]往延时队列中发送消息:{}", event);
		event.setType("order.nonconsumer");
		publisher.publish(event);
	}

	@Override
	public int createOrder(Order order)
	{
		int res = 0;
		AppEvent event = new AppEvent(order);
		try
		{
			res = doCreate(order);
			beforeTermited(event);
		} catch (Exception e)
		{
			log.error("[createOrder]处理订单失败,{}", event);
			return 0;
		}
		return res;
	}

	public StringRedisTemplate getStringRedisTemplate()
	{
		return stringRedisTemplate;
	}

	@Autowired
	public void setStringRedisTemplate(StringRedisTemplate stringRedisTemplate)
	{
		this.stringRedisTemplate = stringRedisTemplate;
	}

	public MQEventPublisher getPublisher()
	{
		return publisher;
	}

	@Autowired
	public void setPublisher(MQEventPublisher publisher)
	{
		this.publisher = publisher;
	}

}
