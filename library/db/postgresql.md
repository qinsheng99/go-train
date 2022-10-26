# sql
```text
CREATE TABLE "public"."boy" (
  "id" int8 NOT NULL DEFAULT nextval('boy_id_seq1'::regclass),
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
  "information" jsonb DEFAULT '{}'::jsonb,
  "arr" int4[],
  CONSTRAINT "boy_pkey" PRIMARY KEY ("id")
)
```

# 一、数组
## 1、常用函数

- array_length(col,int) [表结构](#sql)
``` sql
    SELECT array_length(arr,1) FROM boy //返回每条记录的数组长度
```

- array(int)  [表结构](#sql)
``` sql
    SELECT arr[1] FROM boy  //返回每条记录数组的第一个,索引从1开始
```

- array_to_string(col,str) [表结构](#sql)
``` sql
    SELECT array_to_string(arr, ',') FROM boy //数组以逗号拼接为字符串
```

- array_to_json(col,bool) [表结构](#sql)
``` sql
    SELECT array_to_json(arr, TRUE) FROM boy 
```

- array_position(clo,data) [表结构](#sql)
``` sql
    SELECT array_position(arr,1) FROM boy //返回data在数组第一次出现位置,没有为null
```

- array_positions(clo,data) [表结构](#sql)
``` sql
    SELECT array_positions(arr,1) FROM boy //返回data在数组中出现的位置,没有为null
```

## 2、操作符
- &&(重叠,有相同的元素) [表结构](#sql)
``` sql
    SELECT * FROM boy WHERE '{3}'&&arr // 返回数组中有3的
```

- = <> > >= < <= @>(包含) <@(被包含) [表结构](#sql)
``` sql
    SELECT * FROM boy WHERE '{3}'=arr //返回数组等于3的
```

- ||(连接) arr||arr arr||col col||arr  [表结构](#sql)
``` sql
    SELECT 3||arr FROM boy // 数组和数组连接
```

# 二、jsonb

## 1、常用函数

- jsonb_path_query(col,str) [表结构](#sql)
``` sql
    SELECT jsonb_path_query(information, '$.repo ? (@ starts with "xxx")') FROM boy //repo字段中以xxx开头的
```

## 2、操作符
- -> (-> int获取json数组元素,-> text获取json对象域) [表结构](#sql)
``` sql
    SELECT name,information->'age' FROM boy 
```

- ->>返回文本 [表结构](#sql)
``` sql
    SELECT name,information->>'age' FROM boy 
```

- -> 和 ->> 共用
``` sql
    SELECT name,information->'address'->>'hunan' FROM boy   //{"age": 27, "address": {"hunan": "changsha"}} 返回changsha 
```

- @>
``` sql
    select * from boy where information @> '[{"operator":"OR1"}]'   //jsonb中包含参数的
```