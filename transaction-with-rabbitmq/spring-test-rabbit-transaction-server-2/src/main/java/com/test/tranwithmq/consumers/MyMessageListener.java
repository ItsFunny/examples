/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 下午1:13:53
* 
*/
package com.test.tranwithmq.consumers;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.rabbit.core.ChannelAwareMessageListener;

import com.rabbitmq.client.Channel;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 下午1:13:53
*/
public class MyMessageListener implements ChannelAwareMessageListener
{

	@Override
	public void onMessage(Message message, Channel channel) throws Exception
	{
		channel.basicAck(message.getMessageProperties().getDeliveryTag(), false);
	}

}
