/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:39:47
* 
*/
package com.test.tranwithmq.dao;

import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午9:39:47
*/
@Mapper
public interface AgeDao
{
		@Insert("insert into age (age) values (#{age})")
		Integer insert(Integer age);
}
