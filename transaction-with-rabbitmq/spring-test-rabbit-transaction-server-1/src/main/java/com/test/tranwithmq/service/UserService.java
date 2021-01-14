/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:53:06
* 
*/
package com.test.tranwithmq.service;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.test.tranwithmq.dao.UserDao;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午9:53:06
*/
@Service
public class UserService
{
	@Autowired
	private UserDao userDao;
	
	public Integer insert(String name)
	{
		return userDao.insert(name);
//		throw new RuntimeException("手动抛出异常尝试本地业务失败");
	}

}
