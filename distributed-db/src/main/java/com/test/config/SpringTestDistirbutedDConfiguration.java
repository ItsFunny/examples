/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 上午8:46:01
* 
*/
package com.test.config;

import java.io.IOException;

import org.mybatis.spring.SqlSessionFactoryBean;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.beans.BeansException;
import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Primary;
import org.springframework.core.io.support.PathMatchingResourcePatternResolver;

import com.alibaba.druid.pool.DruidDataSource;
import com.joker.library.sqlextention.SQLExtentionConfigProperty;
import com.joker.library.sqlextention.SQLExtentionHolderV3;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月27日 上午8:46:01
 */
@Configuration
@MapperScan(basePackages = "com.test.dao.user.db0", sqlSessionFactoryRef = "sqlSessionFactoryBean0")
public class SpringTestDistirbutedDConfiguration implements ApplicationContextAware
{
	private ApplicationContext context;

	@Primary
	@Bean(value = "dataSource0")
	public DruidDataSource dataSource0()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setUrl("jdbc:mysql://localhost/test_distribute_db0?characterEncoding=utf-8&useSSL=false");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		return dataSource;
	}

	@Bean(value = "dataSource1")
	public DruidDataSource dataSource1()
	{
		DruidDataSource dataSource = new DruidDataSource();
		dataSource.setUsername("root");
		dataSource.setPassword("123456");
		dataSource.setUrl("jdbc:mysql://localhost/test_distribute_db1?characterEncoding=utf-8&useSSL=false");
		dataSource.setDriverClassName("com.mysql.jdbc.Driver");
		return dataSource;
	}

	@Primary
	@Bean(value = "sqlSessionFactoryBean0")
	public SqlSessionFactoryBean sqlSessionFactoryBean0() throws IOException
	{
		SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
		sqlSessionFactoryBean.setDataSource(dataSource0());
		org.apache.ibatis.session.Configuration configuration = new org.apache.ibatis.session.Configuration();
		configuration.setMapUnderscoreToCamelCase(true);
		sqlSessionFactoryBean.setConfiguration(configuration);
		sqlSessionFactoryBean.setMapperLocations(
				new PathMatchingResourcePatternResolver().getResources("classpath:mapper/db0/*.xml"));
		return sqlSessionFactoryBean;
	}

	@Bean(value = "sqlSessionFactoryBean1")
	public SqlSessionFactoryBean sqlSessionFactoryBean1() throws IOException
	{
		SqlSessionFactoryBean sqlSessionFactoryBean = new SqlSessionFactoryBean();
		sqlSessionFactoryBean.setDataSource(dataSource1());
		org.apache.ibatis.session.Configuration configuration = new org.apache.ibatis.session.Configuration();
		configuration.setMapUnderscoreToCamelCase(true);
		sqlSessionFactoryBean.setConfiguration(configuration);
		sqlSessionFactoryBean.setMapperLocations(
				new PathMatchingResourcePatternResolver().getResources("classpath:mapper/db1/*.xml"));

		return sqlSessionFactoryBean;
	}

	@Bean
	public SQLExtentionHolderV3 holderV3()
	{
		SQLExtentionHolderV3 sqlExtentionHolderV3 = new SQLExtentionHolderV3();
		SQLExtentionConfigProperty property = new SQLExtentionConfigProperty();
		property.setDetailConfigStr("2:userSQLExtentionProxy:3=user0,user1,user2;3=user0,user1,user2");
		property.setTablePrefixNames("user");
		property.setTotalTableCounts(1);
		sqlExtentionHolderV3.config(property, context);
		return sqlExtentionHolderV3;
	}

	@Override
	public void setApplicationContext(ApplicationContext applicationContext) throws BeansException
	{
		this.context = applicationContext;
	}

}
