# YiYi
A cli command tool for reader from DouBan and One App API.


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



1. 