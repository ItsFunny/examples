/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月23日 下午3:14:45
* 
*/
package com.test.consumer;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.ChannelAwareMessageListener;

import com.rabbitmq.client.Channel;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月23日 下午3:14:45
*/
public class OrderFailConsumer implements ChannelAwareMessageListener
{

	@Override
	public void onMessage(Message message, Channel channel) throws Exception
	{
		
	}

}
