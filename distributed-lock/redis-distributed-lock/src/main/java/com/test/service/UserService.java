/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月9日 下午1:53:47
* 
*/
package com.test.service;

import static org.springframework.test.web.servlet.result.MockMvcResultHandlers.log;

import com.alibaba.druid.util.StringUtils;
import com.test.redis.RedisUtil;

import redis.clients.jedis.Jedis;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月9日 下午1:53:47
 */
public class UserService
{
	//失败的时候需要抛出异常的,迫使事务回滚
	public String buy(Long productId, Long userId)
	{
		Jedis jedis = RedisUtil.getJedis();
		boolean lock = RedisUtil.tryLock(jedis, "sec_" + productId, "" + userId);
		if (lock)
		{
			Integer stock = Integer.parseInt(jedis.get("product_" + productId));
			if (--stock < 0)
			{
				return "sold out";
			}
			String ok = jedis.set("product_" + 1, stock + "");
			// 模拟插入数据库
			try
			{
				Thread.sleep(100);
			} catch (InterruptedException e)
			{
				e.printStackTrace();
			}
			try
			{
				if (ok.equals("OK"))
				{
					return ok;
				} else
				{
					// 解锁,然后返回失败
					return "未知原因,更新redis中商品库存失败";
				}
			} finally
			{
				boolean unLock = RedisUtil.unLock(jedis, "sec_" + productId, userId + "");
				if(!unLock)
				{
					System.out.println("解锁失败");  
					return "解锁失败";
				}
			}
		} else
		{
			return "网络拥挤,请稍后再试";
		}
	}
}
