# gorm-mptt
mptt plugin for gorm orm

## TODO

1. ~~move node up and down~~
1. reparent
1. scope
1. add tree configuration
1. level

## DOC

### Database: requirements

- id varchar 36 (uuid v4)
- parent_id varchar 36 (uuid v4) index
- lft int index
- rght int

A simple model example:
```golang 
type Category struct {
	ID          string         `gorm:"primaryKey;type:varchar(36)"`
	ParentID    *string        `gorm:"default:null;index;type:varchar(36)"`
	Name        string         `gorm:"default:null;type:varchar(100)"`
	Lft         int            `gorm:"index"`
	Rght        int
}

```


### Enable Tree

```golang
import mptt "github.com/golgher/gorm-mptt"

# Gorm database connection

dsn := "host=localhost user=root password=1234 dbname=mptt port=5445 sslmode=disable TimeZone=America/Sao_Paulo"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),
})

if err != nil {
    panic("failed to connect database")
}

t := mptt.New(db)
```

### Create node

```golang

Parent := entity.Category{
    Name: "Main",
}

t.CreateNode(&Parent)

Child1 := entity.Category{
    Name:     "Child 1"
    ParentID:    &Parent.ID,
}

t.CreateNode(&Child1)

Child2 := entity.Category{
    Name:     "Child 2"
    ParentID:    &Parent.ID,
}

t.CreateNode(&Child2)
```

### Move up and down

```golang
status, err := t.MoveUp(Child2, 1) //move one place up, if possible
status, err := t.MoveDown(Child2, 1) //move one place down, if possible
```

### Delete node

```golang
t.DeleteNode(&entity.Category{
	ID: "a8c70ff6-6b4e-4caa-aebe-9ad368342c8e",
})
```