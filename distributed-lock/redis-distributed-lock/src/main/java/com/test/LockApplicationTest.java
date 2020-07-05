package com.test;

import javax.sql.DataSource;

import org.apache.commons.dbutils.QueryRunner;
import org.apache.ibatis.session.Configuration;
import org.apache.ibatis.session.SqlSessionFactory;
import org.mybatis.spring.SqlSessionFactoryBean;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.beans.factory.InitializingBean;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.web.client.RestTemplate;

import com.alibaba.druid.pool.DruidDataSource;
import com.alibaba.druid.support.json.JSONUtils;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.test.model.ProductInfo;
import com.test.redis.RedisUtil;

import net.sf.json.JSONObject;
import redis.clients.jedis.Jedis;

/**
 * Hello world!
 *
 */
@SpringBootApplication
@EnableScheduling
public class LockApplicationTest implements InitializingBean
{
	public static void main(String[] args)
	{
		new SpringApplicationBuilder(LockApplicationTest.class).run(args);
	}

//	@Bean
//	public SqlSessionFactoryBean sqlsessionFactorjy()
//	{
//		SqlSessionFactoryBean sqlSessionFactory = new SqlSessionFactoryBean();
//
//		sqlSessionFactory.setDataSource(dataSource());
//		Configuration configuration = new Configuration();
//		configuration.setMapUnderscoreToCamelCase(true);
//		sqlSessionFactory.setConfiguration(configuration);
//		return sqlSessionFactory;
//	}
	@Bean
	public RestTemplate restTemplate()
	{
		return new RestTemplate();
	}
	@Bean
	public DataSource dataSource()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setUrl("jdbc:mysql://localhost/test_distribute_lock?characterEncoding=utf-8&useSSL=false");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		return dataSource;
	}

	@Bean
	public QueryRunner queryRunner()
	{
		return new QueryRunner(dataSource());
	}

	@Override
	public void afterPropertiesSet() throws Exception
	{
		redisPrepare();
	}

	// 填充数据,准备数据
	private void redisPrepare()
	{
		// 假设商品的id为1
//		ProductInfo productInfo = new ProductInfo();
//		productInfo.setProductId(1 + "");
//		productInfo.setProductStock(500);

//		JSONObject jsonObject = JSONObject.fromObject(productInfo);
//		String json = jsonObject.toString();
		Jedis jedis = RedisUtil.getJedis();
		String isOk = jedis.set("product_" + 1, "80");
		System.out.println(isOk);
		
	}
}
