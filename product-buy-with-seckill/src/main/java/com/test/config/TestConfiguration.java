/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午10:39:45
* 
*/
package com.test.config;

import java.io.IOException;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.TimeUnit;

import javax.annotation.PostConstruct;

import org.mybatis.spring.SqlSessionFactoryBean;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.amqp.core.Binding;
import org.springframework.amqp.core.BindingBuilder;
import org.springframework.amqp.core.Queue;
import org.springframework.amqp.core.TopicExchange;
import org.springframework.amqp.rabbit.connection.CachingConnectionFactory;
import org.springframework.amqp.rabbit.connection.ConnectionFactory;
import org.springframework.amqp.rabbit.core.RabbitAdmin;
import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.autoconfigure.condition.ConditionalOnBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.core.io.support.PathMatchingResourcePatternResolver;
import org.springframework.data.redis.core.StringRedisTemplate;

import com.alibaba.druid.pool.DruidDataSource;
import com.google.gson.Gson;
import com.joker.library.service.IdWorkerService;
import com.joker.library.service.IdWorkerServiceTwitter;
import com.test.cache.OrderProductTestCache;
import com.test.constants.ProductConstants;
import com.test.consumer.OrderTimerMessageConsumer;
import com.test.dao.ProductDao;
import com.test.model.ProductDTO;
import com.test.seckill.NormalProductServiceImpl;
import com.test.seckill.ProductBuyService;
import com.test.seckill.SecKillProductServiceImpl;
import com.test.seckill.killinstrategy.ListSerializerKillProduct;
import com.test.seckill.killinstrategy.SecKillInstrantegy;
import com.test.seckill.orderinstrategy.IOrderInstrategy;
import com.test.seckill.orderinstrategy.OrderByDbStrategy;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 上午10:39:45
 */
@Configuration
@MapperScan(basePackages =
{ "com.test.dao" })
public class TestConfiguration
{
	@Autowired
	private ProductDao productDao;
	@Autowired
	private StringRedisTemplate stringRedisTemplate;

	
	@Bean
	public IOrderInstrategy orderStrategy()
	{
		IOrderInstrategy orderInstrategy=new OrderByDbStrategy();
		return orderInstrategy;
	}
	
	@Bean
	public SecKillInstrantegy instrantegy()
	{
		return new ListSerializerKillProduct();
	}

	@Bean
	public IdWorkerService idWorkerService()
	{
		IdWorkerService idWorkerService = new IdWorkerServiceTwitter(0L, 1L);
		return idWorkerService;
	}

