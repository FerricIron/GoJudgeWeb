# GoJudgeWeb数据库

## User

| 字段名      | 类型     | 是否可空 | 约束        | 备注           |
| ----------- | -------- | -------- | ----------- | -------------- |
| uid         | int      | NO       | primary_key | AUTO_INCREMENT |
| username    | varchar  | NO       | unique      | 用户名         |
| nickname    | varchar  | NO       |             | 昵称           |
| password    | char(32) | NO       |             | 32位加盐MD5    |
| description | varchar  | YES      |             | 用户描述       |
| sid         | int      | YES      |             | 学校ID         |
| privilege   | int      | NO       |             | 权限号         |
| submitcount | int      | NO       |             | 总计提交数     |
| solved      | int      | NO       |             | 解决问题数     |

## school

| 字段名     | 类型    | 是否可空 | 约束        | 备注           |
| ---------- | ------- | -------- | ----------- | -------------- |
| sid        | int     | NO       | primary_key | AUTO_INCREMENT |
| schoolname | varchar | NO       | unique      |                |
| shortname  | varchar | NO       |             | 简称           |

## problem

| 字段名      | 类型    | 是否可空 | 约束        | 备注           |
| ----------- | ------- | -------- | ----------- | -------------- |
| problemid   | int     | NO       | primary_key | AUTO_INCREMENT |
| problemname | varchar | NO       |             |                |
| author      | varchar | NO       |             | 创建者         |
| description | varchar | NO       |             | 题面描述       |
| privilege   | int     | NO       |             | 题目权限       |
| property    | int     | NO       |             | 题目类型       |
| submitcount | int     | NO       |             | 总提交数       |
| solved      | int     | NO       |             | ac次数         |
| timelimit   | int     | NO       |             | 时间限制(ms)   |
| memorylimit | int     | NO       |             | 内存限制(M)    |
| status      | int     | NO       |             | 题目状态       |

## submit

| 字段名    | 类型    | 是否可空 | 约束        | 备注           |
| --------- | ------- | -------- | ----------- | -------------- |
| submitid  | int     | NO       | primary_key | AUTO_INCREMENT |
| uid       | int     | NO       |             | 提交用户       |
| time      | int     | NO       |             | 提交时间       |
| language  | int     | NO       |             | 语言           |
| scid      | int     | NO       |             | sourcecodeid   |
| contestid | int     | YES      |             |                |
| problemid | int     | NO       |             |                |
| status    | int     | NO       |             | 当前状态       |
| timecost  | int     | YES      |             | 时间花费       |
| info      | varchar | YES      |             | 编译或运行信息 |

### status

| 状态码 | 含义           | 备注                         |
| ------ | -------------- | ---------------------------- |
| 0      | 收到提交       |                              |
| 1      | 提交至评测队列 |                              |
| 2      | 正在评测       | 此时有收到部分评测机返回数据 |
| 3      | AC             | 评测结束                     |
| 4      | WA             |                              |
| 5      | TLE            |                              |
| 6      | MLE            |                              |
| 7      | OLE            |                              |
| 8      | CE             | 存在info                     |
| 9      | RuntimeError   | 存在info                     |
| 10     | RepresentError |                              |



## sourcecode

| 字段名   | 类型    | 是否可空 | 约束        | 备注           |
| -------- | ------- | -------- | ----------- | -------------- |
| scid     | int     | NO       | primary_key | AUTO_INCREMENT |
| source   | varchar | NO       |             | 源代码         |
| language | int     | NO       |             | 语言           |

## contest

| 字段名      | 类型    | 是否可空 | 约束        | 备注           |
| ----------- | ------- | -------- | ----------- | -------------- |
| contestid   | int     | NO       | primary_key | AUTO_INCREMENT |
| contestname | varchar | NO       |             |                |
| creator     | varchar | NO       |             | 创建者         |
| description | varchar | YES      |             | 描述           |
| starttime   | int     | NO       |             | 开始时间       |
| endtime     | int     | NO       |             | 结束时间       |
| property    | int     | NO       |             | 竞赛类型       |
| privilege   | int     | NO       |             | 可查看权限     |
| status      | int     | NO       |             | 状态           |

## contestinfo

| 字段名      | 类型 | 是否可空 | 约束 | 备注                   |
| ----------- | ---- | -------- | ---- | ---------------------- |
| contestid   | int  | NO       |      |                        |
| problemid   | int  | NO       |      |                        |
| id          | int  | NO       |      | 这道题目在比赛中的序号 |
| submitcount | int  | NO       |      | 竞赛中提交次数         |
| solved      | int  | NO       |      | 竞赛中ac次数           |
| penalty     | int  | NO       |      | 罚时信息               |

## contestregister

| 字段名      | 类型 | 是否可空 | 约束 | 备注           |
| ----------- | ---- | -------- | ---- | -------------- |
| contestid   | int  | NO       |      |                |
| uid         | int  | NO       |      |                |
| penalty     | int  | NO       |      | 本场比赛罚时   |
| solved      | int  | NO       |      | 本场解决问题数 |
| submitcount | int  | NO       |      | 本场提交数     |
| rank        | int  | NO       |      | 排名           |



