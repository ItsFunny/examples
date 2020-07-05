/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 上午8:56:11
* 
*/
package com.test.dto;

import com.joker.library.sqlextention.AbstractSQLExtentionModel;

import lombok.Data;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月27日 上午8:56:11
 */
@Data
public class UserDTO extends AbstractSQLExtentionModel
{
	private long userId;
	private String userName;
	@Override
	public Number getUniquekey()
	{
		return userId;
	}

}
