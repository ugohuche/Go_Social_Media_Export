package main

import(
  "Go_task_3/Exporter"
  "Go_task_3/Facebook"
  "Go_task_3/Twitter"
  "Go_task_3/LinkedIn"
)

func main() {
  fb := new(Facebook.Facebook)
  twtr := new(Twitter.Twitter)
  lkdin := new(LinkedIn.LinkedIn)

  err := Exporter.JSONExport(fb, "Facebookdata.json")
  err = Exporter.JSONExport(twtr, "Twitterdata.json")
  err = Exporter.JSONExport(lkdin, "LinkedIndata.json")

  if err != nil {
    panic(err)
  }
}
