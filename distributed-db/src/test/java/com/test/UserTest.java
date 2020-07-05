/**
 * @Description
 * @author joker
 * @date 创建时间：2018年10月27日 下午12:56:40
 */
package com.test;

import java.util.HashSet;
import java.util.List;
import java.util.Random;
import java.util.Set;
import java.util.concurrent.TimeUnit;

import com.test.model.User;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.context.junit4.SpringJUnit4ClassRunner;

import com.joker.library.page.PageRequestDTO;
import com.joker.library.page.PageResponseDTO;
import com.test.service.IUserService;

import lombok.extern.slf4j.Slf4j;

/**
 * @author joker
 * @When
 * @Description
 * @Detail
 * @date 创建时间：2018年10月27日 下午12:56:40
 */
@SpringBootTest
@RunWith(SpringJUnit4ClassRunner.class)
@Slf4j
public class UserTest
{
    @Autowired
    private IUserService userService;

    @Test
    public void testFindByPage()
    {
        PageRequestDTO pageRequestDTO = new PageRequestDTO();
        pageRequestDTO.setPageNum(1);
        pageRequestDTO.setPageSize(10);
        PageResponseDTO<List<User>> pageResp = userService.findByPage(pageRequestDTO);
        List<User> users = pageResp.getData();
        for (User user : users)
        {
            System.out.println(user);
        }
    }

    @Test
    public void testInsert()
    {
        // 忽略for 循环插入 :)
        long startTime = System.currentTimeMillis();
        log.info("[begin batch insert]");
        for (int i = 0; i < 10000; i++)
        {
            User user = new User();
            user.setUserId(i);
            user.setUserName("joker__" + i);
            userService.insert(user);
        }
        log.info("[batch insert finished] consume :{} ms ", System.currentTimeMillis() - startTime);
    }

    @Test
    public void testRandInsert()
    {
        // 忽略for 循环插入 :)
        long startTime = System.currentTimeMillis();
        log.info("[begin batch insert]");

        Set<Integer> set = new HashSet<>();
        Random random = new Random();
        for (int i = 0; i < 10000; )
        {
            try
            {
                TimeUnit.MILLISECONDS.sleep(1);
            } catch (InterruptedException e)
            {
                e.printStackTrace();
            }
            set.add(i);
            int i1 = random.nextInt(100);
            i += i1;


        }

        for (Integer i : set)
        {
            User user = new User();
            user.setUserId(i);
            user.setUserName("joker__" + i);
            userService.insert(user);
        }

        log.info("[batch insert finished] consume :{} ms ", System.currentTimeMillis() - startTime);
    }

    @Test
    public void pageTest()
    {
        PageRequestDTO pageRequestDTO = new PageRequestDTO();
        pageRequestDTO.setPageNum(1);
        pageRequestDTO.setPageSize(10);
        PageResponseDTO<List<User>> pageResp = userService.findByPage(pageRequestDTO);
        List<User> users = pageResp.getData();
        for (User user : users)
        {
            System.out.println(user);
        }
        pageRequestDTO.setPageNum(2);

        PageResponseDTO<List<User>> pageResp1 = userService.findByPage(pageRequestDTO);
        for (User user : pageResp1.getData())
        {
            System.out.println(user);
        }
    }


    // 特殊插入测试

    /*
        1个库中只有1个表有数据: 第一个库的第一个表有数据,第二个库的第2个表有数据
     */
    @Test
    public void testSpecialInsert()
    {
        // 忽略for 循环插入 :)
        long startTime = System.currentTimeMillis();
        log.info("[begin batch insert]");

        for (int i = 0; i < 10000; i++)
        {
            if ((i % 2 == 0 && i % 3 == 0) || (i % 2 == 1 && i % 3 == 1))
            {
                User user = new User();
                user.setUserId(i);
                user.setUserName("joker__" + i);
                userService.insert(user);
            }

        }


        log.info("[batch insert finished] consume :{} ms ", System.currentTimeMillis() - startTime);
    }

    @Test
    public void specialPageTest()
    {
        PageRequestDTO pageRequestDTO = new PageRequestDTO();
        pageRequestDTO.setPageNum(1);
        pageRequestDTO.setPageSize(10);
        PageResponseDTO<List<User>> pageResp = userService.findByPage(pageRequestDTO);
        List<User> users = pageResp.getData();
        for (User user : users)
        {
            System.out.println(user);
        }

        pageRequestDTO.setPageNum(2);
        PageResponseDTO<List<User>> pageResp1 = userService.findByPage(pageRequestDTO);
        for (User user : pageResp1.getData())
        {
            System.out.println(user);
        }

    }
}