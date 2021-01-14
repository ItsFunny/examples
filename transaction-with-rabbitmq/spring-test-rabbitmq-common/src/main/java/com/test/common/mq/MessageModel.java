/**
*
* @Description
* @author joker 
* @date 创建时间：2018年9月18日 上午9:57:04
* 
*/
package com.test.common.mq;

import java.io.Serializable;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年9月18日 上午9:57:04
*/
public class MessageModel implements Serializable
{
	private Long id;
	private Integer status;
	private String detail;
	public Integer getStatus()
	{
		return status;
	}
	public void setStatus(Integer status)
	{
		this.status = status;
	}
	public String getDetail()
	{
		return detail;
	}
	public void setDetail(String detail)
	{
		this.detail = detail;
	}
	public Long getId()
	{
		return id;
	}
	public void setId(Long id)
	{
		this.id = id;
	}
	
}
