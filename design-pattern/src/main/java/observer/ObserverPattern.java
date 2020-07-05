package observer;

import java.util.Vector;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-02-03 23:38
 */

interface Observer
{
    // 这个用于当观察者检测到变化之后该如何响应
    void update(Observable observable, Object... args);
}

abstract class Observable
{
    protected boolean changed;
    protected Vector<Observer> observers;

    public Observable()
    {
        this.observers = new Vector<>();
    }

    public void addObserver(Observer observer)
    {
        this.observers.add(observer);
    }

    public void notifyWithChanged()
    {
        this.changed = true;
        if (this.changed)
        {
            for (Observer observer : observers)
            {

                observer.update(this, null);
            }
        }
    }


    public void notifyWithChanged(Object... args)
    {
        this.changed = true;
        if (this.changed)
        {
            for (Observer observer : observers)
            {

                observer.update(this, args);
            }
        }
    }
}

public class ObserverPattern extends Observable
{
    public static final Observer OBSERVER_1 = (observable, args) -> System.out.println("观察者1号监听到了消息");
    public static final Observer OBSERVER_2 = (observable, args) -> System.out.println("观察者2号收到了");

    public static void main(String[] args)
    {
        Observable observable = new ObserverPattern();
        observable.addObserver(OBSERVER_1);
        observable.addObserver(OBSERVER_2);
        observable.notifyWithChanged(null);
    }
}
