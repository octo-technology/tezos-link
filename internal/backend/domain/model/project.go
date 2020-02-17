package model

type Project struct {
    ID   int64  `json:"id"`
    Name string `json:"title"`
    Key  string `json:"title"`
}

func NewProject(ID int64, name string, key string) Project {
    return Project{
        ID:   ID,
        Name: name,
        Key:  key,
    }
}
