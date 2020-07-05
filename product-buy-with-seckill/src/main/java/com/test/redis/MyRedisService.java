/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月21日 下午7:27:04
* 
*/
package com.test.redis;

import java.util.Collections;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.ReentrantLock;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.dao.DataAccessException;
import org.springframework.data.redis.connection.RedisConnection;
import org.springframework.data.redis.connection.RedisStringCommands.SetOption;
import org.springframework.data.redis.connection.ReturnType;
import org.springframework.data.redis.core.RedisCallback;
import org.springframework.data.redis.core.StringRedisTemplate;
import org.springframework.data.redis.core.script.RedisScript;
import org.springframework.data.redis.core.types.Expiration;
import org.springframework.stereotype.Service;

import com.alibaba.druid.util.StringUtils;

import io.lettuce.core.RedisAsyncCommandsImpl;
import io.lettuce.core.RedisFuture;
import io.lettuce.core.SetArgs;
import lombok.extern.slf4j.Slf4j;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisCommands;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月21日 下午7:27:04
 */
@Slf4j
@Service
public class MyRedisService
{
	private static final String LOCK_SUCCESS = "OK";
	private static final String SET_IF_NOT_EXIST = "NX";
	private static final String SET_WITH_EXPIRE_TIME = "PX";

	private ReentrantLock lock = new ReentrantLock();
	@Autowired
	private StringRedisTemplate stringRedisTemplate;

	public boolean tryLock(String key, String value, int waitTimes)
	{
		return tryLock(key, value, waitTimes, false);
	}

	public boolean tryLock(String key, String value, int waitTimes, boolean isAlwaysWait)
	{
		long startTime = System.currentTimeMillis();
		Boolean res = false;
		if (isAlwaysWait)
		{
			while (!res)
			{
				res = stringRedisTemplate.execute(new RedisCallback<Boolean>()
				{
					@Override
					public Boolean doInRedis(RedisConnection connection) throws DataAccessException
					{
						return connection.set(key.getBytes(), value.getBytes(), Expiration.seconds(120),
								SetOption.SET_IF_ABSENT);
					}
				});
			}
		} else
		{
			while (!res && System.currentTimeMillis() - startTime <= waitTimes)
			{
				res = stringRedisTemplate.execute(new RedisCallback<Boolean>()
				{
					@Override
					public Boolean doInRedis(RedisConnection connection) throws DataAccessException
					{
						return connection.set(key.getBytes(), value.getBytes(), Expiration.seconds(120),
								SetOption.SET_IF_ABSENT);
					}
				});
			}
		}

		// 2019-01-23 22:00更
		// 既如果某个线程在获取锁然后执行完毕之后,不再占有锁,如果手动删除锁的时候失败了,则这个锁会一直存在
		// 于redis中,等待其自我销毁,很显然这不是我们想要的(因为等待的时候够执行很多任务了),所以我们这时候需要
		// 设定2个超时时间t1和t2,t1作为redis自动过期的时间,t2作为执行业务的时间,并且通常而言t2执行业务的时间
		// 是要小于自动过期时间的(因为如果大于过期时间,都已经自动删除了)
		// 注意点的话就是删除的时候通过双重检测机制防止别人同一时刻刚抢占的锁立马被删除了
		if (!res)
		{
			String tempJson = stringRedisTemplate.opsForValue().get(key);
			if (StringUtils.isEmpty(tempJson))
			{
				res = lockAgain(key, value);
			} else
			{
				String[] arrs = tempJson.split(":");
				long expireTime = Long.parseLong(arrs[1]);
				if (expireTime < System.currentTimeMillis())
				{
					try
					{
						lock.lock();
						// double check 这里的双重检测主要是为了防止其他线程
						// 又获取到了锁,而却被这个线程给删了
						tempJson = stringRedisTemplate.opsForValue().get(key);
						if (!StringUtils.isEmpty(tempJson)
								&& (expireTime = Long.parseLong(tempJson.split(":")[1])) < System.currentTimeMillis())
						{
							Boolean delete = stringRedisTemplate.delete(key);
							if (!delete)
							{
								log.error("[删除过期key]失败,key为:{}", key);
							} else
							{
								log.error("[删除过期的key]成功");
							}
						}
					} finally
					{
						lock.unlock();
					}
				}
			}
		}
		return res;
	}

	public static void main(String[] args)
	{
		long startTime = System.currentTimeMillis();
		while (System.currentTimeMillis() - startTime < 300)
		{
			System.out.println("1");
		}
	}

	public boolean lockAgain(String key, String value)
	{
		return stringRedisTemplate.execute(new RedisCallback<Boolean>()
		{
			@Override
			public Boolean doInRedis(RedisConnection connection) throws DataAccessException
			{
				// Jedis jedis = (Jedis) connection.getNativeConnection();
				return connection.set(key.getBytes(), value.getBytes(), Expiration.seconds(120),
						SetOption.SET_IF_ABSENT);
			}
		});
	}

	/**
	 * @param key
	 *            要解锁的key
	 * @param value
	 *            验证是否是同个用户
	 * @return true release success else fail
	 * @author joker
	 * @date 创建时间：2018年10月21日 下午8:44:01
	 */
	public boolean tryReleaseLock(String key, String value)
	{
		String script = "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end";
		Long res = stringRedisTemplate.execute(new RedisCallback<Long>()
		{
			@Override
			public Long doInRedis(RedisConnection connection) throws DataAccessException
			{
				Object eval = connection.eval(script.getBytes(), ReturnType.INTEGER, 1, key.getBytes(),
						value.getBytes());
				return (Long) eval;
			}
		});
		if (res > 0)
		{
			return true;
		}
		return false;
	}

}
