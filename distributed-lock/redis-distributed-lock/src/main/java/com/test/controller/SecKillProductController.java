/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月9日 下午1:08:59
* 
*/
package com.test.controller;

import org.springframework.http.MediaType;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.test.redis.RedisUtil;
import com.test.service.UserService;

import redis.clients.jedis.Jedis;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月9日 下午1:08:59
 */

@RestController
@RequestMapping(value = "/sec")
public class SecKillProductController
{
	// 显示商品信息
	@RequestMapping(value = "/info", produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
	public String showProductInof()
	{
		// 只是为了演示的话,只显示商品额库存和商品id
		Jedis jedis = RedisUtil.getJedis();
		String json = jedis.get("product_1");
		return json;
	}

	// 这里实际是需要传参的,demo直接固定
	@RequestMapping(value = "/buy/{productId}/{userId}", produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
	public String buyProduct(@PathVariable("productId") long productId, @PathVariable("userId") long userId)
	{
		UserService userService = new UserService();
		return userService.buy(productId, userId);
	}

	@RequestMapping(value = "/buy", produces = MediaType.APPLICATION_JSON_UTF8_VALUE)
	public String buyWithOutParam()
	{
		UserService userService = new UserService();
		return userService.buy(1L,1L);
	}

}