	@Bean
	public DruidDataSource dataSource()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setUrl("jdbc:mysql://localhost/order_stock?characterEncoding=utf-8&useSSL=false");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		return dataSource;
	}

	@Bean
	public SqlSessionFactoryBean sqlSessionFactoryBean() throws IOException
	{
		SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
		org.apache.ibatis.session.Configuration configuration = new org.apache.ibatis.session.Configuration();
		configuration.setMapUnderscoreToCamelCase(true);
		sqlSessionFactoryBean
				.setMapperLocations(new PathMatchingResourcePatternResolver().getResources("classpath:mapper/*.xml"));
		sqlSessionFactoryBean.setConfiguration(configuration);
		sqlSessionFactoryBean.setDataSource(dataSource());
		return sqlSessionFactoryBean;
	}

	@Bean
	public ConnectionFactory connectionFactory()
	{
		CachingConnectionFactory cachingConnectionFactory = new CachingConnectionFactory();
		cachingConnectionFactory.setHost("localhost");
		cachingConnectionFactory.setPort(5672);
		cachingConnectionFactory.setUsername("guest");
		cachingConnectionFactory.setPassword("guest");
		return cachingConnectionFactory;
	}

	@Bean
	public RabbitAdmin rabbitAdmin()
	{
		RabbitAdmin rabbitAdmin = new RabbitAdmin(connectionFactory());
		rabbitAdmin.declareExchange(exchange());
		return rabbitAdmin;
	}

	@Bean
	public RabbitTemplate rabbitTemplate()
	{
		RabbitTemplate rabbitTemplate = new RabbitTemplate();
		rabbitTemplate.setConnectionFactory(connectionFactory());
		rabbitTemplate.setExchange("test_order");
		return rabbitTemplate;
	}

	@Bean
	public TopicExchange exchange()
	{
		return new TopicExchange("test_order");
	}

	@ConditionalOnBean(value = RabbitAdmin.class)
	@Bean
	public Queue nonConsumerOrderQueue()
	{
		Map<String, Object> args = new HashMap<String, Object>();
		// 超时后的转发器 过期转发到 expire_order_queue
		args.put("x-dead-letter-exchange", "test_order");
		// routingKey 转发规则
		args.put("x-dead-letter-routing-key", "order.expire");
		// 过期时间 20 秒
		// 这里设置过期时间,则内部的所有消息的过期时间是一致的,会覆盖消息的过期时间
		args.put("x-message-ttl", 1000*60*24);
		Queue queue = new Queue("nonconsumer_order_queue", false, false, false, args);
		return queue;
	}

	@ConditionalOnBean(value = RabbitAdmin.class)
	@Bean
	public Binding queueBinding()
	{
		Queue queue = nonConsumerOrderQueue();
		return BindingBuilder.bind(queue).to(exchange()).with("order.nonconsumer");
	}

	@ConditionalOnBean(value = RabbitAdmin.class)
	@Bean
	public Queue expireOrderQueue()
	{
		Queue queue = new Queue("expire_order_queue", true);
		return queue;
	}

	@ConditionalOnBean(value = RabbitAdmin.class)
	@Bean
	public Binding expireQueueBinding()
	{
		return BindingBuilder.bind(expireOrderQueue()).to(exchange()).with("order.expire");
	}

	@Bean
	public OrderTimerMessageConsumer consumer()
	{
		return new OrderTimerMessageConsumer();
	}

	// @Bean
	// public SimpleMessageListenerContainer container(List<Queue>queues)
	// {
	// SimpleMessageListenerContainer container=new
	// SimpleMessageListenerContainer(connectionFactory());
	// for (Queue queue : queues)
	// {
	// container.addQueueNames(queue.getName());
	// }
	// MessageListenerAdapter adapter=new MessageListenerAdapter();
	// adapter.setDefaultListenerMethod("process");
	// container.setAcknowledgeMode(AcknowledgeMode.MANUAL);
	// container.setMessageListener(adapter);
	// return container;
	// }
	@Bean
	public ProductBuyService productBuyService()
	{
		NormalProductServiceImpl normalProductServiceImpl = normalProductServiceImpl();
		normalProductServiceImpl.setNextHandler(secKillProductServiceImpl());
		return normalProductServiceImpl;
	}

	@Bean(name = "normalProduct")
	public NormalProductServiceImpl normalProductServiceImpl()
	{
		NormalProductServiceImpl normalProductServiceImpl = new NormalProductServiceImpl();
		normalProductServiceImpl.setType(1);
		return normalProductServiceImpl;
	}

	@Bean(name = "secKillProduct")
	public SecKillProductServiceImpl secKillProductServiceImpl()
	{
		SecKillProductServiceImpl secKillProductServiceImpl = new SecKillProductServiceImpl();
		secKillProductServiceImpl.setType(0);
		return secKillProductServiceImpl;
	}

	@PostConstruct
	public void initProperty()
	{
		List<ProductDTO> products = productDao.findAll();
		OrderProductTestCache.PRODUCTS.addAll(products);
		initStock();
		initProductInfo();
	}
	public void initStock()
	{
		List<ProductDTO> productDTOs = productDao.findAll();
		for (ProductDTO productDTO : productDTOs)
		{
			if (productDTO.getProductLevel() == 0)
			{
				stringRedisTemplate.opsForValue().set(ProductConstants.SEC_PRODUCT_STOCK + productDTO.getProductId(),
						productDTO.getProductStock() + "");
			}
		}
	}
	public void initProductInfo()
	{
		List<ProductDTO> products = OrderProductTestCache.PRODUCTS;
		Gson gson=new Gson();
		for (ProductDTO productDTO : products)
		{
			if (productDTO.getProductLevel() == 0)
			{
				String key = ProductConstants.SEC_PRODUCT_INFO + productDTO.getProductId();
				String value = gson.toJson(productDTO);
				stringRedisTemplate.opsForValue().set(key, value);
			}
		}
	}
}
