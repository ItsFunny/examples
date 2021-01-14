/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午2:57:13
* 
*/
package com.test.model;

import lombok.Data;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月16日 下午2:57:13
*/

@Data
public class OrderDetail
{
	private long orderDetailId;
	private long orderId;
	private long productId;
	private int buyNum;

}
