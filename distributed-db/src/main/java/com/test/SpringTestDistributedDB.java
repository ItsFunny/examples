package com.test;

import org.apache.ibatis.annotations.Mapper;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.FilterType;

/**
 * Hello world!
 *
 */
@SpringBootApplication
@MapperScan(basePackages = "com.test.dao.user.db1", sqlSessionFactoryRef = "sqlSessionFactoryBean1", annotationClass = Mapper.class)
@ComponentScan(basePackages =
{ "com.test" }, excludeFilters =
{ @ComponentScan.Filter(type = FilterType.ANNOTATION, value = Mapper.class) })
public class SpringTestDistributedDB
{
	public static void main(String[] args)
	{
		new SpringApplicationBuilder(SpringTestDistributedDB.class).run(args);
	}

}
