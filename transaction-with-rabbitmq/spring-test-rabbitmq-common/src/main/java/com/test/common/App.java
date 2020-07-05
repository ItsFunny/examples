package com.test.common;

import java.util.concurrent.Executors;
import java.util.concurrent.LinkedBlockingDeque;
import java.util.concurrent.RejectedExecutionHandler;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args )
    {
    	Executors.newFixedThreadPool(3);
    	Executors.newCachedThreadPool();
    	RejectedExecutionHandler
    	
        System.out.println( "Hello World!" );
    }
}
