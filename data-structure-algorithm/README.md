# TODO-LIST

## [GitHub更新的较CSDN勤](https://github.com/ItsFunny/data-structure-algorithm)

## [源码剖析md(现阶段具体的源码暂时没时间展示,只能'理论',有空加源码,不过blog上有我之前的见解,一般)](https://github.com/ItsFunny/data-structure-algorithm/tree/master/src/main/source)

---
* 囊括常见的排序算法,数据结构的实现,具体的代码会有Go也会有Java
* 模块的链接应该最近是都没空了
* 每个算法和数据结构都有对应的测试模块,Java在test下,Go则直接同级目录下**对于为null或者数组长度为0的特殊情况默认不校验**
---
注意点:

* **当一旦涉及递归的时候,最好是先将递归退出的条件先写出来**

目录(目录与实际不符,有些未添加到目录中)
---

* [1.排序算法](#1)
    * [1.1插入排序](#1.1)
    * [1.2冒泡排序](#1.2)
    * [1.3归并排序](#1.3)
    * [1.4快速排序](#1.4)
    * [1.5希尔排序](#1.5)
    * [1.6选择排序](#1.6)
    * [1.7堆排序](#1.7)
    
* [2.数据结构](#2)
    * [2.1链表](#2.1)
        * [2.1.1单链表](#2.1.1)
        * [2.1.2单向循环链表](#2.1.2)
        * [2.1.3双链表](#2.1.3)
        * [2.1.4LinkedList](#2.1.4)
    * [2.2数组](#2.2)
        * [2.2.1CopyOnWriteArrayList](#2.3)
        
* [3.hash](#3)
    * [3.1HashSet的实现](#3.1)
    * [3.2HashMap的实现](#3.2)
    * [3.3解决Hash冲突的方法](#3.3)
        * [3.3.1开放地址法](#3.3.1)
            * [3.3.1.1线性探测法](#3.3.1.1)
        * [3.3.2链地址法](#3.3.2)
        * [3.3.3再散列法](#3.3.3)
        * [3.3.4总结](#3.3.4)

* [4.TOPK解决方案](#4)
    * [4.1全部排序](#4.1)
    * [4.2局部淘汰](#4.2)
    * [4.3分治法](#4.3)
    * [4.4hash过滤法](#4.4)

* [5.map的遍历](#5)
    * [5.1keySet通过key遍历](#5.1)
    * [5.2entrySet的iterator遍历](#5.2)
    * [5.3通过entrySet遍历](#5.3)
    * [5.4通过values直接遍历](#5.4)

* [6.锁](#6)
    * [6.1死锁的实现](#6.1)
    * [6.2生产者消费者的实现](#6.2)
    * [6.3原生的notify/wait/notifyAll实现](#6.3)
    * [6.4通过lock/condition实现](#6.4)
    * [6.5读写锁的实现](#6.5)

* [7.Spring](#7)
    * [7.1动态代理的实现](#7.1)
        * [7.1.1基于JDK的动态代理](#7.1.1)
        * [7.1.2基于Cglib的动态代理](#7.1.2)

---

排序算法
- 稳定与不稳定的意思: 稳定算法与不稳定算法并不是指复杂度是否稳定,而是指前后的位置是否发生了变化,如A原先在B的前面,
当我们经过一系列操作之后A到了B后面,则这个算法就是不稳定的
- 不稳定算法: 快希选堆

- 可能会搞混的记忆点:
    -   插入排序: 有序中插入   (插入既插队)
    -   选择排序: 遍历使得顺序存储(从头到尾暴力匹配)
- 



| 排序方法| 时间复杂度|空间复杂度|
| -------|--------|--------|
| 插入排序(稳定)  | O(n^2)      |  O(1)     |
| 冒泡排序(稳定) | O(n^2)      |   O(1)     |
| 归并排序(稳定)  | O(nlogn)   |  O(n+logn)=O(n) |
| 快速排序(不稳定)| 最好为O(nlogn),最差O(n^2)|最好为O(nlogn),最差为O(n)           |            |
| 希尔排序(不稳定)|O(n^2)      |   O(1)      |
| 选择排序(不稳定)  |O(n^2)      |  O(1)      |
| 堆排序(不稳定)  | O(nlogn)   |  O(1)      |

---
- `WIP` 8中排序算法的实现       **升序的时候要假设比较的值都是大于它的**

    -   **直接插入排序 (稳定)** **:将一个数据插入到已经排序的数组中(分为无序区和有序区)**
        -   时间复杂度: O(n^2) 因为有2层for循环:                                                   一个O(n)找元素,一个O(n)找位置
        -   空间复杂度: O(1) 因为是在原先的数组上操作的  
        -   核心要点:   从下标处往前进行遍历
    ```
    public void insertSort(Integer[] arrs)
    {
        // 插入排序的核心就是假设要插入的元素的下标index 0-index 都是有序的(升序)
        // 所以需要从后往前判断,因为是升序,所以要找到的是坐标是小于他的,而右边是大于他的,因而判断条件是大于
        // 如果条件成立需要将这个值往后移动(不用担心插入的值,因为开始就会保存)
        // 当内层循环跳出的时候也就意味着,下标所在的值是小于temp值的,而我们需要在这index+1处插入
        // 因为之前的数都已经往后移动了
        for (int i = 1; i < arrs.length; i++)
        {
            int temp = arrs[i];
            int j = i - 1;
            for (; j >= 0 && arrs[j] >= temp; j--)
            {
                arrs[j + 1] = arrs[j];
            }
            arrs[j + 1] = temp;

        }
    }
    ```
   - `FINISH`**冒泡排序(稳定)**
        -   时间复杂度: O(n^2) 2层for循环
                        一个O(n)用于遍历查找,一个用于匹配
        -   空间复杂度: O(1) 只在原先的数组中操作
        -   核心要点: 
    ```
     就是2层for循环,对前后进行比较
     public static void popSort(Integer[] arr)
     {
         //  冒泡排序就是暴力遍历比较
         //  如果后者小于则直接进行更换即可
         for (int i = 0; i < arr.length; i++)
         {
             for (int j = i + 1; j < arr.length; j++)
             {
                 if (arr[j] < arr[i])
                 {
                     int temp = arr[i];
                     arr[i] = arr[j];
                     arr[j] = temp;
                 }
             }
         }
     }
    ```
     
   - `FINISH`**归并排序(稳定)**
         ![分割图](https://images2015.cnblogs.com/blog/1024555/201612/1024555-20161218163120151-452283750.png)
         
        -   时间复杂度: O(nlogn) O(n)为需要将待排序的序列都扫描一遍
                         而归并中的分可以认为将数组分成了完全二叉树,
                         所以深度可知为:O(logn)
        -   空间复杂度: O(n+logn) 因为需要相同的额外长度的数组,所以
                         为O(n),而又因为二叉树的性质,所以而O(logn)
        -   核心要点:   归并是指将一个数组分成若干个小的数组,对每个
                         小的数组进行排序,最后统一排序; 记得分的时候
                         退出的条件(下标一致)
                         
     ```
     public static void mergeSort(Integer[] arr)
     {
         // 归并算法分为2个步骤: 分治法  分 + 治
         Integer[] tempArr = new Integer[arr.length];
         mergeSort(arr, 0, arr.length - 1, tempArr);
     }
     // 分代表着将数组分为直至相邻的若干个小数组 ,直至分到了最小情况(既同一下标了),所以最小的数组的长度为2[5,6]
     // 为什么这样子就可以了呢,答案在于merge中,merge会遍历数组,判断大小,按顺序写入到临时队列中
     public static void mergeSort(Integer[] arr, int left, int right, Integer[] temp)
     {
         if (left < right)
         {
             int mid = (left + right) >> 1;
             // 对左边进行分
             mergeSort(arr, left, mid, temp);
             // 对右边进行分
             mergeSort(arr, mid + 1, right, temp);
             // 对左右进行治
             merge(arr, left, mid, right, temp);
         }
     }
     // 治的逻辑::
     // 对左边,右边进行遍历,同时会进行判断,选取小的值放入到临时数组中
     // 之后再进行遍历,此时只会遍历一边
     // 最后则是将临时数组中的元素复制到元数组中
     // 关键点在于: 要建立临时的变量,代替下标去移动
     public static void merge(Integer[] arr, int left, int mid, int right, Integer[] temp)
     {
         int i = left;
         int j = mid+1;
         int k = 0;
         // 进行归并
         while (i <= mid && j <= right)
         {
             if (arr[i] < arr[j])
             {
                 temp[k++] = arr[i++];
             } else
             {
                 temp[k++] = arr[j++];
             }
         }
         // 对数组中剩余的进行复制
         while (i <=mid)
         {
             temp[k++] = arr[i++];
         }
         while (j <=right)
         {
             temp[k++] = arr[j++];
         }
         k = 0;
         while (left <= right)
         {
             arr[left++] = temp[k++];
         }
     }
     ```
    - `WIP` **快速排序(不稳定)**
        -   时间复杂度: 最好O(nlogn),最差O(n^2)最差的情况
        是指当选取的元素恰好是最小或最大的元素,这时候就退化为了冒泡排序
        -   空间复杂度:O(logn),最差为O(n^2)
        -   核心要点:   就是选取一个哨兵值,左边的是小于他的,右边是大于它的
        这样就划分为了2块,再对左右两块进行进行同样的操作
        
        -   `FINISH` 普通快速排序
        
        ```
       qSort 的关键在于有一个标准值,左边的树都小于这个值右边的都大于这个值
        public void qSort(Integer[] arr, Integer start, Integer end)
        {
            if (start < end)
            {
                Integer paration = paration(arr, start, end);
                qSort(arr, start, paration);
                qSort(arr, paration + 1, end);
            }
        }
    
        // 既左边的小于这个值,右边的都大于这个值
        public Integer paration(Integer[] arr, Integer start, Integer end)
        {
            // 取一个标准值作为参考,然后递归进行比较,如果取左边的话,则从右边开始递归
            // 相反如果先取的右边,则先判断左边
            int stanard = arr[start];
            while (start < end)
            {
                // 右边的值都是大于左边的,因此一旦有值小于标准则,则需要将其换到左边去,同时这个时候左边的值是刚好
                // 是临界值,既下一个可能就大于这个标准值了
                while (end > start && arr[end] >= stanard) end--;
                arr[start] = arr[end];
                while (end > start && arr[start] <= stanard) start++;
                arr[end] = arr[start];
            }
            // 因为上述的交换都少了最开始的start值,因而在这里将其补回
            arr[start] = stanard;
            return start;
        }
        ```
        
        -   `WIP` 变更版快速排序(三值排序)
              
    - `FINISH`**希尔排序:不稳定**
    **在原先简单排序的基础,对原先数组通过stride步长分成多块,对每块做简单插入**
        -   时间复杂度: O(n^2) 
    ```
        // 希尔排序是直接插入算法的优化:
        // 将一个数组分成多块,对每块进行插入排序
        // 直接插入排序其实就是步长为1的希尔排序
         public void shellSort(Integer[] arr)
            {
                // 希尔排序是直接插入算法的优化:
                // 将一个数组分成多块,对每块进行插入排序
                // 直接插入排序其实就是步长为1的希尔排序
                int stride = arr.length;
                while (stride != 1)
                {
                    stride >>= 1;
                    // 对每个分组进行排序
                    for (int i = 0; i < stride; i += stride)
                    {
                        for (int j = 0; j < arr.length; j += stride)
                        {
                            int temp = arr[j];
                            int k = j - stride;
                            for (; k >= 0 && arr[k] >= temp; k -= stride)
                            {
                                arr[k + stride] = arr[k];
                            }
                            arr[k + stride] = temp;
                        }
                    }
                }
            }
    ```
    - `FINISH`**简单选择排序(不稳定)**
        -   时间复杂度: O(n^2)  2层for循环,                                                         1层用于下标遍历,1层用于判断匹配
        -   空间复杂度: O(1)    没有申请新的空间
        -   核心思路:   既数组下标与元素是强匹配的:
                        0号存放最小,1号存放次小的,所以需要遍历匹配的
    ```
    简单选择排序就是暴力遍历: 2层for循环 将最小的放在0号,次小的放在1号,也就是说需要从0到length进行遍历
    public static void simpleSelectionSort(Integer[] arr)
    {
        // 简单选择排序的核心就是0号放的是最小的元素,和1号放的是次小             的元素,意味着需要暴力遍历
        for (int i = 0; i < arr.length; i++)
        {
            int min = arr[i];
            int pos = i;
            for (int j = i + 1; j < arr.length; j++)
            {
                if (arr[j] < min)
                {
                    min = arr[j];   //将最小的这个给min
                    pos = j;        // 记录最小的下标,方便更换
                }
            }
            arr[pos] = arr[i];
            arr[i] = min;
        }
    }
    ```
    
    - `FINISH`**选择排序(不稳定)**:每个数都与剩下的所有数比较大小,
    从而选取出最大或者最小的值
        - 时间复杂度: O(n^2),一层O(n)用于起始遍历,另外一层O(n)
        用于遍历剩下的所有数来比较大小
        - 空间复杂度:O(1),不需要额外申请空间
        - 核心实现:核心就是一个数与剩下的所有数比较大小
        ```
        public static void simpleSelectionSort(Integer[] arr)
        {
            // 简单选择排序的核心就是0号放的是最小的元素,和1号放的是次小的元素,意味着需要暴力遍历
            for (int i = 0; i < arr.length; i++)
            {
                int min = arr[i];
                int pos = i;
                for (int j = i + 1; j < arr.length; j++)
                {
                    if (arr[j] < min)
                    {
                        min = arr[j];   //将最小的这个给min
                        pos = j;        // 记录最小的下标,方便更换
                    }
                }
                arr[pos] = arr[i];
                arr[i] = min;
    
            }
        }
        ``` 
           
    - `FINISH`**堆排序(不稳定算法)**
        -   时间复杂度: **O(nlogn)** ,i层有2^(i-1)个节点,而因为每次都要
        进行比较,更换之后子树也要比较,所以时间为:2^(i-1)*(k-i)
        i代表第几层,k代表高度(既这层节点需要比较的次数),提取常量则时间复杂度为
        O(n); 然后重新建堆,重新建堆的方式是:从尾到头,循环n-1次(>0即可),每次基于
        二叉树的特性为logn,所以为nlogn-logn ,所以总共为O(nlogn)
        -   空间复杂度: O(1),不需要额外的申请空间
        -   核心要点: 建堆->排序(再建堆的过程),并且都是从后往前,建堆是中间开始,排序是与0号元素交换位置再建堆,并且建堆过程中当节点更换之后还得将子节点也重新建堆,左孩子下标为2**index+1 ,右孩子下标为2**index+2
        -   实现: 
        
        ```
        // 堆排序:
        // 分为2个步骤: 建堆,排序 (其实排序就是重新建堆的过程,因为构建的是堆,只需要将首位移到最后,剩余的继续建堆即可)
        // 堆排序是建立在完全二叉树上的,堆又分为最大堆和最小堆,如果我们想升序的话则需要构建的是最大堆,并且最大值在第一位
        public void heapSort(Integer[] arr)
        {
            // 建堆
            for (int i = (arr.length >> 1) - 1; i >= 0; i--)
            {
                // 为什么我们要从父节点的上限开始,原因在于当我们从后往前的时候,子节点所处的树较父节点肯定是小的
                // 也就可以省去很多无用的操作 ,想象一下就行,当根节点与某个子节点发生了变化之后,子节点需要重新排序
                // 这时候的树肯定是较当这个节点为父节点时候的子节点的树要大的
                buildHeap(arr, i, arr.length - 1);
            }
            // 当最大堆构建完毕之后,我们只需要不断的将最大的移动到最后然后重新建堆即可
            // 并且因为最大的一直都是在0位,所以我们只需要从后往前更换元素即可
            for (int i = arr.length - 1; i >= 0; i--)
            {
                int temp = arr[i];
                arr[i] = arr[0];
                arr[0] = temp;
                sort(arr, 0, i);
            }
    
        }
    
        // 建堆有2种方式,第一种是通过递归建堆的方式,如下:
        public void buildHeap(Integer[] arr, Integer index, Integer limit)
        {
            // leftChild index
            Integer leftChildIndex = (index << 1) + 1;
            // rightChind index
            Integer rightChindIndex = (index << 1) + 2;
            Integer maxIndex = index;
            if (leftChildIndex < limit && arr[leftChildIndex] > arr[maxIndex])
            {
                maxIndex = leftChildIndex;
            }
            if (rightChindIndex < limit && arr[rightChindIndex] > arr[maxIndex])
            {
                maxIndex = rightChindIndex;
            }
            // 如果根节点不是最小值则交换位置
            if (maxIndex == index)
            {
                return;
            }
            Integer temp = arr[maxIndex];
            arr[maxIndex] = arr[index];
            arr[index] = temp;
            // 因为我们是对整个堆进行排序,所以当更换了值之后,所在的树也基本上变了,所以我们需要重新建堆
            // 既子节点的树很可能是发生了变化
            // 这里就是可以优化的地方,既然变更的只是子节点,大可抽出成为一个for循环实现
            buildHeap(arr, maxIndex, limit);
        }
        public void sort(Integer[] arr, Integer index, Integer limit)
        {
            buildHeap(arr, index, limit);
        }
        
        ```
        
        - `FINISH`  非递归实现:
        
        ```
        // 非递归排序的思路与递归排序的思路是一样的;
        // 选取左右孩子的最大值,然后对交换位置的孩子作为根节点的树继续调整树
        public void buildHeapWithOutRecursion(Integer[] arr, Integer rootIndex, Integer limitIndex)
        {
            Integer temp=arr[rootIndex];
            for(Integer i=(rootIndex<<1)+1;i<limitIndex;i=(rootIndex<<1)+1)
            {
                // 右孩子就是+1所处的位置
                // 选取左右孩子的最大值与根节点进行比较
                if (i+1<=limitIndex&&arr[i+1]>arr[i])
                {
                    i++;
                }
                // 这里可能会有疑问,为什么是不变的temp值去遍历比较底下的叶子节点的值:
                // 因为当我们不满足条件的时候(既根节点与孩子节点更换之后),此时孩子节点的root值就是temp了
                //
                //     5                    7
                //   3   7   -->         3      5    对   5  进行重新建树
                //  2 1 0 6           2    1  0  6       0 6
                //
                if (temp>=arr[i])
                {
                    break;
                }
                // 否则就是孩子节点的值更加大,则需要更换位置,将孩子节点提到root节点上,
                // 然后对子节点所处的进行再建树(rootIndex=i 就使得这个孩子节点变成了根节点),对应上面的就是递归
                arr[rootIndex]=arr[i];
                rootIndex=i;
            }
            // 上面是直接复制的,最先的rootIndex节点的值就丢失了,因而我们需要将其补回,这里其实有点像插入排序
            arr[rootIndex]=temp;
        }    
        ```
        
    -   `WIP`基数排序


数据结构
---

- `WIP` 链表
    -   `FINISH` 单链表(不带尾指针的链表)
        -   定义: 只有root节点的链表
        -   crud: 
            -   增加:增加节点采用遍历到尾节点然后设置即可
            -   删除: 需要判断删除的节点是否是头节点(末尾删除的除外),遍历退出
            的条件是遍历到删除节点的上一个节点即可,然后指向删除节点的下一个节点即可,
            不需要做额外的操作,因为删除的节点已经不被持有引用了(实际情况可能会发生引用泄露)
            -   通过下标删除: 同样需要注意的是删除的是否是头节点,如果不是则需要,
            需要for带次数的循环,同样也是遍历到前一个节点(index-1)即可,之后指向删除节点的下一个节点,
            至于起始顺序是从0开始还是从1开始个人定
         -  `FINISH` 单链表反转
            -   1->2->3->4->5->6->7->8->9 ===>1<-2<-3<-4<-5<-6<-7<-8<-9
            原理很简单,两两交换即可,也就是说需要保存下一个节点,因而我们总共需要申明3个变量
            
            ``
            
                public void reverse()
                {
                    ListNode lCur = this.head, lPrev = null, lNext;
                    while (null != lCur)
                    {   
                        // 保存当前节点的下一个节点
                        lNext = lCur.next;
                        // 将当前节点指向上一个节点
                        lCur.next = lPrev;
                        // 保存当前节点为上一个节点
                        lPrev = lCur;
                        // 往后移动
                        lCur = lNext;
                    }
                }
            
            ``
    -   `FINISH` 单向循环链表
        -   定义: 在单向链表的基础上,尾指针指向头指针形成一个环
        -   注意点: 注意删除元素的时候要对**是否是尾节点进行判断**,如果是需要移动尾节点的指向,另外**强烈建议添加一个size的属性变量**
        **最后则是永远不会有空的值**
    -   `FINISH` 双向链表
        -   定义: 既一个节点有指向前驱节点,也有指向后继节点,双向链表又有双向循环链表,既tail指针next指向了头节点
        -   特点: 每个节点既有前驱节点,又有后继节点
        -   注意点: 注意点无外乎就是头|尾节点操作的时候要比常规的多一步操作(既head:head重新指向,tail.next重新指向head,heaad.previous重新指向tail;而tail:tail重新指向之后,tail.next重新指向head,head.previous重新指向tail,也需要重新指向),如果是删除元素,只需要将引用这个节点
        的节点更改即可,对于前驱而言就是修改next,而对于后继而言就是修改previous,
        ```
           func (d *DoubleSiededList)Add(value interface{}){
            if nil==value{
                return
            }
            newNode:=&DSListNode{
                data:     value,
                }
            if d.size==0{
                d.head,d.tail=newNode,newNode
                d.tail.next=d.head
                d.head.previous=d.tail
            }else{
                d.tail.next=newNode
                newNode.previous=d.tail
                d.tail=newNode
                d.tail.next=d.head
            }
            d.size++
           }
           // @Function: 移除某个下标的元素
           // 下标从0开始
           func (d *DoubleSiededList)Remove(index int)error{
           	if index>d.size {
           		return common.IndexOutOfRangeError
           	}
           	// 判断是否是链头元素
           	if index==0{
           		d.head=d.head.next
           		d.head.previous=d.tail
           		d.tail.next=d.head
           		d.size--
           		return nil
           	}else if index ==d.size-1{
           		// 需要判断是否是链尾元素
           		// 如果是链尾元素,则需要重新调整链尾
           		d.tail=d.tail.previous
           		d.tail.next=d.head
           		d.head.previous=d.tail	// 别忘记头指向的前驱也需要更改
           		d.size--
           		return nil
           	}
           	// 删除指定下标的元素
           	tempNode:=d.head
           	for i:=0;i<index;i++{
           		tempNode=tempNode.next
           	}
           	prevNode:=tempNode.previous
           	prevNode.next=tempNode.next
           	if nil!=tempNode.next{
           		tempNode.next.previous=prevNode
           	}
           	d.size--
           	return nil
           }
        ```
    -   `FINISH` LinkedList的实现
        -   定义: LinkList是一种有序的数据结构,类似于队列先进先出
         -  特点: 基于链表的数据结构,存放了head和tail指针
    -   `WIP` CopyOnWriteArrayList的实现
    
- `WIP` hash
    -   `FINISH`   HashSet的实现
        -   `定义`: HashSet是指存放的元素没有重复的元素
        -   `特点`: 其实内部是Map,并且定义了一个额外的对象用于充当map key value中的value,许多api调用的都是直接map的api
    -   `FINISH`   HashMap的实现 (参考意义不大,在Go下)
        -   `要点`: 内部是数组加链表的形式,有个loadFactor的概念,loadFactor=(size/capilibity),当达到这个阈值就会发生扩容,
        关于扩容,在1.7及其之前会造成死循环的问题,但是1.8之后并不会了,采用临时创建2个对象来复制对象,HashMap是线程不安全的,不安全
        在于,当2个线程写同一个hashCode的对象,且对应的下标bucket为空,则会导致一个线程上写的内容覆盖另外一个线程的内容;还有一个数据重复的问题
        既两个线程写一模一样的2个对象,会导致某个bucket上挂着2个相同的对象,至于死循环,1.8之后就没了
        
    -   `WIP`   解决hash冲突的方法
        -   `WIP`   开放地址法:
            -   定义: 既然当冲突的时候往按某种方式遍历获取不冲突的地址然后赋值

            -   `FINISH`   线性探测法: 
                -   定义: 当index冲突时,从index处往后遍历hash表,如果发现有空的则插入赋值
                    -   查找时根据index开始遍历**直到找到值**或者是**找到一个空槽**(也意味着不可删除元素)或者**整个hash表遍历完毕**
                -   缺陷: 
                    -   因为无法删除元素,所以会造成溢出,解决方法是创建一个溢出表,顺序存储元素
                    -   无法真正的删除元素,**只能通过加上标记来实现**
                    -   聚集现象,既一个大片段都连在一起
        -   `WIP`   链地址法
        -   `WIP`   再散列法
        -   `WIP`   总结:****


- `WIP` TOK 解决方案:
    -   `WIP`   全部排序
    -   `WIP`   局部淘汰法
    -   `WIP`   分治法
    -   `WIP`   hash法


- `FINISH` map的遍历:
    -   `FINISH`   通过keySet来遍历(遍历的都是key)
        -   内部都是key,所以直接遍历然后get即可
    -   `FINISH`   通过entrySet的iterator遍历
        -   entrySet中每个都是Map.Entry对象,其中set接口继承了Iteratable接口,
        又添加了额外的方法,包含了key和value
    -   `FINISH`   通过entrySet来遍历(entrySet内部包含了key和vlaue)
        -   每个entrySet都是Map.Entry对象,内部包含了key和value
    -   `FINISH`   直接通过values遍历值
        -   外抛的一个接口,直接遍历获取value即可
- `WIP` 锁
    -   `FINISH` 死锁的实现
        -   死锁的条件:
            -   循环等待
            -   互斥
            -   占有且等待
            -   不可剥夺
        ```
        public class DeadLock
        {
            private Object a;
            private Object b;
        
            public DeadLock()
            {
                this.a=new Object();
                this.b=new Object();
            }
        
            public void aLock() throws InterruptedException
            {
                synchronized (a)
                {
                    // 为了先让其他线程先获取到b锁所以sleep2s
                    TimeUnit.SECONDS.sleep(2);
                    System.out.println("试图获取到对象b的锁");
                    synchronized (b)
                    {
                        System.out.println("获取到了对象b的锁");
                    }
                }
            }
        
            public void bLock() throws InterruptedException
            {
                synchronized (b)
                {
                    // 为了先让其他线程先获取到a锁所以sleep2s
                    TimeUnit.SECONDS.sleep(2);
                    System.out.println("试图获取到对象a的锁");
                    synchronized (a)
                    {
                        System.out.println("获取到了对象a的锁");
                    }
                }
            }
        
            public static void main(String[] args) throws InterruptedException
            {
                DeadLock lock = new DeadLock();
                new Thread(() ->
                {
                    try
                    {
                        lock.aLock();
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                }).start();
                new Thread(() ->
                {
                    try
                    {
                        lock.bLock();
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                }).start();
                TimeUnit.SECONDS.sleep(100);
            }
        }
        ```
    -   'WIP' 生产者消费者的实现
        -   `WIP`   concurrent组件实现
        ```
        public class ProducerAndConsumerWithConcurrent
        {
            private LinkedBlockingQueue<String> foods;
        
            public ProducerAndConsumerWithConcurrent()
            {
                this.foods = new LinkedBlockingQueue<>();
            }
        
            final Runnable PRODUCER = () ->
            {
                while (!Thread.currentThread().isInterrupted())
                {
                    try
                    {
                        TimeUnit.SECONDS.sleep(1);
                        String food = UUID.randomUUID().toString();
                        foods.put(food);
                        System.out.println("生产者:" + Thread.currentThread().getName() + " 新增食物: " + food);
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                }
            };
            final Runnable CONSUMER = () ->
            {
                while (!Thread.currentThread().isInterrupted())
                {
                    try
                    {
                        String food = foods.take();
                        System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                        TimeUnit.MILLISECONDS.sleep(800);
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                }
            };
        
            public void test()
            {
                // 10个生产者
                for (int i = 0; i < 10; i++)
                {
                    new Thread(PRODUCER).start();
                }
                // 5个消费者
                for (int i = 0; i < 5; i++)
                {
                    new Thread(CONSUMER).start();
                }
            }
        
            public static void main(String[] args) throws InterruptedException
            {
                new ProducerAndConsumerWithConcurrent().test();
                while (true)
                {
                    TimeUnit.SECONDS.sleep(100);
                }
            }
        }
        
        ```
        -   `FINISH`   原生的wait/notify/notifyall实现
        -   `注意点`: 对于原生的notify实现的话,有以下的注意点:
            -   consumer获取食物的时候是需要判断是否为空的,**并且注意是while循环**
            ```
            public class ProducerAndConsumerWithNotify
            {
                private List<String> foods;
            
                public ProducerAndConsumerWithNotify()
                {
                    this.foods = new LinkedList<>();
                }
            
                final Runnable PRODUCER = () ->
                {
                    try
                    {
                        while (!Thread.currentThread().isInterrupted())
                        {
                            String food = UUID.randomUUID().toString();
                            synchronized (foods)
                            {
            //                    while (foods.size() == 16)
            //                    {
            //                        // 如果放不下了,则就需要进行等待了
            //                        foods.wait();
            //                    }
                                foods.add(food);
                                System.out.println("生产者:" + Thread.currentThread().getName() + " 新增食物: " + food);
                                // 提示所有的消费者都可以消费了
                                foods.notifyAll();
                            }
                            TimeUnit.SECONDS.sleep(1);
                        }
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                };
            
                final Runnable CONSUMER = () ->
                {
                    try
                    {
                        while (!Thread.currentThread().isInterrupted())
                        {
                            synchronized (foods)
                            {
                                // 对于消费者而言,如果foods为空,代表着需要等待producer生产食物
                                // 注意这里必须是while循环
                                // 因为当线程重新被唤醒之后,因为程序计数器,从而会继续在这里执行
                                // 而如果producer生产速度慢,当1号consumer消费完毕,2号抢到了之后也会notifyall,
                                // 如果恰巧是consumer获取到了则会跳出了if循环,从而直接remove(0),也就报index错误了,while则会继续先判断
                                // 究其原因是因为只针对foods加锁
                                while (foods.isEmpty()) foods.wait();
                                String food = foods.remove(0);
                                System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                                foods.notifyAll();
                            }
                            TimeUnit.MILLISECONDS.sleep(800);
                        }
                    } catch (InterruptedException e)
                    {
                        e.printStackTrace();
                    }
                };
            
                public void test()
                {
                    for (int i = 0; i < 10; i++)
                    {
                        new Thread(CONSUMER).start();
                    }
                    for (int i = 0; i < 5; i++)
                    {
                        new Thread(PRODUCER).start();
                    }
            
                    while (true) ;
                }
            
                public static void main(String[] args)
                {
                    new ProducerAndConsumerWithNotify().test();
                }
            }

            ```
        -   `FINISH`   lock/condition实现
            -   实现方式其实与上述的类似
        ```
        public class ProducerAndConsumerWIthLockAndCondition
        {
            private Lock lock;
            private Condition takeCondition;
            private Condition putCondition;
            private List<String> foods;
        
            public ProducerAndConsumerWIthLockAndCondition()
            {
                this.lock = new ReentrantLock();
                this.takeCondition = this.lock.newCondition();
                this.putCondition = this.lock.newCondition();
                this.foods = new LinkedList<>();
            }
        
            final Runnable PRODUCER = () ->
            {
                lock.lock();
                try
                {
        
                    while (!Thread.currentThread().isInterrupted())
                    {
                        // 如果引入了putCondition
                        // 就需要判断容量来限制了
                        while (this.foods.size() == 16) this.putCondition.await();
                        String food = UUID.randomUUID().toString();
                        foods.add(food);
                        takeCondition.signalAll();
                    }
        
                } catch (Exception e)
                {
                    e.printStackTrace();
                } finally
                {
                    lock.unlock();
                }
                try
                {
                    TimeUnit.SECONDS.sleep(1);
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            };
            final Runnable CONSUMER = () ->
            {
                lock.lock();
                try
                {
                    while (!Thread.currentThread().isInterrupted())
                    {
                        while (foods.isEmpty()) takeCondition.await();
                        String food = foods.remove(0);
                        System.out.println("消费者:" + Thread.currentThread().getName() + " 消费食物: " + food);
                        putCondition.signalAll();
                    }
                } catch (Exception e)
                {
                    e.printStackTrace();
                } finally
                {
                    lock.unlock();
                }
                try
                {
                    TimeUnit.MILLISECONDS.sleep(800);
                } catch (InterruptedException e)
                {
                    e.printStackTrace();
                }
            };
        
            public void test()
            {
                for (int i = 0; i < 10; i++)
                {
                    new Thread(CONSUMER).start();
                }
                for (int i = 0; i < 5; i++)
                {
                    new Thread(PRODUCER).start();
                }
        
                while (true) ;
            }
        
            public static void main(String[] args)
            {
                new ProducerAndConsumerWIthLockAndCondition().test();
            }
        }
        ```
    -   `WIP`   读写锁的实现
        -   `原理`: 
            -   读的时候判断是否有写,如果有则需要阻塞
            -   写的时候判断是否有读,如果有则需要阻塞
            -   释放的时候都需要notifyAll(`不需要只是单独的notify读或者写,需要提醒全部,因为可能写的时候还有多个请求写,让其自主竞争即可`)

多线程
---

- `WIP` 多线程CountDownLatch编写

- `WIP` 多线程CyclicBarrier 模拟赛马编写

- `WIP` 线程池

- `WIP` 对象池


-   `WIP`   queue
    -   `WIP`   原生queue的实现
    -   `FINISH`   stack实现queue
        -   要点: 栈是先进后出的,如果做到先进先出呢,通过2个栈
        即可,一个栈A用于接收数据,另外一个栈B用于弹出数据,**当然核心在于另外一个栈B
        弹出数据之间将栈A的数据先出栈然后push到栈B中,这样原先a->b->c顺序进的栈就编程了c->
        b->a的顺序,也就是a后进了,然后我们直接弹出即可**
        
        ```
            public T pop()
            {
                // 如果stack2不为空,则先将其内部的元素弹出去(这时候已经是先进先出的了)
                if (!stack2.isEmpty())
                {
                    return stack2.pop();
                }
                // 安全校验
                if (stack1.isEmpty())
                {
                    return null;
                }
                // 这里可以省去一部操作的,如注释所示,可以省去一个入栈
                while (!stack1.isEmpty())
                {
                    stack2.push(stack1.pop());
                }
                // 弹出最后一个
        //        for (int i = 0; i < stack1.size()-1; i++)
        //        {
        //            stack2.push(stack1.pop());
        //        }
        //        return stack1.pop();
                return stack2.pop();
            }
        ```
        
        
    -   `WIP`   LinkedBlockingQueue的实现
    
-   `FINISH`   stack
    -   `FINISH`   原生stack的实现
        -   要点: 要点其实没多少,push的时候从尾巴添加,pop也从尾巴弹出,主要就是记得pop的时候,如果底层是数组形式的话记得要小心数组越界,
        弹出的时候长度或者临时下标--
    -   `FINISH`   queue实现stack
        -   要点: 2个queue实现stack比2个stack实现queue稍微逻辑复杂一点点,核心
        只要记住:**2个队列,push或者pop的时候必然是一个为空队列,另一个为非空队列**,
        **每次添加元素都是往有值的队列中添加元素**,**弹出元素的时候,是将有值的那个队列中的除了最后一个元素全部pop然后push到另外一个队列,剩下的最后一个值就是最新插入的,直接弹出达到后进先出的效果**
        
        ```
        // 注意这里的linkedlist是要用add的,add是添加到末尾,而push是添加到第一个
        // 而我们要的pop是从末尾弹出的
        // push时候的注意点,我们要保持一个队列为空,另外一个队列不为空,所以我们每次添加都是往有值的队列中添加元素
        // 至于第一次添加,就随意了,给第一个还是第二个都可以
        public void push(T value)
        {
            if (queue1.isEmpty())
            {
                queue2.add(value);
            }else if (!queue2.isEmpty())
            {
                throw new RuntimeException("逻辑错误,内部某块逻辑错误");
            }else{
                queue1.add(value);
            }
        }
        // 必须保持2个队列一个为空,另外一个不为空
        // 出队就是将一个队列中的元素移到另外一个队列
        public T pop()
        {
            T temp = null;
            if (!queue1.isEmpty() && queue2.isEmpty())
            {
                // 最后一个元素是我们想要的元素(既最后入队的最先弹出)
                for (int i = 0; i < queue1.size() - 1; i++)
                {
                    temp = queue1.pop();
                    queue2.add(temp);
                }
                temp = queue1.pop();
            } else if (!queue2.isEmpty() && queue1.isEmpty())
            {
                // 最后一个元素是我们想要的元素(既最后入队的最先弹出)
                for (int i = 0; i < queue2.size() - 1; i++)
                {
                    temp = queue2.pop();
                    queue2.add(temp);
                }
                temp = queue2.pop();
            } else if (!queue1.isEmpty() )
            {
                // 说明2个队列都不为空,这是不可能也不该出现的
                throw new RuntimeException("逻辑错误");
            } else
            {
                return null;
            }
            return temp;
        }
        ```

-   `WIP` 树
    -   `WIP`   二叉树的创建
        -   `FINISH`   递归先序创建  
        ```
        // 递归创建普通的树
        public TreeNode loopBuildTree(TreeNode node, Integer[] arr)
        {
            if (index >= arr.length || arr[index] == -1)
            {
                this.index++;
                return null;
            }
            node = new TreeNode();
            node.setData(arr[this.index++]);
            node.setLeftChild(loopBuildTree(node.getLeftChild(), arr));
            node.setRightChild(loopBuildTree(node.getRightChild(), arr));
            return node;
        }
        ```
        -   `FINISH`   非递归先序创建
        ```
        // 非递归创建普通二叉树,因为是给定的数组,所以以-1代表空,流程逻辑如下
        // 核心就是二叉树节点满的时候只有2个节点
        //  判断data==-1   true?说明要么插入右节点,要么是左右节点都没有,如果是插入右节点,则新node入队之后还需要修改方向为左
        //				   false:判断插入的是左节点还是右节点,如果插入的是右节点,则需要修改方向为左(因为二叉树只有2个节点,
        // 						  并且是先序创建:根左右) ,最后元素入队(因为可能后面的值不是-1,是要继续插入的)
        //	ps: 在流程中并不需要对stack进行判空,因为开始之前先将根节点入队了,已经确保了不会为空
        //  ps: 同时,其实二叉树是当遇到连续2个-1的时候会pop,那么n叉树则可以认为是遇到n个-1时是会出队的
        //  ps : 因此,当n叉树时,判断的情况为: if counts!=n {} else{}
        //  总结: 出队的情况: n叉树遇到n个连续的-1(子节点出栈,跳回父节点) | n叉树的节点满了(父节点出栈,子节点开始接收值)
        //       与递归创建不同,栈实现只需要一个for循环即可,也**不需要自己内部++操作**
        func (t *BinaryTree)BuildTreeWithStack(arr []int){
        	if len(arr)==0{
        		return
        	}
        	t.root=&TreeNode{data:arr[0]}
        	left:=true
        	stack := arraystack.New()
        	stack.Push(t.root)
        	for i:=1;i< len(arr);i++{
        		if arr[i]==-1{
        			if left{
        				left=false
        			}else if !stack.Empty() {
        				stack.Pop()
        			}
        		}else{
        			node:=&TreeNode{data:arr[i]}
        				temp, _ := stack.Peek()
        				if left{
        					temp.(*TreeNode).leftChild=node
        				}else{
        					temp.(*TreeNode).rightChild=node
        					left=true
        					stack.Pop()
        				}
        				stack.Push(node)
        		}
        	}
        }
        ```
    -   `WIP`   二叉树的CRUD
    
    -   `FINISH`   满二叉树
        * 定义: 满二叉树是指除了最后一层外,每个节点都有2个孩子
        * 特性:
            -   层数为k,一定有2^k -1个节点
            -   第i层的节点数为2^(i-1) 
            
    -   `FINISH`   完全二叉树   
        * 定义: 是一种效率很高的数据结构,除了最后一层外,其余各层
        的节点数都达到了最大值2,并且最后一层的节点都在左边,当深度为k
        的树,只有当与满二叉树的节点顺序一致的才为完全二叉树
        * 特性: `前提是以数组的0下标为起始判断的`
            -   高度差满足:[0,1] (因为顺序是按照满二叉树的顺序走的)
            -   当n个节点时:
                -   下标为i的左儿子下标为:`2*i+1` (当然范围要<n)
                ,右孩子的下标为2*i+2   既:左0右1,**因此构建树需要通过节点下标来构建**
                同时我们可以发现 **只有[0,n/2-1]有孩子节点**
        -   `FINISH`   完全二叉树的另外一种遍历方式:基于节点的特性,可直接避免构造树而通过对数组进行遍历(写先序和bfs)
        
        **注意点也是相同的,就是基于数组的特殊性,从1开始(从0开始会栈溢出),真正操作要减去1**
                 
        ```
            public void inIteratorByArray(Integer[] arr, Integer index, List<Integer> resultList)
            {
                if (index <= arr.length)
                {
                    // ROOT
                    resultList.add(arr[index - 1]);
                    // LEFT
                    this.inIteratorByArray(arr, 2 * index, resultList);
                    // RIGHT
                    this.inIteratorByArray(arr, 2 * index + 1, resultList);
                }
            }
        ```
        -   `FINISH`   完全二叉树的构建
        
        ```
          // 构建一颗完全二叉树
          // 构件完全二叉树的时候我们需要判断,左孩子节点 2*i+1 是否超过长度  <length,右孩子是否超过长度:2*i+2<length
          // 同时,当左孩子不存在的时候,右孩子就没必要判断了
          public void buildCompleteBinaryTree(Integer[] arr)
          {
      
              List<TreeNode> nodeList = new ArrayList<>();
              this.root = new TreeNode(arr[0]);
              nodeList.add(this.root);
              for (int i = 1; i < arr.length; i++)
              {
                  nodeList.add(new TreeNode(arr[i]));
              }
              Integer length = arr.length >> 1;
              for (int i = 0; i <= length; i++)
              {
                  if (i * 2 + 1 < arr.length)
                  {
                     nodeList.get(i).setLeftChild(nodeList.get(i * 2 + 1));
                     if (i * 2 + 2 < arr.length)
                     {
                         nodeList.get(i).setRightChild(nodeList.get(i * 2 + 2));
                     }
                   }
         
              }
          }
          
        ```
        
    	-	`FINISH`	判断是否是完全二叉树
    	    -   思路: 完全二叉树根据定义来判断即可 
    	        -   如果有右孩子,却没左孩子肯定不是(左:2*i+1,右:2*i+2)
    	        -   如果该节点不存在右孩子,则遍历该层,判断该层的后继节点是否是叶子都
    	        是节点,如果有一个不是,则不是完全二叉树(叶子节点都在左边的特性)
    	        
    	```
    	// 采取标志位的方式:
        // 如果一棵树只有左节点,则标志位为true,代表不完整
        // 若后续的树不完整(意味着left+right|left|right)则不是完全二叉树,既一旦有节点就不是完全二叉树
        // 非二叉树的条件:
        //  1. 当存在右孩子,左孩子却为空
        //  2. 当存在左孩子,右孩子为空,而同层的后续节点有孩子(既可以认为是之前的节点是不完整的)
        // 注意点的话只需要注意下退出条件以及元素入队的条件即可
        public boolean validIfCompleteTree()
        {
            if (this.root == null)
            {
                return false;
            }
            LinkedList<TreeNode> queue = new LinkedList<>();
            boolean previousCompleted = true;
            queue.add(this.root);
            for (TreeNode temp = queue.pop(); temp != null; temp = queue.pop())
            {
                // 1.对第一种条件的判断 和对第二种条件的判断
                if (temp.leftChild == null && temp.rightChild != null ||
                        !previousCompleted && (temp.leftChild != null || temp.rightChild != null))
                {
                    return false;
                } else
                {
    
                    // left=null &right=null
                    // left!=null & right=null || left!=null & right!=null
                    // previousCompleted
                    if (temp.leftChild != null)
                    {
                        queue.add(temp.leftChild);
                    }
                    if (temp.rightChild != null)
                    {
                        previousCompleted = true;
                        queue.add(temp.rightChild);
                    } else
                    {
                        previousCompleted = false;
                    }
                }
                if (queue.isEmpty())
                {
                    break;
                }
            }
            return true;
        }
    	```     
   
    -	`FINISH`	二叉查找树
        * 定义: 左子树的值一定小于根节点的值,右子树的值一定大于等于根节点的值
        - 二叉树的创建: 只需要遍历整棵树,根据大小关系,然后判断是否为空即可(新的元素会插在空的地方)
        
        ```
        // 插入节点逻辑很简单,就是找到插入节点的父节点
        // O(logn)遍历整颗树,如果小于父节点则遍历左子树,否则右子树,只需要判断是否为空即可
        func (t *BinarySearchTree) InsertNode(value int) error {
            //if nil == value {
            //	return common.NilPointerError
            //}
            if nil == t.root {
                t.root = &BinarySearchTreeNode{data: value}
                return nil
            }
            newNode := &BinarySearchTreeNode{data: value}
            tempNode := t.root
            for {
                if newNode.data < tempNode.data {
                    if nil == tempNode.leftChild {
                        tempNode.leftChild = newNode
                        break
                    }
                    tempNode = tempNode.leftChild
                } else {
                    if nil == tempNode.rightChild {
                        tempNode.rightChild = newNode
                        break
                    }
                    tempNode = tempNode.rightChild
                }
            }
            t.size++
            return nil
        }
        ```
        - 注意点: 二叉查找树稍微复杂点的逻辑在于:二叉查找树树的删除操作 
            -   `若删除节点没有右孩子,则左孩子作为新的根节点`
            -   `若删除节点的右孩子没有左孩子,则这个右孩子作为新的根节点`
            -   `若删除节点的右孩子有左孩子,这个右孩子的左孩子作为新的根节点`
            ![可参考网上图](https://images0.cnblogs.com/i/175043/201406/291214353511360.gif)
            
            ```
            func (t *BinarySearchTree) DeleteNode(value int) error {
            	tempNode := t.root
            	var lastVisitNode *BinarySearchTreeNode // 保留的是父节点的地址
            	for nil != tempNode && tempNode.data != value {
            		lastVisitNode = tempNode
            		if value < tempNode.data {
            			tempNode = tempNode.leftChild
            		} else {
            			tempNode = tempNode.rightChild
            		}
            	}
            	// 可能的情况: 到了末尾,从而tempNode为空了
            	// 或者是匹配到了这个值
            	if nil == tempNode {
            		return common.NoSuchElementError
            	}
            	var newRootNode *BinarySearchTreeNode // 减少代码块,这里提前定义一个变量
            	// 说明有这个值,此时lastVisitNode保存的是其父节点的值
            	// 1. 判断右孩子是否存在
            	if nil == tempNode.rightChild {
            		// 不存在则直接将左孩子作为根节点
            		// 此时是不需要对原先根节点的其他节点进行连接的,因为此时的根节点只有左孩子
            		newRootNode = tempNode.leftChild
            		// 2. 判断右孩子的的左孩子是否为空
            	} else if nil == tempNode.rightChild.leftChild {
            		// 如果右孩子的左孩子为空,则直接将这个右孩子作为根节点
            		// 同时原先根节点的左孩子可能不为空,因而我们需要重新连接,但能确保的是原先节点的左孩子节点必定是比这个新节点小的
            		// 因而直接赋值即可
            		newRootNode = tempNode.rightChild
            		newRootNode.leftChild = tempNode.leftChild
            	} else {
            		// 说明右孩子的左孩子不为空,则将这个右孩子的左孩子作为新的根节点
            		newRootNode = tempNode.rightChild.leftChild
            		// 此时他需要连接的元素:原先根节点的左孩子+原先根节点的右孩子
            		// 而这个元素的原先位置也需要进行变更,因此此时需要递归进行处理
            		t.DeleteNode(newRootNode.data)
            		newRootNode.leftChild, newRootNode.rightChild = tempNode.leftChild, tempNode.rightChild
            	}
            	if value < lastVisitNode.data {
            		lastVisitNode.leftChild = newRootNode
            	} else {
            		lastVisitNode.rightChild = newRootNode
            	}
            	return nil
            }
            ```
    -   `WIP`   平衡二叉树
        * 定义: 平衡二叉树是二叉查找树的延伸,因为二叉查找树存在极端情况:如形成的是一个链表
               而平衡二叉树使得高度差最大为1
               
    -   `WIP`   B+ tree Java 实现
    
    -   `WIP`   B tree Java实现
    -   `WIP`   LSM tree Java 实现
    
    -   `WIP`   树的遍历
     
        **关于树的遍历,当涉及递归的时候,内部只需要两个函数,
        如先序:只需要入队node.Left和node.Right即可,不要入node,死循环了,
        至于关于打印的先后顺序,只需要记得,参数中的node都是根节点即可**
        -   `FINISH`   递归先序
        -   `FINISH`   递归中序
        -   `FINISH`   递归后序
        -   `FINISH`   非递归先序
            -   注意点: 注意跳出循环的条件即可:
                -   最外层循环:node为空且栈为空
                -   内层循环: node不为空即可,将其调整到左孩子    
         
        ```
            public List<Integer> preorderTreeByStack()
        {
            Stack<TreeNode> stack = new Stack<>();
            List<Integer> resultList = new ArrayList<>();
            TreeNode tempNode = this.root;
    
            // 为什么这里的限定条件还要有栈不为空
            // 原因在于: 1. 若栈不为空则表明内部还有元素没有遍历完成
            // 2. 如果当方向调整为右边之后,而右孩子为空,则需要继续弹到上一个节点
            while (tempNode != null || !stack.isEmpty())
            {
                // 如果右孩子为空之后,这里就不会执行了,既此时10的右孩子为空
                // 则又会弹出,弹回9,9又会对节点判断,发现右边也是为空,则又会弹回4,8之前已经弹出去了
                while (tempNode != null)
                {
                    resultList.add(tempNode.data);
                    // 根节点入栈
                    stack.push(tempNode);
                    // 如果左孩子不为空,则左孩子也会入栈
                    tempNode = tempNode.leftChild;
                }
                 /*
                        1
                      2    3
                    4   5 6  7
                  8
                   \
                    9
                   /
                  10
    
                 */
                // 上述的退出条件是这个根节点没有左孩子了,如上述树中的8节点
                // 因为是先序,根左右,所以我们需要弹出(弹回8)调整方向为右边
                // 因此
                if (!stack.isEmpty())
                {
                    tempNode = stack.pop();
                    tempNode = tempNode.rightChild;
    
                }
            }
            return resultList;
        }
        ```
         
        -   `FINISH`   非递归中序
            -   注意点: 与非递归先序是相同的
                -   注意外层循环退出的条件:tempNode=null & stack.isEmpty
                -   内层循环退出的条件是tempNode=null
                -   **与先序不同的地方在于:因为中序是先左孩子,因而将元素的操作是放在if代码块中**
           ```
            public List<Integer> inOrderTreeByStack()
            {
                List<Integer> resultList = new ArrayList<>();
                Stack<TreeNode> stack = new Stack<>();
                TreeNode tempNode = this.root;
                while (null != tempNode || !stack.isEmpty())
                {
                    while (null != tempNode)
                    {
                        stack.push(tempNode);
                        tempNode = tempNode.leftChild;
                    }
                    if (!stack.isEmpty())
                    {
                        tempNode = stack.pop();
                        resultList.add(tempNode.data);
                        tempNode=tempNode.rightChild;
                    }
                }
                return resultList;
            }           
           ```
        -   `FINISH`   非递归后序
            -   注意点:
                -   后序遍历与先序和中序遍历不同,是左右根,也就意味着在if阶段需要判断这个节点的有节点
                是否已经遍历过了,当遍历过了才可以进行处理
                -   通过设定一个临时的指针:lastVisitNode 判断上一个访问地址是否这个节点的右节点
                
        ```
        public List<Integer> postOrderTreeByStack()
        {
            List<Integer> resultList = new ArrayList<>();
            Stack<TreeNode> stack = new Stack<>();
            TreeNode tempNode = this.root, lastVisitTopNode = null;
            while (null != tempNode || !stack.isEmpty())
            {
                while (null != tempNode)
                {
                    // 左右根
                    stack.push(tempNode);
                    tempNode = tempNode.leftChild;
                }
    
                // 说明这个节点没有左孩子了,则需要调整方向
                // 但是后序遍历,是左右根,根是最后才会操作的,因而我们需要判断下一个操作的节点是否是右节点
                // 因而有两种方式:
                //      第一种是通过一个指针记录栈头:
                //              用一个指针通过判断上次访问的值与当前值的右节点是否匹配,如果不匹配,则需要先对右节点进行操作
                //      第二种方式是通过对每个节点额外的添加一个变量判断是否遍历过了:
                //              1.我们需要一个标志位来判断这个根节点是否被操作过了
                //      两种方式首先都不能直接pop的,因为我们的右孩子还不知道是否已经被遍历过了
    
                // 如果右节点已经被
                /*
                        1
                    2       3
                     \
                      4
                     / \
                    5   6
                   1入栈->2入栈->2的左孩子为空则跳出while循环->lastVistNode为空则调整方向,调整为2的右孩子->
                   4入栈->5入栈->5的左孩子为空跳出while循环->
                   5的右孩子也为空可以类似表明为5的左右孩子都遍历过了,则表明可以对数据(5的数据)操作了,同时调整
                   之前访问的栈顶元素为5,最后将元素置为空,这样下一轮循环就会在if中判断当前栈顶元素4是否对右孩子遍历过了->
                   栈顶为4->直接进入if判断->4的右孩子不为空且!=lastVisitNode->进入else,对右边操作,6入栈->
                   6的左孩子为空,跳出循环->右孩子也为空,对6的数据操作,设置lastVisitNode为6,弹出6->
                   此时栈顶为4,lastVisitNode=右孩子,因而4的左右孩子也遍历过了,因此同理.....->>>
                 */
                // 上面的while退出的情况为:当这个节点的左孩子为空的时候就会退出while循环
                // 因而我们需要遍历父节点的右孩子,但是需要判断: 是否已经遍历过了
                if (!stack.isEmpty())
                {
                    tempNode = stack.peek();
                    // 判断之前访问的节点是否是右节点
                    if (null == tempNode.rightChild || lastVisitTopNode == tempNode.rightChild)
                    {
                        // 说明已经对右节点访问过了,则直接数据处理即可
                        resultList.add(tempNode.data);
                        lastVisitTopNode = tempNode;
                        stack.pop();
                        tempNode = null;
                    } else
                    {
                        // 则继续遍历右节点
                        tempNode = tempNode.rightChild;
                    }
    
                }
            }
            return resultList;
        }
        ```
        -   `WIP`   非递归遍历总结:
        -   `FINISH`   BFS
        
        ```
           //BFS(又名) 基于队列,较dfs性能高,但是更耗内存
           // 3个陷阱:  1. queue#push的时候要进行判断是否为空,否则空的也放进去了
           //          2. pop的时候要对是否为空进行判断,否则empty也pop了
           //          3. 第三个陷阱其实是第二个陷阱的衍生,就是判空不能放在for循环里(会造成第一次就结束)
           //             只能放在循环体内部,符合条件退出
           public List<Integer> BFSTree()
           {
               LinkedList<TreeNode> queue = new LinkedList<TreeNode>();
               queue.push(this.root);
               TreeNode temp = null;
               LinkedList<Integer> resultList = new LinkedList<>();
               for (temp = queue.pop(); null != temp; temp = queue.pop())
               {
                   // 层级遍历,一层一层打印
                   resultList.add(temp.getData());
                   if (temp.leftChild != null)
                   {
                       queue.add(temp.leftChild);
                   }
                   if (temp.rightChild != null)
                   {
                       queue.add(temp.rightChild);
                   }
       
                   if (queue.isEmpty())
                   {
                       return resultList;
                   }
       
               }
               return resultList;
       
           }
        ```
        
        -   `FINISH`   DFS
        ```
        // DFS 基于栈,较BFS内存占用更少,但是性能较之为低
        // 与BFS相比,不同的地方只有两点:
        // 1.bfs使用的是队列 而dfs使用的是栈
        // 2.bfs的结果是一层一层的顺序,而dfs则是一条连线(在test目录下将结果打印配合自己画图更明确)
        public List<Integer> DFSTree()
        {
            List<Integer> resultList = new ArrayList<>();
            Stack<TreeNode> stack = new Stack<>();
            stack.push(this.root);
            TreeNode temp = null;
            for (temp = stack.pop(); null != temp; temp = stack.pop())
            {
                resultList.add(temp.getData());
                if (null != temp.rightChild)
                {
                    stack.push(temp.rightChild);
                }
                if (null != temp.leftChild)
                {
                    stack.push(temp.leftChild);
                }
    
                if (stack.isEmpty())
                {
                    return resultList;
                }
            }
            return resultList;
        }
        ```
        
        
 算法
 ---   
 
-   `WIP`   查找算法
    -  `WIP`   二分查找     
-   `WIP`   判断一个链表是否有环
    -   `FINISH`   通过两个指针进行判断
        -   要点: 要点就是2个指针,一个快,一个慢,快指的是每次移动到next的时候多移动一次
        ,遇到这个坑过(silly),既快一步不是初始化的时候快一步,而是在`移动的时候快一步`
        ```
        // true :loop else non-loop
        // 遇到的坑: 内部的quickNode 要多走一步,而不是所初始化的时候多走一步,如果内部不多走一步是永远与slowNode相连的
        func (l *LinkedList)ValidIfLoop()bool{
        	if l.size==0 {
        		return false
        	}
        	slowNode:=l.root
        	quickNode:=slowNode.next
        	for {
        		if nil==quickNode{
        			return false
        		}else if quickNode==slowNode{
        			return true
        		}else {
        			// 遇到一个坑,这里不能直接把quickNode的值给他
        			// 不过遇到这个坑的缘故在于少了 内部的if那块代码,如果没有那块代码,quickNode与slowNode相连
        			slowNode=slowNode.next
        			quickNode=quickNode.next
        			if nil!=quickNode {
        				quickNode=quickNode.next
        			}
        		}
        	}
        }
        ```
    -   `FINISH(Go实现)`   通过map来判断并且获取环入口节点
        -   若允许额外申请空间则可以通过map来判断
        -   核心就是通过判断这个节点对应的value是否有值,有值的话就代表着节点重复了,通过map的话并且可以直接获取到这个值,但是需要额外的申请内存空间
        ```
        // 通过map 额外申请内存来判断是否存在环
        func (l LinkedList)ValidIfLoopByMap()bool{
        	validMap:=make(map[*ListNode]struct{})
        	tempNode:=l.root
        	for nil!=tempNode{
        		if _,ok:=validMap[tempNode] ;!ok{
        			validMap[tempNode]= struct{}{}
        		}else{
        			return true
        		}
        	}
        	return false
        }
        ```
    -   `FINISH(Go)`    获取回环节点
        -   通过map的话很简单,直接返回即可,而如果通过快慢指针的话,则需要:
            -   慢节点重新指向root节点,快节点不变
            -   慢节点和快节点都以同样的速率前进,当相等时的节点就是碰撞节点`(这里的方案都是指quickNode和slowNode从同一个节点出发)`
       -    注意点:`我初始化的时候quickNode就比slowNode快了一步,所以当break碰撞之后,quickNode也需要先提前一步(既先跳到next)`     
       ```
       // 通过快慢指针获取回环的节点
       // 遇到的坑: 当break的时候代表发生了碰撞,因为我采用的方法,初始化的时候快节点会比慢节点快一步,所以当break获取
       // 回环节点的时候也需要先快一步
       func (l *LinkedList)GetLoopNode()*ListNode{
       	// 这个方法的测试的目标是:必有回环的链表
       	if l.size==0{
       		return nil
       	}
       	slowNode:=l.root
       	quickNode:=slowNode.next
       	for{
       		slowNode=slowNode.next
       		quickNode=quickNode.next
       		if nil!=quickNode {
       			quickNode=quickNode.next
       		}
       		if nil==quickNode {
       			return nil
       		}else if quickNode==slowNode{
       			// 说明有环碰撞了
       			break
       		}else{
       			slowNode=slowNode.next
       			quickNode=quickNode.next
       			if nil!=quickNode{
       				quickNode=quickNode.next
       			}
       		}
       	}
       	// 如果发生碰撞,则慢节点从表头出发,快节点从碰撞处出发,两者的运动速率是一致的,如果2者相等就是回环的节点
       	// 这里有一个坑,因为最开始quickNode就比slowNode快一步(初始化的时候就快一步,所以这里也需要先提前快一步)
       	slowNode=l.root
       	quickNode=quickNode.next
       	for quickNode!=slowNode{
       		slowNode=slowNode.next
       		quickNode=quickNode.next
       	}
       	return slowNode
       }
       ```

- `WIP` 页面置换算法
    -   `WIP`   最佳置换算法(OPT)
        -   原理: 
    -   `WIP`   先进先出置换算法(FIFO)
        -   原理: 既队列
    -   `FINISH`   最近最久未使用算法(LRU)
        -   原理:
            -   对于访问次数最多的放在表头(或者尾),那么表尾(或者头)就是访问次数最少的元素,也就是移除的元素
            -   注意点:
                -   移动的时候要判断元素是
                    -   `头结点`: 更改头节点的位置,指向next即可
                    -   `尾节点`: 啥也不需要管
                    -   `table的首节点`: 首节点要更换
                    -   `中间节点`: 前驱要指向后继,后继要指向前驱,这个节点指向prev指向tail,tail的next指向这个,然后tail移动到这个
Spring
---

-   `FINISH`   Spring动态代理的实现
    -   `FINISH`   两者区别的总结
        -   JDK是基于接口的实现方式,意味着被代理的对象必须要实现某个接口才行,而Cglib是基于asm字节码技术,子类继承父类复写非final方法在其中添加callback
        的自定义逻辑实现的.意味着被代理的对象不能被final所修饰,并且需要提供默认的无参构造函数
        -   cglib的效率高于jdk动态代理
    -   `FINISH`   基于jdk的动态代理
        -   核心是InvocationHandler 和Proxy.newProxtInstance
        -   自定义类实现InvocationHandler,复写invoke方法
        -   在invoke方法内部添加before或者after或者异常的逻辑,如果要调用原先的方法的话是第二个参数method.invoke(target,args)调用的
        -   注意点:
            -   `自定义handler类中需要有一个Object成员变量`,复写的invoke方法中method.invoke调用的参数是**这个自定义handler中的对象,而不是原先参数中的对象**
        
        ```
        @Data
        public static class MyInvocationHandler implements InvocationHandler
        {
            public MyInvocationHandler(Object target)
            {
                this.target = target;
            }
    
            private Object target;
    
            @Override
            public Object invoke(Object proxy, Method method, Object[] args) throws Throwable
            {
                System.out.println("before ");
                Object result = method.invoke(target, args);
                System.out.println("after");
                if (result instanceof String)
                {
                    return "-----" + result + "+++++";
                }
                return result;
            }
        }
        ```  
            
    -   `FINISH`   基于cglib的动态代理
        -   核心是MethodIntercetor,Enhancer,MethodProxy
        -   自定义类实现MethodInterceptor接口 就已经ok了,但是通常我们还会为其编写一个工厂方法,用于创建对象
        -   调用:创建Enhancer->设置superClass->设置callBack接口的实现类(既我们之前的自定义的那个)->enhancer.create
        
        ```
        public static MethodInterceptor CglibProxyTest = (target, method, args, methodProxy) ->
        {
            System.out.println("before call ");
            String superName = methodProxy.getSuperName();
            System.out.println(superName);
            Object result = methodProxy.invokeSuper(target, args);
            System.out.println("after call");
            if (result instanceof String)
            {
                return "-----" + result + "+++++++++";
            }
            return result;
        };
           ``` 
        
    -   `注意点`:如果是想实现aop的功能,通过JDK的时候必须得传入一个`原先的实现类`,但是如果是想动态生成接口的代理(既原先没有实现类,参考@FeignClient)则可以不需要传入实例对象
     

-   `WIP`   Bean生命周期的探讨




常见问题及其解决方案
---

-   `WIP`   海量数据问题:

    - `WIP` TOK 解决方案:
        -   `WIP`   全部排序
        -   `WIP`   局部淘汰法
        -   `WIP`   分治法
        -   `WIP`   hash法
    -   `WIP`   bitmap寻找重复元素或者判断个别元素是否在海量数据中存在
    

-   `WIP` Java导出excel格式的文件


-   `WIP` 负载均衡算法
    -   `WIP` 轮询法
    -   `WIP` 随机法
    -   `WIP` 源地址hash法
    -   `WIP` 加权轮询法
    -   `WIP` 加权随机法
    -   `WIP` 最小连接数


-   `WIP`   4种引用类型的实例
    -   总结:
        -   软引用和弱引用非常适合作为短暂的存储结构,如我之前dlxy项目中的
        访问每篇文章对应的访问次数都会+1还有其他记录,忘了,当初用的是享源模式,到了
        后期那个map会无比巨大,因为被map持有不会被gc,因此解决方法是通过构建软引用的
        告诉缓存,具体就是通过ConcurrentHashMap存储<key,SoftReference<T>>,同时
        还有一个ReferenceQueue用于清理对象(因为软引用被回收的时候会被收集到这里,而我们就可以
        通过这个ReferenceQueue去Map中清理已经无效的缓存(**ReferenceQueue中的值代表着是无元素引用他,因而完全可**)(当然这里可以交给线程池,因为毕竟是CHM))
        `其实当为弱引用的时候,这个告诉缓存与Go中的sync.Pool类似,都是第一次GC就回收不用的对象,不过Go的sync.Pool是全部因为内部获取的时候会移除`
    -   `FINISH`   强引用
        -   既直接new的方式赋值
    -   `FINISH`   软引用
        -   既通过SoftReference<T>方式保存,通过get获取原先
        对象的引用,如果对象不再被其他对象所持有(或者超出了范围,如作为局部变量),当内存不足,快要发生OOM时会被gc
    -   `FINISH`   弱引用
        -   每次gc都会使得对象被回收(如果对象没有被其他对象所持有)
    -   `WIP`   虚引用
    -   `FINISH(Java)`   软|需引用构造高速缓存 

-   `WIP`   时间轮算法实现延迟处理

MQ
---

