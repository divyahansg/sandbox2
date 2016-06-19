package rect

type Rectangle struct {
    length, width int 
    name string
}

func MakeRect() Rectangle {
    r := Rectangle{1, 2, "test"}
    return r
}

