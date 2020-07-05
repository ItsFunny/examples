/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月22日 上午11:25:35
* 
*/
package com.test;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingQueue;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import com.test.seckill.ProductBuyService;

import lombok.extern.slf4j.Slf4j;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月22日 上午11:25:35
 */
@Slf4j
@RunWith(SpringJUnit4ClassRunner.class)
@SpringBootTest(classes = OrderTestApplication.class)
public class SecKillTest
{
	@Autowired
	private ProductBuyService productBuyService;

	/*
	 * 高并发下的测试
	 */
	@Test
	public void testManyThreads()
	{
		ExecutorService service = Executors.newFixedThreadPool(100);
		long stTime = System.currentTimeMillis();
		LinkedBlockingQueue<Integer>list=new LinkedBlockingQueue<>();
		CountDownLatch countDownLatch = new CountDownLatch(10000);
		for (int i = 0; i < 10000; i++)
		{
			final int userId = i;
			service.execute(new Runnable()
			{
				@Override
				public void run()
				{
					Map<String, Object> params = new HashMap<>();
					productBuyService.killProduct(1L, userId, 1, params);
					if (params.containsKey("error"))
					{
//						log.error("[失败]userId:{},error:{}",userId,params.get("error"));
					}else {
						list.add(userId);
					}
					countDownLatch.countDown();
				}
			});
		}
		try
		{
			countDownLatch.await();
			log.info("[所有任务结束]总耗时:{}",System.currentTimeMillis()-stTime);
			log.info("[总共有{}是成功购买到的]",list.size());
		} catch (InterruptedException e)
		{
			e.printStackTrace();
		}
	}
}
