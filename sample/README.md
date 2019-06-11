应用程序的项目结构
-sample
   -data
      data.json         -- 包含一组数据源
   -matchers
      rss.go            -- 搜索rss源的匹配器
   -search
      default.go        -- 搜索数据用的默认匹配器
      feed.go           -- 用于读取json数据文件
      match.go          -- 用于支持不同匹配器的接口
      search.go         -- 执行搜索的主控制逻辑
   main.go              -- 程序的入口