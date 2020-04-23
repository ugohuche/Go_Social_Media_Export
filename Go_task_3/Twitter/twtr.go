package Twitter

type Twitter struct {
  URL      string
  Username string
  Followers int
}

func (t *Twitter) Feed() []string {
  return []string{
    " Twitter feeds",
    " Coding is actually fun",
  }
}

func (t *Twitter) Fame() int {
  return t.Followers
}
