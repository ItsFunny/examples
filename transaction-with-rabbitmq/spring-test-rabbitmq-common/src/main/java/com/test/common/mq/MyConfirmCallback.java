/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 下午12:49:34
* 
*/
package com.test.common.mq;

import org.springframework.amqp.rabbit.core.RabbitTemplate.ConfirmCallback;
import org.springframework.amqp.rabbit.support.CorrelationData;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.util.StringUtils;


/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 下午12:49:34
*/
public class MyConfirmCallback implements ConfirmCallback
{
	@Autowired
	private MessageContainer container;
	@Override
	public void confirm(CorrelationData correlationData, boolean ack, String cause)
	{
		if(ack)
		{
			String id = correlationData.getId();
			if(StringUtils.isEmpty(id))
			{
				//提示格式不正确的同时
				//格式校验应该不能在这里校验,更正确的应该是组转消息的时候判断是否组转成功
				throw new RuntimeException("发送的消息格式不正确");
			}
			//2代表成功
			//更新成功,这里的操作更应该都放在db中,而不是在内存中
			container.updateMessageStatus(2, Long.parseLong(id));
		}else {
			//尝试重新发送,如果发送不成功好像只能人为处理了
			//如果重新发送的次数失败,则需要额外记录到一张表中,人为处理(意味着db中的消息表需要一个重发次数的属性列)
		}
	}

}
