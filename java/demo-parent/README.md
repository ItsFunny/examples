-  如果要打包,添加如下build
    ```$xslt
     <build>
            <plugins>
    
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-jar-plugin</artifactId>
                </plugin>
    
                <plugin>
                    <groupId>org.apache.maven.plugins</groupId>
                    <artifactId>maven-assembly-plugin</artifactId>
                </plugin>
    
                <plugin>
                    <groupId>org.springframework.boot</groupId>
                    <artifactId>spring-boot-maven-plugin</artifactId>
                    <configuration>
                        <mainClass>cn.bidsun.BenchMarkApplication</mainClass>
                    </configuration>
                    <executions>
                        <execution>
                            <goals>
                                <goal>repackage</goal><!--可以把依赖的包都打包到生成的Jar包中-->
                            </goals>
                        </execution>
                    </executions>
                </plugin>
            </plugins>
        </build>
    ```


