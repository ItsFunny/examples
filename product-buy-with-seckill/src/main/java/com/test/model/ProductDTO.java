/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月16日 上午10:54:31
* 
*/
package com.test.model;

import java.math.BigDecimal;

import lombok.Data;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月16日 上午10:54:31
 */
@Data
public class ProductDTO
{
	private long productId;
	private String productName;
	private BigDecimal productPrice;
	private Integer productStock;
	private int productLevel;
	//
	private int buyNum;
}
