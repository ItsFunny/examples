/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 下午12:56:40
* 
*/
package com.test;

import java.util.List;

import com.test.model.User;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import com.joker.library.page.PageRequestDTO;
import com.joker.library.page.PageResponseDTO;
import com.test.service.IUserService;

import lombok.extern.slf4j.Slf4j;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月27日 下午12:56:40
*/
@SpringBootTest
@RunWith(SpringJUnit4ClassRunner.class)
@Slf4j
public class UserTest
{
	@Autowired
	private IUserService userService;
	
	@Test
	public void testFindByPage()
	{
		PageRequestDTO pageRequestDTO=new PageRequestDTO();
		pageRequestDTO.setPageNum(2);
		pageRequestDTO.setPageSize(1);
		PageResponseDTO<List<User>> pageResp = userService.findByPage(pageRequestDTO);
		List<User> users = pageResp.getData();
		for (User user : users)
		{
			System.out.println(user);
		}
		
	}
	@Test
	public void testInsert()
	{
		long startTime=System.currentTimeMillis();
		log.info("[begin batch insert]");
		for(int i=0;i<10000;i++)
		{
			User user=new User();
			user.setUserId(i);
			user.setUserName("joker__"+i);
			int validCount=userService.insert(user);
			System.out.println(validCount);
		}
		log.info("[batch insert finished] consume :{} ms ",System.currentTimeMillis()-startTime);
	}
}
