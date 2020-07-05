/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午3:32:24
* 
*/
package com.test.mq;

import java.io.Serializable;
import java.util.Map;

import org.springframework.amqp.core.Message;
import org.springframework.amqp.core.MessageProperties;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import lombok.Data;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月16日 下午3:32:24
*/
@Service
public class MQEventPublisher
{
	@Data
	public static class AppEvent implements Serializable
	{
		public AppEvent(Serializable data)
		{
			this.data=data;
		}
		private static final long serialVersionUID = 9136622927279113117L;
		private String type;
		private Serializable data;
		private Map<String, Object>extProps;
	}
	@Autowired
	private RabbitTemplate rabbitTemplate;
	public void publish(AppEvent event)
	{
		this.rabbitTemplate.convertAndSend(event.getType(), event);
	}
	
}
