package responsibility;

/**
 * @author joker
 * @When
 * @Description 责任链模式
 * 几个重要的角色:
 * 1. 接口,用于外抛给外部调用
 * 2. 事件对象,被处理的角色
 * 2. handler角色,用于处理不同的事件角色
 * @Detail
 * @date 创建时间：2019-01-30 16:51
 */

interface IUserService
{
    void login(UserBO userBO);
}

class UserBO
{
    byte level;
    String userName;
    String passWord;

    public UserBO(byte level, String userName, String passWord)
    {
        this.level = level;
        this.userName = userName;
        this.passWord = passWord;
    }


}


abstract class AbstractUserServiceHandler implements IUserService
{
    protected byte level;    // 用于判断不同的实现类处理不同的角色
    protected IUserService nextHandler;

    protected abstract void doLogin(String userName, String password);

    protected AbstractUserServiceHandler(byte level)
    {
        this.level = level;
    }

    public void login(UserBO userBO)
    {
        if (userBO.level == this.level)
        {
            this.doLogin(userBO.userName, userBO.passWord);
        } else if (null != this.nextHandler)
        {
            this.nextHandler.login(userBO);
        } else
        {
            throw new RuntimeException("no concrete handler to handle the request");
        }
    }
}

class NormalUserHandler extends AbstractUserServiceHandler
{


    protected NormalUserHandler(byte level)
    {
        super(level);
        this.level = level;
    }

    protected void doLogin(String userName, String password)
    {
        System.out.println("normal user login:" + userName);
    }
}

class VIPUserHandler extends AbstractUserServiceHandler
{

    protected VIPUserHandler(byte level)
    {
        super(level);
        this.level = level;
    }

    protected void doLogin(String userName, String password)
    {
        System.out.println("vipUserLogin:" + userName);
    }
}

// 简单的builder模式
class HandlerBuilder
{
    public static IUserService buildUserServiceHandler()
    {
        NormalUserHandler normalUserHandler = new NormalUserHandler((byte) 0);
        VIPUserHandler vipUserHandler = new VIPUserHandler((byte) 1);
        normalUserHandler.nextHandler = vipUserHandler;
        return normalUserHandler;
    }
}

public class ResponsibilityPattern
{
    public static void main(String[] args)
    {
        IUserService userService = HandlerBuilder.buildUserServiceHandler();
        UserBO userBO = new UserBO((byte) 0, "joker", "123");
        userService.login(userBO);
        userBO = new UserBO((byte) 1, "clown", "123");
        userService.login(userBO);

    }
}
