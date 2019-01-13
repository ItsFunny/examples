/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 下午1:18:47
* 
*/
package com.test.common.mq;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.concurrent.ConcurrentHashMap;

import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年9月18日 下午1:18:47
 */
public class AppEventPublisher
{
	@Autowired
	private RabbitTemplate rabbitTemplate;
	
	public void publishe(AppEvent event)
	{
		this.rabbitTemplate.convertAndSend(event.getType(), event);
		
	}
	
	public static class AppEvent implements Serializable
	{
		/**
		* @Description
		* @author joker 
		* @date 创建时间：2018年10月3日 下午7:47:56
		*/
		private static final long serialVersionUID = -6403831571484111761L;
		private Serializable data;
		private Long id;
		private String type;
		//0 新建 1准备 2发送成功 3发送失败 4消费成功
		private Integer status;
		
		public Serializable getData()
		{
			List<Integer>list=new ArrayList<>();
			return data;
		}

		public void setData(Serializable data)
		{
			this.data = data;
		}


		public String getType()
		{
			return type;
		}

		public void setType(String type)
		{
			this.type = type;
		}

		public Integer getStatus()
		{
			return status;
		}

		public void setStatus(Integer status)
		{
			this.status = status;
		}

		public Long getId()
		{
			return id;
		}

		public void setId(Long id)
		{
			this.id = id;
		}

		public static long getSerialversionuid()
		{
			return serialVersionUID;
		}

	}

	public RabbitTemplate getRabbitTemplate()
	{
		return rabbitTemplate;
	}

	public void setRabbitTemplate(RabbitTemplate rabbitTemplate)
	{
		this.rabbitTemplate = rabbitTemplate;
	}


	
}
