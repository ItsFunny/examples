package template;

/**
 * @author joker
 * @When
 * @Description
 * 模板模式有三个角色:
 * 1. service接口,用于提供具体的方法
 * 2. 抽象类,抽象类会复写接口中的方法,将相同的部分代码在抽象类中实现,这层对外部而言是不可见的
 * 3. 具体实现类,继承抽象类,实现具体的细节
 * 如下所示IObjectService用于生成各种对象
 * AbstractObjectService对于一个参数都会有前置校验,这是相同的部分,因而可以抽出来放到抽象类中
 * 至于具体怎么找,找什么交给具体的实现类来写
 * <p>
 * 这样在调用的时候就可以解耦了 IObjectService objectService = new UserServiceImpl();
 * @Detail
 * @date 创建时间：2019-01-29 08:18
 */

interface IObjectService
{
    void findByName(String name);
}
abstract class AbstractObjectService implements IObjectService
{
    protected abstract String doFindObject(String user);
    public void findByName(String obj)
    {
        if (null == obj || "".equals(obj))
        {
            return;
        }
    }
}
class UserServiceImpl extends AbstractObjectService
{
    protected String doFindObject(String user)
    {
        return "find user:" + user;
    }
}

class AnimalServiceImpl extends AbstractObjectService
{
    protected String doFindObject(String animal)
    {
        return "find animal:" + animal;
    }
}
public class TemplatePattern
{
    public static void main(String[] args)
    {
        IObjectService objectService = new UserServiceImpl();
    }
}
