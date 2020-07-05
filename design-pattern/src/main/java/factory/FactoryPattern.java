package factory;

/**
 * @author joker
 * @When
 * @Description 静态工厂模式
 * @Detail
 * @date 创建时间：2019-02-02 07:38
 */
interface IHumanService
{
    String type();
}

class Student implements IHumanService
{
    public String name;

    @Override
    public String type()
    {
        return "student";
    }

    public Student(String name)
    {
        this.name = name;
    }
}

class Teacher implements IHumanService
{
    public String name;

    @Override
    public String type()
    {
        return "teacher";
    }

    public Teacher(String name)
    {
        this.name = name;
    }
}

public class FactoryPattern
{


}

//////////////////////////////////////////////////////////////////////
// 静态工厂模式
class SimpleFactory
{
    public static IHumanService CreateHuman(String type)
    {
        if ("student".equals(type))
        {
            return new Student("joker");
        } else if ("teacher".equals(type))
        {
            return new Teacher("clown");
        } else
        {
            return null;
        }
    }
}


/////////////////////////////////////////////////////////////////////////////
// 普通工厂模式

interface IHumanFactoryService
{
    IHumanService createHuman();
}

class StudentFactory implements IHumanFactoryService
{

    @Override
    public IHumanService createHuman()
    {
        return new Student("joker");
    }
}

class TeacherFactory implements IHumanFactoryService
{

    @Override
    public IHumanService createHuman()
    {
        return new Teacher("clown");
    }
}

////////////////////////////////////////////////////////////////////
// 抽象工厂模式
interface  IHuman
{

}
interface IStudent extends IHuman
{

}
interface ITeacher extends IHuman
{

}

class ZheJiangStudent implements IStudent
{
    String name;

    public ZheJiangStudent(String name)
    {
        this.name = name;
    }
}

class BeiJingStudent implements IStudent
{
    String name;

    public BeiJingStudent(String name)
    {
        this.name = name;
    }
}

class ZheJiangTeacher implements ITeacher
{
    String name;

    public ZheJiangTeacher(String name)
    {
        this.name = name;
    }
}

class BeiJingTeacher implements ITeacher
{
    String name;

    public BeiJingTeacher(String name)
    {
        this.name = name;
    }
}

abstract class AbstractHumanServiceFactory
{
    public abstract IStudent createStudent();

    public abstract ITeacher createTeacher();
}

class ZheJiangHumanFactory extends AbstractHumanServiceFactory
{

    @Override
    public IStudent createStudent()
    {
        return new ZheJiangStudent("joker");
    }

    @Override
    public ITeacher createTeacher()
    {
        return new ZheJiangTeacher("joker");
    }
}

class BeijingHumanFactory extends AbstractHumanServiceFactory
{

    @Override
    public IStudent createStudent()
    {
        return new BeiJingStudent("clown");
    }
    @Override
    public ITeacher createTeacher()
    {
        return new BeiJingTeacher("clown");
    }

}