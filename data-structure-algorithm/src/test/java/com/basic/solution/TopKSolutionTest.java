package com.basic.solution;


import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.sql.*;
import java.util.UUID;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2019-01-18 23:00
 */
public class TopKSolutionTest
{
    private static Logger log = LoggerFactory.getLogger(TopKSolutionTest.class);
    private static String sqlUserName = "root";
    private static String password = "123456";
    private static String driverClass = "com.mysql.jdbc.Driver";
    private static String url = "jdbc:mysql://localhost:3306/test?characterEncoding=utf-8&useSSL=false";
    private static AtomicInteger limit = new AtomicInteger();
    private static ThreadLocal<PreparedStatement> threadLocal = new ThreadLocal<>();

    public static class ThreadHelper extends Thread
    {
        private PreparedStatement preparedStatement;
        String insertDataSql = "insert into topk_test (name) ";

        CountDownLatch countDownLatch;

        public ThreadHelper(PreparedStatement preparedStatement, CountDownLatch countDownLatch)
        {
            this.countDownLatch = countDownLatch;
            this.preparedStatement = preparedStatement;
        }

        @Override
        public void run()
        {
            threadLocal.set(preparedStatement);
            System.out.println(Thread.currentThread().getName() + "---begin");
            PreparedStatement ps = threadLocal.get();
            try
            {
                while (limit.get() < 10000000)
                {
                    for (int i = 0; i < 1000; i++)
                    {
                        String newSql = insertDataSql + " values (" + "'" + UUID.randomUUID().toString().substring(0, 7) + "')";
                        ps.addBatch(newSql);
                    }
                    int[] ints = ps.executeBatch();
                    System.out.println(Thread.currentThread().getName() + "add " + ints.length + "条记录");
                    limit.addAndGet(ints.length);
                }
            } catch (Exception e)
            {
                e.printStackTrace();
                countDownLatch.countDown();
                try
                {
                    if (null != ps) ps.close();
                } catch (SQLException e1)
                {
                    e1.printStackTrace();
                }
            }
        }
    }

    @Test
    public void initData() throws Exception
    {
        String dropTableSql = "Drop table if exists topk_test";

        String createTableSql = "create table topk_test (" +
                "id int(10) primary key auto_increment ," +
                "name varchar(32)" +
                ")ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8";
        String insertDataSql = "insert into topk_test (name) ";
        // 初始化数据
        Class.forName(driverClass);
        Connection connection = DriverManager.getConnection(url, sqlUserName, password);
        Statement statement = connection.createStatement();
        boolean execute = statement.execute(dropTableSql);
        if (execute)
        {
            System.out.println("删除原先表成功");
        }
        execute = statement.execute(createTableSql);
        statement.close();
        PreparedStatement preparedStatement = connection.prepareStatement(insertDataSql);
        if (execute)
        {
            System.out.println("创建表成功");
        }
        long startTimemills = System.currentTimeMillis();
        log.info("开始批量插入数据");
        CountDownLatch countDownLatch = new CountDownLatch(8);
        // 本机为4核8线程,让其跑满,
        ExecutorService executorService = Executors.newFixedThreadPool(8);
        for (int i = 0; i < 8; i++)
        {
            executorService.execute(new ThreadHelper(preparedStatement, countDownLatch));
        }
        countDownLatch.await();
        log.info("初始化完毕,耗时:{} ms", System.currentTimeMillis() - startTimemills);
    }

    @Test
    public void testTopkByQsort()
    {
    }


}
