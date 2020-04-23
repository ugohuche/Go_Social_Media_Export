package Facebook

type Facebook struct {
  URL    string
  Name   string
  Friends int
}

func (f *Facebook) Feed() []string {
  return []string{
    " Facebook feeds",
    " Checkout my cool selfie",
  }
}

func (f *Facebook) Fame() int {
  return f.Friends
}
