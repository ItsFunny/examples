/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午2:33:19
* 
*/
package com.test.seckill.orderinstrategy;

import org.springframework.beans.factory.annotation.Autowired;

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
 * @date 创建时间：2018年10月23日 下午2:33:19
 */
@Slf4j
public class OrderByMQStrategy extends AbstractOrderInstrategy
{

	@Override
	protected int doCreate(Order order)
	{
		AppEvent event = new AppEvent(order);
		publisher.publish(event);
		log.info("[OrderByMQStrategy生成订单],往订单处理队列发送消息:{}", event);
		return 1;
	}

}
