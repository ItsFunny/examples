/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午11:14:30
* 
*/
package com.test.cache;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import com.test.model.ProductDTO;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月16日 上午11:14:30
*/
public class OrderProductTestCache
{
	public static final  Map<Long, ProductDTO>PRODUCT_CACHE=new HashMap<>();
	
	public static final List<ProductDTO>PRODUCTS=new ArrayList<ProductDTO>();
}
