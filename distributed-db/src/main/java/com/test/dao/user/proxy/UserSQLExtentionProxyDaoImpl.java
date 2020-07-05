/**
*
* @Description
* @author joker 
* @date 创建时间：2018年10月27日 下午2:06:08
* 
*/
package com.test.dao.user.proxy;

import java.util.ArrayList;
import java.util.List;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import com.joker.library.sqlextention.AbstractSQLExtentionModel;
import com.joker.library.sqlextention.AbstractSQLExtentionpProxyBaseCRUDDao;
import com.joker.library.sqlextention.ISQLExtentionBaseCRUDDao;
import com.joker.library.sqlextention.SQLExtentionDaoWrapper;
import com.test.dao.user.IUserBaseDao;
import com.test.model.User;

/**
* 
* @When
* @Description
* @Detail
* @author joker 
* @date 创建时间：2018年10月27日 下午2:06:08
*/
@Component(value="userSQLExtentionProxy")
public class UserSQLExtentionProxyDaoImpl extends AbstractSQLExtentionpProxyBaseCRUDDao<User>
{
	//better user reflect 
	@Autowired
	private List<IUserBaseDao>daos;
	
	private final String USER_PREFIX_NAME="user";
	
	
	@Override
	public ISQLExtentionBaseCRUDDao<User> getDetailConfigDao(Number uniqueKey)
	{
		ISQLExtentionBaseCRUDDao<?> dao = daos.get((int) (uniqueKey.longValue()%daos.size()));
		return (ISQLExtentionBaseCRUDDao<User>) dao;
	}

	@Override
	public List<? extends ISQLExtentionBaseCRUDDao<User>> getAllDaos()
	{
		return (List<? extends ISQLExtentionBaseCRUDDao<User>>) this.daos;
	}

	@Override
	public Integer insertSelective(User t)
	{
		ISQLExtentionBaseCRUDDao<User> dao = getDetailConfigDao(t.getUniquekey());
		return dao.insertSelective(t);
	}

	@Override
	public int updateByExampleSelective(User record, Object example)
	{
		return 0;
	}

	@Override
	public List<User> selectByExample(Object example)
	{
		List<User>users=new ArrayList<User>();
		for (IUserBaseDao dao : daos)
		{
			users.addAll(dao.selectByExample(example));
		}
		return users;
	}

	@Override
	public long countByExample(Object example)
	{
		return 0;
	}

	@Override
	public int deleteByPrimaryKey(Integer uniquekey)
	{
		return 0;
	}

	@Override
	public int insert(User record)
	{
		SQLExtentionDaoWrapper<AbstractSQLExtentionModel> wrapper = holder.getConcreteDao(USER_PREFIX_NAME, record.getUniquekey());
		record.setTableName(wrapper.getTableName());
		return wrapper.getDao().insert(record);
	}

	@Override
	public User selectByPrimaryKey(String tableConcreteName, Number primaryKey)
	{
		return null;
	}

}
