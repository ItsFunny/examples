package strategy;

import java.io.File;

/**
 * @author joker
 * @When
 * @Description
 * 策略模式: 策略模式非常适合用lambda
 * 两个角色:
 * 1. 接口:封装具体的方法
 * 2. 实现类: 不同的方案有不同的实现方式
 * @Detail
 * @date 创建时间：2019-02-01 06:00
 */

interface FileStrategy
{
    boolean upload(File file);
}

public class StrategyPattern
{
    public static FileStrategy FTPFileStrategy = (file) ->
    {
        System.out.println("这是文件策略中的ftp策略,文件会上传到远程的ftp服务器");
        return true;
    };
    public static FileStrategy LocalFileStrategy = (file) ->
    {
        System.out.println("本地策略:文件上传到本地");
        return true;
    };
}
