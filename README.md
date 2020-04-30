# kiwi-basket

- /users

アカウント作成

`POST`
```
{
  "username": "gleam",
  "password": "abcdefg"
}
```

アカウント削除

`DELETE`
```
{
  "username": "gleam",
  "password": "abcdefg"
}
```


- /tokens

Token生成

`POST`
```
{
  "username": "gleam",
  "password": "abcdefg"
}
```
```
{
  "token": "1234567890"
}
```

- /timetables

時間割の作成

`POST`
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

時間割の取得

`GET`
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
  "task": {
    "id": "-1",
    "date": "2020-01-01",
    "title": "task"
  }
}
```

課題の削除

`DELETE`
```
{
  "id": "1"
}
```


課題の取得

`GET`
```
{
  "tasks": [
    {
      "id": "1",
      "date": "2020-01-01",
      "title": "task1"
    },
    {
      "id": "2",
      "date": "2020-01-02",
      "title": "task2"
    },
    ...
  ]
}
```
