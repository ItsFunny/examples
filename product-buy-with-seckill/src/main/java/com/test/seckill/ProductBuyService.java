/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午3:21:55
* 
*/
package com.test.seckill;

import java.util.Map;

import com.test.model.ProductDTO;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 下午3:21:55
 */
public interface ProductBuyService
{
	void killProduct(long productId,long userId,int buyNum,Map<String, Object>res);
}
