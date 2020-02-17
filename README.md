# kiwi-basket

- /users

アカウント作成

`POST`
```
{
  "username": "gleam",
  "hashed_password": "abcdefg"
}
```

アカウント削除

`DELETE`
```
{
  "username": "gleam",
  "hashed_password": "abcdefg"
}
```


- /tokens

Token生成

`GET`
```
{
  "username": "gleam",
  "hashed_password": "abcdefg"
}
```
```
{
  "token": "1234567890"
}
```

Token削除

`DELETE`
```
{  
  "username": "gleam",
  "token": "1234567890"
}
```

- /timetables

時間割の作成

`POST`
```
{
  "username": "gleam",
  "token": "1234567890".
  "timetable": {
    "mon": {
      "1": {
        "subject": "A",
        "room": "100"
      },
      // 空きコマ
      "2": null,
      ...
      "5": {...}
    },
    "tue": {...},
    ...
    "fri": {...}
  }
}
```

時間割の取得

`GET`
```
{  
  "username": "gleam",
  "token": "1234567890"
}
```
```
{
  "timetable": {
    "mon": {
      "1": {
        "subject": "A",
        "room": "100"
      },
      // 空きコマ
      "2": null,
      ...
      "5": {...}
    },
    "tue": {...},
    ...
    "fri": {...}
  }
}
```

- /tasks

課題の作成

`POST`
```
{
  "username": "gleam",
  "token": "1234567890",
  "task": {
    "date": "2020-01-01",
    "title": "task"
  }
}
```

課題の取得

`GET`
```
{
  "username": "gleam",
  "token": "1234567890"
}
```
```
{
  "tasks": [
    {
      "date": "2020-01-01",
      "title": "task1"
    },
    {
      "date": "2020-01-02",
      "title": "task2"
    },
    ...
  ]
}
```
