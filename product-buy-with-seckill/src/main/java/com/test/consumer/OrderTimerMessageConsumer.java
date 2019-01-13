/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午3:43:30
* 
*/
package com.test.consumer;

import org.springframework.amqp.rabbit.annotation.RabbitHandler;
import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.test.dao.OrderDao;
import com.test.dao.OrderDetailDao;
import com.test.model.Order;
import com.test.mq.MQEventPublisher.AppEvent;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description 这个是处理没有付钱订单过期的
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午3:43:30
 */
@Service
@RabbitListener(queues =
{ "expire_order_queue" })
@Slf4j
public class OrderTimerMessageConsumer implements MessageConsumer
{
	@Autowired
	private OrderDao orderDao;
	@Autowired
	private OrderDetailDao orderDetailDao;

	@RabbitHandler
	@Override
	public void process(AppEvent event)
	{
		log.info("[删除过期订单]收到消息:{}", event);
		Order order = (Order) event.getData();
		long orderId = order.getOrderId();
		int validCOunt = orderDao.deleteByIdAndStatus(orderId, 0);
		if (validCOunt >= 1)
		{
			log.info("[删除过期订单],删除了{}条数据", validCOunt);
			int detailValidCount = orderDetailDao.deleteByOrderId(orderId);
			log.info("[删除过期订单详情],删除{}条数据", detailValidCount);
		}
	}

	@Override
	public String getType()
	{
		return "";
	}

}
