/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 上午8:55:16
* 
*/
package com.test.dao.user.db1;

import org.apache.ibatis.annotations.Mapper;
import org.springframework.core.annotation.Order;
import org.springframework.stereotype.Component;

import com.test.dao.user.IUserBaseDao;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月27日 上午8:55:16
*/
@Component(value="db1UserDao")
@Mapper
@Order(1)
public interface UserDao extends IUserBaseDao
{
	
}
