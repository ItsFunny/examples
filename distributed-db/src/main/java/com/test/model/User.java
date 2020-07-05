/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 上午10:24:20
* 
*/
package com.test.model;

import com.joker.library.sqlextention.AbstractSQLExtentionModel;

import lombok.Data;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月27日 上午10:24:20
*/
@Data
public class User extends AbstractSQLExtentionModel
{
	private long userId;
	private String userName;
	@Override
	public Number getUniquekey()
	{
		return userId;
	}
}
