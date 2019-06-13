/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 上午10:24:52
* 
*/
package com.test.service.impl;

import java.util.List;
import java.util.Map;

import javax.annotation.PostConstruct;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.joker.library.page.AbstractMultipartDBPageService;
import com.joker.library.page.PageRequestDTO;
import com.joker.library.page.PageResponseDTO;
import com.joker.library.sqlextention.AbstractSQLExtentionModel;
import com.joker.library.sqlextention.AbstractSQLExtentionpProxyBaseCRUDDao;
import com.joker.library.sqlextention.ISQLExtentionBaseCRUDDao;
import com.joker.library.sqlextention.ISQLExtentionConfigBaseCRUDDao;
import com.joker.library.sqlextention.SQLExtentionDaoWrapper;
import com.joker.library.sqlextention.SQLExtentionHolderV3;
import com.test.dao.user.IUserBaseDao;
import com.test.model.User;
import com.test.model.example.UserExample;
import com.test.service.IUserService;

/**
 * 
 * @When
 * @Description
 * @Detail
 * @author joker
 * @date 创建时间：2018年10月27日 上午10:24:52
 */
@Service
public class UserServiceImpl extends AbstractMultipartDBPageService<User, UserExample> implements IUserService
{
	private final String USER_TABLE_CONSTANT_PREFIX_NAME = "user";
	@Autowired
	private SQLExtentionHolderV3 holder;

	private AbstractSQLExtentionpProxyBaseCRUDDao<User> proxyDao;

	@PostConstruct
	public void afterPropertiesSet()
	{
		AbstractSQLExtentionpProxyBaseCRUDDao<User> proxy = holder.getProxyDao(USER_TABLE_CONSTANT_PREFIX_NAME);
		if (null == proxy)
			throw new RuntimeException(
					"配置错误," + USER_TABLE_CONSTANT_PREFIX_NAME + "对应的proxyDao不存在,confirm the proxy dao has registed");
		this.proxyDao = proxy;
	}

	@Override
	public int insert(User t)
	{
		return this.proxyDao.insert(t);
	}

	@Override
	protected UserExample getExample(Map<String, Object> condition)
	{
		UserExample example = new UserExample();
		if (condition.isEmpty())
			return example;
		return null;
	}

	@Override
	protected List<User> doFindByExample(String tableName, ISQLExtentionBaseCRUDDao<User> dao, Integer avgStart,
			Integer end, UserExample exampleObj)
	{
		exampleObj.setStart(avgStart);
		exampleObj.setEnd(end);
		exampleObj.setTableName(tableName);
		return dao.selectByExample(exampleObj);
	}

	@Override
	protected Long getMinId(List<List<User>> list)
	{
		long minId = 0l;
		// 默认情况下是第一个是最小的
		for (List<User> list2 : list)
		{
			long id = 0l;
			if (null != list2 && !list2.isEmpty() && (minId > (id = list2.get(0).getUserId())))
				minId = id;
		}
		return minId;
	}

	@Override
	protected void getMaxId(List<Long> maxId, List<List<User>> totalList)
	{
		// default the max is the last value
		for (List<User> list : totalList)
		{
			if (null != list && !list.isEmpty())
				maxId.add(list.get(list.size() - 1).getUserId());
		}
	}

	@Override
	protected List<User> secondFindByBetween(String concreteTableName, ISQLExtentionBaseCRUDDao<User> dao, long min,
			long max, Map<String, Object> condition)
	{
		// in case ,ignore the condition
		return ((IUserBaseDao) dao).findByUserIdBetween(concreteTableName, min, max);
	}

	@Override
	protected void wrapRequest(PageRequestDTO pageRequestDTO)
	{
		pageRequestDTO.setTablePrefixName(USER_TABLE_CONSTANT_PREFIX_NAME);
	}

	@Override
	public PageResponseDTO<List<User>> findByPage(int pageSize, int pageNum, Map<String, Object> conditions)
	{
		PageRequestDTO pageRequestDTO=new PageRequestDTO();
		pageRequestDTO.setPageNum(pageNum);
		pageRequestDTO.setPageSize(pageSize);
		//这里自行根据业务判断是否是依据主键查询
		return findByPage(pageRequestDTO);
	}
}
