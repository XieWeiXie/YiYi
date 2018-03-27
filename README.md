# YiYi
A cli command tool for reader from DouBan and One App API.

![](https://img.shields.io/badge/YiYi-v0.0.1-519dd9.svg)
![](https://img.shields.io/badge/language-golang-orange.svg)
[![](https://img.shields.io/badge/weibo-@谢小小路-red.svg)](https://weibo.com/1948244870/profile?topnav=1&wvr=6)
[![](https://img.shields.io/badge/jianshu-@谢小路-F59581.svg)](https://www.jianshu.com/u/58f0817209aa)



### API

**图书**: 豆瓣图书API

|URL|Content|
|---|---|
|https://api.douban.com/v2/book/series/:id/books|获取一系列书，:id 可以是任意数字|
|https://api.douban.com/v2/book/isbn/:name|根据isbn获取书的详细信息，其中:name可以替换成isbn号码|
|https://api.douban.com/v2/book/search| 搜索书，可以设置参数搜索相关|


搜索参数:

|参数|意义|说明|
|---|---|---|
|q|查询关键字|q 和 tag 必须选择一个|
|tag|查询标签|q 和 tag 必须选择一个|
|start| 结果的offset| 默认为 0|
|count| 结果的条数 | 默认为 20, 最大的为 100|

---

**电影**: 豆瓣电影API

|URL|Content|
|---|---|
|https://api.douban.com/v2/movie/in_theaters|上映的电影|
|http://api.douban.com/v2/movie/subject/:id|电影的详细信息|
|https://api.douban.com/v2/movie/coming_soon|即将上映的电影|
|http://api.douban.com/v2/movie/top250|Top 250 排行榜电影|





---


**故事**: ONE APP API

|URL|描述|
|---|---|
|http://v3.wufazhuce.com:8000/api/onelist/idlist|获取最新 idlist, 以获取今日或往日的 onelist 信息|
|http://v3.wufazhuce.com:8000/api/onelist/:data/0|获取某一天的onelist,其中:data替换成上面的idlist中的数据|
|http://v3.wufazhuce.com:8000/api/essay/:item_id|	获取故事详细信息，其中:item_id替换成onelist中的item_id值|

---






### Usage

```text

YiYi --help

```

```text
YiYi is a tool for reading with DouBan and One APP api.


NAME:
   YiYi - An application for book, movie, and story from DouBan and One App.

USAGE:
   YiYi [global options] command [command options] [arguments...]

VERSION:

    ___       ___       ___       ___
   /\__\     /\  \     /\__\     /\  \
  |::L__L   _\:\  \   |::L__L   _\:\  \
  |:::\__\ /\/::\__\  |:::\__\ /\/::\__\
  /:;;/__/ \::/\/__/  /:;;/__/ \::/\/__/
  \/__/     \:\__\    \/__/     \:\__\
             \/__/               \/__/   v0.0.1


DESCRIPTION:
   An application for book, movie, and story from DouBan and One App.

AUTHOR:
   xieWei <wuxiaoshen@shu.edu.cn>

COMMANDS:
     book     get book info from DouBan API
     story    get story info from One API
     movie    get movie info for DouBan API
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```


### For example 


命令集合：

<img src="https://github.com/wuxiaoxiaoshen/YiYi/blob/master/doc/image/YiYi.png?raw=true" alt="支付宝" align=center />

1. YiYi.exe book random
```text
Show Book from DouBan Api:
    AllCollections:

            Title: 希尼诗文集
            Subtitle:
            URL: https://api.douban.com/v2/book/1042723
            Isbn10: 7506320193
            Isbn13: 9787506320191
            Price: 38.00元

            Title: 帕斯选集（上下）
            Subtitle:
            URL: https://api.douban.com/v2/book/1854313
            Isbn10: 7506335263
            Isbn13: 9787506335263
            Price: 98.00元

            Title: 田园交响曲
            Subtitle:
            URL: https://api.douban.com/v2/book/1941575
            Isbn10: 7506337266
            Isbn13: 9787506337267
            Price: 29.00

            Title: 魔山
            Subtitle:
            URL: https://api.douban.com/v2/book/1945985
            Isbn10: 750633741X
            Isbn13: 9787506337410
            Price: 65.00元

            Title: 幻象
            Subtitle: 生命的阐释
            URL: https://api.douban.com/v2/book/1950662
            Isbn10: 7506337878
            Isbn13: 9787506337878
            Price: 28.00元

            Title: 萧伯纳戏剧选
            Subtitle: 16开
            URL: https://api.douban.com/v2/book/1950749
            Isbn10: 7506337584
            Isbn13: 9787506337588
            Price: 33.00元

            Title: 万延元年的Football
            Subtitle:
            URL: https://api.douban.com/v2/book/1951744
            Isbn10: 7506337932
            Isbn13: 9787506337939
            Price: 25.00

            Title: 英国旗
            Subtitle:
            URL: https://api.douban.com/v2/book/1953719
            Isbn10: 7506337657
            Isbn13: 9787506337656
            Price: 36.00元

            Title: 雪国.古都
            Subtitle:
            URL: https://api.douban.com/v2/book/1954706
            Isbn10: 7506337398
            Isbn13: 9787506337397
            Price: 33.00

            Title: 阿尔谢尼耶夫的一生
            Subtitle:
            URL: https://api.douban.com/v2/book/1955395
            Isbn10: 7506338041
            Isbn13: 9787506338042
            Price: 39.00元

            Title: 巴比特
            Subtitle:
            URL: https://api.douban.com/v2/book/1955935
            Isbn10: 7506337401
            Isbn13: 9787506337403
            Price: 33.00元

            Title: 泰戈尔诗文精粹
            Subtitle: 16开
            URL: https://api.douban.com/v2/book/1970307
            Isbn10: 7506337258
            Isbn13: 9787506337250
            Price: 35.0



``` 

2. YiYi.exe book detail 7506337258

```text

Show BookDetail for DonBan Api.
    ID: 1970307
    Title: 泰戈尔诗文精粹
    Average: 8.8
    Author: [&#34;拉宾德拉纳斯·泰戈&#34;]
    Publisher: 作家
    PublicationDate: 2006-11
    ImageURL: https://img1.doubanio.com/mpic/s10081887.jpg
    Translator: [&#34;冰心&#34;,&#34;郑振铎&#34;,&#34;西蒙&#34;,&#34;倪培耕&#34;]
    URL: https://api.douban.com/v2/book/1970307
      Summary: 泰戈尔多才多艺，才华超人。既是作品浩繁的文学艺术大师、学识渊博的哲人、成就卓著的社会活动家，也是锐意革新的教育家。他在印度文化的各个方面都产生了广泛而深远的影响。而他最突出的天才的表现，恐怕就是他惊人的创作量了。他12岁开始
写诗，在60余年的笔耕生涯中，创作了大量作品，其中有诗歌上千首，歌词1200余首，并为其中大多数歌词谱了曲； 中长篇小说12部，短篇小说200多篇，戏剧38部，还有许多有关哲学、文学、政治的论文及回忆录、书简、游记等；此外还创作了2700余幅画。他给印度和
世界留下了一笔异常丰富的文化遗产。
    Catalog:
    Price: 35.0
```

3. YiYi.exe book search query Golang

```text
Show Book from DouBan Api:
    AllCollections:

            Title: Cloud Native programming with Golang: Develop microservice-based high performance web apps for the cloud with Go
            Subtitle: Discover practical techniques to build cloud-native apps that are scalable, reliable, and always available.
            URL: https://api.douban.com/v2/book/30154908
            Isbn10: 178712598X
            Isbn13: 9781787125988
            Price: USD 44.99

            Title: Building RESTful Web services with Go: Learn how to build powerful RESTful APIs with Golang that scale gracefully
            Subtitle: Explore the necessary concepts of REST API development by building few real world services from scratch.
            URL: https://api.douban.com/v2/book/30154905
            Isbn10: 1788294289
            Isbn13: 9781788294287
            Price: USD 44.99

            Title: Go Standard Library Cookbook: Over 120 specific ways to make full use of the standard library components in Golang
            Subtitle: Implement solutions by leveraging the power of the GO standard library and reducing dependency on external crates
            URL: https://api.douban.com/v2/book/30179004
            Isbn10: 1788475275
            Isbn13: 9781788475273
            Price: USD 49.99

            Title: Go语言编程
            Subtitle:
            URL: https://api.douban.com/v2/book/11577300
            Isbn10: 7115290369
            Isbn13: 9787115290366
            Price: 49.00元

            Title: The Go Programming Language
            Subtitle:
            URL: https://api.douban.com/v2/book/26337545
            Isbn10: 0134190440
            Isbn13: 9780134190440
            Price: USD 39.99

            Title: Go Web编程
            Subtitle:
            URL: https://api.douban.com/v2/book/24316255
            Isbn10: 7121200910
            Isbn13: 9787121200915
            Price: 65.00元

            Title: Go语言学习笔记
            Subtitle:
            URL: https://api.douban.com/v2/book/26832468
            Isbn10: 7121291606
            Isbn13: 9787121291609
            Price: 89

            Title: Go 语言程序设计
            Subtitle:
            URL: https://api.douban.com/v2/book/24869910
            Isbn10: 7115317909
            Isbn13: 9787115317902
            Price: CNY 69.00

            Title: Go程序设计语言
            Subtitle:
            URL: https://api.douban.com/v2/book/27044219
            Isbn10: 7111558421
            Isbn13: 9787111558422
            Price: 79

            Title: Go并发编程实战
            Subtitle: Go并发编程实战
            URL: https://api.douban.com/v2/book/26244729
            Isbn10: 7115373981
            Isbn13: 9787115373984
            Price: 89元

            Title: Go in Action
            Subtitle:
            URL: https://api.douban.com/v2/book/25858023
            Isbn10: 1617291781
            Isbn13: 9781617291784
            Price: USD 39.99

            Title: Docker源码分析
            Subtitle:
            URL: https://api.douban.com/v2/book/26581184
            Isbn10: 7111510720
            Isbn13: 9787111510727
            Price: 59.00

            Title: Go语言·云动力
            Subtitle:
            URL: https://api.douban.com/v2/book/10770080
            Isbn10: 7115283079
            Isbn13: 9787115283078
            Price: 39.00元

            Title: Go语言实战
            Subtitle:
            URL: https://api.douban.com/v2/book/27015617
            Isbn10: 7115445354
            Isbn13: 9787115445353
            Price: CNY 59.00

            Title: An Introduction to Programming in Go
            Subtitle:
            URL: https://api.douban.com/v2/book/19897704
            Isbn10: 1478355824
            Isbn13: 9781478355823
            Price: USD 10.00

            Title: Go程序设计语言(英文版)
            Subtitle:
            URL: https://api.douban.com/v2/book/26859123
            Isbn10: 7111526287
            Isbn13: 9787111526285
            Price: CNY 79.00

            Title: Go并发编程实战（第2版）
            Subtitle:
            URL: https://api.douban.com/v2/book/27016236
            Isbn10: 7115452512
            Isbn13: 9787115452511
            Price: 79.00元

            Title: The Go Programming Language Phrasebook
            Subtitle:
            URL: https://api.douban.com/v2/book/7952516
            Isbn10: 0321817141
            Isbn13: 9780321817143
            Price: USD 29.99

            Title: Microservices in Go
            Subtitle: Use Go to Build Scalable Backends
            URL: https://api.douban.com/v2/book/26650145
            Isbn10: 149194255X
            Isbn13: 9781491942550
            Price: USD 49.99

            Title: Go Web Programming
            Subtitle:
            URL: https://api.douban.com/v2/book/26340005
            Isbn10: 1617292567
            Isbn13: 9781617292569
            Price: USD 44.99




```

4. YiYi.exe story id-list

```text
You Get One List From One App:

     Response: 0,
     Data:     ["4727","4722","4703","4721","4720","4710","4719","4718","4715","4700"]

```

5. YiYi.exe story one-list weather 4727

```text
Weather Info for One App:
Informatations:
    CityName:      地球,
    Date:          2018-03-27,
    Temperature:   -275,
    Humidity:      120,
    Climate:       对流层,
    WindDirection: 一阵妖风,
    Hurricane:     36级

```

6. YiYi.exe story one-list menu 4727

```text

Show One List Menu form One App:
    Vol: 1998

    ContentType: 1
    ContentID:   3168
    Title:       炉斧酒吧之夜
    Tag:         7 ONE STORY

    ContentType: 2
    ContentID:   585
    Title:       酒神的玫瑰 · 第二章 · 骗子去死吧
    Tag:

    ContentType: 3
    ContentID:   2050
    Title:       心理成熟的人是如何化解悲伤的？
    Tag:

    ContentType: 1
    ContentID:   3164
    Title:       今天天气真好，我们分手吧
    Tag:

    ContentType: 4
    ContentID:   2593
    Title:       被世界遗弃不可怕，喜欢你有时还可怕
    Tag:

    ContentType: 5
    ContentID:   1363
    Title:       你不面对现实，现实就会面对你
    Tag:


```
7. YiYi.exe story detail info 1363

```text
                                                                                                                                          
Show ContentInfo from one App:                                                                                                            
    ID: 1363                                                                                                                              
    Title: 我的傻瓜表叔                                                                                                                         
    Auth: 青年作家，现任光线影业副总裁。已出版《谁的青春不迷茫》《你的孤独，虽败犹荣》，新书《向着光亮那方》正在预售中。                                                                         
    AuthorIntroduce: （责任编辑：金子棋 ）                                                                                                          
    MakeTime: 2016-03-30 23:00:00                                                                                                         
    LastUpdateDate: 2016-03-30 17:01:53                                                                                                   
    WebURL: http://m.wufazhuce.com/article/1363                                                                                           
    GuideWord: 俗语说，傻人有傻福。 一个人能获得多少，上天早已注定。争也无益。 机关算尽太聪明，聪明反被聪明误。 傻，有时是一种福气。                                                               
    Audio: http://music.wufazhuce.com/lh2vtqxvv89co7WNJjF4LmTQnB9x                                                                        
    Anchor:                                                                                                                               
    EditorEmail:                                                                                                                          
    AuthorInfo:                                                                                                                           
       UserID: 4814672                                                                                                                    
       UserName: 刘同                                                                                                                       
       Desc: 青年作家，现任光线影业副总裁。已出版《谁的青春不迷茫》《你的孤独，虽败犹荣》《向着光亮那方》。                                                                              
       WbName: @刘同                                                                                                                        
       Summary: 青年作家，现任光线影业副总裁。                                                                                                           
       FansTotal: 3562                                                                                                                    
                                                                                                                                          
```

8. YiYi.exe story detail content 1363

```text

福田比我小两岁，却是我的远房表叔。小时候，每次回老家过年，长辈们总是让我管福田叫表叔，我很不好意思，因为不能理解为什么要对一个比自己小两岁的人叫表叔。我不懂家长们之间的辈分——据说是因为福田的爸爸按辈分是我的爷爷，所以自然而然我要叫福田表
叔。 福田表叔总是乐呵呵的，小时候我们每年只在过年的时候才能见一次，但每一次听说我们要回老家过年，福田表叔都会早早地在村口的山坡下等着我们，然后远远地一看到我们就开始乐呵呵地笑。不熟的人，初次见面总是感到陌生，即使一年见一次也需要熟络的过
程，但福田表叔好像完全没有这样的障碍，帮我背书包，带我在村里到处逛，田埂上有一只狗，我害怕不敢动，福田表叔就会冲上去一脚把狗踢到田里去，然后笑嘻嘻地对我说：“狗有什么好怕的。” 读初中之前，我一直觉得这个比我小的表叔很好打交道，直到初一那
年春节。福田表叔接上我们，扛着大包小包在前面走，我跟在后面一直盯着他，我觉得福田表叔走路好像很奇怪，每一步都一拐一拐的，总在快要失去平衡的时候才迈开另外一步。在印象中，他好像一直都是这么走路，只是那时我才意识到奇怪。 我问爸爸：“福田表叔
走路是不是歪了？”爸爸说：“他走路一直是这样。”我觉得我爸没有理解我的意思，又追问了一句：“为什么福田表叔走路是歪的？”爸爸有些尴尬地笑了笑，大概想了一种最恰当的方式来回答我：“因为你小爷爷和小奶奶是堂兄妹，所以福田生下来和其他人不太一样
。” 我默念了几遍这句话，才理解，原来福田表叔是近亲结婚生育的小孩。因为近亲的影响，所以他的大脑发育比一般小孩慢，所以走路总是有些踉跄，所以总是对人抱以热情和信任，对外界没有防备…… 那时的我正在读初中，同学们说一个人傻就会用“近亲结婚”
来攻击对方，乐此不疲。我怎么也没有想到，我的亲戚，一个每年都见上一次的福田表叔，居然是近亲结婚生育的小孩。我根本无法理解为什么小爷爷要和小奶奶结婚，也没法判断福田未来的人生究竟会怎样，我只是突然产生一种害怕，以及一种羞耻感，我的表叔居然是
个近亲结婚生育的傻子！ 那个春节，无论福田表叔怎么跟我聊天，我都不敢搭理。他说带我去看刚出生的狗崽，我很嫌弃地说不去。他问我要不要去挖茨菰，我很冷淡地说不要。他给我看各种昆虫标本，我也没兴趣。福田表叔很难过，吃饭的时候就问大家：“为什么同
同不理我？”福田表叔说完那句话，大家都愣住了。小爷爷立刻笑着说：“没有没有，他哪有不理你，是吧？”我不知道怎么回答这个问题，一下愣在那儿。我从来没有想过福田表叔会当着所有人的面说出为什么我不理他这种话。换做是我，我肯定把被排挤的尴尬隐藏起
来。什么样的人才会不掩饰自己的情感呢？这个问题困扰了我很多年。不显露自己的情绪，是不想被人瞧不起，也不想给人造成麻烦。也是在那一刻，我告诉自己：这辈子，绝对！再也不叫福田表叔了！他真是一个令人讨厌的家伙啊。爸爸出来救场说：“福田表叔不开心
了，你一会儿吃完饭买一些摔炮，跟福田表叔一起玩。”随即掏出十块钱给福田，福田立刻就忘记了难过。 吃完饭，福田等着我一起去买摔炮。我实在不想和他待在一起，害怕自己也变傻。突然，我想看看福田究竟有多傻。我从兜里拿出了五张一块钱的人民币，对福田
说：“哎，你看你有一张钱，我有五张钱，我们交换吧？这样你还多了四张。”福田看看我手上的五块钱，又看看自己手上的十块钱，想了一会儿，小心翼翼地问我：“那你不是会少一些吗？”我说：“没事，你不是要买摔炮吗？钱多一点儿比较好买。”然后福田点点头
，很开心地把十块钱给我，收下了我那五张一块钱的人民币。 福田还有个妹妹，比福田乖巧很多。我也偷偷问过爸爸，那个比我小很多岁的姑姑也是近亲结婚生育的吗？爸爸说因为福田的关系，所以小姑姑是小爷爷他们收养的。因为知道了福田的秘密，初中的我自然而
然就跟小姑姑走得近了。我拿着到手的十块钱对她说：“你看，我拿五块钱换的。”小姑姑那时还在读小学五年级，看我居然能赚那么多钱，很是羡慕。我说：“我拿五张一块钱的跟福田换的，他那还有五块钱。”突然我灵机一动，问：“你有没有五毛钱的？你拿六张五
毛钱的，跟他换五张一块钱的，他肯定会同意。”小姑姑那时是个小姑娘，一听可以立刻赚两块钱，开心死了。翻箱倒柜才找到四张五毛钱的，死活凑不齐六张。我又出馊主意：“你去试试，求求他，没准同意了呢？”小姑姑也是演技派，打着是妹妹的旗号，四张五毛钱
换了福田五张一块钱。福田乐呵呵的，我们也乐呵呵的，一点儿愧疚也没有。 第二天是大年初一，一大早孩子们就给大人拜年，就是要拿到那个等了一年的红包。福田也是，跟着磕头，家里给小孩的红包都是五十块一个，福田不一会儿就拿到了好几百块钱，然后就消失
了。福田下午回来的时候，特别开心。他很主动地跟大家说：“我换了好多钱回来！”然后从兜里掏出一大把一块、两块的零钱，五块、十块的都很少。小爷爷不明白什么意思。我一看，觉得福田完蛋了。小爷爷问福田：“你这些钱是从哪里来的？”福田乐呵呵地说，是
跟其他村里的人一家一家换的。“那你拿什么换的？”“红包啊。”小爷爷把福田兜里的钱全部掏出来，数了数，只有六十多块钱。而他的红包里有十几张五十元钞票，五六百块钱。那时一百块钱不是一个小数目，是福田一年读书的学费。小爷爷问他具体是跟哪些人换的
，福田说不上来。小爷爷气得给了他一个重重的耳光，那一巴掌就好像扇在了我的脸上，但血却是从福田的鼻子里流下来的。福田被一个耳光扇蒙了，越着急越语无伦次，只能呜呜地哭。小爷爷问福田：“是谁让你这么做的？”福田流着鼻血捂着脸哭着说：“我换回更多
的钱了啊，为什么要打我？”我和小姑姑大气不敢喘，更不敢主动说是我们把福田给害了。爸爸扯住小爷爷，小爷爷也很无奈，谁让他生了一个这样的儿子呢？本以为福田的鼻血流一会儿就能止住了，谁知道一直流一直流，怎么止也止不住。大家慌了神，赶紧联系车把福
田送到县城的医院里。我躲在爸爸后面，听他跟不同的人打电话，什么都不敢说。那一天，我才知道小爷爷生福田之前还有一个大儿子，同样因为近亲，大儿子天生带着血友病，一旦流血，就很难止住。他六七岁时，不小心摔倒，最后因大量出血死掉了。后来才生了福田
。血友病最怕的就是出血，福田那一次流鼻血就相当于去了一趟鬼门关，折腾了两天才止住。他从县城医院回来之后，我很想主动打招呼，却又怕大人看见。我特别不好意思，但又不知道该怎么道歉。等到福田一个人又准备出去逛时，我趁着没人，追上福田：“福田……
呃……福田表叔，这个给你。”我拿出了自己的二百块钱。福田不肯要，还问我为什么要给他。我解释他也听不懂。我想着他因为上了我的当，损失了好几百块钱，还因此挨了一个大巴掌，差点活不过来，便越发前言不搭后语。实在没办法，只好向小姑姑求助。小姑姑哭
丧着脸给福田道歉，福田立刻就乐呵呵的了。小姑姑比我了解福田，她让福田拿出他的存钱罐，说我要帮他存钱。福田这时才明白，立刻从房间里拿出存钱罐，让我把钱放进去。存钱罐里都是各种零钱，我问福田存那么多钱干吗。福田不好意思地笑了，挠挠头说：“以后
娶媳妇用。”“这么早就存钱了？”福田说：“钱越多，娶的媳妇越好。”我也不知道谁跟他说的这个道理，当时听起来好荒谬，现在想起来，对福田来说，好像也并非离谱。总之，经过那一次，我深刻反省了自己，又发了一次誓——小时候常常喜欢发誓，把它看得很重
，每一次都信誓旦旦的。我发誓的内容是：再也不欺负福田了。虽然已经发誓不叫他表叔，但那并不会影响我和他的关系。不叫表叔，更像朋友。那次走了之后，我听说隔壁村子的年轻人围着福田准备狂揍他一顿，幸好被亲戚看见了，不然一出血肯定完蛋。我问为什么他
们要打福田，爸爸说因为他们又拿着零钱去逗福田换整钞，福田一生气就冲上去打人家。 因为中考，过年回老家也少了。但那几年，我总是想起福田，想起我干的傻事，如果福田真是因为我的错误而离开，我可能一辈子都有阴影吧。那几年，有好看的动画片我都买两套
，我一套，他一套。好看的漫画书就算看完也不借给别人，打算过年的时候带给他看。爸爸给我买了一个相机和两个胶卷，拍完一个之后，我留着另一个。福田很喜欢拍照，他唯一照相的地点就是县城里的照相馆。他看电视的时候曾说：“我也要拍一张好看的照片，放在
钱包里。”爸爸问：“你怎么突然什么事情都想着福田了啊？”我很尴尬，硬着头皮说：“反正我也不喜欢，干吗浪费？” “反正我也不喜欢，我想着不能浪费，所以就给你带来了。”又一年过年，我对兴高采烈的福田说。我也不知道自己为什么不能直接和他分享，也
许是怕被别人看出内疚。“走，我给你照相去。”我举着相机，朝福田晃了晃。福田可开心了，像个小孩一样手舞足蹈，说：“照相去喽，照相去喽，你等一下，我上去换件衣服。”那时我很羡慕福田，因为他这样的人随时都很开心，不用去思考成长中必须面对的那些艰
难，比如同学之间的关系，比如对自己未来的规划，比如应付父母对自己的期望。这些福田都不会想，似乎，他只需要好好地活着，对谁都可以不负责任。我很羡慕他，单纯地羡慕他，羡慕他可以成绩不好，但大家能理解；羡慕他可以犯错误，大家也能当作什么都没发生
。 “我们去水塘南边拍照好不好，那里有水。”福田往前面带路，我和小姑姑跟在后面。一卷胶卷只能拍30张照片，福田什么都想拍，又想拍田埂上的花，又想拍路边懒洋洋的狗，又边走边回过头来比一个手势，希望我能给他拍。这时，小姑姑看到路边有一棵长得很挺
拔的树，希望能靠在上面拍一张。我看了看四周，打算爬到旁边的坡上给她俯拍。没有想到山坡上的泥又滑又松，我一脚踩上去便打了滑，整个人重重地仰面摔倒，从坡上滚了下来，手里的相机整个飞了出去，掉进了积着水的烂泥潭里。福田赶紧跑过来把我扶起，然后很
着急地站在泥田旁边，不知道如何是好。我看着整个相机都淹没在了烂泥潭里，觉得相机肯定毁了，就对福田和小姑姑说：“算了，相机肯定坏了，捡上来也没用了，而且那么远，怎么捡得着。”小姑姑觉得是因为她，相机才坏掉，默默地流眼泪。福田则很焦躁地走来走
去，很生气的样子。 回到家之后，爸爸问我怎么了，我说因为自己不小心，把相机掉进了泥潭里。说着的时候，没有人注意到福田在杂房里换上了去田里挖藕的连体防水衣。因为是冬天，衣服又多又厚，穿不进防水衣，福田就脱了外面的大衣和毛衣，只剩了一件秋衣，
然后套着防水衣就出去了。等他回来的时候，浑身冻得发抖，右手拎着沾满了泥巴的相机带子，还一路滴着水，大家都骂他傻，我也骂他傻，小奶奶赶紧烧上热水让他好好洗个澡暖和暖和。我又生气又心疼，真是傻到家了啊。 大家围着火炉烤火，一言一语在聊福田。一
个亲戚说：小时候，小爷爷去县城卖东西，让福田在一个地方待着等他，小爷爷卖完东西后就忘记了，自己径直回了家。回到家才想起福田没有回来，这时已经凌晨两三点了。等再回到县城，去到那个小爷爷让福田等的地方，发现福田蜷缩在地上靠着墙睡着了……还有亲
戚说：有段时间福田每天都躺在床上，怎么也不肯起床，一开始以为他生病了，后来发现他没病，拼死拼活把他拽起来后，才发现福田的被子里藏着十几个鸡蛋。因为不知道听谁说的鸡蛋可以孵出小鸡，就偷了好多鸡蛋放在自己被子里孵……所有人都哈哈大笑，我也跟着
哈哈大笑。 福田对我来说就是一个奇怪的存在，一开始是长辈表叔，后来是害怕被传染低智商的敌人，再后来是想好好爱护的朋友，现在又变成了全家人的笑料。因为他和我们都不一样，也许并不需要我们的理解吧。 福田洗完澡，站在厨房的门口很神秘地叫我过去。
他笑嘻嘻地看着我，一直看着我，当我问他“干吗”的时候，他突然把手从身后伸出来，掏出一个东西吓了我一跳。相机已经被他洗得非常干净了，机身上一点儿泥巴都没有，相机绳也洗得很干净，不说的话，没有人能看出相机掉进过泥潭。唯一的破绽是，机身一直朝外
滴着水。“嘿嘿，你看，我洗得干不干净？你看，是不是和新的一样？”福田扬起头看着我，一副很得意的样子。我看着那个相机，又看着他的表情，突然觉得鼻子一酸，眼泪就流出来了。福田一愣，问：“你怎么哭了？”我擦擦眼泪，说：“没有，就是隐形眼镜太涩了
。”我接过相机对他说：“谢谢你啊，真的跟新的一样。”福田很开心地说：“我认真洗了几遍，还用刷子刷了。”我问他：“你还要不要照相？”他说：“要！”然后我就陪着他拍照，再也没有提过只能拍30张照片的事，只要他摆好姿势，我就认真拍。我看着远远的福
田，觉得那一刻我比任何人都懂他。他根本不是一个没有智商、不善思考的人。相反，他是一个真真正正在自己的世界里用所有的热情去爱别人的人。他听信任何人的话，从未怀疑。他相信只要自己努力就能孵出小鸡，就能给予它们新的生命。他把相机刷干净，让它恢复
成以前的样子，他相信我一定会开心。他比我见过的大多数人，更希望别人好。曾经的我以为，他只要自己没心没肺地活着，对谁都可以不负责。其实他每天活得比我们更认真，因为他想对每一个人都负责。 后来，我、福田还有小姑姑都长大了，家里免不了会聊到福田
结婚的问题。我问福田：“你的钱能不能借给我啊，我以后挣更多的钱还给你。”福田摇摇头，断然拒绝，他说：“我的钱是娶老婆用的，谁都不给。”我听后哈哈大笑：“好的好的，祝你早日娶到一个好老婆。”坦白讲，要给福田介绍对象不是一件容易的事，正常姑娘
都不会嫁给他，愿意见见面的姑娘，身体又总是有一些缺陷。这件事越拖越久，最后村子里有人给福田介绍了隔壁村子的女孩，对方父母也是近亲结婚。媒人说：“虽然傻是傻了一点儿，但做事跟正常人一样。”那天是大年初二，福田起得很早，穿得很帅，意气风发地跟
着小爷爷他们出去了。还没到中午，福田就回来了，怒气冲冲地，一进院子就大骂：“为什么要给我介绍一个傻子?!我不要娶一个傻子当老婆！”然后在院子里拿起各种东西乱摔，搞得鸡飞狗跳。换作以前，如果看见福田说“为什么我要娶一个傻子”，我一定会笑出来。
那一天，我远远地看着福田，虽然他的智商只有十几岁。但他毕竟25岁了，像个真正的男人了。摔着摔着，他自己就蹲在院子里哭，我过去劝他，他呜呜地一边哭一边说：“我一直被人说成傻子，我不想再娶一个傻子老婆，我不想再生一个儿子也是傻子……” 福田不傻
，只是有些事情凭他自己一个人想不明白。福田不傻，他知道自己哪些地方做不到，使了劲也不行，能力不够。而那一次，是我最后见到福田的样子。 后来我读大学了，很少回郴州，更少回老家。有一天老家来人了，我突然想起福田，问福田结婚了没？老家人摇摇头说
：“福田啊，前两年走了。”走了？什么意思？我愣了一下。 老家的亲戚说他在楼顶帮晒黄豆，不小心踩空了，从三楼摔下来，血流不止。后来去医院抢救也没用，伤口太大，血根本止不住。然后就走了。我问爸爸是否知道，爸爸说：“知道，那时你在外地读书，就没
有告诉你。”后来听说，福田走的那天，小姑姑去看他，他一直说着“存钱罐……存钱罐……”，因为离家太远了，怎么劝都没用，然后福田就一直吵一直吵。等福田走了之后，家里人去收福田的东西，发现他房间里的一个存钱罐变成了两个，一个写了自己的名字，一个
写了小姑姑的名字。那时大家才反应过来，写小姑姑名字的存钱罐是福田把自己存了27年的钱分了一大半给小姑姑结婚用的。那天，所有人，尤其是小姑姑，都哭成了傻子。 回忆起福田的时候，爸爸说，有一年，山路泥泞，爸爸他们开的车进不来，只能把车停在进山的
路口。晚上打牌的时候提了一句“担心车放在外面不安全”，当时谁都没当回事。第二天一大早，大家看见福田抱着一大堆被子和尼龙布回来，便问他去哪了，福田说：“昨天哥哥怕车停在外面有事，我就在车的旁边睡了一晚，好冷哦。” 现在想起来，其实，福田一点
儿都不傻。他只是太好了。好傻，好傻。 <strong>后来</strong> 事情已经过去了好些年，每次回老家提到福田的时候，爸爸总会说起一些关于他的新故事。每次听，都觉得很想哭，觉得过去没有珍惜福田的好。爸爸说福田离开前的最后那几年，身上总会带一个本子
，每次听到大家说什么成语，就记下来，不明白意思就回去查字典，然后自己也用成语说话。福田说用成语说话显得很像大人。你看他，做的每一件事，说的每一句话，都没有任何掩饰，让人一眼就能看到他的心思，听懂他想说的话。他是一个如此透明的人，和他交往起
来丝毫不费力。有时我会觉得做一个傻子多好，不纠结、不计较，自己活得快乐，还被那么多人喜欢。而福田为了让自己变得更好，那种没皮没脸、奋不顾身的态度与决心，每每想起，也让我很羡慕。本文选自作者新书《向着光亮那方》。


```

9. YiYi.exe movie now
```text

Show Movie Info from DouBan Api.
    MovieCollections In theaters:

        ID: 20435622
        Title: 环太平洋：雷霆再起
        OriginalTitle: Pacific Rim: Uprising
        Average: 5.8
        CollectionCount: 60611
        Stars: 30
        Casts:

              Alt: https://movie.douban.com/celebrity/1339915/
              Name: 约翰·博耶加
              ID: 1339915

              Alt: https://movie.douban.com/celebrity/1000188/
              Name: 斯科特·伊斯特伍德
              ID: 1000188

              Alt: https://movie.douban.com/celebrity/1362560/
              Name: 卡莉·史派妮
              ID: 1362560


        ID: 3445906
        Title: 古墓丽影：源起之战
        OriginalTitle: Tomb Raider
        Average: 6.4
        CollectionCount: 45967
        Stars: 35
        Casts:

              Alt: https://movie.douban.com/celebrity/1233154/
              Name: 艾丽西亚·维坎德
              ID: 1233154

              Alt: https://movie.douban.com/celebrity/1027173/
              Name: 多米尼克·威斯特
              ID: 1027173

              Alt: https://movie.douban.com/celebrity/1098551/
              Name: 沃尔顿·戈金斯
              ID: 1098551


        ID: 30152451
        Title: 厉害了，我的国
        OriginalTitle: 厉害了，我的国
        Average: 0
        CollectionCount: 124
        Stars: 00
        Casts:


        ID: 26861685
        Title: 红海行动
        OriginalTitle: 红海行动
        Average: 8.5
        CollectionCount: 352563
        Stars: 45
        Casts:

              Alt: https://movie.douban.com/celebrity/1274761/
              Name: 张译
              ID: 1274761

              Alt: https://movie.douban.com/celebrity/1354442/
              Name: 黄景瑜
              ID: 1354442

              Alt: https://movie.douban.com/celebrity/1272245/
              Name: 海清
              ID: 1272245


        ID: 26752852
        Title: 水形物语
        OriginalTitle: The Shape of Water
        Average: 7.3
        CollectionCount: 179376
        Stars: 40
        Casts:

              Alt: https://movie.douban.com/celebrity/1044915/
              Name: 莎莉·霍金斯
              ID: 1044915

              Alt: https://movie.douban.com/celebrity/1019031/
              Name: 道格·琼斯
              ID: 1019031

              Alt: https://movie.douban.com/celebrity/1144415/
              Name: 迈克尔·珊农
              ID: 1144415


        ID: 6390825
        Title: 黑豹
        OriginalTitle: Black Panther
        Average: 6.7
        CollectionCount: 131082
        Stars: 35
        Casts:

              Alt: https://movie.douban.com/celebrity/1327680/
              Name: 查德维克·博斯曼
              ID: 1327680

              Alt: https://movie.douban.com/celebrity/1334862/
              Name: 露皮塔·尼永奥
              ID: 1334862

              Alt: https://movie.douban.com/celebrity/1107320/
              Name: 迈克尔·B·乔丹
              ID: 1107320


        ID: 26393561
        Title: 小萝莉的猴神大叔
        OriginalTitle: Bajrangi Bhaijaan
        Average: 8.6
        CollectionCount: 140260
        Stars: 45
        Casts:

              Alt: https://movie.douban.com/celebrity/1017831/
              Name: 萨尔曼·汗
              ID: 1017831

              Alt: https://movie.douban.com/celebrity/1350825/
              Name: 哈莎莉·马洛特拉
              ID: 1350825

              Alt: https://movie.douban.com/celebrity/1049635/
              Name: 卡琳娜·卡普尔
              ID: 1049635


        ID: 26846031
        Title: 萌犬好声音
        OriginalTitle: Pup Star
        Average: 5.8
        CollectionCount: 448
        Stars: 30
        Casts:

              Alt: https://movie.douban.com/celebrity/1326635/
              Name: 麦肯泽·摩斯
              ID: 1326635

              Alt: https://movie.douban.com/celebrity/1268703/
              Name: 凯特林·马希尔
              ID: 1268703

              Alt: https://movie.douban.com/celebrity/1148334/
              Name: 杰德·瑞斯
              ID: 1148334


        ID: 26649604
        Title: 比得兔
        OriginalTitle: Peter Rabbit
        Average: 7.4
        CollectionCount: 23320
        Stars: 40
        Casts:

              Alt: https://movie.douban.com/celebrity/1017966/
              Name: 詹姆斯·柯登
              ID: 1017966

              Alt: https://movie.douban.com/celebrity/1313116/
              Name: 多姆纳尔·格里森
              ID: 1313116

              Alt: https://movie.douban.com/celebrity/1022562/
              Name: 萝丝·拜恩
              ID: 1022562


        ID: 26698897
        Title: 唐人街探案2
        OriginalTitle: 唐人街探案2
        Average: 7
        CollectionCount: 303656
        Stars: 35
        Casts:

              Alt: https://movie.douban.com/celebrity/1274388/
              Name: 王宝强
              ID: 1274388

              Alt: https://movie.douban.com/celebrity/1336305/
              Name: 刘昊然
              ID: 1336305

              Alt: https://movie.douban.com/celebrity/1274979/
              Name: 肖央
              ID: 1274979


        ID: 26661194
        Title: 脱皮爸爸
        OriginalTitle: 脱皮爸爸
        Average: 5.7
        CollectionCount: 1322
        Stars: 30
        Casts:

              Alt: https://movie.douban.com/celebrity/1037098/
              Name: 吴镇宇
              ID: 1037098

              Alt: https://movie.douban.com/celebrity/1027577/
              Name: 古天乐
              ID: 1027577

              Alt: https://movie.douban.com/celebrity/1339442/
              Name: 春夏
              ID: 1339442


        ID: 26611804
        Title: 三块广告牌
        OriginalTitle: Three Billboards Outside Ebbing, Missouri
        Average: 8.7
        CollectionCount: 239338
        Stars: 45
        Casts:

              Alt: https://movie.douban.com/celebrity/1010548/
              Name: 弗兰西斯·麦克多蒙德
              ID: 1010548

              Alt: https://movie.douban.com/celebrity/1053560/
              Name: 伍迪·哈里森
              ID: 1053560

              Alt: https://movie.douban.com/celebrity/1047972/
              Name: 山姆·洛克威尔
              ID: 1047972


        ID: 27042405
        Title: 大坏狐狸的故事
        OriginalTitle: Le Grand Méchant Renard et autres contes...
        Average: 8.2
        CollectionCount: 15746
        Stars: 40
        Casts:

              Alt: https://movie.douban.com/celebrity/1385549/
              Name: 纪尧姆·达尔诺
              ID: 1385549

              Alt: https://movie.douban.com/celebrity/1385552/
              Name: 席琳·荣特
              ID: 1385552

              Alt: https://movie.douban.com/celebrity/1385555/
              Name: 达米安·维特卡
              ID: 1385555


        ID: 26972275
        Title: 恋爱回旋
        OriginalTitle: ミックス。
        Average: 7.5
        CollectionCount: 26418
        Stars: 40
        Casts:

              Alt: https://movie.douban.com/celebrity/1018562/
              Name: 新垣结衣
              ID: 1018562

              Alt: https://movie.douban.com/celebrity/1037363/
              Name: 瑛太
              ID: 1037363

              Alt: https://movie.douban.com/celebrity/1098533/
              Name: 广末凉子
              ID: 1098533


        ID: 26575103
        Title: 捉妖记2
        OriginalTitle: 捉妖记2
        Average: 5.2
        CollectionCount: 107816
        Stars: 25
        Casts:

              Alt: https://movie.douban.com/celebrity/1115918/
              Name: 梁朝伟
              ID: 1115918

              Alt: https://movie.douban.com/celebrity/1275542/
              Name: 白百何
              ID: 1275542

              Alt: https://movie.douban.com/celebrity/1274628/
              Name: 井柏然
              ID: 1274628


        ID: 27076492
        Title: 虎皮萌企鹅
        OriginalTitle: Les As de la Jungle
        Average: 6
        CollectionCount: 1315
        Stars: 30
        Casts:

              Alt: https://movie.douban.com/celebrity/1389534/
              Name: 菲利普·波佐
              ID: 1389534

              Alt: https://movie.douban.com/celebrity/1389535/
              Name: 劳伦·默度
              ID: 1389535

              Alt: https://movie.douban.com/celebrity/1389536/
              Name: 帕斯卡尔·卡萨诺瓦
              ID: 1389536


        ID: 27176717
        Title: 熊出没·变形记
        OriginalTitle: 熊出没·变形记
        Average: 6.8
        CollectionCount: 6379
        Stars: 35
        Casts:

              Alt: https://movie.douban.com/celebrity/1336920/
              Name: 张伟
              ID: 1336920

              Alt: https://movie.douban.com/celebrity/1336919/
              Name: 张秉君
              ID: 1336919

              Alt: https://movie.douban.com/celebrity/1336930/
              Name: 谭笑
              ID: 1336930


        ID: 26984234
        Title: 玲珑井
        OriginalTitle: 玲珑井
        Average: 0
        CollectionCount: 177
        Stars: 00
        Casts:

              Alt: https://movie.douban.com/celebrity/1329270/
              Name: 宋睿
              ID: 1329270

              Alt: https://movie.douban.com/celebrity/1341569/
              Name: 罗翔
              ID: 1341569

              Alt: https://movie.douban.com/celebrity/1354039/
              Name: 曾漪莲
              ID: 1354039


        ID: 26603666
        Title: 妈妈咪鸭
        OriginalTitle: 妈妈咪鸭
        Average: 6.8
        CollectionCount: 2061
        Stars: 35
        Casts:

              Alt: https://movie.douban.com/celebrity/1349631/
              Name: 赵路
              ID: 1349631

              Alt: https://movie.douban.com/celebrity/1385413/
              Name: 武皓栋
              ID: 1385413

              Alt: https://movie.douban.com/celebrity/1350671/
              Name: 殷筱瑜
              ID: 1350671


        ID: 26967920
        Title: 遇见你真好
        OriginalTitle: 遇见你真好
        Average: 0
        CollectionCount: 446
        Stars: 00
        Casts:

              Alt: https://movie.douban.com/celebrity/1332932/
              Name: 白客
              ID: 1332932

              Alt: https://movie.douban.com/celebrity/1318954/
              Name: 蓝盈莹
              ID: 1318954

              Alt: https://movie.douban.com/celebrity/1364842/
              Name: 张海宇
              ID: 1364842


```

10. YiYi.exe movie detail 26967920

```text
                                                                                                                                                                                                                                              
Show Get Movie In theaters from DouBan Api:                                                                                                                                                                                                   
    ID: 26967920                                                                                                                                                                                                                              
    Title: 遇见你真好                                                                                                                                                                                                                              
    OriginalTitle: 遇见你真好                                                                                                                                                                                                                      
    Directors: 顾长卫                                                                                                                                                                                                                            
    Average: 0                                                                                                                                                                                                                                
    ReviewsCount: 30                                                                                                                                                                                                                          
    WishCount: 5109                                                                                                                                                                                                                           
    CommentsCount:                                                                                                                                                                                                                            
    RatingsCount: 430                                                                                                                                                                                                                         
    Aka: 飞火流星 | Nice to Meet You                                                                                                                                                                                                              
    Year: 2018                                                                                                                                                                                                                                
    MobileURL: https://movie.douban.com/subject/26967920/mobile                                                                                                                                                                               
    ShareURL: http://m.douban.com/movie/subject/26967920                                                                                                                                                                                      
    Countries: 中国大陆                                                                                                                                                                                                                           
    Genres: 剧情 | 喜剧 | 爱情                                                                                                                                                                                                                      
    CollectCount: 446                                                                                                                                                                                                                         
    Casts:                                                                                                                                                                                                                                    
                                                                                                                                                                                                                                              
        Alt: https://movie.douban.com/celebrity/1332932/                                                                                                                                                                                      
        Name: 白客                                                                                                                                                                                                                              
        ID: 1332932                                                                                                                                                                                                                           
                                                                                                                                                                                                                                              
        Alt: https://movie.douban.com/celebrity/1318954/                                                                                                                                                                                      
        Name: 蓝盈莹                                                                                                                                                                                                                             
        ID: 1318954                                                                                                                                                                                                                           
                                                                                                                                                                                                                                              
        Alt: https://movie.douban.com/celebrity/1364842/                                                                                                                                                                                      
        Name: 张海宇                                                                                                                                                                                                                             
        ID: 1364842                                                                                                                                                                                                                           
                                                                                                                                                                                                                                              
        Alt: https://movie.douban.com/celebrity/1275152/                                                                                                                                                                                      
        Name: 周楚濋                                                                                                                                                                                                                             
        ID: 1275152                                                                                                                                                                                                                           
                                                                                                                                                                                                                                              
    Summary: 当遇见爱情的时候才会发现，它就像万有引力，苹果总归会落地。在被称为“高四”的紫荆复读学校是一所爱情集中营，梦想成为作家的张文生和珊妮在互损中萌生情愫，好友陈奇误认为文生横刀夺爱，与“坏学生”青龙在小酒馆相遇并商议各自的复仇计                                                                                                                 
划。不料陈奇带去的巨型烟花被青龙向酒馆发泄中引发火灾，带出人命。于是，这所学校里的一个又一个神秘搞笑的爱情故事浮出水面，一群热血而欢乐的男女演绎了另类的青春和爱情……                                                                                                                                                           

```


11. YiYi.exe movie future

```text
Show Get Movie In theaters from DouBan Api:
    ID: 26967920
    Title: 遇见你真好
    OriginalTitle: 遇见你真好
    Directors: 顾长卫
    Average: 0
    ReviewsCount:
    WishCount:
    CommentsCount:
    RatingsCount:
    Aka:
    Year: 2018
    MobileURL:
    ShareURL:
    Countries:
    Genres: 剧情 | 喜剧 | 爱情
    CollectCount: 446
    Casts:

        Alt: https://movie.douban.com/celebrity/1332932/
        Name: 白客
        ID: 1332932

        Alt: https://movie.douban.com/celebrity/1318954/
        Name: 蓝盈莹
        ID: 1318954

        Alt: https://movie.douban.com/celebrity/1364842/
        Name: 张海宇
        ID: 1364842

    Summary:

Show Get Movie In theaters from DouBan Api:
    ID: 27072634
    Title: 河间圣手
    OriginalTitle: 河间圣手
    Directors: 孙爱国
    Average: 0
    ReviewsCount:
    WishCount:
    CommentsCount:
    RatingsCount:
    Aka:
    Year: 2018
    MobileURL:
    ShareURL:
    Countries:
    Genres: 剧情 | 传记 | 历史
    CollectCount: 12
    Casts:

        Alt: https://movie.douban.com/celebrity/1322769/
        Name: 杨子骅
        ID: 1322769

        Alt: https://movie.douban.com/celebrity/1389385/
        Name: 孟晓涵
        ID: 1389385

        Alt: https://movie.douban.com/celebrity/1389386/
        Name: 杨千里
        ID: 1389386

    Summary:



```

12. YiYi.exe movie top 2

```text
Show Get Movie In theaters from DouBan Api:
    ID: 1292063
    Title: 美丽人生
    OriginalTitle: La vita è bella
    Directors: 罗伯托·贝尼尼
    Average: 9.5
    ReviewsCount:
    WishCount:
    CommentsCount:
    RatingsCount:
    Aka:
    Year: 1997
    MobileURL:
    ShareURL:
    Countries:
    Genres: 剧情 | 喜剧 | 爱情
    CollectCount: 586592
    Casts:

        Alt: https://movie.douban.com/celebrity/1041004/
        Name: 罗伯托·贝尼尼
        ID: 1041004

        Alt: https://movie.douban.com/celebrity/1000375/
        Name: 尼可莱塔·布拉斯基
        ID: 1000375

        Alt: https://movie.douban.com/celebrity/1000368/
        Name: 乔治·坎塔里尼
        ID: 1000368

    Summary:


```


