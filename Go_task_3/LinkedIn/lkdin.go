package LinkedIn

type LinkedIn struct {
  URl    string
  Name   string
  Connections  int
}

func (l *LinkedIn) Feed() []string  {
  return []string{
    " LinkedIn feeds",
    " I just got a new Software Developer job",
  }
}

func (l *LinkedIn) Fame() int {
  return l.Connections
}
