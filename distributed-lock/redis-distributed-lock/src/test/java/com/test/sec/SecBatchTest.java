/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月9日 下午3:08:20
* 
*/
package com.test.sec;

import java.util.concurrent.Executors;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;
import org.springframework.web.client.RestTemplate;

import com.test.service.UserService;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月9日 下午3:08:20
 */
@SpringBootTest
@RunWith(SpringJUnit4ClassRunner.class)
public class SecBatchTest
{

	@Test
	public void testBatchBuy()
	{
//		UserService userService=new UserService();
//		String res = userService.buy(1L, 1l);
//		System.out.println(res);
		// 200个线程去请求buy
		for (int i = 0; i < 200; i++)
		{
			Thread thread=new Thread(new BuyRunnable());
			thread.start();
		}
	}

	private class BuyRunnable implements Runnable
	{


		@Override
		public void run()
		{
//			String res = restTemplate.getForObject(url, String.class);
//			System.out.println(res);
			UserService userService=new UserService();
			String res = userService.buy(1L, 1L);
			System.out.println(res);
		}

	}
}
