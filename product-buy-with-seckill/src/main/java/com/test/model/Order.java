/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 下午1:50:27
* 
*/
package com.test.model;

import java.io.Serializable;
import java.math.BigDecimal;

import lombok.Data;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月16日 下午1:50:27
*/
@Data
public class Order implements Serializable
{
	/**
	* 
	* @Description
	* @author joker 
	* @date 创建时间：2018年10月16日 下午4:05:45
	*/
	private static final long serialVersionUID = 8465393415970140403L;
	private long orderId;
	private long userId;
	private BigDecimal orderPayment;
	private int orderStatus;
	

}
