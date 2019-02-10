package com.test.tranwithmq;

import javax.sql.DataSource;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cloud.client.discovery.EnableDiscoveryClient;
import org.springframework.cloud.openfeign.EnableFeignClients;
import org.springframework.context.annotation.Bean;
import org.springframework.jdbc.datasource.DataSourceTransactionManager;

import com.alibaba.druid.pool.DruidDataSource;

@SpringBootApplication
@EnableDiscoveryClient
@EnableFeignClients(basePackages="com.test.message")
public class SpringTestRabbitTransactionApplication
{
	@Bean
	public DataSource dataSource()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		dataSource.setUrl("jdbc:mysql://localhost/rt_server_1?characterEncoding=utf-8&useSSL=false");
		return dataSource;
	}
	@Bean
	public DataSourceTransactionManager transactionManager()
	{
		DataSourceTransactionManager manager=new DataSourceTransactionManager();
		manager.setDataSource(dataSource());
		return manager;
	}
	public static void main(String[] args)
	{
		SpringApplication.run(SpringTestRabbitTransactionApplication.class, args);
	}
}
