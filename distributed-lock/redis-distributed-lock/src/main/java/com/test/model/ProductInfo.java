/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月9日 下午1:12:40
* 
*/
package com.test.model;

import java.io.Serializable;

import lombok.Data;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月9日 下午1:12:40
 */
@Data
public class ProductInfo implements Serializable
{
	/**
	* 
	* @Description
	* @author joker 
	* @date 创建时间：2018年10月9日 下午1:13:06
	*/
	private static final long serialVersionUID = 7406961254680572959L;
	private String productId;
	private Integer productStock;

}
