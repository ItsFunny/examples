/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:41:18
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
* @date 创建时间：2018年9月18日 上午9:41:18
*/
@Mapper
public interface UserDao
{
	@Insert("insert into user (name ) values (#{name})")
	Integer insert(String name);
	

}
