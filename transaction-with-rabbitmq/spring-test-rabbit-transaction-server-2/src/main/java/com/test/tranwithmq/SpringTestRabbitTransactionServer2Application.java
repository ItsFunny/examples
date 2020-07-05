package com.test.tranwithmq;

import javax.sql.DataSource;

import org.springframework.amqp.core.AcknowledgeMode;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.amqp.rabbit.listener.MessageListenerContainer;
import org.springframework.amqp.rabbit.listener.SimpleMessageListenerContainer;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.context.annotation.Bean;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;

import com.alibaba.druid.pool.DruidDataSource;
import com.test.tranwithmq.consumers.MyMessageListener;


@SpringBootApplication
@EnableDiscoveryClient
public class SpringTestRabbitTransactionServer2Application
{

	@Bean
	public DataSource dataSource()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		dataSource.setUrl("jdbc:mysql://localhost/rt_server_2?characterEncoding=utf-8&useSSL=false");
		return dataSource;
	}

	@Bean
	public DataSourceTransactionManager transactionManager()
	{
		DataSourceTransactionManager manager = new DataSourceTransactionManager();
		manager.setDataSource(dataSource());
		return manager;
	}
	
	@Bean
	public ConnectionFactory connectionFactory()
	{
		CachingConnectionFactory cachingConnectionFactory=new CachingConnectionFactory();
		cachingConnectionFactory.setHost("localhost");
		cachingConnectionFactory.setUsername("guest");
		cachingConnectionFactory.setPassword("123456");		
		return cachingConnectionFactory;
	}
	@Bean
	public RabbitTemplate rabbitTemplate()
	{
		RabbitTemplate rabbitTemplate=new RabbitTemplate();
		rabbitTemplate.setConnectionFactory(connectionFactory());
		return rabbitTemplate;
	}

	@Bean
	public MessageListenerContainer container()
	{
		SimpleMessageListenerContainer container=new SimpleMessageListenerContainer(connectionFactory());
		container.setAcknowledgeMode(AcknowledgeMode.MANUAL);
		container.setMessageListener(new MyMessageListener());
		container.setQueueNames("queue-test");
		
		return container;
	}
	
	public static void main(String[] args)
	{
		SpringApplication.run(SpringTestRabbitTransactionServer2Application.class, args);
	}
}
