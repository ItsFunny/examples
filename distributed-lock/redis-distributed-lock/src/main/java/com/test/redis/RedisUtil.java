/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月9日 下午1:05:09
* 
*/
package com.test.redis;

import java.util.Collections;

import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

/**
 * 
 * 没有关闭操作,懒得写,判断下即可
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月9日 下午1:05:09
 */
public class RedisUtil
{
	private static JedisPool jedisPool;

	private static final String SET_IF_NOT_EXIST = "NX";

	private static final String PREFIX_EXPIRE_TIME_MILLSECS = "PX";

	private static final Long LOCK_SELF_TIME = 3000 * 1L;

	public static final String SUCCESS_STRING = "OK";
	static
	{
		JedisPoolConfig jedisPoolConfig = new JedisPoolConfig();
		jedisPoolConfig.setMaxIdle(10);
		jedisPoolConfig.setMaxTotal(1024);
		jedisPoolConfig.setMaxWaitMillis(10000);
		jedisPoolConfig.setTestOnBorrow(true);
		// JedisPool jedisPool=new
		jedisPool = new JedisPool(jedisPoolConfig, "localhost", 6379, 10000, "123456");
		// JedisPool jedisPool = new JedisPool("localhost", 6379);
	}

	// 尝试设置锁
	public static boolean tryLock(Jedis jedis, String key, String value)
	{
		long startTime = System.currentTimeMillis();
		boolean getLock = false;
		while (System.currentTimeMillis() - startTime < LOCK_SELF_TIME && !getLock)
		{
			String res = jedis.set(key, value, SET_IF_NOT_EXIST, PREFIX_EXPIRE_TIME_MILLSECS, LOCK_SELF_TIME);
			if (SUCCESS_STRING.equals(res))
			{
				getLock = true;
			}
		}
		return getLock;
	}

	public static boolean unLock(Jedis jedis, String key, String identify)
	{
		String script = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end";
		Object result = jedis.eval(script, Collections.singletonList(key), Collections.singletonList(identify));
		if ((Long) result > 0)
		{
			return true;
		}
		return false;
	}

	public static Jedis getJedis()
	{
		return jedisPool.getResource();
	}

}
