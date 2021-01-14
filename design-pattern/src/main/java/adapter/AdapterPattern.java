package adapter;

/**
 * @author joker
 * @When
 * @Description 适配器模式
 * @Detail 适配器模式适用于如下情况
 * 1.已经上线的接口, 需求新增, 但是原先的接口又不能发生变动, 这时候就可以使用适配器模式了
 * 2.或者两个不相关的接口,通过适配器模式兼容
 * 所以适配器模式起的是一个兼容的作用
 * <p>
 * 核心: 适配器模式的实现可以通过组合或者继承都是可以的
 * @date 创建时间：2019-02-04 21:28
 */

interface IUserService
{
    String name();
}

class UserServiceImpl implements IUserService
{

    @Override
    public String name()
    {
        return "joker";
    }
}

// 1. 新增的需求,但是假设上方的IUserService不能发生变动了,则引入适配器类
interface IUserAdapterService
{
    String ageWithName();
}

class UserServiceAdapter implements IUserAdapterService
{
    IUserService userService;

    public UserServiceAdapter(IUserService userService)
    {
        this.userService = userService;
    }

    @Override
    public String ageWithName()
    {
        return userService.name() + 20;
    }
}

public class AdapterPattern
{
    public static void main(String[] args)
    {
        IUserAdapterService userAdapterService = new UserServiceAdapter(new UserServiceImpl());
        System.out.println(userAdapterService.ageWithName());
    }

}
