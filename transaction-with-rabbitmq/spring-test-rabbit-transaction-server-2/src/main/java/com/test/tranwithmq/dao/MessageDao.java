/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:32:15
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
* @date 创建时间：2018年9月18日 上午9:32:15
*/
@Mapper
public interface MessageDao
{
	@Insert("insert into message (detail) values (#{detail})")
	Integer insert(String detail);
}
