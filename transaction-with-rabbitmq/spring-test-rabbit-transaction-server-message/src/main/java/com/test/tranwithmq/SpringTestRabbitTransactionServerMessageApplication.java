package com.test.tranwithmq;

import org.springframework.amqp.core.Binding;
import org.springframework.amqp.core.BindingBuilder;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.core.TopicExchange;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Bean;

import com.test.common.mq.AppEventPublisher;
import com.test.common.mq.MessageContainer;
import com.test.common.mq.MyConfirmCallback;

@SpringBootApplication
@EnableDiscoveryClient
public class SpringTestRabbitTransactionServerMessageApplication
{

	
	@Bean
	public AppEventPublisher eventPublisher()
	{
		return new AppEventPublisher();
	}
	@Bean
	public ConnectionFactory connectionFactory()
	{
		CachingConnectionFactory cachingConnectionFactory = new CachingConnectionFactory();
		cachingConnectionFactory.setHost("localhost");
		cachingConnectionFactory.setUsername("guest");
		cachingConnectionFactory.setPassword("guest");
		cachingConnectionFactory.setPublisherConfirms(true);
		return cachingConnectionFactory;
	}

	@Bean
	public RabbitTemplate rabbitTemplate()
	{
		RabbitTemplate rabbitTemplate = new RabbitTemplate();
		rabbitTemplate.setConfirmCallback(new MyConfirmCallback());
		rabbitTemplate.setConnectionFactory(connectionFactory());
		rabbitTemplate.setExchange("test");
		return rabbitTemplate;
	}

	@Bean
	public RabbitAdmin RabbitAdmin()
	{
		RabbitAdmin rabbitAdmin = new RabbitAdmin(connectionFactory());
		rabbitAdmin.declareExchange(testExchange());
		return rabbitAdmin;
	}

	@Bean
	public TopicExchange testExchange()
	{
		return new TopicExchange("test");
	}

	@Bean
	public Queue testQueue()
	{
		return new Queue("test-queue");
	}

	@Bean
	public Binding testBinding()
	{
		return BindingBuilder.bind(testQueue()).to(testExchange()).with("test".toUpperCase());
	}

	@Bean
	public MessageContainer container()
	{
		return new MessageContainer();
	}

	public static void main(String[] args)
	{
		SpringApplication.run(SpringTestRabbitTransactionServerMessageApplication.class, args);
	}
}
